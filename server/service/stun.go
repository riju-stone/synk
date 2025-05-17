package service

import (
	"context"
	"errors"

	"gortc.io/stun"
)

var PublicServerList = []string{
	"stun.l.google.com:19302",
	"stun.l.google.com:5349",
	"stun1.l.google.com:3478",
	"stun1.l.google.com:5349",
	"stun2.l.google.com:19302",
	// "stun:stun2.l.google.com:5349",
	// "stun:stun3.l.google.com:3478",
	// "stun:stun3.l.google.com:5349",
	"stun4.l.google.com:19302",
	// "stun:stun4.l.google.com:5349",
}

type PeerResult struct {
	server string
	res    string
}

type PeerErrorResult struct {
	server string
	err    error
}

type StunResponse struct {
	Results map[string]string
	Errors  map[string]error
	Quorum  uint
}

// QueryServer queries the STUN server to get the public IP address and port.
// It returns the public IP address and port as a string in the format "IP:Port".
// If ipv6 is true, it queries the IPv6 address.
// If ipv6 is false, it queries the IPv4 address.
func QueryServer(ctx context.Context, server string, ipv6 bool) (string, error) {
	queryType := "udp4"
	if ipv6 {
		queryType = "udp6"
	}

	conn, err := stun.Dial(queryType, server)
	if err != nil {
		return "", err
	}
	defer conn.Close()

	mssg, err := stun.Build(stun.TransactionID, stun.BindingRequest)
	if err != nil {
		return "", err
	}

	// Make two separate channels to handle the response and error
	clientResp := make(chan stun.Event)
	clientErr := make(chan error)

	go func() {
		err := conn.Do(mssg, func(res stun.Event) {
			clientResp <- res
		})

		if err != nil {
			clientErr <- err
		}
	}()

	select {
	case res := <-clientResp:
		if res.Error != nil {
			return "", res.Error
		}

		var xorAddr stun.XORMappedAddress
		if err := xorAddr.GetFrom(res.Message); err == nil {
			return xorAddr.IP.String(), nil
		} else {
			var mappedAddr stun.MappedAddress
			if err := mappedAddr.GetFrom(res.Message); err == nil {
				return mappedAddr.IP.String(), nil
			} else {
				return "", err
			}
		}
	case err := <-clientErr:
		return "", err
	case <-ctx.Done():
		return "", ctx.Err()
	}
}

// QueryServerList queries the STUN server list to get the public IP address and port.
// It returns the public IP address and port as a string in the format "IP:Port".
func QueryServerList(ctx context.Context, ipv6 bool, quorum uint) (string, error) {
	serverCount := len(PublicServerList)
	if serverCount == 0 {
		return "", errors.New("no STUN Servers Found")
	}

	if serverCount < int(quorum) {
		return "", errors.New("not enough STUN Servers to reach quorum")
	}

	peerResp := make(chan PeerResult)
	peerErr := make(chan PeerErrorResult)

	for _, server := range PublicServerList {
		go func(serv string) {
			res, err := QueryServer(ctx, serv, ipv6)
			if err != nil {
				peerErr <- PeerErrorResult{server: serv, err: err}
			} else {
				peerResp <- PeerResult{server: serv, res: res}
			}
		}(server)
	}

	resultCount := make(map[string]uint)
	resultMap := make(map[string]string)
	errorMap := make(map[string]error)

	for range serverCount {
		select {
		case resp := <-peerResp:
			resultCount[resp.res]++
			if resultCount[resp.res] >= quorum {
				return resp.res, nil
			}
			resultMap[resp.server] = resp.res
		case err := <-peerErr:
			errorMap[err.server] = err.err
		}
	}

	if len(resultCount) == 0 {
		return "", errors.New("no STUN Servers responded")
	}

	if len(resultCount) < int(quorum) {
		return "", errors.New("not enough STUN Servers to reach quorum")
	}

	return "", nil
}
