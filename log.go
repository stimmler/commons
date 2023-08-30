package commons

import (
	"log"
)

func LogError(msg any) {
	log.Printf("| ERROR | %s", msg)
}

func LogInfo(msg any) {
	log.Printf("| INFO | %s", msg)
}
