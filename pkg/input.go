package pkg

import (
	"fmt"
	"os"
	"regexp"
	"strconv"

	"github.com/manifoldco/promptui"
)

type InputError struct {
	Message string
}

func (e *InputError) Error() string {
	return e.Message
}

func InputText(text string, validation string) string {

	validate := func(input string) error {
		_, err := strconv.ParseFloat(input, 64)
		if validation == "number" && err != nil {
			return &InputError{Message: "Tolong berikan angka yang valid"}
		}

		matched, _ := regexp.MatchString(`^[0-9+\-*/\s]+$`, input)
		if validation == "math" && !matched {
			return &InputError{Message: "Hanya bisa menerima angka dan operator matematika"}
		}

		return nil
	}

	prompt := promptui.Prompt{
		Label:    text,
		Validate: validate,
		Stdin: os.Stdin,
	}

	result, err := prompt.Run()
	if err != nil {
		panic(fmt.Sprintf("Prompt failed %v\n", err))
	}

	return result

}

func InputSelect(question string, option []string) string {

	prompt := promptui.Select{
		Label: question,
		Items: option,
	}

	_, result, err := prompt.Run()
	if err != nil {
		panic(fmt.Sprintf("Prompt failed %v\n", err))
	}

	return result

}
