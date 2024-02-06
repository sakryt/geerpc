package codec

import "io"

// 一次普通的RPC调用如下
// err = client.Call("Arith.Multiply", args, &reply)

type Header struct {
	ServiceMethod string // format "Service.Method"
	Seq           uint64 // sequence number chosen by client
	Error         string
}

type Codec interface {
	io.Closer                         // golang接口组合，这里组合了io.Closer接口
	ReadHeader(*Header) error         // 读取Header
	ReadBody(interface{}) error       // 读取Body
	Write(*Header, interface{}) error // 写
}

type NewCodecFunc func(io.ReadWriteCloser) Codec

// 给string类型起了个别名Type
type Type string

const (
	// 定义字符串常量时用Type别名指定类型
	GobType  Type = "application/gob"
	JsonType Type = "application/json"
)

var NewCodecFuncMap map[Type]NewCodecFunc

func init() {
	// 创建一个字典用于存放 类型与构造函数 的映射关系
	NewCodecFuncMap = make(map[Type]NewCodecFunc)
	NewCodecFuncMap[GobType] = NewGobCodec
}
