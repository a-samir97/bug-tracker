package helpers

import "log"

func handleError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s : %s", err, msg)
	}
}
