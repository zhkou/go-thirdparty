package thridparty

import "fmt"

// Sprint formats string
func Sprint(sum int) string {
	return fmt.Sprintf("in thirdparty: %v\n", sum)
}
