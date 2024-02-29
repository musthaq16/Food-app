package utils

import (
	"math/rand"
	"strconv"
)

// generate 6 Digit otp pin.
func OtpGenerate() string {
	randNumber := rand.Intn(1000000)
	return strconv.Itoa(randNumber)
}

func OtpVerication(otp string) {
	
}
