package client

import (
	"fmt"
	"godis-client/interface/resp"
	"godis-client/lib/logger"
	"godis-client/resp/reply"
	"io"
)

// Response 把resp格式的reply转化成阅读性高的消息写回给用户
func Response(writer io.Writer, data resp.Reply) (err error) {
	switch data := data.(type) {
	case *reply.MultiBulkReply:
		for i, b := range data.Args {
			fmt.Fprintf(writer, "(%d) ", i+1)
			// result = append(result, []byte("("+strconv.Itoa(i+1)+") ")...)
			if len(b) == 0 {
				fmt.Fprintln(writer, "nil")
				// result = append(result, []byte("nil")...)
			} else {
				fmt.Fprintf(writer, "%s\n", b)
				// result = append(result, b...)
			}
		}
	case *reply.BulkReply:
		fmt.Fprintf(writer, "%s\n", data.Arg)
		// _, err = writer.Write(append(data.Arg, '\n'))
	case *reply.IntReply:
		fmt.Fprintf(writer, "(int) %d\n", data.Code())
	case *reply.OkReply:
		fmt.Fprintln(writer, "OK")
	case *reply.EmptyMultiBulkReply:
		// logger.Debug("EmptyMultiBulkReply:", data)
		fmt.Fprintln(writer, "empty list or set")
	case *reply.NullBulkReply:
		fmt.Fprintln(writer, "nil")
		// logger.Debug("NullBulkReply:", data)
	case *reply.PongReply:
		fmt.Fprintln(writer, "PONG")
		// logger.Debug("PongReply:", data)
	case *reply.StatusReply:
		fmt.Fprintf(writer, "%s\n", data.Status)
		// logger.Debug("StatusReply:", data)
	case resp.ErrorReply:
		fmt.Fprintln(writer, data.Error())
		// logger.Debug("ErrorReply:", data)
	default:
		logger.Error("unknown reply")
	}
	return err
}
