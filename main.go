package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"

	"godis-client/client"
	"godis-client/lib/logger"
	"godis-client/lib/utils"
	"godis-client/resp/parser"
)

func main() {
	// 1. connect Server
	godisClient, err := client.NewClient(client.ServerAddr)
	if err != nil {
		logger.Error(err)
	}
	godisClient.Start()
	logger.Info("connect server:", client.ServerAddr, "success")
	localAddr := godisClient.Conn.LocalAddr()
	logger.Info("local address is:", localAddr)

	// 2. scan command from cmdline
	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	for {
		fmt.Print(">> ")

		if !scanner.Scan() {
			break
		}

		// 2.1 read bytes
		readBytes := scanner.Bytes()
		cmd := utils.ToCmdLine3(readBytes)
		// 2.2 send command to server
		r := godisClient.Send(cmd)
		// 2.3 parse reply from server
		stream := parser.ParseStream(bytes.NewReader(r.Bytes()))
		payload := <-stream
		// 2.4 write to user cmd
		if err = client.Response(writer, payload.Data); err != nil {
			logger.Error(err)
			return
		}
		err = writer.Flush()
		if err != nil {
			return
		}
	}
	// 3. scan end
	if err = scanner.Err(); err != nil {
		logger.Error(err)
		return
	}
}
