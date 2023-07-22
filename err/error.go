package err

import "log"

func ErrorHandler(err error, message string) {
	if err != nil {
		log.Fatalf("%s: %s\n", message, err.Error())
	}
}
