package thridparty

import "fmt"

// Sprint formats string
func Sprint(s string) string {
	return fmt.Sprintf("in thirdparty: %v\n", s)
}
