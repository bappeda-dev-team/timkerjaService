package helper

import (
	"fmt"
	"time"
)

func GenerateKodeTim(lastNumber int) string {
	year := time.Now().Year()
	return fmt.Sprintf("TIM-%d-%d", year, lastNumber+1)
}
