// Code generated : DO NOT EDIT.
// Copyright (c) 2022 Jean-Francois SMIGIELSKI
// Distributed under the MIT License

package ptz

import (
	"context"
	"git.taservs.net/federal-devices/go-onvif"
	"git.taservs.net/federal-devices/go-onvif/sdk"
	"git.taservs.net/federal-devices/go-onvif/ptz"
)

// Call_GotoPreset forwards the call to dev.CallMethod() then parses the payload of the reply as a GotoPresetResponse.
func Call_GotoPreset(ctx context.Context, dev *onvif.Device, request ptz.GotoPreset) (ptz.GotoPresetResponse, error) {
	type Envelope struct {
		Header struct{}
		Body   struct {
			GotoPresetResponse ptz.GotoPresetResponse
		}
	}
	var reply Envelope
	if httpReply, err := dev.CallMethod(request); err != nil {
		return reply.Body.GotoPresetResponse, err
	} else {
		err = sdk.ReadAndParse(ctx, httpReply, &reply, "GotoPreset")
		return reply.Body.GotoPresetResponse, err
	}
}
