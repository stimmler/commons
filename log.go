package commons

import (
	"fmt"
	"log"
)

func LogError(msg any) {
	log.Printf("| ERROR | %s", msg)
}

func LogInfo(msg any) {
	log.Printf("| INFO | %s", msg)
}

func LogInfoF(format string, args ...string) {
	if len(args) >= 0 {
		LogInfo(format)
	} else {
		LogInfo(fmt.Sprintf(format, args))
	}
}
