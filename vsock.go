package vsock

import (
    "golang.org/x/sys/unix"
)

const (
    AF_VSOCK       = 40         // Address family for VSOCK sockets
    SOCK_STREAM    = 1          // Provides sequenced, reliable, two-way, connection-based byte streams
    VMADDR_CID_ANY = 0xFFFFFFFF // Bind to any CID
    VMADDR_CID_SERVER = 0x4
)

// SockaddrVM defines the sockaddr_vm struct
type SockaddrVM struct {
    CID  uint32
    Port uint32
}

func (sa *SockaddrVM) sockaddr() *unix.SockaddrVM {
    return &unix.SockaddrVM{
        CID:  sa.CID,
        Port: sa.Port,
    }
}

func Socket() (int, error) {
    return unix.Socket(AF_VSOCK, SOCK_STREAM, 0)
}

func Connect(fd int, sa *SockaddrVM) error {
    return unix.Connect(fd, sa.sockaddr())
}

func Bind(fd int, sa *SockaddrVM) error {
    return unix.Bind(fd, sa.sockaddr())
}

func Listen(fd int, backlog int) error {
    return unix.Listen(fd, backlog)
}

func Accept(fd int) (int, *SockaddrVM, error) {
    nfd, sa, err := unix.Accept(fd)
    if err != nil {
        return -1, nil, err
    }
    sockaddr := sa.(*unix.SockaddrVM)
    return nfd, &SockaddrVM{CID: sockaddr.CID, Port: sockaddr.Port}, nil
}

func Close(fd int) error {
    return unix.Close(fd)
}

func Recv(fd int) ([]byte, error) {
    buf := make([]byte, 1024)
    n, err := unix.Read(fd, buf)
    if err != nil {
        return nil, err
    }
    return buf[:n], nil
}

func Send(fd int, data []byte) error {
    _, err := unix.Write(fd, data)
    return err
}
