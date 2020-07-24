package Connection

import (
	"crypto/tls"
	"fmt"
	"io"
	"net"
	"os"
)

type DataConnection struct {
	TransferType    TransferType
	Protocol        string
	LocalAddress    string
	RemoteAddress   string
	Connection      net.Conn
	PassivePort     int
	PassiveListener net.Listener
	Security        Security
	IsSetup         bool
}

type DataConnectionError struct {
	Err string
}

func (dce *DataConnectionError) Error() string {
	if dce == nil {
		return "<nil>"
	}
	return fmt.Sprintf("DataConnectionError: %v", dce.Err)
}

func (dc *DataConnection) Setup() error {
	ActiveSetup := func() error {
		if dc.LocalAddress == "" {
			return &DataConnectionError{Err: "Local address not set for Active Connection"}
		}
		dc.IsSetup = true
		return nil
	}
	PassiveSetup := func() error {
		if dc.LocalAddress == "" {
			return &DataConnectionError{Err: "Local address not set for Active Connection"}
		}
		listenSocket, err := net.Listen(dc.Protocol, dc.LocalAddress)
		if err != nil {
			return err
		}
		tcpAddr, err := net.ResolveTCPAddr(dc.Protocol, listenSocket.Addr().String())
		if err != nil {
			return err
		}
		dc.PassivePort = tcpAddr.Port
		dc.PassiveListener = listenSocket
		dc.IsSetup = true
		return nil
	}

	switch dc.TransferType {
	case TransferType(Active):
		return ActiveSetup()
	case TransferType(Passive):
		return PassiveSetup()
	}
	return &DataConnectionError{"Invalid TransferType"}
}

func (dc *DataConnection) FinishSetup() error {
	ActiveFinishSetup := func() error {
		local, _ := net.ResolveTCPAddr(dc.Protocol, dc.LocalAddress)
		remote, _ := net.ResolveTCPAddr(dc.Protocol, dc.RemoteAddress)
		connection, err := net.DialTCP(dc.Protocol, local, remote)
		if err != nil {
			return err
		}
		if dc.Security.ProtectionLevel == DataChannelProtectionLevel(Private) {
			dc.Connection = tls.Server(connection, dc.Security.Config)
		} else {
			dc.Connection = connection
		}
		return nil
	}

	PassiveFinishSetup := func() error {
		connection, err := dc.PassiveListener.Accept()
		if err != nil {
			return err
		}
		err = dc.PassiveListener.Close()
		if err != nil {
			return err
		}
		if dc.Security.ProtectionLevel == DataChannelProtectionLevel(Private) {
			dc.Connection = tls.Server(connection, dc.Security.Config)
		} else {
			dc.Connection = connection
		}
		return nil
	}

	if !dc.IsSetup {
		return &DataConnectionError{Err: "DataConnection has not been setup"}
	}

	switch dc.TransferType {
	case TransferType(Active):
		return ActiveFinishSetup()
	case TransferType(Passive):
		return PassiveFinishSetup()
	}

	return &DataConnectionError{Err: "Invalid TransferType"}
}

func (dc DataConnection) Send(data string) (int, error) {
	return dc.Connection.Write([]byte(data))
}

func (dc DataConnection) Receive() (string, error) {
	var buffer []byte
	_, err := dc.Connection.Read(buffer)
	if err != nil {
		return "", err
	}
	return string(buffer[:]), nil
}

func (dc DataConnection) CopyIn(file *os.File) (int64, error) {
	return io.Copy(file, dc.Connection)
}

func (dc DataConnection) CopyOut(file *os.File) (int64, error) {
	return io.Copy(dc.Connection, file)
}

func (dc DataConnection) Close() {
	if dc.Connection != nil {
		dc.Connection.Close()
	}
	if dc.PassiveListener != nil {
		_ = dc.PassiveListener.Close()
	}
}
