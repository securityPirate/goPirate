package instances

import (
	"net"
)

//Instance hold machine related data
type Instance struct {
	IP []net.IP
}

func reboot() {}
func delete() {}
