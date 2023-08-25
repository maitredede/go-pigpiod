package pigpiod

import (
	"context"
	"net"
)

func Connect(ctx context.Context, address string) (*Conn, error) {
	var dialer net.Dialer
	tcp, err := dialer.DialContext(ctx, "tcp", address)
	if err != nil {
		return nil, err
	}
	c := Conn{
		tcp:              tcp,
		duttyCycleRanges: make(map[int]uint32),
	}
	return &c, nil
}

type Conn struct {
	tcp net.Conn

	duttyCycleRanges map[int]uint32
}

func (c *Conn) Close() error {
	return c.tcp.Close()
}
