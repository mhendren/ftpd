package Commands

import (
	"FTPserver/Connection"
	"FTPserver/Replies"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

type LPRT struct {
	cs *Connection.Status
}

type AddressFamily int

func (af AddressFamily) String() string {
	switch {
	case (af >= 1 && af <= 3) || (af >= 10 && af <= 14):
		return "unassigned"
	case af == 0 || af == 15:
		return "reserved"
	default:
		switch af {
		case 4:
			return "Internet Protocol"
		case 5:
			return "ST Datagram Mode"
		case 6:
			return "SIP"
		case 7:
			return "TP/IX"
		case 8:
			return "PIP"
		case 9:
			return "TUBA"
		case 16:
			return "Novell IPX"
		default:
			return "unknown address family"
		}
	}
}

func IPv6LPRTExecute(segments []string, cs *Connection.Status) Replies.FTPReply {
	getSegments := func(strings []string, expectedLength uint8) []uint8 {
		length, err := strconv.ParseUint(strings[0], 10, 8)
		if err != nil || uint8(length) != expectedLength || len(strings[1:]) != int(length) {
			return nil
		}
		segments := make([]uint8, length)
		for i, octet := range strings[1:] {
			segment, err := strconv.ParseUint(octet, 10, 8)
			if err != nil {
				return nil
			}
			segments[i] = uint8(segment)
		}
		return segments
	}

	getIP6Addr := func(addressSegmentStrings []string) net.IP {
		as := getSegments(addressSegmentStrings, 16)
		if as == nil {
			return net.IPv6unspecified
		}
		return net.IP{as[0], as[1], as[2], as[3], as[4], as[5], as[6], as[7],
			as[8], as[9], as[10], as[11], as[12], as[13], as[14], as[15]}
	}

	getIP6Port := func(portSegmentStrings []string) uint16 {
		ps := getSegments(portSegmentStrings, 2)
		if ps == nil {
			return 0
		}
		return (uint16(ps[0]) * 256) + uint16(ps[1])
	}

	ip6Addr := getIP6Addr(segments[0:17])
	if ip6Addr.IsUnspecified() {
		return Replies.CreateReplySyntaxErrorInParameters()
	}

	ip6Port := getIP6Port(segments[17:20])
	if ip6Port == 0 {
		return Replies.CreateReplySyntaxErrorInParameters()
	}

	destinationSocket := fmt.Sprintf("[%v]:%v", ip6Addr, ip6Port)
	localIP := strings.Split(cs.CommandConnection.LocalAddr().String(), ":")
	localPort, _ := strconv.Atoi(localIP[len(localIP)-1])
	localIPAddress := cs.CommandConnection.LocalAddr().(*net.TCPAddr).IP
	localSocket := fmt.Sprintf("[%v]:%v", localIPAddress, (uint16)(localPort-1))
	local, _ := net.ResolveTCPAddr("tcp6", localSocket)
	remote, _ := net.ResolveTCPAddr("tcp6", destinationSocket)
	conn, err := net.DialTCP("tcp6", local, remote)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "error opening data connection: %v\n", err)
		return Replies.CreateReplyCantOpenDataConnection()
	}
	cs.DataConnect(conn)
	return Replies.CreateReplyCommandOkay()
}

func (cmd LPRT) Execute(args string) Replies.FTPReply {
	if cmd.cs.EPSVAll {
		return Replies.CreateReplyBadCommandSequence()
	}
	cmd.cs.Type = Connection.TransferType(Connection.Active)
	segments := strings.Split(args, ",")
	afCode, err := strconv.ParseUint(segments[0], 10, 8)
	if err != nil {
		return Replies.CreateReplySyntaxErrorInParameters()
	}
	af := AddressFamily(afCode)
	switch af {
	case 6:
		return IPv6LPRTExecute(segments[1:], cmd.cs)
	}
	return Replies.CreateReplySupportedAddressFamilies([]int{6})
}
