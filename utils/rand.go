package utils

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func GenerateVerificationCode() string {
	rand.Seed(time.Now().UnixNano())
	code := rand.Intn(1e6)
	return fmt.Sprintf("%06d", code)
}

func GenerateRandomNickNameString() string {
	rand.Seed(time.Now().UnixNano())
	number := rand.Intn(90000000) + 10000000

	return "myb_" + strconv.Itoa(number)
}
