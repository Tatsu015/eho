package eho

import "fmt"

type Degree int
type Lang string

const (
	ENE Degree = 75
	SSE Degree = 165
	WSW Degree = 225
	NNW Degree = 345
)

const (
	JA Lang = "ja"
	EN Lang = "en"
)

type DegreeName map[Degree]string

var JA_STR = DegreeName{
	ENE: "東北東",
	SSE: "南南東",
	WSW: "西南西",
	NNW: "北北西",
}

var EN_STR = DegreeName{
	ENE: "ENE",
	SSE: "SSE",
	WSW: "WSW",
	NNW: "NNW",
}

var MULTILINGUAL = map[Lang]DegreeName{
	"ja": JA_STR,
	"en": EN_STR,
}

func GetDegree(year int) Degree {
	onesPlace5Rem := (year % 10) % 5

	if onesPlace5Rem == 0 {
		return WSW
	} else if onesPlace5Rem == 2 {
		return NNW
	} else if onesPlace5Rem == 4 {
		return ENE
	} else {
		return SSE
	}
}

func DegreeToString(degree Degree, lang Lang) (string, error) {
	if _, ok := MULTILINGUAL[lang]; !ok {
		return "", fmt.Errorf("lang not exist: %v", lang)
	}
	if _, ok := MULTILINGUAL[lang][degree]; !ok {
		return "", fmt.Errorf("degree not exist: %v", degree)
	}
	dmap := MULTILINGUAL[lang]
	d := dmap[degree]
	return d, nil
}
