package client

import "pb"

func NewBackendClient(name string, port int) (pb.TwitterCloneClient, func() error, error) {
	c, err := gRPCConn(name, port)
	if err != nil {
		return nil, nil, err
	}

	return pb.NewTwitterCloneClient(c), c.Close, nil
}
