// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package device

import (
	"context"
	"github.com/gordcurrie/onvif"
	"github.com/gordcurrie/onvif/sdk"
	"github.com/gordcurrie/onvif/device"
)

// Call_GetRemoteDiscoveryMode forwards the call to dev.CallMethod() then parses the payload of the reply as a GetRemoteDiscoveryModeResponse.
func Call_GetRemoteDiscoveryMode(ctx context.Context, dev *onvif.Device, request device.GetRemoteDiscoveryMode) (device.GetRemoteDiscoveryModeResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetRemoteDiscoveryModeResponse device.GetRemoteDiscoveryModeResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetRemoteDiscoveryModeResponse, err
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetRemoteDiscoveryMode")
		return reply.Body.GetRemoteDiscoveryModeResponse, err
	}
}
