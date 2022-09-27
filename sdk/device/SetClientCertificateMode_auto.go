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

// Call_SetClientCertificateMode forwards the call to dev.CallMethod() then parses the payload of the reply as a SetClientCertificateModeResponse.
func Call_SetClientCertificateMode(ctx context.Context, dev *onvif.Device, request device.SetClientCertificateMode) (device.SetClientCertificateModeResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			SetClientCertificateModeResponse device.SetClientCertificateModeResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.SetClientCertificateModeResponse, err
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "SetClientCertificateMode")
		return reply.Body.SetClientCertificateModeResponse, err
	}
}
