package target

import (
	"github.com/transhift/transhift/transhift/puncher"
	"github.com/transhift/transhift/common/protocol"
	"crypto/tls"
	"fmt"
	"net"
)

func punchHole(host string, port int, cert tls.Certificate) (laddr net.Addr, sourceAddr string, err error) {
	p := puncher.New(host, port, protocol.TargetNode, cert)
	if laddr, err = p.Connect(); err != nil {
		return
	}

	// Expect ID.
	var id string
	if err = p.Dec().Decode(&id); err != nil {
		return
	}

	fmt.Printf("your ID is '%s'\n", id)

	// Expect target address.
	err = p.Dec().Decode(&sourceAddr)
	return
}
