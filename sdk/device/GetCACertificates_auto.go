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

// Call_GetCACertificates forwards the call to dev.CallMethod() then parses the payload of the reply as a GetCACertificatesResponse.
func Call_GetCACertificates(ctx context.Context, dev *onvif.Device, request device.GetCACertificates) (device.GetCACertificatesResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetCACertificatesResponse device.GetCACertificatesResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetCACertificatesResponse, err
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetCACertificates")
		return reply.Body.GetCACertificatesResponse, err
	}
}
