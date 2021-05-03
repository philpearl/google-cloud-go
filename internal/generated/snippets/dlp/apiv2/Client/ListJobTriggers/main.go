// Copyright 2021 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// [START dlp_v2_generated_DlpService_ListJobTriggers_sync]

package main

import (
	"context"

	dlp "cloud.google.com/go/dlp/apiv2"
	"google.golang.org/api/iterator"
	dlppb "google.golang.org/genproto/googleapis/privacy/dlp/v2"
)

func main() {
	// import dlppb "google.golang.org/genproto/googleapis/privacy/dlp/v2"
	// import "google.golang.org/api/iterator"

	ctx := context.Background()
	c, err := dlp.NewClient(ctx)
	if err != nil {
		// TODO: Handle error.
	}

	req := &dlppb.ListJobTriggersRequest{
		// TODO: Fill request struct fields.
	}
	it := c.ListJobTriggers(ctx, req)
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

// [END dlp_v2_generated_DlpService_ListJobTriggers_sync]