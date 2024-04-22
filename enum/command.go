package enum

// Command 命令枚举
type Command struct {
	name       string // 命令名称
	paramCount int    // 除去命令本身后的参数数量
}

// Name 返回命令名称
func (cmd *Command) Name() string {
	return cmd.name
}

// String equals to Name().
func (cmd *Command) String() string {
	return cmd.name
}

// ParamCount 返回命令参数数量
func (cmd *Command) ParamCount() int {
	return cmd.paramCount
}

// Arity 返回命令带命令本身的参数数量, 即 ParamCount() + 1
func (cmd *Command) Arity() int {
	return cmd.paramCount + 1
}

// keys
var (
	DEL         = &Command{name: "DEL", paramCount: -2}
	PING        = &Command{name: "PING", paramCount: 0}
	EXISTS      = &Command{name: "EXISTS", paramCount: -2}
	FLUSHDB     = &Command{name: "FLUSHDB", paramCount: 0}
	TYPE        = &Command{name: "TYPE", paramCount: 1}
	RENAME      = &Command{name: "RENAME", paramCount: 2}
	RENAMENX    = &Command{name: "RENAMENX", paramCount: 2}
	KEYS        = &Command{name: "KEYS", paramCount: 1}
	SELECT      = &Command{name: "SELECT", paramCount: 1}
	EXPIRE      = &Command{name: "EXPIRE", paramCount: 2}
	EXPIREAT    = &Command{name: "EXPIREAT", paramCount: 2}
	EXPIRETIME  = &Command{name: "EXPIRETIME", paramCount: 1}
	TTL         = &Command{name: "TTL", paramCount: 1}
	PEXPIRE     = &Command{name: "PEXPIRE", paramCount: 2}
	PEXPIREAT   = &Command{name: "PEXPIREAT", paramCount: 2}
	PEXPIRETIME = &Command{name: "PEXPIRETIME", paramCount: 1}
	PTTL        = &Command{name: "PTTL", paramCount: 1}
)

// string
var (
	GET    = &Command{name: "GET", paramCount: 1}
	SET    = &Command{name: "SET", paramCount: 2}
	SETNX  = &Command{name: "SETNX", paramCount: 2}
	STRLEN = &Command{name: "STRLEN", paramCount: 1}
	GETSET = &Command{name: "GETSET", paramCount: 2}
	INCR   = &Command{name: "INCR", paramCount: 1}
	DECR   = &Command{name: "DECR", paramCount: 1}
)

var (
	ZADD             = &Command{name: "ZADD", paramCount: -4}
	ZSCORE           = &Command{name: "ZSCORE", paramCount: 2}
	ZINCRBY          = &Command{name: "ZINCRBY", paramCount: 3}
	ZRANK            = &Command{name: "ZRANK", paramCount: 2}
	ZCOUNT           = &Command{name: "ZCOUNT", paramCount: 3}
	ZREVRANK         = &Command{name: "ZREVRANK", paramCount: 2}
	ZCARD            = &Command{name: "ZCARD", paramCount: 1}
	ZRANGE           = &Command{name: "ZRANGE", paramCount: -3}
	ZRANGEBYSCORE    = &Command{name: "ZRANGEBYSCORE", paramCount: -3}
	ZREVRANGE        = &Command{name: "ZREVRANGE", paramCount: -3}
	ZREVRANGEBYSCORE = &Command{name: "ZREVRANGEBYSCORE", paramCount: -3}
	ZPOPMIN          = &Command{name: "ZPOPMIN", paramCount: -1}
	ZPOPMAX          = &Command{name: "ZPOPMAX", paramCount: -1}
	ZREM             = &Command{name: "ZREM", paramCount: -2}
	ZREMRANGEBYSCORE = &Command{name: "ZREMRANGEBYSCORE", paramCount: 3}
	ZREMRANGEBYRANK  = &Command{name: "ZREMRANGEBYRANK", paramCount: 3}
)

const (
	WITH_SCORES = "WITHSCORES"
	LIMIT       = "LIMIT"
)
