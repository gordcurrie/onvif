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

// Call_GetDot1XConfiguration forwards the call to dev.CallMethod() then parses the payload of the reply as a GetDot1XConfigurationResponse.
func Call_GetDot1XConfiguration(ctx context.Context, dev *onvif.Device, request device.GetDot1XConfiguration) (device.GetDot1XConfigurationResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetDot1XConfigurationResponse device.GetDot1XConfigurationResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetDot1XConfigurationResponse, err
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetDot1XConfiguration")
		return reply.Body.GetDot1XConfigurationResponse, err
	}
}
