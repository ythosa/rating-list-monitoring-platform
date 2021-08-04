package validation_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ythosa/rating-list-monitoring-platform-api/pkg/validation"
)

func TestSnils(t *testing.T) {
	t.Parallel()

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
			name:  "valid snils",
			snils: "16691218387",
			ok:    true,
		},
		{
			name:  "valid snils",
			snils: "12358325649",
			ok:    true,
		},
		{
			name:  "invalid check sum",
			snils: "11223344594",
			ok:    false,
		},
		{
			name:  "invalid check sum format",
			snils: "1122334459a",
			ok:    false,
		},
		{
			name:  "invalid snils format",
			snils: "11a23344595",
			ok:    false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			if tc.ok {
				assert.NoError(t, validation.Snils(tc.snils))
			} else {
				assert.Error(t, validation.Snils(tc.snils))
			}
		})
	}
}
