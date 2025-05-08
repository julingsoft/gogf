package guidx

import (
    "github.com/julingsoft/gogf/contrib/netx"
)

func MachineID() (uint16, error) {
	ip, err := netx.PrivateIPv4()
	if err != nil {
		return 0, err
	}

	return uint16(ip[2])<<8 + uint16(ip[3]), nil
}
