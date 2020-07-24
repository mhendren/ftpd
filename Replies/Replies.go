package Replies

import (
	"encoding/binary"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
)

type FTPReply struct {
	Code    int16
	Message string
	Interim bool
}

func CreateReplyNoAction() FTPReply {
	return FTPReply{
		Code:    -1,
		Message: "No action",
	}
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

func CreateReplyCommandOkayCustom(message string) FTPReply {
	return FTPReply{
		Code:    200,
		Message: message,
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
		Interim: true,
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

func CreateReplyEnteringPassiveMode(addr net.Addr) FTPReply {
	address := strings.Split(addr.String(), ":")
	if len(address) != 2 {
		return CreateReplySyntaxError()
	}
	netPort := make([]byte, 2)
	portNum, err := strconv.Atoi(address[1])
	if err != nil {
		return CreateReplySyntaxError()
	}
	binary.BigEndian.PutUint16(netPort, uint16(portNum))
	ipAddress := strings.Split(address[0], ".")
	if len(ipAddress) != 4 {
		return CreateReplySyntaxError()
	}
	_, _ = fmt.Fprintf(os.Stderr, "Entering Passive Mode (%v,%v,%v,%v,%v,%v)", ipAddress[0], ipAddress[1], ipAddress[2], ipAddress[3], netPort[0], netPort[1])
	return FTPReply{
		Code: 227,
		Message: fmt.Sprintf("Entering Passive Mode (%v,%v,%v,%v,%v,%v)",
			ipAddress[0], ipAddress[1], ipAddress[2], ipAddress[3], netPort[0], netPort[1]),
	}
}

func CreateReplyEnteringLongPassiveMode(address string) FTPReply {
	_, _ = fmt.Fprintf(os.Stderr, "Entering Passive Mode (%v)\n", address)
	return FTPReply{
		Code:    228,
		Message: fmt.Sprintf("Entering Passive Mode (%v)", address),
	}
}

func CreateReplyEnteringExtendedPassiveMode(addr net.Addr) FTPReply {
	tcpAddr, _ := net.ResolveTCPAddr("tcp", addr.String())
	return FTPReply{
		Code:    229,
		Message: fmt.Sprintf("Entering Extended Passive Mode (|||%v|)", tcpAddr.Port),
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

func CreateReplyUserLoggedInBySecurityExchange() FTPReply {
	return FTPReply{
		Code:    232,
		Message: "User logged in, authorized by security data exchange",
	}
}

func CreateReplySecurityDataExchangeComplete() FTPReply {
	return FTPReply{
		Code:    234,
		Message: "Security data exchange complete",
	}
}

func CreateReplySecurityDataExchangeCompleteWithData(data string) FTPReply {
	return FTPReply{
		Code:    235,
		Message: data,
	}
}

func CreateReplyFileActionCompleted(filename string) FTPReply {
	return FTPReply{
		Code:    250,
		Message: fmt.Sprintf("\"%v\" Requested file action okay, completed", filename),
	}
}

func CreateReplyPathnameCreated(pathname string, isCreated bool) FTPReply {
	createdFunc := func(str string) string {
		if isCreated {
			return fmt.Sprintf("%v created", str)
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

func CreateReplySecurityOkay(data string) FTPReply {
	return FTPReply{
		Code:    334,
		Message: data,
	}
}

func CreateReplySecurityAcceptable(data string) FTPReply {
	return FTPReply{
		Code:    335,
		Message: data,
	}
}

func CreateReplyNeedPasswordWithChallenge(challenge string) FTPReply {
	return FTPReply{
		Code:    336,
		Message: fmt.Sprintf("Username okay, need password. Challenge is \"%v\"", challenge),
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

func CreateReplySecurityResourceUnavailable() FTPReply {
	return FTPReply{
		Code:    431,
		Message: "Need some unavailable resource to process security",
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

func CreateReplySupportedAddressFamilies(list []int) FTPReply {
	message := "Supported address families are ("
	for i, af := range list {
		if i != 0 {
			message += ", "
		}
		message += fmt.Sprintf("%v", af)
	}
	message += ")"
	return FTPReply{
		Code:    521,
		Message: message,
	}
}

func CreateReplyUnsupportedExtendedPortProtocol(supportedProtocols string) FTPReply {
	message := fmt.Sprintf("Unsupported network protocol, use (%v)", supportedProtocols)
	return FTPReply{
		Code:    522,
		Message: message,
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

func CreateReplyCommandProtectionDeniedPolicy() FTPReply {
	return FTPReply{
		Code:    533,
		Message: "Command protection level denied for policy reasons",
	}
}

func CreateReplyRequestDeniedForPolicyReasons() FTPReply {
	return FTPReply{
		Code:    534,
		Message: "Request denied for policy reasons",
	}
}

func CreateReplyFailedSecurityCheck() FTPReply {
	return FTPReply{
		Code:    535,
		Message: "Failed security check (hash, sequence, etc)",
	}
}

func CreateReplyRequestedPROTNotSupported() FTPReply {
	return FTPReply{
		Code:    536,
		Message: "Requested PROT level not supported by mechanism",
	}
}

func CreateReplyCommandProtectionNotSupported() FTPReply {
	return FTPReply{
		Code:    537,
		Message: "Command protection level not supported by security mechanism",
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
		line := strings.TrimSuffix(messageSplit[0], "\n")
		str := fmt.Sprintf("%v- %v\r\n", ftpReply.Code, line)
		for i := 1; i < (count - 1); i++ {
			line = strings.TrimSuffix(messageSplit[i], "\n")
			if !ftpReply.Interim {
				str += fmt.Sprintf("  %v %v\r\n", ftpReply.Code, line)
			} else {
				str += fmt.Sprintf("%v\r\n", line)
			}
		}
		line = strings.TrimSuffix(messageSplit[count-1], "\n")
		str += fmt.Sprintf("%v %v\r\n", ftpReply.Code, line)
		return str
	}
	return fmt.Sprintf("%v %v\r\n", ftpReply.Code, ftpReply.Message)
}
