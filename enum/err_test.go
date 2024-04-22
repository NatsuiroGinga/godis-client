package enum

import (
	"godis-client/lib/utils"
	"testing"
)

func TestDemo(t *testing.T) {
	cmd := toCmdLine2("lpush", []byte("list"), []byte("a"), []byte("c"))
	for _, line := range cmd {
		t.Log(string(line))
	}
}

func toCmdLine2(commandName string, args ...[]byte) [][]byte {
	result := make([][]byte, len(args)+1)
	result[0] = utils.String2Bytes(commandName)
	copy(result[1:], args)
	return result
}
