# ocagent_structs_no_grpc

A standalone collection of all the structs needed
to export Traces and Metrics to the OpenCensus Agent
as OpenCensus-Proto.

The reason for its existence to gut out the gRPC
dependencies in order for golang.org/x/tools/*
to use OpenCensus and have performance monitoring.

## Note

Make sure to use

    `"github.com/golang/protobuf/jsonpb"`

instead of

    `"encoding/json"`

for JSON serialization of Proto messages.
