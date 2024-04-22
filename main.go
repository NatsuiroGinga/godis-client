package main

import (
	"bufio"
	"bytes"
	"os"

	"godis-client/client"
	"godis-client/lib/logger"
	"godis-client/lib/utils"
	"godis-client/resp/parser"
)

func main() {
	// 1. connect Server
	godisClient, err := client.NewClient("127.0.0.1:6379")
	if err != nil {
		logger.Error(err)
	}
	godisClient.Start()

	localAddr := godisClient.Conn.LocalAddr()
	logger.Info("localAddr:", localAddr)
	/* listener, err := net.ListenTCP("tcp", localAddr.(*net.TCPAddr))
	if err != nil {
		logger.Error(err)
	}

	conn, err := listener.Accept()
	if err != nil {
		logger.Error(err)
	}
	logger.Info("conn:", conn) */
	/* reader := bufio.NewReader(conn)

	for {
		readBytes, err := reader.ReadBytes('\n')
		// handle the error
		if err != nil {
			if err == io.EOF { // if client closed, close the connection
				logger.Info("client closed")
			} else {
				logger.Warn("read error:", err)
			}

			return
		}
		logger.Info("readBytes:", string(readBytes))
		// trim suffix '\n'
		readBytes = readBytes[:len(readBytes)-1]
		cmd := utils.ToCmdLine3(readBytes) */

	// 2. receiver cmdline
	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	for scanner.Scan() {
		// 打印前缀

		readBytes := scanner.Bytes()
		cmd := utils.ToCmdLine3(readBytes)
		r := godisClient.Send(cmd)
		stream := parser.ParseStream(bytes.NewReader(r.Bytes()))
		payload := <-stream

		if err := client.Response(writer, payload.Data); err != nil {
			logger.Error(err)
			continue
		}

		writer.Flush()
	}
	if err := scanner.Err(); err != nil {
		logger.Error(err)
		return
	}

	logger.Info("connection closed")
	// }
}
