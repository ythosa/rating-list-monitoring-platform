package ratingparser

import (
	"github.com/PuerkitoBio/goquery"
	"gopkg.in/errgo.v2/fmt/errors"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

func SPBGU(url string, userSnils string) (*ParsingResult, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, errors.Newf("getting %s by HTML: %v", url, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, errors.Newf("analise %s by HTML: %v", url, err.Error())
	}

	title := doc.Find("p").Text()
	budgetPlacesRe := regexp.MustCompile(`КЦП по конкурсу: (\d+)`)
	budgetPlaces, err := strconv.Atoi(strings.Split(string(budgetPlacesRe.Find([]byte(title))), " ")[3])
	if err != nil {
		return nil, err
	}

	var (
		userScore        uint
		userPosition     uint
		priorityOneUpper uint
	)
	isUserFound := false
	doc.Find("tr").Each(func(_ int, s *goquery.Selection) {
		if _, exists := s.Attr("id"); exists {
			if isUserFound {
				return
			}

			parts := strings.Split(s.Text(), "\n")
			pos := strings.TrimSpace(parts[1])
			priority := strings.TrimSpace(parts[4])
			score := strings.TrimSpace(parts[5])
			snils := strings.TrimSpace(parts[2])
			if snils == userSnils {
				s, _ := strconv.Atoi(strings.Split(score, ",")[0])
				userScore = uint(s)
				p, _ := strconv.Atoi(pos)
				userPosition = uint(p)
				isUserFound = true

				return
			}

			p, _ := strconv.Atoi(priority)
			if p == 1 {
				priorityOneUpper += 1
			}
		}
	})

	if !isUserFound {
		return nil, errors.New("user not found")
	}

	return &ParsingResult{
		Position:         userPosition,
		Score:            userScore,
		PriorityOneUpper: priorityOneUpper,
		BudgetPlaces:     uint(budgetPlaces),
	}, nil
}
