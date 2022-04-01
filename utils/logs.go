package utils

import (
	"log"

	"github.com/fatih/color"
)

func StatusOk(statusCode string, method string, handler string) {
	log.Printf("[%s] %s %s", color.BlueString(method), color.GreenString(statusCode), handler)
}

func StatusError(statusCode string, method string, handler string) {
	log.Printf("[%s] %s %s", color.BlueString(method), color.RedString(statusCode), handler)
}
