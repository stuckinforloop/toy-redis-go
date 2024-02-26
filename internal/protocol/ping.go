package protocol

import "fmt"

func (c Command) ping() ([]byte, error) {
	msg := fmt.Sprintf("%sPONG\r\n", string(RespString))
	return []byte(msg), nil
}
