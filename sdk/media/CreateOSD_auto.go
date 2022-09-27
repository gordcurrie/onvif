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

// Call_CreateOSD forwards the call to dev.CallMethod() then parses the payload of the reply as a CreateOSDResponse.
func Call_CreateOSD(ctx context.Context, dev *onvif.Device, request media.CreateOSD) (media.CreateOSDResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			CreateOSDResponse media.CreateOSDResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.CreateOSDResponse, err
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "CreateOSD")
		return reply.Body.CreateOSDResponse, err
	}
}
