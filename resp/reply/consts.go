package reply

import (
	"sync"

	"godis-client/enum"
	"godis-client/interface/resp"
	"godis-client/lib/utils"
)

// 用于存储所有的回复, 使用懒加载的方式, 只有在需要的时候才会初始化且只会初始化一次
var replies = map[resp.Reply][]byte{}

// 用于保证只初始化一次
var (
	storePongReplyOnce           sync.Once
	storeOKReplyOnce             sync.Once
	storeNullBulkReplyOnce       sync.Once
	storeEmptyMultiBulkReplyOnce sync.Once
	storeNoReplyOnce             sync.Once
)

// 优化: 使用单例模式, 保证只有一个实例, 且只有在需要的时候才会初始化
var (
	thePongReply           *PongReply
	theOKReply             *OkReply
	theNullBulkReply       *NullBulkReply
	theEmptyMultiBulkReply *EmptyMultiBulkReply
	theNoReply             *NoReply
)

// PongReply 用于表示PONG的回复
type PongReply struct {
}

func NewPongReply() resp.Reply {
	storePongReplyOnce.Do(func() {
		thePongReply = new(PongReply)
		replies[thePongReply] = utils.String2Bytes(enum.PONG)
	})
	return thePongReply
}

func (reply *PongReply) Bytes() []byte {
	return replies[reply]
}

// OKReply 用于表示OK的回复
type OkReply struct {
}

// NewOKReply 用于创建OK的回复
func NewOKReply() resp.Reply {
	storeOKReplyOnce.Do(func() {
		theOKReply = new(OkReply)
		replies[theOKReply] = utils.String2Bytes(enum.OK)
	})
	return theOKReply
}

func (reply *OkReply) Bytes() []byte {
	return replies[reply]
}

// nullBulkReply 用于表示空的回复字符串
type NullBulkReply struct {
}

// NewNullBulkReply 用于创建空的回复字符串
func NewNullBulkReply() resp.Reply {
	storeNullBulkReplyOnce.Do(func() {
		theNullBulkReply = new(NullBulkReply)
		replies[theNullBulkReply] = utils.String2Bytes(enum.NIL)
	})
	return theNullBulkReply
}

func (reply *NullBulkReply) Bytes() []byte {
	return replies[reply]
}

// emptyMultiBulkReply 用于表示空的多条批量回复数组
type EmptyMultiBulkReply struct {
}

// NewEmptyMultiBulkReply 用于创建空的多条批量回复数组
func NewEmptyMultiBulkReply() resp.Reply {
	storeEmptyMultiBulkReplyOnce.Do(func() {
		theEmptyMultiBulkReply = new(EmptyMultiBulkReply)
		replies[theEmptyMultiBulkReply] = utils.String2Bytes(enum.EMPTY_BULK_REPLY)
	})
	return theEmptyMultiBulkReply
}

func (reply *EmptyMultiBulkReply) Bytes() []byte {
	return replies[reply]
}

// noReply 用于表示没有回复
type NoReply struct {
}

func NewNoReply() resp.Reply {
	storeNoReplyOnce.Do(func() {
		theNoReply = new(NoReply)
		replies[theNoReply] = utils.String2Bytes(enum.NO_REPLY)
	})
	return theNoReply
}

func (reply *NoReply) Bytes() []byte {
	return replies[reply]
}
