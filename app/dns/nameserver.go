package dns

import (
	"context"

	"v2ray.com/core/common/net"
)

type NameServer interface {
	QueryIP(ctx context.Context, domain string) ([]net.IP, error)
}

type LocalNameServer struct {
	resolver net.Resolver
}

func (s *LocalNameServer) QueryIP(ctx context.Context, domain string) ([]net.IP, error) {
	ipAddr, err := s.resolver.LookupIPAddr(ctx, domain)
	if err != nil {
		return nil, err
	}
	var ips []net.IP
	for _, addr := range ipAddr {
		ips = append(ips, addr.IP)
	}
	return ips, nil
}

func NewLocalNameServer() *LocalNameServer {
	return &LocalNameServer{
		resolver: net.Resolver{
			PreferGo: true,
		},
	}
}
