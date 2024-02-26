package protocol

import "fmt"

func (c Command) echo(msg string) ([]byte, error) {
	msg = fmt.Sprintf("%s%d\r\n%s\r\n", string(RespBulkString), len(msg), msg)
	return []byte(msg), nil
}
