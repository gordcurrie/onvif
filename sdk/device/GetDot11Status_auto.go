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

// Call_GetDot11Status forwards the call to dev.CallMethod() then parses the payload of the reply as a GetDot11StatusResponse.
func Call_GetDot11Status(ctx context.Context, dev *onvif.Device, request device.GetDot11Status) (device.GetDot11StatusResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetDot11StatusResponse device.GetDot11StatusResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetDot11StatusResponse, err
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetDot11Status")
		return reply.Body.GetDot11StatusResponse, err
	}
}
