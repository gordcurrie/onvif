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

// Call_GetDynamicDNS forwards the call to dev.CallMethod() then parses the payload of the reply as a GetDynamicDNSResponse.
func Call_GetDynamicDNS(ctx context.Context, dev *onvif.Device, request device.GetDynamicDNS) (device.GetDynamicDNSResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetDynamicDNSResponse device.GetDynamicDNSResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetDynamicDNSResponse, err
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetDynamicDNS")
		return reply.Body.GetDynamicDNSResponse, err
	}
}
