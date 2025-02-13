// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package device

import (
	"context"
	"git.taservs.net/federal-devices/go-onvif"
	"git.taservs.net/federal-devices/go-onvif/sdk"
	"git.taservs.net/federal-devices/go-onvif/device"
)

// Call_GetNetworkDefaultGateway forwards the call to dev.CallMethod() then parses the payload of the reply as a GetNetworkDefaultGatewayResponse.
func Call_GetNetworkDefaultGateway(ctx context.Context, dev *onvif.Device, request device.GetNetworkDefaultGateway) (device.GetNetworkDefaultGatewayResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetNetworkDefaultGatewayResponse device.GetNetworkDefaultGatewayResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetNetworkDefaultGatewayResponse, err
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetNetworkDefaultGateway")
		return reply.Body.GetNetworkDefaultGatewayResponse, err
	}
}
