package protocol

type RespDataType string

const (
	SimpleString RespDataType = "simple_string"
)

var DataTypeToFirstByte = map[RespDataType]string{
	SimpleString: "+",
}
