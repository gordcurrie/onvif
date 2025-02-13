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

// Call_SetNetworkDefaultGateway forwards the call to dev.CallMethod() then parses the payload of the reply as a SetNetworkDefaultGatewayResponse.
func Call_SetNetworkDefaultGateway(ctx context.Context, dev *onvif.Device, request device.SetNetworkDefaultGateway) (device.SetNetworkDefaultGatewayResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetNetworkDefaultGatewayResponse device.SetNetworkDefaultGatewayResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.SetNetworkDefaultGatewayResponse, err
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "SetNetworkDefaultGateway")
		return reply.Body.SetNetworkDefaultGatewayResponse, err
	}
}
