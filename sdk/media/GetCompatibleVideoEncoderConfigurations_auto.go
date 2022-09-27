// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package media

import (
	"context"
	"git.taservs.net/federal-devices/go-onvif"
	"git.taservs.net/federal-devices/go-onvif/sdk"
	"git.taservs.net/federal-devices/go-onvif/media"
)

// Call_GetCompatibleVideoEncoderConfigurations forwards the call to dev.CallMethod() then parses the payload of the reply as a GetCompatibleVideoEncoderConfigurationsResponse.
func Call_GetCompatibleVideoEncoderConfigurations(ctx context.Context, dev *onvif.Device, request media.GetCompatibleVideoEncoderConfigurations) (media.GetCompatibleVideoEncoderConfigurationsResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetCompatibleVideoEncoderConfigurationsResponse media.GetCompatibleVideoEncoderConfigurationsResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetCompatibleVideoEncoderConfigurationsResponse, err
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetCompatibleVideoEncoderConfigurations")
		return reply.Body.GetCompatibleVideoEncoderConfigurationsResponse, err
	}
}
