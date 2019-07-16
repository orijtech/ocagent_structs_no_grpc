// Copyright 2019, OpenCensus Authors
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

package ocagent_test

import (
	"bytes"
	"log"
	"net/http"
	"net/http/httputil"
	"time"

	// We use JSONPb instead of "encoding/json" for
	// JSON serialization of Proto messages.
	"github.com/golang/protobuf/jsonpb"

	"go.opencensus.io/stats"
	"go.opencensus.io/stats/view"
	"go.opencensus.io/tag"
	"go.opencensus.io/trace"
	"go.opencensus.io/trace/tracestate"

	ocagent "github.com/orijtech/ocagent_structs_no_grpc"
)

var (
	startTime = time.Now()
	endTime   = startTime.Add(17 * time.Second)
)

// The example shows how to JSON export Traces from OpenCensus-Go trace.SpanData to OpenCensus-Proto trace.Span requests.
func ExampleTrace_jsonExport() {
	ocTracestate, err := tracestate.New(new(tracestate.Tracestate), tracestate.Entry{Key: "foo", Value: "bar"},
		tracestate.Entry{Key: "a", Value: "b"})
	if err != nil || ocTracestate == nil {
		log.Fatalf("Failed to create ocTracestate: %v", err)
	}
	// This trace.SpanData will typically be obtained after binding to a trace.Exporter.
	ocSpanData := &trace.SpanData{
		SpanContext: trace.SpanContext{
			TraceID:    trace.TraceID{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0A, 0x0B, 0x0C, 0x0D, 0x0E, 0x0F},
			SpanID:     trace.SpanID{0xFF, 0xFE, 0xFD, 0xFC, 0xFB, 0xFA, 0xF9, 0xF8},
			Tracestate: ocTracestate,
		},
		SpanKind:     trace.SpanKindServer,
		ParentSpanID: trace.SpanID{0xEF, 0xEE, 0xED, 0xEC, 0xEB, 0xEA, 0xE9, 0xE8},
		Name:         "End-To-End Here",
		StartTime:    startTime,
		EndTime:      endTime,
		Annotations: []trace.Annotation{
			{
				Time:    startTime,
				Message: "start",
				Attributes: map[string]interface{}{
					"timeout_ns": int64(12e9),
					"agent":      "ocagent",
					"cache_hit":  true,
				},
			},
		},
		MessageEvents: []trace.MessageEvent{
			{Time: startTime, EventType: trace.MessageEventTypeSent, UncompressedByteSize: 1024, CompressedByteSize: 512},
			{Time: endTime, EventType: trace.MessageEventTypeRecv, UncompressedByteSize: 1024, CompressedByteSize: 1000},
		},
		Links: []trace.Link{
			{
				TraceID: trace.TraceID{0xE0, 0xE1, 0xE2, 0xE3, 0xE4, 0xE5, 0xE6, 0xE7, 0xE8, 0xE9, 0xEA, 0xEB, 0xEC, 0xED, 0xEE, 0xEF},
				SpanID:  trace.SpanID{0xD0, 0xD1, 0xD2, 0xD3, 0xD4, 0xD5, 0xD6, 0xD7},
				Type:    trace.LinkTypeChild,
			},
		},
		Status: trace.Status{
			Code:    trace.StatusCodeInternal,
			Message: "This is not a drill!",
		},
		HasRemoteParent: true,
		Attributes: map[string]interface{}{
			"timeout_ns": int64(12e9),
			"agent":      "ocagent",
			"cache_hit":  true,
			"ping_count": int(25),
		},
	}

	protoTraceReq := ocagent.OpenCensusSpanDataToProtoSpans([]*trace.SpanData{ocSpanData})
	// Ensure that you ALWAYS pass in the node as the first message.
	protoTraceReq.Node = ocagent.NodeWithStartTime("example", time.Now())

	ts := &jsonpb.Marshaler{}
	buf := new(bytes.Buffer)
	if err := ts.Marshal(buf, protoTraceReq); err != nil {
		log.Fatalf("Failed to JSONPb marshal: %v", err)
	}

	ocagentAddr := "http://localhost:55678" // The address of the running OpenCensus Agent.
	req, err := http.NewRequest("POST", ocagentAddr+"/v1/trace", buf)
	if err != nil {
		log.Fatalf("Failed to create HTTP request: %v", err)
	}

	tr := &http.Transport{}
	client := &http.Client{
		Transport: tr,
	}
	req.Header.Set("Content-Type", "application/json")
	res, err := client.Do(req)
	if err != nil {
		log.Fatalf("Failed to get a successful response: %v", err)
	}
	out, _ := httputil.DumpResponse(res, true)
	log.Printf("Res: %s\n", out)
}

