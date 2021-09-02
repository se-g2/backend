package utils

import "log"

func SendMessage (number string, msg string) {
	log.Println("Message to ", number, ": ", msg)
}
