package helper

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateKodeTim(_ int) string {
	year := time.Now().Year()
	// Generate random 5-digit number between 10000-99999
	randomNum := 10000 + rand.Intn(90000)
	return fmt.Sprintf("TIM-%d-%d", year, randomNum)
}
