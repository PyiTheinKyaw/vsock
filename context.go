package vsock

import (
  "os"
  "golang.org/x/sys/unix"
)

// contextID retrieves the local context ID for this system.
var devVsock = "/dev/vsock"

func ContextID() (uint32, error) {
  f, err := os.Open(devVsock)
  if err != nil {
    return 0, err
  }
  defer f.Close()

  return unix.IoctlGetUint32(int(f.Fd()), unix.IOCTL_VM_SOCKETS_GET_LOCAL_CID)
}