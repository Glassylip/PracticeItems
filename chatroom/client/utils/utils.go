package utils

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"moudle/common/message"
	"net"
)

type Transfer struct {
	Conn net.Conn
	Buf  [8096]byte // When transmitting, use buffering
}

func (this *Transfer) ReadPkg() (mes message.Message, err error) {

	fmt.Println("Read the data sent by the client...")
	// conn.Read will block if conn is not closed
	// If the client closes conn, it will not block
	_, err = this.Conn.Read(this.Buf[:4])
	if err != nil {
		return
	}
	// convert to unit32
	pkgLen := binary.BigEndian.Uint32(this.Buf[0:4])

	// Read the message content according to pkgLen
	n, err := this.Conn.Read(this.Buf[:pkgLen])
	if n != int(pkgLen) || err != nil {
		return
	}
	//Deserialize pkgLen to -> message.Message
	err = json.Unmarshal(this.Buf[:pkgLen], &mes)
	if err != nil {
		fmt.Println("json.Unmarsha err=", err)
		return
	}
	return
}

func (this *Transfer) WritePkg(data []byte) (err error) {

	// Send a length to the other party first
	pkgLen := uint32(len(data))
	binary.BigEndian.PutUint32(this.Buf[0:4], pkgLen)
	// Send length
	n, err := this.Conn.Write(this.Buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(bytes) fail", err)
		return
	}

	// Send data
	n, err = this.Conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("conn.Write(bytes) fail", err)
		return
	}
	return
}
