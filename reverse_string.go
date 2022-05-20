package reverse_string
import (
	"strings"
)

func ReverseString(input string) (output string) {
	k := strings.Split(input, "")
	last := len(k) - 1
	for i := 0; i < len(k)/2; i++ {
		k[i], k[last-i] = k[last-i], k[i]
	}
	output = strings.Join(k, "")
	return output
}
