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
		lang      string
		expectStr string
		expectErr assert.ErrorAssertionFunc
	}{
		// ja
		"ENE-ja": {
			degree:    75,
			lang:      "ja",
			expectStr: "東北東",
			expectErr: assert.NoError,
		},
		"SSE-ja": {
			degree:    165,
			lang:      "ja",
			expectStr: "南南東",
			expectErr: assert.NoError,
		},
		"WSW-ja": {
			degree:    225,
			lang:      "ja",
			expectStr: "西南西",
			expectErr: assert.NoError,
		},
		"NNW-ja": {
			degree:    345,
			lang:      "ja",
			expectStr: "北北西",
			expectErr: assert.NoError,
		},
		// en
		"ENE-en": {
			degree:    75,
			lang:      "en",
			expectStr: "ENE",
			expectErr: assert.NoError,
		},
		"SSE-en": {
			degree:    165,
			lang:      "en",
			expectStr: "SSE",
			expectErr: assert.NoError,
		},
		"WSW-en": {
			degree:    225,
			lang:      "en",
			expectStr: "WSW",
			expectErr: assert.NoError,
		},
		"NNW-en": {
			degree:    345,
			lang:      "en",
			expectStr: "NNW",
			expectErr: assert.NoError,
		},
		"lang error": {
			degree:    345,
			lang:      "a",
			expectStr: "北北西",
			expectErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.EqualError(t, err, "lang not exist: a")
			},
		},
		"degree error": {
			degree:    0,
			lang:      "ja",
			expectStr: "北北西",
			expectErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return assert.EqualError(t, err, "degree not exist: 0")
			},
		},
	}
	for name, tt := range cases {
		t.Run(name, func(t *testing.T) {
			ja, err1 := DegreeToString(tt.degree, Lang(tt.lang))
			if !tt.expectErr(t, err1) || err1 != nil {
				return
			}
			assert.Equal(t, tt.expectStr, ja)
		})
	}
}
