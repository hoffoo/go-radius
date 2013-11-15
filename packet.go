package radius

import (
	"math/rand"
	"net"
	"time"
	"encoding/binary"
)

// incrementing packet identification
var ident uint8

type Packet struct {
	buff []byte
	read int
	addr *net.UDPAddr
	err  error
}

func NewPacket(code, ident uint8, auth []byte, ats ...Attribute) Packet {
	buff := make([]byte, 6)

	buff[0] = byte(code)
	buff[1] = byte(ident)
	buff[3] = auth[0]
	buff[4] = auth[1]
}

func AuthPacket(ats ...Attribute) Packet {
	ident += 1

	c := Access_Request
	i := ident
	authenticator := DefaultAuthenticator()

	return NewPacket(c, i, authenticator)
}

const (
	Invalid_Type     uint8 = 0
	Access_Request   uint8 = 1
	Access_Accept    uint8 = 2
	Access_Reject    uint8 = 3
	Access_Challange uint8 = 11
)

const (
	code                int = 0
	identifier              = 1
	len_start               = 2
	len_end                 = 3
	authenticator_start     = 4
	authenticator_end       = 20
)

func (p *Packet) Code() uint8 {
	pt := packetType(p.buff[0])
	return pt
}

func (p *Packet) Identifier() uint8 {
	return uint8(p.buff[1])
}

func (p *Packet) Len() uint8 {
	return uint8(p.buff[2])
}

func packetType(code byte) uint8 {
	switch uint8(code) {
	case Access_Request:
		return Access_Request
	case Access_Accept:
		return Access_Accept
	case Access_Reject:
		return Access_Reject
	default:
		return Invalid_Type
	}
}

// Utility function to generate a pseudo random 2 byte auth
// Uses time.Now().Unix() as a Source
func DefaultAuthenticator() []byte {
	r := rand.New(rand.NewSource(time.Now().Unix()))

	res := make([]byte, 2)

	binary.LittleEndian.PutUint32(res, r.Uint32())

	return res
}
