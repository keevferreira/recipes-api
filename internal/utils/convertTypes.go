package utils

import (
	"fmt"
	"os"
	"strconv"
)

func StringToInt(str string) int {
	int, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("Erro ao converter string para inteiro:", err)
		os.Exit(1)
	}
	return int
}
