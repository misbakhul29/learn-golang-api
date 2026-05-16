package services

import (
	"fmt"
	"time"
)

type Logger struct {
	Name    string
	Message string
}

func Log(l Logger) {
	fmt.Printf("[%s]: %s - %s\n", l.Name, time.Now().Format(time.RFC1123), l.Message)
}
