package utils

import "github.com/sirupsen/logrus"

func DefaultErrorHandler(err error, fnArr ...func()) {
	if err != nil {
		logrus.Println(err)

		for _, eachFn := range fnArr {
			eachFn()
		}
	}
}
