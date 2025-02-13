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

// Call_GetClientCertificateMode forwards the call to dev.CallMethod() then parses the payload of the reply as a GetClientCertificateModeResponse.
func Call_GetClientCertificateMode(ctx context.Context, dev *onvif.Device, request device.GetClientCertificateMode) (device.GetClientCertificateModeResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetClientCertificateModeResponse device.GetClientCertificateModeResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetClientCertificateModeResponse, err
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetClientCertificateMode")
		return reply.Body.GetClientCertificateModeResponse, err
	}
}
