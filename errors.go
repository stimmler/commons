package commons

import (
	"os"
)

func checkForError(err error) bool {
	return err != nil
}

func ExitOnError(err error, statusCode ...int) {
	if checkForError(err) {
		LogError(err.Error())

		if len(statusCode) > 0 {
			os.Exit(statusCode[0])
		} else {
			os.Exit(1)
		}
	}
}
