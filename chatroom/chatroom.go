package main

import (
	. "deepgou.com/google_deep_go/util/simpleloger"
	"github.com/google/uuid"
	"io"
	"log"
	"net"
	"strings"
)

var onlineClients = make(map[string]net.Conn)  // map of Online Clients, the Connections
var messageChan = make(chan string, 1000) // channel of messages
var quitChan = make(chan bool)

func main() {

	listener, err := net.Listen("tcp", "127.0.0.1:8080")
	PanicError(err)
	defer listener.Close()

	go handleClientMsg()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Panicln(err)
			continue
		}

		go process(conn)

	}

	log.Println("Done")
}

func handleClientMsg() {
	for {
		select {
		case message := <-messageChan:
			//对消息进行解析
			doProcessMessage(message)
		case <-quitChan:
			break
		}
	}
}

func doProcessMessage(message string) {
	contents := strings.Split(message, "#")
	if len(contents) > 1 {
		uuidOfConn := contents[0]
		sendMessage := contents[1]

		uuidOfConn = strings.Trim(uuidOfConn, " ")

		//通过addr查看是否有链接对象
		if conn, ok := onlineClients[uuidOfConn]; ok {
			_, err := conn.Write([]byte(sendMessage))
			if err != nil {

			}
		}
	}
}

func storeConn(conn net.Conn) ( *string , error) {
	newUUID, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}
	//assign uuid to each conn
	uuidOfConn := newUUID.String()
	onlineClients[uuidOfConn] = conn
	return &uuidOfConn, nil

}

func writeACKtoClient(conn net.Conn, uuidOfConn string) error {
	for {
		n, err := conn.Write([]byte(uuidOfConn))
		if err != nil {
			return err
		}
		if len(uuidOfConn) == n {
			break
		}
		continue
	}
	return nil
}

func process(conn net.Conn) {
	defer conn.Close()
	// store each Conn with UUID
	uuidOfConn, err2 := storeConn(conn)
	if err2 != nil {
		log.Panicln(err2)
		return
	}
	// send UUID to each Conn
	err3 := writeACKtoClient(conn, *uuidOfConn)
	if err3 != nil {
		log.Panicln(err3)
		return
	}

	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)

		if err != nil {
			if err == io.EOF {
				continue
			}
			break
		}

		if n > 0{
			messageChan <- string(buf[:n])
		}
	}
}
