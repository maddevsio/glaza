package lib

import "log"

// CheckError is a helper func for error checking
func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
