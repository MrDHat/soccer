package utils

import (
	"math/rand"
	"soccer-manager/constants"
	"time"

	"github.com/goombaio/namegenerator"
)

func RandomName() string {
	seed := time.Now().UTC().UnixNano()
	nameGenerator := namegenerator.NewNameGenerator(seed)

	return nameGenerator.Generate()
}

func RandomCountry() string {
	rand.Seed(time.Now().Unix()) // initialize global pseudo random generator
	return constants.Countries[rand.Intn(len(constants.Countries))]
}

func RandomAge() int64 {
	rand.Seed(time.Now().UnixNano())
	return int64(rand.Intn(constants.PlayerMaxAge-constants.PlayerMinAge+1) + constants.PlayerMinAge)
}

func RandomValuePercentage() int64 {
	rand.Seed(time.Now().UnixNano())
	return int64(rand.Intn(constants.PlayerMaxPercentageIncrease-constants.PlayerMinPercentageIncrease+1) + constants.PlayerMinPercentageIncrease)
}
