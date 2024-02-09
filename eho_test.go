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
		lang      Lang
		expectStr string
	}{
		"ENE-ja": {
			degree:    75,
			lang:      "ja",
			expectStr: "東北東",
		},
		"SSE-ja": {
			degree:    165,
			lang:      "ja",
			expectStr: "南南東",
		},
		"WSW-ja": {
			degree:    225,
			lang:      "ja",
			expectStr: "西南西",
		},
		"NNW-ja": {
			degree:    345,
			lang:      "ja",
			expectStr: "北北西",
		},
	}
	for name, tt := range cases {
		t.Run(name, func(t *testing.T) {
			s := DegreeToString(tt.degree, tt.lang)
			assert.Equal(t, tt.expectStr, s)
		})
	}
}
