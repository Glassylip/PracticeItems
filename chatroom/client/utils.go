package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"moudle/common/message"
	"net"
)

func readPkg(conn net.Conn) (mes message.Message, err error) {

	buf := make([]byte, 8096)
	fmt.Println("Read the data sent by the client~~~...")
	// conn.Read will block if conn is not closed
	// If the client closes conn, it will not block
	_, err = conn.Read(buf[:4])
	if err != nil {
		return
	}

	pkgLen := binary.BigEndian.Uint32(buf[0:4])
	// Read the message content according to pkgLen
	n, err := conn.Read(buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		return
	}
	// Deserialize pkgLen to -> message.Message
	err = json.Unmarshal(buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("json.Unmarsha err=", err)
		return
	}
	return
}

func writePkg(conn net.Conn, data []byte) (err error) {

	// Send a length to the other party first
	var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkgLen)
	// Send length
	n, err := conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(bytes) fail", err)
		return
	}

	// Send data
	n, err = conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Write(bytes) fail", err)
		return
	}
	return
}
