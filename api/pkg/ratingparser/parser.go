package ratingparser

import "gopkg.in/errgo.v2/fmt/errors"

func ParseRating(university string, ratingURL string, userSnils string) (*ParsingResult, error) {
	switch university {
	case "ЛЭТИ":
		return leti(ratingURL, userSnils)
	case "СПБГУ":
		return spbgu(ratingURL, userSnils)
	default:
		return nil, errors.Newf("invalid university: %s", university)
	}
}
