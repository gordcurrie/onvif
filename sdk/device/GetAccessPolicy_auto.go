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

// Call_GetAccessPolicy forwards the call to dev.CallMethod() then parses the payload of the reply as a GetAccessPolicyResponse.
func Call_GetAccessPolicy(ctx context.Context, dev *onvif.Device, request device.GetAccessPolicy) (device.GetAccessPolicyResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetAccessPolicyResponse device.GetAccessPolicyResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetAccessPolicyResponse, err
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetAccessPolicy")
		return reply.Body.GetAccessPolicyResponse, err
	}
}
