package error

import (
	"fmt"
	"strings"
	"errors"
)

// FormatError returns a new formatted error given an error string and an input
func FormatError(input interface{}, err string) error {

	s := []string{
		err,
		`:`,
		`%v`,
	}
	f := strings.Join(s, " ")
	r := fmt.Sprintf(f, input)

	return errors.New(r)

}
