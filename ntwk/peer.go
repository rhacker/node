package ntwk

import (
	"bufio"
	"log"
	"net"
	"strconv"

	"github.com/pkg/errors"
)

type Peer struct {
	Address string `json:"Address"`

	Req_chan     chan Message
	Rep_chan     chan Message
	Out_req_chan chan Message
	Out_rep_chan chan Message
	Name         string //can set name
}

//peer functions
//get peers
//onReceiveBlock
//validateBlockSlot
//generateBlock
//loadBlocksFromPeer
//loadBlocksOffset
//getCommonBlock //Performs chain comparison with remote peer

func OpenConn(addr string) net.Conn {
	// Dial the remote process.
	log.Println("Dial " + addr)
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		//return nil, errors.Wrap(err, "Dialing "+addr+" failed")
	}
	if err != nil {
		log.Println("Error:", errors.WithStack(err))
	}
	return conn
}

// Open connects to a TCP Address
// It returns a TCP connection with a timeout wrapped into a buffered ReadWriter.
func Open(addr string) (*bufio.ReadWriter, error) {
	// Dial the remote process.
	log.Println("Dial " + addr)
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		return nil, errors.Wrap(err, "Dialing "+addr+" failed")
	}
	if err != nil {
		log.Println("Error:", errors.WithStack(err))
	}
	return bufio.NewReadWriter(bufio.NewReader(conn), bufio.NewWriter(conn)), nil
}

func OpenOut(ip string, Port int) *bufio.ReadWriter {
	rw, _ := Open(ip + strconv.Itoa(Port))
	return rw
}
