package validation_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/ythosa/rating-list-monitoring-platfrom-api/pkg/validation"
	"testing"
)

func TestSnils(t *testing.T) {
	testCases := []struct {
		name  string
		snils string
		ok    bool
	}{
		{
			name:  "valid snils",
			snils: "11223344595",
			ok:    true,
		},
		{
			name: "invalid check sum",
			snils: "11223344594",
			ok: false,
		},
		{
			name: "invalid check sum format",
			snils: "1122334459a",
			ok: false,
		},
		{
			name: "invalid snils format",
			snils: "11a23344595",
			ok: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			if tc.ok {
				assert.NoError(t, validation.Snils(tc.snils))
			} else {
				assert.Error(t, validation.Snils(tc.snils))
			}
		})
	}
}
