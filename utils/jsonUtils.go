package utils

import (
	"strings"
)

func IsJson(content string) bool {

	if strings.HasPrefix(content, "{") || strings.HasPrefix(content, "[") {
		return true
	}

	return false
}

func IsJsonByte(content []byte) bool {

	if strings.HasPrefix(string(content), "{") || strings.HasPrefix(string(content), "[") {
		return true
	}

	return false
}

func IsJsonByteObject(content []byte) bool {

	if strings.HasPrefix(string(content), "{")  {
		return true
	}

	return false
}

func IsJsonByteArray(content []byte) bool {

	if strings.HasPrefix(string(content), "[")  {
		return true
	}

	return false
}
