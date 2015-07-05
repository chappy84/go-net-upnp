// Copyright 2015 Satoshi Konno. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ssdp

import (
	"net"
)

// A UDPSocket represents a socket for UDP.
type UDPSocket struct {
	Conn      *net.UDPConn
	readBuf   []byte
	Interface *net.Interface
}

// NewUDPSocket returns a new UDPSocket.
func NewUDPSocket() *UDPSocket {
	uppSock := &UDPSocket{}
	uppSock.readBuf = make([]byte, MAX_PACKET_SIZE)
	return uppSock
}

// Close closes the current opened socket.
func (self *UDPSocket) Close() error {
	if self.Conn == nil {
		return nil
	}
	err := self.Conn.Close()
	if err != nil {
		return err
	}

	self.Conn = nil
	self.Interface = nil

	return nil
}

// Read reads from the current opend socket.
func (self *UDPSocket) Read() (*Packet, error) {
	n, from, err := self.Conn.ReadFrom(self.readBuf)
	if err != nil {
		return nil, err
	}

	ssdpPkt, err := NewPacketFromBytes(self.readBuf[:n])
	if err != nil {
		return nil, err
	}

	ssdpPkt.Interface = *self.Interface
	ssdpPkt.From = from

	return ssdpPkt, nil
}
