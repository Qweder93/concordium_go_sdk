package v2

import (
	"context"

	"github.com/Qweder93/concordium_go_sdk/v2/pb"
)

// PeerDisconnect disconnect from the peer and remove them from the given addresses list if they are on it.
// Return if the request was processed successfully. Otherwise return a GRPC error.
func (c *Client) PeerDisconnect(ctx context.Context, req *pb.IpSocketAddress) (err error) {
	_, err = c.grpcClient.PeerDisconnect(ctx, req)
	if err != nil {
		return err
	}

	return nil
}
