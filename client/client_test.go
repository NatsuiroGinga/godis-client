package client_test

import (
	"bufio"
	"godis-client/lib/logger"
	"os"
	"testing"
)

func TestDemo(t *testing.T) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		readBytes := scanner.Bytes()
		logger.Info(string(readBytes))
	}
}
