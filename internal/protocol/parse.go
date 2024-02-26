package protocol

import (
	"bytes"
	"fmt"
	"strconv"
)

type RESP byte

const (
	RespString     RESP = '+'
	RespError      RESP = '-'
	RespInteger    RESP = ':'
	RespBulkString RESP = '$'
	RespArray      RESP = '*'
)

const Delimiter string = "\r\n"

func parseBulkString(data []byte) (string, error) {
	parts := bytes.Split(data, []byte(Delimiter))
	if len(parts) < 2 {
		return "", fmt.Errorf("invalid bulk string size: %v", len(parts))
	}

	strLength, err := strconv.Atoi(string(parts[0]))
	if err != nil {
		return "", fmt.Errorf("string to int conversion: %w", err)
	}

	str := parts[1]
	if len(str) < strLength {
		return "", fmt.Errorf("invalid bulk string: %v", data)
	}

	return string(str), nil
}

func parseArray(data []byte) ([]string, error) {
	parts := bytes.Split(data, []byte(Delimiter))
	if len(parts) < 1 {
		return nil, fmt.Errorf("invalid array: %s", data)
	}

	arrLength, err := strconv.Atoi(string(parts[0]))
	if err != nil {
		return nil, fmt.Errorf("string to int conversion: %w", err)
	}

	if arrLength == 0 {
		return []string{}, nil
	}

	if arrLength != (len(parts)-1)/2 {
		return nil, fmt.Errorf("no. of elements don't match array length: %v", data)
	}

	elements := []string{}
	for i := 1; i < len(parts)-1; i = i + 2 {
		parts[i] = append(parts[i], []byte(Delimiter)...)
		parts[i] = append(parts[i], parts[i+1]...)
		parts[i] = append(parts[i], []byte(Delimiter)...)

		data := parts[i]

		parsedElement, err := Parse(data)
		if err != nil {
			return nil, fmt.Errorf("parse array element: %w", err)
		}

		element, ok := parsedElement.(string)
		if !ok {
			return nil, fmt.Errorf("invalid element type: %v", parsedElement)
		}

		elements = append(elements, element)
	}

	return elements, nil
}

func Parse(data []byte) (any, error) {
	firstByte := data[0]
	data = data[1:]

	switch RESP(firstByte) {
	case RespBulkString:
		str, err := parseBulkString(data)
		if err != nil {
			return "", err
		}
		return str, nil
	case RespArray:
		elements, err := parseArray(data)
		if err != nil {
			return nil, err
		}
		return elements, nil
	default:
		return nil, nil
	}
}
