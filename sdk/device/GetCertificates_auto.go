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

// Call_GetCertificates forwards the call to dev.CallMethod() then parses the payload of the reply as a GetCertificatesResponse.
func Call_GetCertificates(ctx context.Context, dev *onvif.Device, request device.GetCertificates) (device.GetCertificatesResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetCertificatesResponse device.GetCertificatesResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetCertificatesResponse, err
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetCertificates")
		return reply.Body.GetCertificatesResponse, err
	}
}
