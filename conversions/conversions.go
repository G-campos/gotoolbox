package conversions

import (
	"fmt"
	"strconv"
)

func StrToHex(hexString string) string {
	number, err := strconv.Atoi(hexString)
	if err != nil {
		fmt.Println("Erro ao converter:", err)
		return ""
	}

	result := fmt.Sprintf("%X", number)
	//fmt.Println(result)

	return result
}
