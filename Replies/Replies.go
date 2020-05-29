package Replies

import (
	"encoding/binary"
	"fmt"
	"net"
	"strings"
)

type FTPReply struct {
	Code    int16
	Message string
}

func CreateReplyRestartMarker(userMarker uint32, serverMarker uint32) FTPReply {
	return FTPReply{
		Code:    110,
		Message: fmt.Sprintf("MARK %v = %v", userMarker, serverMarker),
	}
}

func CreateReplyServiceReadyInMinutes(minutes uint32) FTPReply {
	return FTPReply{
		Code:    120,
		Message: fmt.Sprintf("Service ready in %v minutes", minutes),
	}
}

func CreateReplyDataConnectionOpen() FTPReply {
	return FTPReply{
		Code:    125,
		Message: "Data connection already open; transfer starting",
	}
}

func CreateReplyFileStatusOkay() FTPReply {
	return FTPReply{
		Code:    150,
		Message: "File status okay; about to open data connection",
	}
}

func CreateReplyCommandOkay() FTPReply {
	return FTPReply{
		Code:    200,
		Message: "Command okay",
	}
}

func CreateReplyCommandNotImplementedSuperfluous() FTPReply {
	return FTPReply{
		Code:    202,
		Message: "Command not implemented, superfluous at this site",
	}
}

func CreateReplySystemStatus(message string) FTPReply {
	return FTPReply{
		Code:    211,
		Message: message,
	}
}

func CreateReplyDirectoryStatus(message string) FTPReply {
	return FTPReply{
		Code:    212,
		Message: message,
	}
}

func CreateReplyFileStatus(message string) FTPReply {
	return FTPReply{
		Code:    213,
		Message: message,
	}
}

func CreateReplyHelpMessage(message string) FTPReply {
	return FTPReply{
		Code:    214,
		Message: message,
	}
}

func CreateReplyNameSystemType(message string) FTPReply {
	return FTPReply{
		Code:    215,
		Message: message,
	}
}

func CreateReplyReadyForNewUser(message string) FTPReply {
	if message == "" {
		message = "Service ready for new user"
	}
	return FTPReply{
		Code:    220,
		Message: message,
	}
}

func CreateReplyClosingControlConnection(message string) FTPReply {
	if message == "" {
		message = "Service closing control connection"
	}
	return FTPReply{
		Code:    221,
		Message: message,
	}
}

func CreateReplyDataConnectionOpenNoTransfer() FTPReply {
	return FTPReply{
		Code:    225,
		Message: "Data connection open; no transfer in progress",
	}
}

func CreateReplyClosingDataConnection() FTPReply {
	return FTPReply{
		Code:    226,
		Message: "Closing data connection",
	}
}

func CreateReplyEnteringPassiveMode(addr net.IPAddr, port uint16) FTPReply {
	netPort := make([]byte, 2)
	binary.BigEndian.PutUint16(netPort, port)
	address := addr.IP.To4()
	return FTPReply{
		Code: 227,
		Message: fmt.Sprintf("Entering Passive Mode (%v,%v,%v,%v,%v,%v)",
			address[0], address[1], address[2], address[3], netPort[0], netPort[1]),
	}
}

func CreateReplyUserLoggedIn(message string) FTPReply {
	if message == "" {
		message = "User logged in, proceed"
	}
	return FTPReply{
		Code:    230,
		Message: message,
	}
}

func CreateReplyFileActionCompleted() FTPReply {
	return FTPReply{
		Code:    250,
		Message: "Requested file action okay, completed",
	}
}

func CreateReplyPathnameCreated(pathname string, isCreated bool) FTPReply {
	createdFunc := func(str string) string {
		if isCreated {
			return fmt.Sprintf("%v created")
		}
		return str
	}
	return FTPReply{
		Code:    257,
		Message: createdFunc(fmt.Sprintf("\"%v\"", pathname)),
	}
}

func CreateReplyNeedPassword() FTPReply {
	return FTPReply{
		Code:    331,
		Message: "User name okay, need password",
	}
}

func CreateReplyNeedAccount() FTPReply {
	return FTPReply{
		Code:    332,
		Message: "Need account for login",
	}
}

func CreateReplyPendingFurtherInformation() FTPReply {
	return FTPReply{
		Code:    350,
		Message: "Requested file action pending further information",
	}
}

func CreateReplyNotAvailableClosingConnection() FTPReply {
	return FTPReply{
		Code:    421,
		Message: "Service not available, closing control connection",
	}
}

func CreateReplyCantOpenDataConnection() FTPReply {
	return FTPReply{
		Code:    425,
		Message: "Can't open data connection",
	}
}

func CreateReplyConnectionClosedTransferAborted() FTPReply {
	return FTPReply{
		Code:    426,
		Message: "Connection closed; transfer aborted",
	}
}

func CreateReplyRequestedFileActionNotAvailable() FTPReply {
	return FTPReply{
		Code:    450,
		Message: "Requested file action not taken",
	}
}

func CreateReplyRequestedFileActionAborted() FTPReply {
	return FTPReply{
		Code:    451,
		Message: "Requested action aborted: local error in processing",
	}
}

func CreateReplyRequestedActionNotTaken() FTPReply {
	return FTPReply{
		Code:    452,
		Message: "Requested action not taken",
	}
}

func CreateReplySyntaxError() FTPReply {
	return FTPReply{
		Code:    500,
		Message: "Syntax error, command unrecognized",
	}
}

func CreateReplySyntaxErrorInParameters() FTPReply {
	return FTPReply{
		Code:    501,
		Message: "Syntax error in parameters or arguments",
	}
}

func CreateReplyCommandNotImplemented() FTPReply {
	return FTPReply{
		Code:    502,
		Message: "Command not implemented",
	}
}

func CreateReplyBadCommandSequence() FTPReply {
	return FTPReply{
		Code:    503,
		Message: "Bad sequence of commands",
	}
}

func CreateReplyCommandNotImplementedForParameter() FTPReply {
	return FTPReply{
		Code:    504,
		Message: "Command not implemented for that parameter",
	}
}

func CreateReplyNotLoggedIn() FTPReply {
	return FTPReply{
		Code:    530,
		Message: "Not logged in",
	}
}

func CreateReplyNeedAccountForStoring() FTPReply {
	return FTPReply{
		Code:    532,
		Message: "Need account for storing files",
	}
}

func CreateReplyRequestedActionNotTakenError() FTPReply {
	return FTPReply{
		Code:    550,
		Message: "Requested action not taken",
	}
}

func CreateReplyRequestedActionAborted() FTPReply {
	return FTPReply{
		Code:    551,
		Message: "Requested action aborted: page type unknown",
	}
}

func CreateReplyRequestedFileActionAbortedError() FTPReply {
	return FTPReply{
		Code:    552,
		Message: "Requested file action aborted",
	}
}

func CreateReplyRequestedActionNotTakenErrorFilename() FTPReply {
	return FTPReply{
		Code:    553,
		Message: "Requested action not taken",
	}
}

func (ftpReply FTPReply) String() string {
	messageSplit := strings.SplitAfter(ftpReply.Message, "\n")
	count := len(messageSplit)
	if count > 1 {
		str := fmt.Sprintf("%v- %v\r\n", ftpReply.Code, messageSplit[0])
		for i := 1; i < (count - 1); i++ {
			str += fmt.Sprintf("  %v %v\r\n", ftpReply.Code, messageSplit[i])
		}
		str += fmt.Sprintf("%v %v\r\n", ftpReply.Code, messageSplit[count-1])
		return str
	}
	return fmt.Sprintf("%v %v\r\n", ftpReply.Code, ftpReply.Message)
}
