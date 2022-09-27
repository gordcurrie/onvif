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

// Call_CreateProfile forwards the call to dev.CallMethod() then parses the payload of the reply as a CreateProfileResponse.
func Call_CreateProfile(ctx context.Context, dev *onvif.Device, request media.CreateProfile) (media.CreateProfileResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			CreateProfileResponse media.CreateProfileResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.CreateProfileResponse, err
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "CreateProfile")
		return reply.Body.CreateProfileResponse, err
	}
}
