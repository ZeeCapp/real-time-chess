package loggers

import "fmt"

type ConsoleLogger struct {
}

func (this ConsoleLogger) LogInfo(text string) {
	fmt.Printf("INFO | %s\n", text)
}

func (this ConsoleLogger) LogError(text string) {
	fmt.Printf("!!ERROR!! | %s\n", text)
}
