package reverse

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
)

const DatagramLen = 2 + 1 + 1
const DatagramFlag = uint16(3553)
const DatagramVersion = uint8(1)
const (
	DatagramHeart = uint8(1) + iota
	DatagramSyn
	DatagramSynAck
	DatagramAck
)

type datagramStream struct {
	rw net.Conn
	r  [DatagramLen]byte
	w  [DatagramLen]byte
}

func (s *datagramStream) Flag() uint16 {
	return binary.BigEndian.Uint16(s.r[:])
}
func (s *datagramStream) Version() uint8 {
	return uint8(s.r[2])
}
func (s *datagramStream) Event() uint8 {
	return uint8(s.r[3])
}
func (s *datagramStream) Recv(events ...uint8) (e error) {
	_, e = io.ReadAtLeast(s.rw, s.r[:], DatagramLen)
	if e != nil {
		return
	}
	flag := s.Flag()
	if flag != DatagramFlag {
		e = fmt.Errorf(`%w: not supported flag=%v`, ErrProtocol, flag)
		return
	}
	version := s.Version()
	if version > DatagramVersion {
		e = fmt.Errorf(`%w: not supported version=%v`, ErrProtocol, version)
		return
	}
	event := s.Event()
	if event > DatagramAck || event < DatagramHeart {
		e = fmt.Errorf(`%w: not supported event=%v`, ErrProtocol, event)
		return
	}
	if len(events) != 0 {
		for _, evt := range events {
			if evt == event {
				return
			}
		}
		e = fmt.Errorf(`%w: unexpected event=%v`, ErrProtocol, event)
	}
	return
}
func (s *datagramStream) Send(evt uint8) (e error) {
	if evt > DatagramAck || evt < DatagramHeart {
		e = fmt.Errorf(`%w: not supported event=%v`, ErrProtocol, evt)
		return
	}

	if s.w[2] == 0 {
		binary.BigEndian.PutUint16(s.w[:], DatagramFlag)
		s.w[2] = DatagramVersion
	}
	s.w[3] = evt
	_, e = s.rw.Write(s.w[:])
	return
}
