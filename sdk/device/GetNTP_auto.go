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

// Call_GetNTP forwards the call to dev.CallMethod() then parses the payload of the reply as a GetNTPResponse.
func Call_GetNTP(ctx context.Context, dev *onvif.Device, request device.GetNTP) (device.GetNTPResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetNTPResponse device.GetNTPResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetNTPResponse, err
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetNTP")
		return reply.Body.GetNTPResponse, err
	}
}
