package util

import (
	"fmt"
	"os"
	"strconv"
)

func Try() {
	if r := recover(); r != nil {
		fmt.Println("Recovered from panic:", r)
	}
}

func StrToInt(str string) int {
	out, err := strconv.Atoi(str)

	if err != nil {
		panic("String Converter Error")
	}

	return out
}

func GetLimitIndex(limit string, index string) (int, int) {
	limits := StrToInt(limit)
	indexs := StrToInt(index)

	if limits >= 30 {
		limits = 30
	} else if limits <= 0 {
		limits = 1
	}

	if indexs >= 30 {
		indexs = 0
	}

	return limits, indexs
}

func EnvPortOr(port string) string {
	// If `PORT` variable in environment exists, return it
	if envPort := os.Getenv("PORT"); envPort != "" {
		return ":" + envPort
	}
	// Otherwise, return the value of `port` variable from function argument
	return ":" + port
}
