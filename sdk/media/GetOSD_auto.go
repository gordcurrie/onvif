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

// Call_GetOSD forwards the call to dev.CallMethod() then parses the payload of the reply as a GetOSDResponse.
func Call_GetOSD(ctx context.Context, dev *onvif.Device, request media.GetOSD) (media.GetOSDResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GetOSDResponse media.GetOSDResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GetOSDResponse, err
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GetOSD")
		return reply.Body.GetOSDResponse, err
	}
}
