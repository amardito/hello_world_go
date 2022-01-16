package randomize

import (
	"math/rand"
	"time"
)

func RandomGreetings() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}
	return formats[rand.Intn(len(formats))]
}

func init() {
	rand.Seed(time.Now().UnixNano())
}