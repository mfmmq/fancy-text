package multiplier

import (
	"fmt"
	"math/rand"
	"time"
)

var randSource = rand.NewSource(time.Now().Unix())


func GetMultiplierFromSuperSlowService() int {
	r := rand.New(randSource)
	randomInt := r.Intn(2) + 1
	randomTime := r.Intn(2) + 1
	time.Sleep(time.Duration(randomInt) * time.Second)
	fmt.Printf("Multiplier service sent the value '%v'. Took %v seconds \n", randomInt, randomTime)
	return randomInt
}
