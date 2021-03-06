// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by protoc-gen-go_gapic. DO NOT EDIT.

package speech_test

import (
	"context"

	speech "cloud.google.com/go/speech/apiv1p1beta1"
	"google.golang.org/api/iterator"
	speechpb "google.golang.org/genproto/googleapis/cloud/speech/v1p1beta1"
)

func ExampleNewAdaptationClient() {
	ctx := context.Background()
	c, err := speech.NewAdaptationClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use client.
	_ = c
}

func ExampleAdaptationClient_CreatePhraseSet() {
	// import speechpb "google.golang.org/genproto/googleapis/cloud/speech/v1p1beta1"

	ctx := context.Background()
	c, err := speech.NewAdaptationClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &speechpb.CreatePhraseSetRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.CreatePhraseSet(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleAdaptationClient_GetPhraseSet() {
	// import speechpb "google.golang.org/genproto/googleapis/cloud/speech/v1p1beta1"

	ctx := context.Background()
	c, err := speech.NewAdaptationClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &speechpb.GetPhraseSetRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetPhraseSet(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleAdaptationClient_ListPhraseSet() {
	// import speechpb "google.golang.org/genproto/googleapis/cloud/speech/v1p1beta1"
	// import "google.golang.org/api/iterator"

	ctx := context.Background()
	c, err := speech.NewAdaptationClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &speechpb.ListPhraseSetRequest{
		// TODO: Fill request struct fields.
	}
	it := c.ListPhraseSet(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleAdaptationClient_UpdatePhraseSet() {
	// import speechpb "google.golang.org/genproto/googleapis/cloud/speech/v1p1beta1"

	ctx := context.Background()
	c, err := speech.NewAdaptationClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &speechpb.UpdatePhraseSetRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.UpdatePhraseSet(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleAdaptationClient_DeletePhraseSet() {
	ctx := context.Background()
	c, err := speech.NewAdaptationClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &speechpb.DeletePhraseSetRequest{
		// TODO: Fill request struct fields.
	}
	err = c.DeletePhraseSet(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}

func ExampleAdaptationClient_CreateCustomClass() {
	// import speechpb "google.golang.org/genproto/googleapis/cloud/speech/v1p1beta1"

	ctx := context.Background()
	c, err := speech.NewAdaptationClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &speechpb.CreateCustomClassRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.CreateCustomClass(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleAdaptationClient_GetCustomClass() {
	// import speechpb "google.golang.org/genproto/googleapis/cloud/speech/v1p1beta1"

	ctx := context.Background()
	c, err := speech.NewAdaptationClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &speechpb.GetCustomClassRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.GetCustomClass(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleAdaptationClient_ListCustomClasses() {
	// import speechpb "google.golang.org/genproto/googleapis/cloud/speech/v1p1beta1"
	// import "google.golang.org/api/iterator"

	ctx := context.Background()
	c, err := speech.NewAdaptationClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &speechpb.ListCustomClassesRequest{
		// TODO: Fill request struct fields.
	}
	it := c.ListCustomClasses(ctx, req)
	for {
		resp, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			// TODO: Handle error.
		}
		// TODO: Use resp.
		_ = resp
	}
}

func ExampleAdaptationClient_UpdateCustomClass() {
	// import speechpb "google.golang.org/genproto/googleapis/cloud/speech/v1p1beta1"

	ctx := context.Background()
	c, err := speech.NewAdaptationClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &speechpb.UpdateCustomClassRequest{
		// TODO: Fill request struct fields.
	}
	resp, err := c.UpdateCustomClass(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
	// TODO: Use resp.
	_ = resp
}

func ExampleAdaptationClient_DeleteCustomClass() {
	ctx := context.Background()
	c, err := speech.NewAdaptationClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &speechpb.DeleteCustomClassRequest{
		// TODO: Fill request struct fields.
	}
	err = c.DeleteCustomClass(ctx, req)
	if err != nil {
		// TODO: Handle error.
	}
}
