package protocol

type Command string

const (
	Ping Command = "ping"
	Echo Command = "echo"
)

func (c Command) Run(args []string) ([]byte, error) {
	switch c {
	case Ping:
		return c.ping()
	case Echo:
		return c.echo(args[0])
	default:
		return []byte("unkown command"), nil
	}
}
