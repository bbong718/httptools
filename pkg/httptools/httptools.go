package httptools

import "fmt"

// HTTPTools struct
type HTTPTools struct {
}

// New creates a new object to httptools
func New(http interface{}) *HTTPTools {

	fmt.Printf("Fuck is going on?")
	return &HTTPTools{}
}
