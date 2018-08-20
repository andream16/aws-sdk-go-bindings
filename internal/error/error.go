package error

import (
	"errors"
	"fmt"
	"strings"
)

// Format returns a new formatted error given an error string and an input
func Format(input interface{}, err string) error {

	s := []string{
		err,
		`:`,
		`%v`,
	}
	f := strings.Join(s, " ")
	r := fmt.Sprintf(f, input)

	return errors.New(r)

}
