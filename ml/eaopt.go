package main

import (
	"fmt"
	m "math"
	"math/rand"

	"github.com/MaxHalford/eaopt"
)

type Vector []float64

func (X Vector) Evaluate() (float64, error) {
	var (
		numerator   = 1 + m.Cos(12*m.Sqrt(m.Pow(X[0], 2)+m.Pow(X[1], 2)))
		denominator = 0.5*(m.Pow(X[0], 2)+m.Pow(X[1], 2)) + 2
	)
	return -numerator / denominator, nil
}

func (X Vector) Mutate(rng *rand.Rand) {
	eaopt.MutNormalFloat64(X, 0.8, rng)
}

func (X Vector) Crossover(Y eaopt.Genome, rng *rand.Rand) {
	eaopt.CrossUniformFloat64(X, Y.(Vector), rng)
}

func (X Vector) Clone() eaopt.Genome {
	var XX = make(Vector, len(X))
	copy(XX, X)
	return XX
}

func MakeVector(rng *rand.Rand) eaopt.Genome {
	return Vector(eaopt.InitUnifFloat64(2, -10, 10, rng))
}

func main() {
	var conf = eaopt.NewDefaultGAConfig()
	conf.NPops = 1
	var ga, err = conf.NewGA()
	if err != nil {
		fmt.Println(err)
		return
	}

	ga.Callback = func(ga *eaopt.GA) {
		fmt.Printf("Best fitness at generation %d: %f\n", ga.Generations, ga.HallOfFame[0].Fitness)
	}

	ga.Minimize(MakeVector)
}
