package concatenate

import "bytes"

func ConcCopy() []byte {
	str := []string{"testing ", "string ", "for ", "test ", "concatenate "}
	concatenatedString := make([]byte, 100)
	byteItem := 0
	for _, item := range str {
		byteItem += copy(concatenatedString[byteItem:], []byte(item))
	}
	return concatenatedString
}

func ConcBuffer() string {
	strings := []string{"testing ", "string ", "for ", "test ", "concatenate "}
	buffer := bytes.Buffer{}
	for _, element := range strings {
		buffer.WriteString(element)
	}
	return buffer.String()
}
