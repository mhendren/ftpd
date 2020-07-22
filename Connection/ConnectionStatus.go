package Connection

import (
	"FTPserver/Replies"
	"fmt"
	"net"
	"net/textproto"
	"os"
)

type TransferType int

const (
	Active int = iota
	Passive
)

type Structure int

const (
	File int = iota
	Record
	Page
)

type TransferMode int

const (
	Stream int = iota
	Block
	Compressed
)

func (tt TransferType) String() string {
	return [...]string{"Active", "Passive"}[tt]
}

func (stru Structure) String() string {
	return [...]string{"File", "Record", "Page"}[stru]
}

func (tm TransferMode) String() string {
	return [...]string{"Stream", "Block", "Compressed"}[tm]
}

type Status struct {
	Connected          bool
	CommandConnection  net.Conn
	TextProto          *textproto.Conn
	StandardConnection net.Conn
	DataConnected      bool
	DataConnection     net.Conn
	Remote             string
	Authenticated      bool
	Anonymous          bool
	User               string
	Account            string
	CurrentPath        string
	AcceptableCommands map[string]bool
	TypeCode           string
	FormCode           string
	LocalByteSize      uint8
	Type               TransferType
	RenameFrom         string
	Structure          Structure
	Mode               TransferMode
	Umask              os.FileMode
	IdleTimeout        int
	PreferredEProtocol int
	EPSVAll            bool
	Security           Security
}

func (cs *Status) Connect() {
	cs.Connected = true
}

func (cs *Status) Disconnect() {
	cs.Connected = false
}

func (cs *Status) DataConnect(conn net.Conn) {
	cs.DataConnected = true
	cs.DataConnection = conn
}

func (cs *Status) DataDisconnect() {
	if cs.DataConnection != nil {
		cs.DataConnection.Close()
	}
	cs.DataConnected = false
	cs.DataConnection = nil
}

func (cs *Status) IsConnected() bool {
	return cs.Connected
}

func (cs *Status) SetRemote(addr string) {
	cs.Remote = addr
}

func (cs *Status) SetAuthenticated(isAuth bool) {
	cs.Authenticated = isAuth
}

func (cs *Status) SetAnonymous(isAnonymous bool) {
	cs.Anonymous = isAnonymous
}

func (cs *Status) SetUser(username string) {
	cs.User = username
}

func (cs *Status) SetPath(path string) {
	cs.CurrentPath = path
}

func (cs *Status) SetCommands(commands map[string]bool) {
	cs.AcceptableCommands = commands
}

func (cs *Status) CanUseCommand(command string) bool {
	if cs.AcceptableCommands == nil {
		return true
	}
	_, ok := cs.AcceptableCommands[command]
	return ok
}

func (cs Status) SendFTPReply(reply Replies.FTPReply) {
	if reply.Code < 0 {
		_, _ = fmt.Fprintf(os.Stderr, "SendFTPReply: No action requested (communication handled by Command)\n")
		return
	}
	var err error
	_, err = fmt.Fprintf(cs.CommandConnection, "%v", reply)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "SendFTPReply: %v\n", err)
	}
}
