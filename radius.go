package radius

import (
	"net"
)

const (
	LocalAuth string = "127.0.0.1:1812"
	LocalAcct string = "127.0.0.1:1813"
)

type Server struct {
	in     *chan Packet
	secret string
	conn   *net.UDPConn
	open   bool
}

// Sets up a Server starting it listening on the specified addr.
// Pass a *chan Packet to have full control over how the Packets read
// from the server are handled
func Start(c *chan Packet, addr, secret string) (s *Server, err error) {
	s = &Server{c, secret, nil, true}
	err = s.listen(addr)

	return
}

// Open a udp connection on addr, and loop input in a goroutine
func (s *Server) listen(addr string) error {
	udpAddr, err := net.ResolveUDPAddr("udp", addr)

	if err != nil {
		return err
	}

	s.conn, err = net.ListenUDP("udp", udpAddr)

	if err != nil {
		return err
	}

	// loop reading
	go func() {
		for {
			buff := make([]byte, 1024)
			n, addr, err := s.conn.ReadFromUDP(buff)
			if (s.open) {
				*s.in <- Packet{buff, n, addr, err}
			} else {
				close(*s.in)
				return
			}
		}
	}()

	return nil
}

func (s *Server) Close() error {
	s.open = false
	err := s.conn.Close()
	return err
}
