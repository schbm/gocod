package gocod

import (
	"errors"
	"math"
)

// [bit]
func DecisionContent(n int) (float64, error) {
	if n < 0 {
		return 0.0, errors.New("n cannot be smaller 0")
	}
	return math.Log2(float64(n)), nil
}

// [bit/s]
func DecisionFlow(n int, timeSpan float64) (float64, error) {
	if n < 0 {
		return 0.0, errors.New("n cannot be smaller 0")
	}
	if timeSpan <= 0 {
		return 0.0, errors.New("t cannot be smaller equal 0")
	}
	decisionContent, err := DecisionContent(n)
	if err != nil {
		return 0.0, err
	}
	return decisionContent / timeSpan, nil
}

// [bit]
func InformationContent(probability float64) (float64, error) {
	if probability <= 0 {
		return 0.0, errors.New("propability cannot be zero or below zero")
	}
	if probability > 1 {
		return 0.0, errors.New("probability cannot be greater than 1")
	}

	return math.Log2((1 / probability)), nil
}

func Entropy(probabilities []float64) (float64, error) {
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

func AvgCodewordLength(probabilities []float64, length []int) (float64, error) {
	if len(probabilities) != len(length) {
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

	for _, l := range length {
		if l <= 0 {
			return 0.0, errors.New("length must be greater than 0")
		}
	}

	result := 0.0
	for i := 0; i < len(probabilities); i++ {
		result += probabilities[i] * float64(length[i])
	}
	return result, nil
}

func SourceRedundancy(decisionContent float64, entropie float64) (float64, error) {
	if decisionContent <= 0 {
		return 0.0, errors.New("entscheidungsgehalt kann nicht 0 sein")
	}
	if decisionContent < entropie {
		return 0.0, errors.New("entscheidungsgehalt einer quelle kann nicht kleiner sein als entropie")
	}

	return decisionContent - entropie, nil
}

func CodeRedundancy(avgCodewordLength float64, entropie float64) (float64, error) {
	if avgCodewordLength <= 0 {
		return 0.0, errors.New("mittlere codewort länge kann nicht 0 sein")
	}
	if avgCodewordLength < entropie {
		return 0.0, errors.New("nach shannon kann die mcwl nicht kleiner sein als entropie")
	}

	return avgCodewordLength - entropie, nil
}
