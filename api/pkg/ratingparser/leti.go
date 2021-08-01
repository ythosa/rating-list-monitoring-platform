package ratingparser

import (
	"github.com/PuerkitoBio/goquery"
	"gopkg.in/errgo.v2/fmt/errors"
	"net/http"
	"strconv"
	"strings"
)

func leti(ratingURL string, userSnils string) (*ParsingResult, error) {
	res, err := http.Get(ratingURL)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()
	if res.StatusCode != http.StatusOK {
		return nil, errors.Newf("getting %s by HTML: %v", ratingURL, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, errors.Newf("analise %s by HTML: %v", ratingURL, err.Error())
	}

	var (
		userScore        uint
		userPosition     uint
		priorityOneUpper uint
	)
	isUserFound := false
	doc.Find("tbody tr").Each(func(_ int, s *goquery.Selection) {
		if isUserFound {
			return
		}

		if strings.TrimSpace(s.Text()) == "" {
			return
		}

		data := strings.TrimSpace(s.Text())
		parts := strings.Split(data, "\n")
		snils := strings.TrimSpace(parts[1])
		priority, _ := strconv.Atoi(strings.TrimSpace(parts[2]))

		if snils != userSnils {
			if priority == 1 {
				priorityOneUpper += 1
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
		return nil, UserNotFoundErr
	}

	return &ParsingResult{
		Position:         userPosition,
		Score:            userScore,
		PriorityOneUpper: priorityOneUpper,
		BudgetPlaces:     0,
	}, nil
}
