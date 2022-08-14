package logger

import "fmt"

// Print Info
func I(v ...interface{}) {
	fmt.Println("Info:  ", v)
}

// Print for Debug purpose
func D(v ...interface{}) {
	fmt.Println("Debug: ", v)
}

// Print for Error
func E(v ...interface{}) {
	fmt.Println("Error: ", v)
}
