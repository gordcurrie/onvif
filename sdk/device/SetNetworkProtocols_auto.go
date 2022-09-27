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

// Call_SetNetworkProtocols forwards the call to dev.CallMethod() then parses the payload of the reply as a SetNetworkProtocolsResponse.
func Call_SetNetworkProtocols(ctx context.Context, dev *onvif.Device, request device.SetNetworkProtocols) (device.SetNetworkProtocolsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetNetworkProtocolsResponse device.SetNetworkProtocolsResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.SetNetworkProtocolsResponse, err
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "SetNetworkProtocols")
		return reply.Body.SetNetworkProtocolsResponse, err
	}
}
