package client

import (
	"godis-client/interface/resp"
	"godis-client/lib/logger"
	"godis-client/resp/reply"
	"net"
	"strconv"
)

// Response 把resp格式的reply转化成阅读性高的消息写回给用户
func Response(conn net.Conn, data resp.Reply) (err error) {
	switch data := data.(type) {
	case *reply.MultiBulkReply:
		var result []byte
		for i, b := range data.Args {
			result = append(result, []byte("("+strconv.Itoa(i+1)+") ")...)
			if len(b) == 0 {
				result = append(result, []byte("nil")...)
			} else {
				result = append(result, b...)
			}
			result = append(result, '\n')
		}
		_, err = conn.Write(result)
	case *reply.BulkReply:
		_, err = conn.Write(data.Arg)
	case *reply.IntReply:
		codeStr := strconv.FormatInt(data.Code(), 10)
		_, err = conn.Write([]byte("(int)" + codeStr))
	case *reply.OkReply:
		_, err = conn.Write([]byte("OK"))
	case *reply.EmptyMultiBulkReply:
		logger.Debug("EmptyMultiBulkReply:", data)
		_, err = conn.Write([]byte("empty list or set"))
	case *reply.NullBulkReply:
		logger.Debug("NullBulkReply:", data)
		_, err = conn.Write([]byte("nil"))
	case *reply.PongReply:
		logger.Debug("PongReply:", data)
		_, err = conn.Write([]byte("PONG"))
	case *reply.StatusReply:
		logger.Debug("StatusReply:", data)
		_, err = conn.Write([]byte(data.Status))
	case resp.ErrorReply:
		logger.Debug("ErrorReply:", data)
		_, err = conn.Write([]byte(data.Error()))
	default:
		logger.Error("unknown reply")
	}

	return err
}
