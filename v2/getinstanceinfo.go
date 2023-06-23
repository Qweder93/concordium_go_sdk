package v2

import (
	"context"

	"concordium_go_sdk/v2/pb"
)

// GetInstanceInfo get info about a smart contract instance as it appears at the end of the given block.
func (c *Client) GetInstanceInfo(ctx context.Context, req *pb.InstanceInfoRequest) (_ *pb.InstanceInfo, err error) {
	instanceInfo, err := c.grpcClient.GetInstanceInfo(ctx, req)
	if err != nil {
		return &pb.InstanceInfo{}, err
	}

	return instanceInfo, nil
}
