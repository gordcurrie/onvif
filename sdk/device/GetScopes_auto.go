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

// Call_GetScopes forwards the call to dev.CallMethod() then parses the payload of the reply as a GetScopesResponse.
func Call_GetScopes(ctx context.Context, dev *onvif.Device, request device.GetScopes) (device.GetScopesResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetScopesResponse device.GetScopesResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetScopesResponse, err
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetScopes")
		return reply.Body.GetScopesResponse, err
	}
}
