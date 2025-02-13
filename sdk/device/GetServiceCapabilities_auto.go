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

// Call_GetServiceCapabilities forwards the call to dev.CallMethod() then parses the payload of the reply as a GetServiceCapabilitiesResponse.
func Call_GetServiceCapabilities(ctx context.Context, dev *onvif.Device, request device.GetServiceCapabilities) (device.GetServiceCapabilitiesResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetServiceCapabilitiesResponse device.GetServiceCapabilitiesResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetServiceCapabilitiesResponse, err
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetServiceCapabilities")
		return reply.Body.GetServiceCapabilitiesResponse, err
	}
}
