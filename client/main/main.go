package main

import (
	"bufio"
	"godis-client/lib/logger"
	"godis-client/lib/utils"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		readBytes := scanner.Bytes()
		cmd := utils.ToCmdLine3(readBytes)
		for _, line := range cmd {
			logger.Info(string(line))
		}
	}
	if err := scanner.Err(); err != nil {
		logger.Error(err)
		return
	}
}
