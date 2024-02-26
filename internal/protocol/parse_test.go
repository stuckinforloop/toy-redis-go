package protocol

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseBulkString(t *testing.T) {
	testCases := []struct {
		bulkString []byte
		str        string
	}{
		{
			bulkString: []byte("$0\r\n\r\n"),
			str:        "",
		},
		{
			bulkString: []byte("$5\r\nhello\r\n"),
			str:        "hello",
		},
	}

	for _, testCase := range testCases {
		str, err := Parse(testCase.bulkString)
		assert.NoError(t, err)
		assert.Equal(t, testCase.str, str)
	}
}

func TestParseArray(t *testing.T) {
	testCases := []struct {
		array    []byte
		elements []string
	}{
		{
			array:    []byte("*0\r\n"),
			elements: []string{},
		},
		{
			array:    []byte("*2\r\n$5\r\nhello\r\n$5\r\nworld\r\n"),
			elements: []string{"hello", "world"},
		},
	}

	for _, testCase := range testCases {
		elements, err := Parse(testCase.array)
		assert.NoError(t, err)
		assert.ObjectsAreEqual(testCase.elements, elements)
	}
}
