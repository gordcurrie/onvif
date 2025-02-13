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

// Call_DeleteOSD forwards the call to dev.CallMethod() then parses the payload of the reply as a DeleteOSDResponse.
func Call_DeleteOSD(ctx context.Context, dev *onvif.Device, request media.DeleteOSD) (media.DeleteOSDResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			DeleteOSDResponse media.DeleteOSDResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.DeleteOSDResponse, err
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "DeleteOSD")
		return reply.Body.DeleteOSDResponse, err
	}
}