// The example shows how to JSON export view.Data from OpenCensus-Go to OpenCensus-Proto metrics.Metric requests.
func ExampleMetrics_jsonExport() {
	// This view.Data will typically be obtained after binding to a view.Exporter.
	keyField, _ := tag.NewKey("field")
	keyName, _ := tag.NewKey("name")

	mSprinterLatencyMs := stats.Float64("sprint_latency", "The time in which a sprinter completes the course", "ms")

	vd := &view.Data{
		Start: startTime,
		End:   endTime,
		View: &view.View{
			Name:        "ocagent.io/latency",
			Description: "latency of runners for a 100m dash",
			Aggregation: view.Distribution(0, 10, 20, 30, 40),
			TagKeys:     []tag.Key{keyField, keyName},
			Measure:     mSprinterLatencyMs,
		},
		Rows: []*view.Row{
			{
				Tags: []tag.Tag{
					{Key: keyField, Value: "main-field"},
					{Key: keyName, Value: "sprinter-#10"},
				},
				Data: &view.DistributionData{
					// Points: [11.9]
					Count:           1,
					Min:             11.9,
					Max:             11.9,
					Mean:            11.9,
					CountPerBucket:  []int64{0, 1, 0, 0, 0},
					SumOfSquaredDev: 0,
				},
			},
			{
				Tags: []tag.Tag{
					{Key: keyField, Value: "small-field"},
					{Key: keyName, Value: ""},
				},
				Data: &view.DistributionData{
					// Points: [20.2]
					Count:           1,
					Min:             20.2,
					Max:             20.2,
					Mean:            20.2,
					CountPerBucket:  []int64{0, 0, 1, 0, 0},
					SumOfSquaredDev: 0,
				},
			},
			{
				Tags: []tag.Tag{
					{Key: keyField, Value: "small-field"},
					{Key: keyName, Value: "sprinter-#yp"},
				},
				Data: &view.DistributionData{
					// Points: [28.9]
					Count:           1,
					Min:             28.9,
					Max:             28.9,
					Mean:            28.9,
					CountPerBucket:  []int64{0, 0, 1, 0, 0},
					SumOfSquaredDev: 0,
				},
			},
		},
	}

	protoMetricsReq := ocagent.OpenCensusViewDataToProtoMetrics([]*view.Data{vd})
	// Ensure that you ALWAYS pass in the node as the first message.
	protoMetricsReq.Node = ocagent.NodeWithStartTime("example", time.Now())

	ts := &jsonpb.Marshaler{}
	buf := new(bytes.Buffer)
	if err := ts.Marshal(buf, protoMetricsReq); err != nil {
		log.Fatalf("Failed to JSONPb marshal: %v", err)
	}

	ocagentAddr := "http://localhost:55678" // The address of the running OpenCensus Agent.
	req, err := http.NewRequest("POST", ocagentAddr+"/v1/metrics", buf)
	if err != nil {
		log.Fatalf("Failed to create HTTP request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalf("Failed to get a successful response: %v", err)
	}
	out, _ := httputil.DumpResponse(res, true)
	log.Printf("Res: %s\n", out)
}
