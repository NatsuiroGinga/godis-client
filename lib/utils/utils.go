package utils

import (
	"bytes"
)

func ToCmdLine3(cmd []byte) [][]byte {
	if len(cmd) > 0 && cmd[len(cmd)-1] == '\n' {
		cmd = cmd[:len(cmd)-1]
	}
	// trim front and back space
	cmd = bytes.TrimSpace(cmd)
	// split bytes
	params := bytes.Split(cmd, String2Bytes(" "))
	result := make([][]byte, 0, len(params))
	for _, param := range params {
		if len(param) > 0 { // delete empty bytes
			result = append(result, param)
		}
	}
	return result
}

// BytesEquals check whether the given bytes is equal
func BytesEquals(a, b []byte) bool {
	return bytes.Equal(a, b)
}

// If returns trueVal if condition is true, otherwise falseVal.
func If[T any](condition bool, trueVal, falseVal T) T {
	if condition {
		return trueVal
	}
	return falseVal
}

// If2Kinds returns trueVal if condition is true, otherwise falseVal.
//
// This function is used to avoid the type of trueVal and falseVal is not the same.
func If2Kinds(condition bool, trueVal, falseVal any) any {
	if condition {
		return trueVal
	}
	return falseVal
}
