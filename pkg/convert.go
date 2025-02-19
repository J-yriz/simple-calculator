package pkg

import "strconv"

type ConvertError struct {
	Message string
}

func (e *ConvertError) Error() string {
	return e.Message
}

func StrToNmbr(input string) (int64, error) {

	result, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		return 0, &ConvertError{Message: "Tolong berikan angka yang valid"}
	}

	return result, nil

}

func NmbrToStr(input int64) (string, error) {
	
	result := strconv.FormatInt(input, 10)
	return result, nil

}
