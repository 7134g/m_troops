package connect

import (
	"log"
	"net"
	"time"
)

// NewDial 优化通过tcp下载文件超时逻辑
func NewDial(network, address string, interval time.Duration) (Conn, error) {
	var err error
	c := Conn{}
	c.timeout = interval
	c.ticker = time.NewTicker(c.timeout)
	c.nc, err = net.Dial(network, address)

	go func() {
		select {
		case <-c.ticker.C:
			log.Println("interval timeout")
			_ = c.nc.Close()
		}
	}()

	return c, err
}

type Conn struct {
	nc net.Conn

	timeout time.Duration
	ticker  *time.Ticker
}

func (c *Conn) Read(b []byte) (n int, err error) {
	c.ticker.Reset(c.timeout)
	return c.nc.Read(b)
}

func (c *Conn) Write(b []byte) (n int, err error) {
	c.ticker.Reset(c.timeout)
	return c.nc.Write(b)
}

func (c *Conn) Close() error {
	return c.nc.Close()
}

func (c *Conn) LocalAddr() net.Addr {
	return c.nc.LocalAddr()
}

func (c *Conn) RemoteAddr() net.Addr {
	return c.nc.RemoteAddr()
}

func (c *Conn) SetDeadline(t time.Time) error {
	return c.nc.SetDeadline(t)
}

func (c *Conn) SetReadDeadline(t time.Time) error {
	return c.nc.SetReadDeadline(t)
}

func (c *Conn) SetWriteDeadline(t time.Time) error {
	return c.nc.SetWriteDeadline(t)
}
