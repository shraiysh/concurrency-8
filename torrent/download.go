package torrent

import (
	"bytes"
	"encoding/binary"
	"net"
	"fmt"
	"github.com/concurrency-8/tracker"
)

type handler func([]byte)

// MakeHandshake is this.
func MakeHandshake(){
	report := tracker.GetRandomClientReport()
	_, err := BuildHandshake(*report)
	if err!=nil{
		return
	}
	service := "127.0.0.1:6000"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	if err!=nil{
		return
	}
	
	conn, err := net.DialTCP("tcp", nil, tcpAddr)


}

// onWholeMessage sends complete messages to callback function
func onWholeMessage(conn *net.IPConn, msgHandler handler) { // TODO add an extra argument for callback function i.e msgHandler
	buffer := new(bytes.Buffer)
	handshake := true
	resp := make([]byte, 100)

	for {

		respLen, err := conn.Read(resp)

		if err != nil {
			// TODO : close the connection and return
		}

		binary.Write(buffer, binary.BigEndian, resp[:respLen])

		var msgLen int

		if handshake {

			var length uint8
			binary.Read(buffer, binary.BigEndian, &length)
			msgLen = int(length + 49)
		} else {

			var length int32
			binary.Read(buffer, binary.BigEndian, &length)
			msgLen = int(length + 4)
		}

		for len(buffer.Bytes()) >= 4 && len(buffer.Bytes()) >= msgLen {
			// TODO implement msgHandler
			msgHandler((buffer.Bytes())[:msgLen])
			buffer = bytes.NewBuffer((buffer.Bytes())[msgLen:])
			handshake = false
		}
	}

}
