package mail

import (
	"errors"
	"fmt"
	"github.com/emersion/go-imap/client"
)

func Dial(address string) (*client.Client, error) {
	c, errTls := client.DialTLS(address, nil)
	if errTls != nil {
		// tls failed, try to dial without TLS
		var errInsecure error
		c, errInsecure = client.Dial(address)
		if errInsecure != nil {
			return nil, errors.Join(fmt.Errorf("failed to dial %s", address), errTls, errInsecure)
		}
	}
	return c, nil
}
