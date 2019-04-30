package main

import (
	"log"
	"net"
	"strconv"

	. "deepgou.com/google_deep_go/util/simpleloger"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	PanicError(err)
	defer conn.Close()

	uuid := readUUID(conn, err)

	go sendMsg(conn, uuid)

	buf := make([]byte, 1024)
	for {
		numOfBytes, err := conn.Read(buf)
		PanicError(err)

		/*结尾buf[0:numOfBytes]的原因是：numOfBytes是指接收的字节数 如果只用string(buf)
	    	可能会导致接收字符串中有其他之前接收的字符
	    */
		log.Println("receive server message content:" + string(buf[0:numOfBytes]))
	}

	log.Println("Client program end!")

}

func readUUID(conn net.Conn, err error) string {
	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	PanicError(err)
	uuidOfConn := string(buf[:n])
	log.Println("client recieve uuid: " + uuidOfConn)
	return uuidOfConn
}


func sendMsg(conn net.Conn, uuidOfConn string) {
	msg := "Hello No.\n"
	for i := 1; i < 6; i++ {
		conn.Write([]byte(uuidOfConn + "#" + msg + strconv.Itoa(i)))
	}

	log.Println("Has sent the message!")
}
