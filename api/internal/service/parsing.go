package service

import (
	"errors"
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/valyala/fasthttp"

	"github.com/ythosa/rating-list-monitoring-platform-api/internal/cache"
	"github.com/ythosa/rating-list-monitoring-platform-api/internal/config"
	"github.com/ythosa/rating-list-monitoring-platform-api/internal/dto"
	"github.com/ythosa/rating-list-monitoring-platform-api/internal/logging"
)

type ParsingImpl struct {
	client fasthttp.Client
	cache  cache.RatingList
	logger *logging.Logger
}

func NewParsingImpl(cache cache.RatingList) *ParsingImpl {
	return &ParsingImpl{
		client: fasthttp.Client{
			ReadTimeout:         config.Get().Server.ReadTimeout,
			MaxConnsPerHost:     config.Get().Parsing.MaxConnsPerHost,
			ReadBufferSize:      config.Get().Parsing.ReadBufferSize,
			MaxResponseBodySize: config.Get().Parsing.MaxResponseBodySize,
		},
		cache:  cache,
		logger: logging.NewLogger("parsing service"),
	}
}

var ErrUserNotFoundInRatingList = errors.New("user not found in rating list")

const (
	consentStatusSubmitted = "Да"
	priorityOne            = 1
)

func (p *ParsingImpl) ParseRating(university string, ratingURL string, userSnils string) (*dto.ParsingResult, error) {
	ratingList, err := p.cache.Get(ratingURL)
	if err != nil {
		res, req := fasthttp.AcquireResponse(), fasthttp.AcquireRequest()
		req.SetRequestURI(ratingURL)

		if err := p.client.Do(req, res); err != nil {
			return nil, fmt.Errorf("error while getting rating list page: %w", err)
		}

		if res.StatusCode() != fasthttp.StatusOK {
			return nil, fmt.Errorf("getting %s by HTML: %v", ratingURL, res.StatusCode())
		}

		ratingList = string(res.Body())
		if err := p.cache.Save(ratingURL, ratingList, config.Get().Parsing.RatingListTTL); err != nil {
			return nil, fmt.Errorf("error while caching rating list: %w", err)
		}
	}

	parsedRatingList, err := goquery.NewDocumentFromReader(ioutil.NopCloser(strings.NewReader(ratingList)))
	if err != nil {
		return nil, fmt.Errorf("analise by HTML error: %w", err)
	}

	switch university {
	case "ЛЭТИ":
		return p.parseLETI(parsedRatingList, p.formatSnils(userSnils))
	case "СПБГУ":
		return p.parseSPBGU(parsedRatingList, p.formatSnils(userSnils))
	default:
		return nil, fmt.Errorf("invalid university: %s", university)
	}
}

func (p *ParsingImpl) formatSnils(snils string) string {
	return fmt.Sprintf("%s-%s-%s %s", snils[:3], snils[3:6], snils[6:9], snils[9:11])
}

func (p *ParsingImpl) parseLETI(ratingList *goquery.Document, userSnils string) (*dto.ParsingResult, error) {
	var (
		userScore             uint
		userPosition          uint
		priorityOneUpper      uint
		submittedConsentUpper uint
	)

	isUserFound := false

	ratingList.Find("tbody tr").Each(func(_ int, s *goquery.Selection) {
		if isUserFound {
			return
		}

		if strings.TrimSpace(s.Text()) == "" {
			return
		}

		data := strings.TrimSpace(s.Text())
		parts := strings.Split(data, "\n")
		priority, _ := strconv.Atoi(strings.TrimSpace(parts[2]))
		consentStatus := strings.TrimSpace(parts[11])

		snils := strings.TrimSpace(parts[1])
		if snils != userSnils {
			if priority == priorityOne {
				priorityOneUpper++
			}

			if consentStatus == consentStatusSubmitted {
				submittedConsentUpper++
			}

			return
		}

		isUserFound = true

		position, _ := strconv.Atoi(parts[0])
		score, _ := strconv.Atoi(strings.TrimSpace(parts[4]))
		userPosition = uint(position)
		userScore = uint(score)
	})

	if !isUserFound {
		return nil, ErrUserNotFoundInRatingList
	}

	return &dto.ParsingResult{
		Position:              userPosition,
		Score:                 userScore,
		PriorityOneUpper:      priorityOneUpper,
		SubmittedConsentUpper: submittedConsentUpper,
		BudgetPlaces:          0,
	}, nil
}

func (p *ParsingImpl) parseSPBGU(ratingList *goquery.Document, userSnils string) (*dto.ParsingResult, error) {
	title := ratingList.Find("p").Text()
	budgetPlacesRe := regexp.MustCompile(`КЦП по конкурсу: (\d+)`)

	budgetPlaces, err := strconv.Atoi(strings.Split(string(budgetPlacesRe.Find([]byte(title))), " ")[3])
	if err != nil {
		return nil, fmt.Errorf("error while parsing budget places: %w", err)
	}

	var (
		userScore             uint
		userPosition          uint
		priorityOneUpper      uint
		submittedConsentUpper uint
	)

	isUserFound := false

	ratingList.Find("tr").Each(func(_ int, s *goquery.Selection) {
		if _, exists := s.Attr("id"); !exists {
			return
		}

		if isUserFound {
			return
		}

		parts := strings.Split(s.Text(), "\n")
		priority := strings.TrimSpace(parts[4])
		consentStatus := strings.TrimSpace(parts[11])

		snils := strings.TrimSpace(parts[2])
		if snils != userSnils {
			p, _ := strconv.Atoi(priority)
			if p == priorityOne {
				priorityOneUpper++
			}

			if consentStatus == consentStatusSubmitted {
				submittedConsentUpper++
			}

			return
		}

		isUserFound = true

		score, _ := strconv.Atoi(strings.Split(strings.TrimSpace(parts[5]), ",")[0])
		userScore = uint(score)
		pos, _ := strconv.Atoi(strings.TrimSpace(parts[1]))
		userPosition = uint(pos)
	})

	if !isUserFound {
		return nil, ErrUserNotFoundInRatingList
	}

	return &dto.ParsingResult{
		Position:              userPosition,
		Score:                 userScore,
		PriorityOneUpper:      priorityOneUpper,
		SubmittedConsentUpper: submittedConsentUpper,
		BudgetPlaces:          uint(budgetPlaces),
	}, nil
}
