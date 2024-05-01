package validates

import "fmt"

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func ErrParamIsRequired(name, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}
