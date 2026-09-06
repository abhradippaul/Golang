package fileops

import (
	"os"
	"fmt"
	"errors"
	"github.com/fatih/color"
)

func WriteToFile(fileName string, investmentAmount, fv, rfv float64) {
	balanceText := fmt.Sprintf("The investment amount is %.2f. Your future value wil be %.2f and actual future value will be %.2f.\n", investmentAmount, fv, rfv)
	os.WriteFile(fileName, []byte(balanceText), 0644)
}

func ReadFromFile(fileName string) {
	text, err := os.ReadFile(fileName)
	if err != nil {
		customError := errors.New("Failed to read file")
		color.Red(customError.Error())
	}
	color.Cyan(string(text))
}