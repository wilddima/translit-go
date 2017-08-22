package translit

import (
	"bytes"
)

func Translit(s string) string {
	var buffer bytes.Buffer

	for _, v := range s {
    if char, ok := dictionary[string(v)]; ok {
      buffer.WriteString(char)
    } else {
      buffer.WriteString(string(v))
    }
	}

	return buffer.String()
}
