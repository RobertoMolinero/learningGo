package main

import (
	"fmt"
	"github.com/pkg/errors"
	"log"
)

type mathOperationError struct {
	lat  string
	long string
	err  error
}

func (me mathOperationError) Error() string {
	return fmt.Sprintf("math operation error: %v | %v | %v", me.lat, me.long, me.err)
}

func main() {
	_, e := mathOperation(-10.23)

	if e != nil {
		log.Println(e)
	}
}

func mathOperation(f float64) (float64, error) {
	if f < 0 {
		e := errors.New("Wrong!")
		return 0, mathOperationError{
			lat:  "50.2289 N",
			long: "99.4656 W",
			err:  e,
		}
	}
	return 42, nil
}
