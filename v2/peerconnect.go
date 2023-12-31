package v2

import (
	"context"

	"github.com/Qweder93/concordium_go_sdk/v2/pb"
)

// PeerConnect suggest to a peer to connect to the submitted peer details. This, if successful,
// adds the peer to the list of given addresses. Otherwise return a GRPC error.
// Note. The peer might not be connected to instantly, in that case the node will try to establish the connection in near future.
// This function returns a GRPC status 'Ok' in this case.
func (c *Client) PeerConnect(ctx context.Context, req *pb.IpSocketAddress) (err error) {
	_, err = c.grpcClient.PeerConnect(ctx, req)
	if err != nil {
		return err
	}

	return nil
}
