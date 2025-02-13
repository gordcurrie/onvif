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

// Call_RemoveScopes forwards the call to dev.CallMethod() then parses the payload of the reply as a RemoveScopesResponse.
func Call_RemoveScopes(ctx context.Context, dev *onvif.Device, request device.RemoveScopes) (device.RemoveScopesResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			RemoveScopesResponse device.RemoveScopesResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.RemoveScopesResponse, err
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "RemoveScopes")
		return reply.Body.RemoveScopesResponse, err
	}
}
