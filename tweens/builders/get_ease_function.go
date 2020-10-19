package builders

import "github.com/gravestench/pho/phomath/easing"

func GetEaseFunction(ease interface{}, params []float64) func(float64) float64 {
	switch e := ease.(type) {
	case string:
		provider, found := easing.EaseMap[e]
		if found {
			return provider.New(params)
		}
	case func(float64) float64:
		return e
	}

	return easing.EaseMap[easing.Default].New(params)
}