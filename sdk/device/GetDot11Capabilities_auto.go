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

// Call_GetDot11Capabilities forwards the call to dev.CallMethod() then parses the payload of the reply as a GetDot11CapabilitiesResponse.
func Call_GetDot11Capabilities(ctx context.Context, dev *onvif.Device, request device.GetDot11Capabilities) (device.GetDot11CapabilitiesResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetDot11CapabilitiesResponse device.GetDot11CapabilitiesResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetDot11CapabilitiesResponse, err
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetDot11Capabilities")
		return reply.Body.GetDot11CapabilitiesResponse, err
	}
}
