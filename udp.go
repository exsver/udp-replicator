package main

import "net"

func NewSourceListener(addr string) (*net.UDPConn, error) {
	udpAddress, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		return nil, err
	}

	return net.ListenUDP("udp", udpAddress)
}

func MakeDestinationConnections(conf *Config) ([]*net.UDPConn, error) {
	dstConnections := make([]*net.UDPConn, 0, len(conf.Destinations))
	for _, destination := range conf.Destinations {
		DstAddr, err := net.ResolveUDPAddr("udp", destination)
		if err != nil {
			return nil, err
		}

		conn, err := net.DialUDP("udp", nil, DstAddr)
		if err != nil {
			return nil, err
		}

		dstConnections = append(dstConnections, conn)
	}

	return dstConnections, nil
}
