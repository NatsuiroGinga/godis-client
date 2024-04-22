package client

import (
	"fmt"
	"io"

	"godis-client/interface/resp"
	"godis-client/lib/logger"
	"godis-client/resp/reply"
)

// Response 把resp格式的reply转化成阅读性高的消息写回给用户
func Response(writer io.Writer, data resp.Reply) (err error) {
	switch data := data.(type) {
	case *reply.MultiBulkReply:
		for i, b := range data.Args {
			_, err = fmt.Fprintf(writer, "(%d) ", i+1)
			if len(b) == 0 {
				_, err = fmt.Fprintln(writer, "nil")
			} else {
				_, err = fmt.Fprintf(writer, "%s\n", b)
			}
		}
	case *reply.BulkReply:
		_, err = fmt.Fprintf(writer, "%s\n", data.Arg)
	case *reply.IntReply:
		_, err = fmt.Fprintf(writer, "(int) %d\n", data.Code())
	case *reply.OkReply:
		_, err = fmt.Fprintln(writer, "OK")
	case *reply.EmptyMultiBulkReply:
		_, err = fmt.Fprintln(writer, "empty list or set")
	case *reply.NullBulkReply:
		_, err = fmt.Fprintln(writer, "nil")
	case *reply.PongReply:
		_, err = fmt.Fprintln(writer, "PONG")
	case *reply.StatusReply:
		_, err = fmt.Fprintf(writer, "%s\n", data.Status)
	case resp.ErrorReply:
		_, err = fmt.Fprintln(writer, data.Error())
	default:
		logger.Error("unknown reply")
	}
	return err
}
