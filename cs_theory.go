package gocod

import (
	"errors"
	"math"
)

// [bit]
func Entscheidungsgehalt(n int) (float64, error) {
	if n < 0 {
		return 0.0, errors.New("n cannot be smaller 0")
	}

	return math.Log2(float64(n)), nil
}

// [bit/s]
func Entscheidungsfluss(n int, t float64) (float64, error) {
	if n < 0 {
		return 0.0, errors.New("n cannot be smaller 0")
	}
	if t <= 0 {
		return 0.0, errors.New("t cannot be smaller equal 0")
	}
	eg, err := Entscheidungsgehalt(n)
	if err != nil {
		return 0.0, err
	}
	return eg / t, nil
}

// [bit]
func Informationsgehalt(p float64) (float64, error) {
	if p <= 0 {
		return 0.0, errors.New("propability cannot be zero or below zero")
	}
	if p > 1 {
		return 0.0, errors.New("probability cannot be greater than 1")
	}

	return math.Log2((1 / p)), nil
}

func Entropie(probabilities []float64) (float64, error) {
	sum := 0.0
	for _, p := range probabilities {
		if p <= 0 {

			return 0.0, errors.New("propability cannot be zero or below zero")
		}
		if p > 1 {
			return 0.0, errors.New("probability cannot be greater than 1")
		}
		sum += p
	}
	if sum != 1.0 {
		return 0.0, errors.New("probabilities must add up to 1")
	}

	//calculate
	result := 0.0
	for _, p := range probabilities {
		result += p * math.Log2(p)
	}
	return (-1.0) * result, nil
}

func MittlereCodeWortLänge(probabilities []float64, länge []int) (float64, error) {
	if len(probabilities) != len(länge) {
		return 0.0, errors.New("amount of probabilities and codewort length have to match")
	}

	sum := 0.0
	for _, p := range probabilities {
		if p <= 0 {

			return 0.0, errors.New("propability cannot be zero or below zero")
		}
		if p > 1 {
			return 0.0, errors.New("probability cannot be greater than 1")
		}
		sum += p
	}
	if sum != 1.0 {
		return 0.0, errors.New("probabilities must add up to 1")
	}

	for _, l := range länge {
		if l <= 0 {
			return 0.0, errors.New("length must be greater than 0")
		}
	}

	result := 0.0
	for i := 0; i < len(probabilities); i++ {
		result += probabilities[i] * float64(länge[i])
	}
	return result, nil
}

func RedundanzQuelle(entscheidungsGehalt float64, entropie float64) (float64, error) {
	if entscheidungsGehalt <= 0 {
		return 0.0, errors.New("entscheidungsgehalt kann nicht 0 sein")
	}
	if entscheidungsGehalt < entropie {
		return 0.0, errors.New("entscheidungsgehalt einer quelle kann nicht kleiner sein als entropie")
	}

	return entscheidungsGehalt - entropie, nil
}

func RedundanzCode(mittlereCodeWLänge float64, entropie float64) (float64, error) {
	if mittlereCodeWLänge <= 0 {
		return 0.0, errors.New("mittlere codewort länge kann nicht 0 sein")
	}
	if mittlereCodeWLänge < entropie {
		return 0.0, errors.New("nach shannon kann die mcwl nicht kleiner sein als entropie")
	}

	return mittlereCodeWLänge - entropie, nil
}
