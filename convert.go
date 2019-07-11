// Copyright 2018, OpenCensus Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package ocagent

import (
	"context"

	resourcepb "github.com/orijtech/ocagent_structs_no_grpc/pb/resource/v1"
	"go.opencensus.io/resource"
)

func resourceProtoFromEnv() *resourcepb.Resource {
	rs, _ := resource.FromEnv(context.Background())
	if rs == nil {
		return nil
	}
	return resourceToResourcePb(rs)
}

func resourceToResourcePb(rs *resource.Resource) *resourcepb.Resource {
	rprs := &resourcepb.Resource{
		Type: rs.Type,
	}
	if rs.Labels != nil {
		rprs.Labels = make(map[string]string)
		for k, v := range rs.Labels {
			rprs.Labels[k] = v
		}
	}
	return rprs
}
