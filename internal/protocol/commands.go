package protocol

type Command string

const (
	Ping Command = "ping"
)

func (c Command) Valid() bool {
	switch c {
	case Ping:
		return true
	default:
		return false
	}
}
