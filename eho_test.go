package eho

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetDirection(t *testing.T) {
	cases := map[string]struct {
		year         int
		expectDegree int
	}{
		"2024": {
			year:         2024,
			expectDegree: 75,
		},
		"2025": {
			year:         2025,
			expectDegree: 225,
		},
		"2026": {
			year:         2026,
			expectDegree: 165,
		},
		"2027": {
			year:         2027,
			expectDegree: 345,
		},
		"2028": {
			year:         2028,
			expectDegree: 165,
		},
		"2029": {
			year:         2029,
			expectDegree: 75,
		},
	}

	for name, tt := range cases {
		t.Run(name, func(t *testing.T) {
			degree := GetDegree(tt.year)
			assert.Equal(t, tt.expectDegree, int(degree))
		})
	}
}

func TestDegreeToString(t *testing.T) {
	cases := map[string]struct {
		degree    Degree
		expectJa  string
		expectEn  string
		expectErr assert.ErrorAssertionFunc
	}{
		"ENE-ja": {
			degree:    75,
			expectJa:  "東北東",
			expectEn:  "ENE",
			expectErr: assert.NoError,
		},
		"SSE-ja": {
			degree:    165,
			expectJa:  "南南東",
			expectEn:  "SSE",
			expectErr: assert.NoError,
		},
		"WSW-ja": {
			degree:    225,
			expectJa:  "西南西",
			expectEn:  "WSW",
			expectErr: assert.NoError,
		},
		"NNW-ja": {
			degree:    345,
			expectJa:  "北北西",
			expectEn:  "NNW",
			expectErr: assert.NoError,
		},
	}
	for name, tt := range cases {
		t.Run(name, func(t *testing.T) {
			ja, err1 := DegreeToString(tt.degree, "ja")
			en, err2 := DegreeToString(tt.degree, "en")
			if !tt.expectErr(t, err1) || err1 != nil {
				return
			}
			if !tt.expectErr(t, err2) || err2 != nil {
				return
			}
			assert.Equal(t, tt.expectJa, ja)
			assert.Equal(t, tt.expectEn, en)
		})
	}
}
