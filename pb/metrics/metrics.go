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

package metrics

import (
	fmt "fmt"
	v1 "github.com/orijtech/ocagent_structs_no_grpc/pb/resource/v1"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// The kind of metric. It describes how the data is reported.
//
// A gauge is an instantaneous measurement of a value.
//
// A cumulative measurement is a value accumulated over a time interval. In
// a time series, cumulative measurements should have the same start time,
// increasing values and increasing end times, until an event resets the
// cumulative value to zero and sets a new start time for the following
// points.
type MetricDescriptor_Type int32

const (
	// Do not use this default value.
	MetricDescriptor_UNSPECIFIED MetricDescriptor_Type = 0
	// Integer gauge. The value can go both up and down.
	MetricDescriptor_GAUGE_INT64 MetricDescriptor_Type = 1
	// Floating point gauge. The value can go both up and down.
	MetricDescriptor_GAUGE_DOUBLE MetricDescriptor_Type = 2
	// Distribution gauge measurement. The count and sum can go both up and
	// down. Recorded values are always >= 0.
	// Used in scenarios like a snapshot of time the current items in a queue
	// have spent there.
	MetricDescriptor_GAUGE_DISTRIBUTION MetricDescriptor_Type = 3
	// Integer cumulative measurement. The value cannot decrease, if resets
	// then the start_time should also be reset.
	MetricDescriptor_CUMULATIVE_INT64 MetricDescriptor_Type = 4
	// Floating point cumulative measurement. The value cannot decrease, if
	// resets then the start_time should also be reset. Recorded values are
	// always >= 0.
	MetricDescriptor_CUMULATIVE_DOUBLE MetricDescriptor_Type = 5
	// Distribution cumulative measurement. The count and sum cannot decrease,
	// if resets then the start_time should also be reset.
	MetricDescriptor_CUMULATIVE_DISTRIBUTION MetricDescriptor_Type = 6
	// Some frameworks implemented Histograms as a summary of observations
	// (usually things like request durations and response sizes). While it
	// also provides a total count of observations and a sum of all observed
	// values, it calculates configurable percentiles over a sliding time
	// window. This is not recommended, since it cannot be aggregated.
	MetricDescriptor_SUMMARY MetricDescriptor_Type = 7
)

var MetricDescriptor_Type_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "GAUGE_INT64",
	2: "GAUGE_DOUBLE",
	3: "GAUGE_DISTRIBUTION",
	4: "CUMULATIVE_INT64",
	5: "CUMULATIVE_DOUBLE",
	6: "CUMULATIVE_DISTRIBUTION",
	7: "SUMMARY",
}

var MetricDescriptor_Type_value = map[string]int32{
	"UNSPECIFIED":             0,
	"GAUGE_INT64":             1,
	"GAUGE_DOUBLE":            2,
	"GAUGE_DISTRIBUTION":      3,
	"CUMULATIVE_INT64":        4,
	"CUMULATIVE_DOUBLE":       5,
	"CUMULATIVE_DISTRIBUTION": 6,
	"SUMMARY":                 7,
}

func (x MetricDescriptor_Type) String() string {
	return proto.EnumName(MetricDescriptor_Type_name, int32(x))
}

func (MetricDescriptor_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0ee3deb72053811a, []int{1, 0}
}

// Defines a Metric which has one or more timeseries.
type Metric struct {
	// The descriptor of the Metric.
	// TODO(issue #152): consider only sending the name of descriptor for
	// optimization.
	MetricDescriptor *MetricDescriptor `protobuf:"bytes,1,opt,name=metric_descriptor,json=metricDescriptor,proto3" json:"metric_descriptor,omitempty"`
	// One or more timeseries for a single metric, where each timeseries has
	// one or more points.
	Timeseries []*TimeSeries `protobuf:"bytes,2,rep,name=timeseries,proto3" json:"timeseries,omitempty"`
	// The resource for the metric. If unset, it may be set to a default value
	// provided for a sequence of messages in an RPC stream.
	Resource             *v1.Resource `protobuf:"bytes,3,opt,name=resource,proto3" json:"resource,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Metric) Reset()         { *m = Metric{} }
func (m *Metric) String() string { return proto.CompactTextString(m) }
func (*Metric) ProtoMessage()    {}
func (*Metric) Descriptor() ([]byte, []int) {
	return fileDescriptor_0ee3deb72053811a, []int{0}
}

func (m *Metric) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Metric.Unmarshal(m, b)
}
func (m *Metric) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Metric.Marshal(b, m, deterministic)
}
func (m *Metric) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Metric.Merge(m, src)
}
func (m *Metric) XXX_Size() int {
	return xxx_messageInfo_Metric.Size(m)
}
func (m *Metric) XXX_DiscardUnknown() {
	xxx_messageInfo_Metric.DiscardUnknown(m)
}

var xxx_messageInfo_Metric proto.InternalMessageInfo

func (m *Metric) GetMetricDescriptor() *MetricDescriptor {
	if m != nil {
		return m.MetricDescriptor
	}
	return nil
}

func (m *Metric) GetTimeseries() []*TimeSeries {
	if m != nil {
		return m.Timeseries
	}
	return nil
}

func (m *Metric) GetResource() *v1.Resource {
	if m != nil {
		return m.Resource
	}
	return nil
}

// Defines a metric type and its schema.
type MetricDescriptor struct {
	// The metric type, including its DNS name prefix. It must be unique.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// A detailed description of the metric, which can be used in documentation.
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	// The unit in which the metric value is reported. Follows the format
	// described by http://unitsofmeasure.org/ucum.html.
	Unit string                `protobuf:"bytes,3,opt,name=unit,proto3" json:"unit,omitempty"`
	Type MetricDescriptor_Type `protobuf:"varint,4,opt,name=type,proto3,enum=opencensus.proto.metrics.v1.MetricDescriptor_Type" json:"type,omitempty"`
	// The label keys associated with the metric descriptor.
	LabelKeys            []*LabelKey `protobuf:"bytes,5,rep,name=label_keys,json=labelKeys,proto3" json:"label_keys,omitempty"`
	XXX_NoUnkeyedLiteral struct{}    `json:"-"`
	XXX_unrecognized     []byte      `json:"-"`
	XXX_sizecache        int32       `json:"-"`
}

func (m *MetricDescriptor) Reset()         { *m = MetricDescriptor{} }
func (m *MetricDescriptor) String() string { return proto.CompactTextString(m) }
func (*MetricDescriptor) ProtoMessage()    {}
func (*MetricDescriptor) Descriptor() ([]byte, []int) {
	return fileDescriptor_0ee3deb72053811a, []int{1}
}

func (m *MetricDescriptor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MetricDescriptor.Unmarshal(m, b)
}
func (m *MetricDescriptor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MetricDescriptor.Marshal(b, m, deterministic)
}
func (m *MetricDescriptor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MetricDescriptor.Merge(m, src)
}
func (m *MetricDescriptor) XXX_Size() int {
	return xxx_messageInfo_MetricDescriptor.Size(m)
}
func (m *MetricDescriptor) XXX_DiscardUnknown() {
	xxx_messageInfo_MetricDescriptor.DiscardUnknown(m)
}

var xxx_messageInfo_MetricDescriptor proto.InternalMessageInfo

func (m *MetricDescriptor) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *MetricDescriptor) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *MetricDescriptor) GetUnit() string {
	if m != nil {
		return m.Unit
	}
	return ""
}

func (m *MetricDescriptor) GetType() MetricDescriptor_Type {
	if m != nil {
		return m.Type
	}
	return MetricDescriptor_UNSPECIFIED
}

func (m *MetricDescriptor) GetLabelKeys() []*LabelKey {
	if m != nil {
		return m.LabelKeys
	}
	return nil
}

// Defines a label key associated with a metric descriptor.
type LabelKey struct {
	// The key for the label.
	Key string `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	// A human-readable description of what this label key represents.
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LabelKey) Reset()         { *m = LabelKey{} }
func (m *LabelKey) String() string { return proto.CompactTextString(m) }
func (*LabelKey) ProtoMessage()    {}
func (*LabelKey) Descriptor() ([]byte, []int) {
	return fileDescriptor_0ee3deb72053811a, []int{2}
}

func (m *LabelKey) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LabelKey.Unmarshal(m, b)
}
func (m *LabelKey) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LabelKey.Marshal(b, m, deterministic)
}
func (m *LabelKey) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LabelKey.Merge(m, src)
}
func (m *LabelKey) XXX_Size() int {
	return xxx_messageInfo_LabelKey.Size(m)
}
func (m *LabelKey) XXX_DiscardUnknown() {
	xxx_messageInfo_LabelKey.DiscardUnknown(m)
}

var xxx_messageInfo_LabelKey proto.InternalMessageInfo

func (m *LabelKey) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *LabelKey) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

// A collection of data points that describes the time-varying values
// of a metric.
type TimeSeries struct {
	// Must be present for cumulative metrics. The time when the cumulative value
	// was reset to zero. Exclusive. The cumulative value is over the time interval
	// (start_timestamp, timestamp]. If not specified, the backend can use the
	// previous recorded value.
	StartTimestamp *timestamp.Timestamp `protobuf:"bytes,1,opt,name=start_timestamp,json=startTimestamp,proto3" json:"start_timestamp,omitempty"`
	// The set of label values that uniquely identify this timeseries. Applies to
	// all points. The order of label values must match that of label keys in the
	// metric descriptor.
	LabelValues []*LabelValue `protobuf:"bytes,2,rep,name=label_values,json=labelValues,proto3" json:"label_values,omitempty"`
	// The data points of this timeseries. Point.value type MUST match the
	// MetricDescriptor.type.
	Points               []*Point `protobuf:"bytes,3,rep,name=points,proto3" json:"points,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TimeSeries) Reset()         { *m = TimeSeries{} }
func (m *TimeSeries) String() string { return proto.CompactTextString(m) }
func (*TimeSeries) ProtoMessage()    {}
func (*TimeSeries) Descriptor() ([]byte, []int) {
	return fileDescriptor_0ee3deb72053811a, []int{3}
}

func (m *TimeSeries) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TimeSeries.Unmarshal(m, b)
}
func (m *TimeSeries) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TimeSeries.Marshal(b, m, deterministic)
}
func (m *TimeSeries) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TimeSeries.Merge(m, src)
}
func (m *TimeSeries) XXX_Size() int {
	return xxx_messageInfo_TimeSeries.Size(m)
}
func (m *TimeSeries) XXX_DiscardUnknown() {
	xxx_messageInfo_TimeSeries.DiscardUnknown(m)
}

var xxx_messageInfo_TimeSeries proto.InternalMessageInfo

func (m *TimeSeries) GetStartTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.StartTimestamp
	}
	return nil
}

func (m *TimeSeries) GetLabelValues() []*LabelValue {
	if m != nil {
		return m.LabelValues
	}
	return nil
}

func (m *TimeSeries) GetPoints() []*Point {
	if m != nil {
		return m.Points
	}
	return nil
}

type LabelValue struct {
	// The value for the label.
	Value string `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
	// If false the value field is ignored and considered not set.
	// This is used to differentiate a missing label from an empty string.
	HasValue             bool     `protobuf:"varint,2,opt,name=has_value,json=hasValue,proto3" json:"has_value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LabelValue) Reset()         { *m = LabelValue{} }
func (m *LabelValue) String() string { return proto.CompactTextString(m) }
func (*LabelValue) ProtoMessage()    {}
func (*LabelValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_0ee3deb72053811a, []int{4}
}

func (m *LabelValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LabelValue.Unmarshal(m, b)
}
func (m *LabelValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LabelValue.Marshal(b, m, deterministic)
}
func (m *LabelValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LabelValue.Merge(m, src)
}
func (m *LabelValue) XXX_Size() int {
	return xxx_messageInfo_LabelValue.Size(m)
}
func (m *LabelValue) XXX_DiscardUnknown() {
	xxx_messageInfo_LabelValue.DiscardUnknown(m)
}

var xxx_messageInfo_LabelValue proto.InternalMessageInfo

func (m *LabelValue) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *LabelValue) GetHasValue() bool {
	if m != nil {
		return m.HasValue
	}
	return false
}

// A timestamped measurement.
type Point struct {
	// The moment when this point was recorded. Inclusive.
	// If not specified, the timestamp will be decided by the backend.
	Timestamp *timestamp.Timestamp `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// The actual point value.
	//
	// Types that are valid to be assigned to Value:
	//	*Point_Int64Value
	//	*Point_DoubleValue
	//	*Point_DistributionValue
	//	*Point_SummaryValue
	Value                isPoint_Value `protobuf_oneof:"value"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Point) Reset()         { *m = Point{} }
func (m *Point) String() string { return proto.CompactTextString(m) }
func (*Point) ProtoMessage()    {}
func (*Point) Descriptor() ([]byte, []int) {
	return fileDescriptor_0ee3deb72053811a, []int{5}
}

func (m *Point) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Point.Unmarshal(m, b)
}
func (m *Point) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Point.Marshal(b, m, deterministic)
}
func (m *Point) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Point.Merge(m, src)
}
func (m *Point) XXX_Size() int {
	return xxx_messageInfo_Point.Size(m)
}
func (m *Point) XXX_DiscardUnknown() {
	xxx_messageInfo_Point.DiscardUnknown(m)
}

var xxx_messageInfo_Point proto.InternalMessageInfo

func (m *Point) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

type isPoint_Value interface {
	isPoint_Value()
}

type Point_Int64Value struct {
	Int64Value int64 `protobuf:"varint,2,opt,name=int64_value,json=int64Value,proto3,oneof"`
}

type Point_DoubleValue struct {
	DoubleValue float64 `protobuf:"fixed64,3,opt,name=double_value,json=doubleValue,proto3,oneof"`
}

type Point_DistributionValue struct {
	DistributionValue *DistributionValue `protobuf:"bytes,4,opt,name=distribution_value,json=distributionValue,proto3,oneof"`
}

type Point_SummaryValue struct {
	SummaryValue *SummaryValue `protobuf:"bytes,5,opt,name=summary_value,json=summaryValue,proto3,oneof"`
}

func (*Point_Int64Value) isPoint_Value() {}

func (*Point_DoubleValue) isPoint_Value() {}

func (*Point_DistributionValue) isPoint_Value() {}

func (*Point_SummaryValue) isPoint_Value() {}

func (m *Point) GetValue() isPoint_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *Point) GetInt64Value() int64 {
	if x, ok := m.GetValue().(*Point_Int64Value); ok {
		return x.Int64Value
	}
	return 0
}

func (m *Point) GetDoubleValue() float64 {
	if x, ok := m.GetValue().(*Point_DoubleValue); ok {
		return x.DoubleValue
	}
	return 0
}

func (m *Point) GetDistributionValue() *DistributionValue {
	if x, ok := m.GetValue().(*Point_DistributionValue); ok {
		return x.DistributionValue
	}
	return nil
}

func (m *Point) GetSummaryValue() *SummaryValue {
	if x, ok := m.GetValue().(*Point_SummaryValue); ok {
		return x.SummaryValue
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Point) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Point_Int64Value)(nil),
		(*Point_DoubleValue)(nil),
		(*Point_DistributionValue)(nil),
		(*Point_SummaryValue)(nil),
	}
}

// Distribution contains summary statistics for a population of values. It
// optionally contains a histogram representing the distribution of those
// values across a set of buckets.
type DistributionValue struct {
	// The number of values in the population. Must be non-negative. This value
	// must equal the sum of the values in bucket_counts if a histogram is
	// provided.
	Count int64 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	// The sum of the values in the population. If count is zero then this field
	// must be zero.
	Sum float64 `protobuf:"fixed64,2,opt,name=sum,proto3" json:"sum,omitempty"`
	// The sum of squared deviations from the mean of the values in the
	// population. For values x_i this is:
	//
	//     Sum[i=1..n]((x_i - mean)^2)
	//
	// Knuth, "The Art of Computer Programming", Vol. 2, page 323, 3rd edition
	// describes Welford's method for accumulating this sum in one pass.
	//
	// If count is zero then this field must be zero.
	SumOfSquaredDeviation float64 `protobuf:"fixed64,3,opt,name=sum_of_squared_deviation,json=sumOfSquaredDeviation,proto3" json:"sum_of_squared_deviation,omitempty"`
	// Don't change bucket boundaries within a TimeSeries if your backend doesn't
	// support this.
	// TODO(issue #152): consider not required to send bucket options for
	// optimization.
	BucketOptions *DistributionValue_BucketOptions `protobuf:"bytes,4,opt,name=bucket_options,json=bucketOptions,proto3" json:"bucket_options,omitempty"`
	// If the distribution does not have a histogram, then omit this field.
	// If there is a histogram, then the sum of the values in the Bucket counts
	// must equal the value in the count field of the distribution.
	Buckets              []*DistributionValue_Bucket `protobuf:"bytes,5,rep,name=buckets,proto3" json:"buckets,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *DistributionValue) Reset()         { *m = DistributionValue{} }
func (m *DistributionValue) String() string { return proto.CompactTextString(m) }
func (*DistributionValue) ProtoMessage()    {}
func (*DistributionValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_0ee3deb72053811a, []int{6}
}

func (m *DistributionValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DistributionValue.Unmarshal(m, b)
}
func (m *DistributionValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DistributionValue.Marshal(b, m, deterministic)
}
func (m *DistributionValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DistributionValue.Merge(m, src)
}
func (m *DistributionValue) XXX_Size() int {
	return xxx_messageInfo_DistributionValue.Size(m)
}
func (m *DistributionValue) XXX_DiscardUnknown() {
	xxx_messageInfo_DistributionValue.DiscardUnknown(m)
}

var xxx_messageInfo_DistributionValue proto.InternalMessageInfo

func (m *DistributionValue) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *DistributionValue) GetSum() float64 {
	if m != nil {
		return m.Sum
	}
	return 0
}

func (m *DistributionValue) GetSumOfSquaredDeviation() float64 {
	if m != nil {
		return m.SumOfSquaredDeviation
	}
	return 0
}

func (m *DistributionValue) GetBucketOptions() *DistributionValue_BucketOptions {
	if m != nil {
		return m.BucketOptions
	}
	return nil
}

func (m *DistributionValue) GetBuckets() []*DistributionValue_Bucket {
	if m != nil {
		return m.Buckets
	}
	return nil
}

// A Distribution may optionally contain a histogram of the values in the
// population. The bucket boundaries for that histogram are described by
// BucketOptions.
//
// If bucket_options has no type, then there is no histogram associated with
// the Distribution.
type DistributionValue_BucketOptions struct {
	// Types that are valid to be assigned to Type:
	//	*DistributionValue_BucketOptions_Explicit_
	Type                 isDistributionValue_BucketOptions_Type `protobuf_oneof:"type"`
	XXX_NoUnkeyedLiteral struct{}                               `json:"-"`
	XXX_unrecognized     []byte                                 `json:"-"`
	XXX_sizecache        int32                                  `json:"-"`
}

func (m *DistributionValue_BucketOptions) Reset()         { *m = DistributionValue_BucketOptions{} }
func (m *DistributionValue_BucketOptions) String() string { return proto.CompactTextString(m) }
func (*DistributionValue_BucketOptions) ProtoMessage()    {}
func (*DistributionValue_BucketOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_0ee3deb72053811a, []int{6, 0}
}

func (m *DistributionValue_BucketOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DistributionValue_BucketOptions.Unmarshal(m, b)
}
func (m *DistributionValue_BucketOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DistributionValue_BucketOptions.Marshal(b, m, deterministic)
}
func (m *DistributionValue_BucketOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DistributionValue_BucketOptions.Merge(m, src)
}
func (m *DistributionValue_BucketOptions) XXX_Size() int {
	return xxx_messageInfo_DistributionValue_BucketOptions.Size(m)
}
func (m *DistributionValue_BucketOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_DistributionValue_BucketOptions.DiscardUnknown(m)
}

var xxx_messageInfo_DistributionValue_BucketOptions proto.InternalMessageInfo

type isDistributionValue_BucketOptions_Type interface {
	isDistributionValue_BucketOptions_Type()
}

type DistributionValue_BucketOptions_Explicit_ struct {
	Explicit *DistributionValue_BucketOptions_Explicit `protobuf:"bytes,1,opt,name=explicit,proto3,oneof"`
}

func (*DistributionValue_BucketOptions_Explicit_) isDistributionValue_BucketOptions_Type() {}

func (m *DistributionValue_BucketOptions) GetType() isDistributionValue_BucketOptions_Type {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *DistributionValue_BucketOptions) GetExplicit() *DistributionValue_BucketOptions_Explicit {
	if x, ok := m.GetType().(*DistributionValue_BucketOptions_Explicit_); ok {
		return x.Explicit
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*DistributionValue_BucketOptions) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*DistributionValue_BucketOptions_Explicit_)(nil),
	}
}

// Specifies a set of buckets with arbitrary upper-bounds.
// This defines size(bounds) + 1 (= N) buckets. The boundaries for bucket
// index i are:
//
// [0, bucket_bounds[i]) for i == 0
// [bucket_bounds[i-1], bucket_bounds[i]) for 0 < i < N-1
// [bucket_bounds[i], +infinity) for i == N-1
type DistributionValue_BucketOptions_Explicit struct {
	// The values must be strictly increasing and > 0.
	Bounds               []float64 `protobuf:"fixed64,1,rep,packed,name=bounds,proto3" json:"bounds,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *DistributionValue_BucketOptions_Explicit) Reset() {
	*m = DistributionValue_BucketOptions_Explicit{}
}
func (m *DistributionValue_BucketOptions_Explicit) String() string { return proto.CompactTextString(m) }
func (*DistributionValue_BucketOptions_Explicit) ProtoMessage()    {}
func (*DistributionValue_BucketOptions_Explicit) Descriptor() ([]byte, []int) {
	return fileDescriptor_0ee3deb72053811a, []int{6, 0, 0}
}

func (m *DistributionValue_BucketOptions_Explicit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DistributionValue_BucketOptions_Explicit.Unmarshal(m, b)
}
func (m *DistributionValue_BucketOptions_Explicit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DistributionValue_BucketOptions_Explicit.Marshal(b, m, deterministic)
}
func (m *DistributionValue_BucketOptions_Explicit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DistributionValue_BucketOptions_Explicit.Merge(m, src)
}
func (m *DistributionValue_BucketOptions_Explicit) XXX_Size() int {
	return xxx_messageInfo_DistributionValue_BucketOptions_Explicit.Size(m)
}
func (m *DistributionValue_BucketOptions_Explicit) XXX_DiscardUnknown() {
	xxx_messageInfo_DistributionValue_BucketOptions_Explicit.DiscardUnknown(m)
}

var xxx_messageInfo_DistributionValue_BucketOptions_Explicit proto.InternalMessageInfo

func (m *DistributionValue_BucketOptions_Explicit) GetBounds() []float64 {
	if m != nil {
		return m.Bounds
	}
	return nil
}

type DistributionValue_Bucket struct {
	// The number of values in each bucket of the histogram, as described in
	// bucket_bounds.
	Count int64 `protobuf:"varint,1,opt,name=count,proto3" json:"count,omitempty"`
	// If the distribution does not have a histogram, then omit this field.
	Exemplar             *DistributionValue_Exemplar `protobuf:"bytes,2,opt,name=exemplar,proto3" json:"exemplar,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                    `json:"-"`
	XXX_unrecognized     []byte                      `json:"-"`
	XXX_sizecache        int32                       `json:"-"`
}

func (m *DistributionValue_Bucket) Reset()         { *m = DistributionValue_Bucket{} }
func (m *DistributionValue_Bucket) String() string { return proto.CompactTextString(m) }
func (*DistributionValue_Bucket) ProtoMessage()    {}
func (*DistributionValue_Bucket) Descriptor() ([]byte, []int) {
	return fileDescriptor_0ee3deb72053811a, []int{6, 1}
}

func (m *DistributionValue_Bucket) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DistributionValue_Bucket.Unmarshal(m, b)
}
func (m *DistributionValue_Bucket) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DistributionValue_Bucket.Marshal(b, m, deterministic)
}
func (m *DistributionValue_Bucket) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DistributionValue_Bucket.Merge(m, src)
}
func (m *DistributionValue_Bucket) XXX_Size() int {
	return xxx_messageInfo_DistributionValue_Bucket.Size(m)
}
func (m *DistributionValue_Bucket) XXX_DiscardUnknown() {
	xxx_messageInfo_DistributionValue_Bucket.DiscardUnknown(m)
}

var xxx_messageInfo_DistributionValue_Bucket proto.InternalMessageInfo

func (m *DistributionValue_Bucket) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *DistributionValue_Bucket) GetExemplar() *DistributionValue_Exemplar {
	if m != nil {
		return m.Exemplar
	}
	return nil
}

// Exemplars are example points that may be used to annotate aggregated
// Distribution values. They are metadata that gives information about a
// particular value added to a Distribution bucket.
type DistributionValue_Exemplar struct {
	// Value of the exemplar point. It determines which bucket the exemplar
	// belongs to.
	Value float64 `protobuf:"fixed64,1,opt,name=value,proto3" json:"value,omitempty"`
	// The observation (sampling) time of the above value.
	Timestamp *timestamp.Timestamp `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// Contextual information about the example value.
	Attachments          map[string]string `protobuf:"bytes,3,rep,name=attachments,proto3" json:"attachments,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *DistributionValue_Exemplar) Reset()         { *m = DistributionValue_Exemplar{} }
func (m *DistributionValue_Exemplar) String() string { return proto.CompactTextString(m) }
func (*DistributionValue_Exemplar) ProtoMessage()    {}
func (*DistributionValue_Exemplar) Descriptor() ([]byte, []int) {
	return fileDescriptor_0ee3deb72053811a, []int{6, 2}
}

func (m *DistributionValue_Exemplar) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DistributionValue_Exemplar.Unmarshal(m, b)
}
func (m *DistributionValue_Exemplar) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DistributionValue_Exemplar.Marshal(b, m, deterministic)
}
func (m *DistributionValue_Exemplar) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DistributionValue_Exemplar.Merge(m, src)
}
func (m *DistributionValue_Exemplar) XXX_Size() int {
	return xxx_messageInfo_DistributionValue_Exemplar.Size(m)
}
func (m *DistributionValue_Exemplar) XXX_DiscardUnknown() {
	xxx_messageInfo_DistributionValue_Exemplar.DiscardUnknown(m)
}

var xxx_messageInfo_DistributionValue_Exemplar proto.InternalMessageInfo

func (m *DistributionValue_Exemplar) GetValue() float64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *DistributionValue_Exemplar) GetTimestamp() *timestamp.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *DistributionValue_Exemplar) GetAttachments() map[string]string {
	if m != nil {
		return m.Attachments
	}
	return nil
}

// The start_timestamp only applies to the count and sum in the SummaryValue.
type SummaryValue struct {
	// The total number of recorded values since start_time. Optional since
	// some systems don't expose this.
	Count *wrappers.Int64Value `protobuf:"bytes,1,opt,name=count,proto3" json:"count,omitempty"`
	// The total sum of recorded values since start_time. Optional since some
	// systems don't expose this. If count is zero then this field must be zero.
	// This field must be unset if the sum is not available.
	Sum *wrappers.DoubleValue `protobuf:"bytes,2,opt,name=sum,proto3" json:"sum,omitempty"`
	// Values calculated over an arbitrary time window.
	Snapshot             *SummaryValue_Snapshot `protobuf:"bytes,3,opt,name=snapshot,proto3" json:"snapshot,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *SummaryValue) Reset()         { *m = SummaryValue{} }
func (m *SummaryValue) String() string { return proto.CompactTextString(m) }
func (*SummaryValue) ProtoMessage()    {}
func (*SummaryValue) Descriptor() ([]byte, []int) {
	return fileDescriptor_0ee3deb72053811a, []int{7}
}

func (m *SummaryValue) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SummaryValue.Unmarshal(m, b)
}
func (m *SummaryValue) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SummaryValue.Marshal(b, m, deterministic)
}
func (m *SummaryValue) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SummaryValue.Merge(m, src)
}
func (m *SummaryValue) XXX_Size() int {
	return xxx_messageInfo_SummaryValue.Size(m)
}
func (m *SummaryValue) XXX_DiscardUnknown() {
	xxx_messageInfo_SummaryValue.DiscardUnknown(m)
}

var xxx_messageInfo_SummaryValue proto.InternalMessageInfo

func (m *SummaryValue) GetCount() *wrappers.Int64Value {
	if m != nil {
		return m.Count
	}
	return nil
}

func (m *SummaryValue) GetSum() *wrappers.DoubleValue {
	if m != nil {
		return m.Sum
	}
	return nil
}

func (m *SummaryValue) GetSnapshot() *SummaryValue_Snapshot {
	if m != nil {
		return m.Snapshot
	}
	return nil
}

// The values in this message can be reset at arbitrary unknown times, with
// the requirement that all of them are reset at the same time.
type SummaryValue_Snapshot struct {
	// The number of values in the snapshot. Optional since some systems don't
	// expose this.
	Count *wrappers.Int64Value `protobuf:"bytes,1,opt,name=count,proto3" json:"count,omitempty"`
	// The sum of values in the snapshot. Optional since some systems don't
	// expose this. If count is zero then this field must be zero or not set
	// (if not supported).
	Sum *wrappers.DoubleValue `protobuf:"bytes,2,opt,name=sum,proto3" json:"sum,omitempty"`
	// A list of values at different percentiles of the distribution calculated
	// from the current snapshot. The percentiles must be strictly increasing.
	PercentileValues     []*SummaryValue_Snapshot_ValueAtPercentile `protobuf:"bytes,3,rep,name=percentile_values,json=percentileValues,proto3" json:"percentile_values,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                                   `json:"-"`
	XXX_unrecognized     []byte                                     `json:"-"`
	XXX_sizecache        int32                                      `json:"-"`
}

func (m *SummaryValue_Snapshot) Reset()         { *m = SummaryValue_Snapshot{} }
func (m *SummaryValue_Snapshot) String() string { return proto.CompactTextString(m) }
func (*SummaryValue_Snapshot) ProtoMessage()    {}
func (*SummaryValue_Snapshot) Descriptor() ([]byte, []int) {
	return fileDescriptor_0ee3deb72053811a, []int{7, 0}
}

func (m *SummaryValue_Snapshot) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SummaryValue_Snapshot.Unmarshal(m, b)
}
func (m *SummaryValue_Snapshot) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SummaryValue_Snapshot.Marshal(b, m, deterministic)
}
func (m *SummaryValue_Snapshot) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SummaryValue_Snapshot.Merge(m, src)
}
func (m *SummaryValue_Snapshot) XXX_Size() int {
	return xxx_messageInfo_SummaryValue_Snapshot.Size(m)
}
func (m *SummaryValue_Snapshot) XXX_DiscardUnknown() {
	xxx_messageInfo_SummaryValue_Snapshot.DiscardUnknown(m)
}

var xxx_messageInfo_SummaryValue_Snapshot proto.InternalMessageInfo

func (m *SummaryValue_Snapshot) GetCount() *wrappers.Int64Value {
	if m != nil {
		return m.Count
	}
	return nil
}

func (m *SummaryValue_Snapshot) GetSum() *wrappers.DoubleValue {
	if m != nil {
		return m.Sum
	}
	return nil
}

func (m *SummaryValue_Snapshot) GetPercentileValues() []*SummaryValue_Snapshot_ValueAtPercentile {
	if m != nil {
		return m.PercentileValues
	}
	return nil
}

// Represents the value at a given percentile of a distribution.
type SummaryValue_Snapshot_ValueAtPercentile struct {
	// The percentile of a distribution. Must be in the interval
	// (0.0, 100.0].
	Percentile float64 `protobuf:"fixed64,1,opt,name=percentile,proto3" json:"percentile,omitempty"`
	// The value at the given percentile of a distribution.
	Value                float64  `protobuf:"fixed64,2,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SummaryValue_Snapshot_ValueAtPercentile) Reset() {
	*m = SummaryValue_Snapshot_ValueAtPercentile{}
}
func (m *SummaryValue_Snapshot_ValueAtPercentile) String() string { return proto.CompactTextString(m) }
func (*SummaryValue_Snapshot_ValueAtPercentile) ProtoMessage()    {}
func (*SummaryValue_Snapshot_ValueAtPercentile) Descriptor() ([]byte, []int) {
	return fileDescriptor_0ee3deb72053811a, []int{7, 0, 0}
}

func (m *SummaryValue_Snapshot_ValueAtPercentile) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SummaryValue_Snapshot_ValueAtPercentile.Unmarshal(m, b)
}
func (m *SummaryValue_Snapshot_ValueAtPercentile) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SummaryValue_Snapshot_ValueAtPercentile.Marshal(b, m, deterministic)
}
func (m *SummaryValue_Snapshot_ValueAtPercentile) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SummaryValue_Snapshot_ValueAtPercentile.Merge(m, src)
}
func (m *SummaryValue_Snapshot_ValueAtPercentile) XXX_Size() int {
	return xxx_messageInfo_SummaryValue_Snapshot_ValueAtPercentile.Size(m)
}
func (m *SummaryValue_Snapshot_ValueAtPercentile) XXX_DiscardUnknown() {
	xxx_messageInfo_SummaryValue_Snapshot_ValueAtPercentile.DiscardUnknown(m)
}

var xxx_messageInfo_SummaryValue_Snapshot_ValueAtPercentile proto.InternalMessageInfo

func (m *SummaryValue_Snapshot_ValueAtPercentile) GetPercentile() float64 {
	if m != nil {
		return m.Percentile
	}
	return 0
}

func (m *SummaryValue_Snapshot_ValueAtPercentile) GetValue() float64 {
	if m != nil {
		return m.Value
	}
	return 0
}

var fileDescriptor_0ee3deb72053811a = []byte{
	// 1118 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x56, 0xdd, 0x6e, 0x1b, 0xc5,
	0x17, 0xcf, 0xda, 0x8e, 0xe3, 0x9c, 0x75, 0xd2, 0xf5, 0xa8, 0xed, 0xdf, 0x72, 0xfe, 0x0a, 0x61,
	0x11, 0x90, 0x0a, 0x65, 0xad, 0x98, 0xd2, 0x56, 0x15, 0x2a, 0x8a, 0x63, 0x37, 0x31, 0x24, 0xb1,
	0x35, 0xb6, 0x23, 0xd1, 0x1b, 0x6b, 0xbd, 0x9e, 0x24, 0x4b, 0xbc, 0x1f, 0xdd, 0x99, 0x0d, 0xf8,
	0x05, 0x78, 0x02, 0xc4, 0x35, 0xb7, 0x88, 0xe7, 0xe0, 0x8a, 0x27, 0xe0, 0x15, 0xb8, 0x41, 0xbc,
	0x01, 0xda, 0x99, 0xd9, 0x8f, 0xc4, 0x60, 0xea, 0x22, 0x71, 0x77, 0xe6, 0xcc, 0x39, 0xbf, 0xfd,
	0x9d, 0xcf, 0x1d, 0x78, 0xe4, 0xf9, 0xc4, 0xb5, 0x88, 0x4b, 0x43, 0x5a, 0xf7, 0x03, 0x8f, 0x79,
	0x75, 0x87, 0xb0, 0xc0, 0xb6, 0x68, 0xfd, 0x66, 0x3f, 0x16, 0x0d, 0x7e, 0x81, 0xb6, 0x52, 0x53,
	0xa1, 0x31, 0xe2, 0xfb, 0x9b, 0xfd, 0xda, 0x3b, 0x97, 0x9e, 0x77, 0x39, 0x25, 0x02, 0x63, 0x1c,
	0x5e, 0xd4, 0x99, 0xed, 0x10, 0xca, 0x4c, 0xc7, 0x17, 0xb6, 0xb5, 0xed, 0xbb, 0x06, 0x5f, 0x07,
	0xa6, 0xef, 0x93, 0x40, 0x62, 0xd5, 0x3e, 0x9a, 0x23, 0x12, 0x10, 0xea, 0x85, 0x81, 0x45, 0x22,
	0x26, 0xb1, 0x2c, 0x8c, 0xf5, 0x3f, 0x14, 0x28, 0x9e, 0xf2, 0x8f, 0xa3, 0x57, 0x50, 0x11, 0x34,
	0x46, 0x13, 0x42, 0xad, 0xc0, 0xf6, 0x99, 0x17, 0x54, 0x95, 0x1d, 0x65, 0x57, 0x6d, 0xec, 0x19,
	0x0b, 0x18, 0x1b, 0xc2, 0xbf, 0x95, 0x38, 0x61, 0xcd, 0xb9, 0xa3, 0x41, 0x47, 0x00, 0x3c, 0x0c,
	0x12, 0xd8, 0x84, 0x56, 0x73, 0x3b, 0xf9, 0x5d, 0xb5, 0xf1, 0xe1, 0x42, 0xd0, 0x81, 0xed, 0x90,
	0x3e, 0x37, 0xc7, 0x19, 0x57, 0xd4, 0x84, 0x52, 0x1c, 0x41, 0x35, 0xcf, 0xb9, 0x7d, 0x30, 0x0f,
	0x93, 0xc4, 0x78, 0xb3, 0x6f, 0x60, 0x29, 0xe3, 0xc4, 0x4f, 0xff, 0x3e, 0x0f, 0xda, 0x5d, 0xce,
	0x08, 0x41, 0xc1, 0x35, 0x1d, 0xc2, 0x03, 0x5e, 0xc7, 0x5c, 0x46, 0x3b, 0xa0, 0xc6, 0xa9, 0xb0,
	0x3d, 0xb7, 0x9a, 0xe3, 0x57, 0x59, 0x55, 0xe4, 0x15, 0xba, 0x36, 0xe3, 0x54, 0xd6, 0x31, 0x97,
	0xd1, 0x4b, 0x28, 0xb0, 0x99, 0x4f, 0xaa, 0x85, 0x1d, 0x65, 0x77, 0xb3, 0xd1, 0x58, 0x2a, 0x75,
	0xc6, 0x60, 0xe6, 0x13, 0xcc, 0xfd, 0x51, 0x0b, 0x60, 0x6a, 0x8e, 0xc9, 0x74, 0x74, 0x4d, 0x66,
	0xb4, 0xba, 0xca, 0x73, 0xf6, 0xfe, 0x42, 0xb4, 0x93, 0xc8, 0xfc, 0x0b, 0x32, 0xc3, 0xeb, 0x53,
	0x29, 0x51, 0xfd, 0x47, 0x05, 0x0a, 0x11, 0x28, 0xba, 0x07, 0xea, 0xf0, 0xac, 0xdf, 0x6b, 0x1f,
	0x76, 0x5e, 0x76, 0xda, 0x2d, 0x6d, 0x25, 0x52, 0x1c, 0x1d, 0x0c, 0x8f, 0xda, 0xa3, 0xce, 0xd9,
	0xe0, 0xc9, 0x63, 0x4d, 0x41, 0x1a, 0x94, 0x85, 0xa2, 0xd5, 0x1d, 0x36, 0x4f, 0xda, 0x5a, 0x0e,
	0x3d, 0x04, 0x24, 0x35, 0x9d, 0xfe, 0x00, 0x77, 0x9a, 0xc3, 0x41, 0xa7, 0x7b, 0xa6, 0xe5, 0xd1,
	0x7d, 0xd0, 0x0e, 0x87, 0xa7, 0xc3, 0x93, 0x83, 0x41, 0xe7, 0x3c, 0xf6, 0x2f, 0xa0, 0x07, 0x50,
	0xc9, 0x68, 0x25, 0xc8, 0x2a, 0xda, 0x82, 0xff, 0x65, 0xd5, 0x59, 0xa4, 0x22, 0x52, 0x61, 0xad,
	0x3f, 0x3c, 0x3d, 0x3d, 0xc0, 0x5f, 0x6a, 0x6b, 0xfa, 0x0b, 0x28, 0xc5, 0x21, 0x20, 0x0d, 0xf2,
	0xd7, 0x64, 0x26, 0xcb, 0x11, 0x89, 0xff, 0x5c, 0x0d, 0xfd, 0x57, 0x05, 0x20, 0xed, 0x1b, 0x74,
	0x08, 0xf7, 0x28, 0x33, 0x03, 0x36, 0x4a, 0x26, 0x48, 0xb6, 0x73, 0xcd, 0x10, 0x23, 0x64, 0xc4,
	0x23, 0xc4, 0xbb, 0x8d, 0x5b, 0xe0, 0x4d, 0xee, 0x92, 0x9c, 0xd1, 0xe7, 0x50, 0x16, 0x55, 0xb8,
	0x31, 0xa7, 0xe1, 0x1b, 0xf6, 0x2e, 0x0f, 0xe2, 0x3c, 0xb2, 0xc7, 0xea, 0x34, 0x91, 0x29, 0x7a,
	0x0e, 0x45, 0xdf, 0xb3, 0x5d, 0x46, 0xab, 0x79, 0x8e, 0xa2, 0x2f, 0x44, 0xe9, 0x45, 0xa6, 0x58,
	0x7a, 0xe8, 0x9f, 0x01, 0xa4, 0xb0, 0xe8, 0x3e, 0xac, 0x72, 0x3e, 0x32, 0x3f, 0xe2, 0x80, 0xb6,
	0x60, 0xfd, 0xca, 0xa4, 0x82, 0x29, 0xcf, 0x4f, 0x09, 0x97, 0xae, 0x4c, 0xca, 0x5d, 0xf4, 0x9f,
	0x73, 0xb0, 0xca, 0x21, 0xd1, 0x33, 0x58, 0x5f, 0x26, 0x23, 0xa9, 0x31, 0x7a, 0x17, 0x54, 0xdb,
	0x65, 0x4f, 0x1e, 0x67, 0x3e, 0x91, 0x3f, 0x5e, 0xc1, 0xc0, 0x95, 0x82, 0xd9, 0x7b, 0x50, 0x9e,
	0x78, 0xe1, 0x78, 0x4a, 0xa4, 0x4d, 0x34, 0x19, 0xca, 0xf1, 0x0a, 0x56, 0x85, 0x56, 0x18, 0x8d,
	0x00, 0x4d, 0x6c, 0xca, 0x02, 0x7b, 0x1c, 0x46, 0x85, 0x93, 0xa6, 0x05, 0x4e, 0xc5, 0x58, 0x98,
	0x94, 0x56, 0xc6, 0x8d, 0x63, 0x1d, 0xaf, 0xe0, 0xca, 0xe4, 0xae, 0x12, 0xf5, 0x60, 0x83, 0x86,
	0x8e, 0x63, 0x06, 0x33, 0x89, 0xbd, 0xca, 0xb1, 0x1f, 0x2d, 0xc4, 0xee, 0x0b, 0x8f, 0x18, 0xb6,
	0x4c, 0x33, 0xe7, 0xe6, 0x9a, 0xcc, 0xb8, 0xfe, 0x4b, 0x11, 0x2a, 0x73, 0x2c, 0xa2, 0x82, 0x58,
	0x5e, 0xe8, 0x32, 0x9e, 0xcf, 0x3c, 0x16, 0x87, 0xa8, 0x89, 0x69, 0xe8, 0xf0, 0x3c, 0x29, 0x38,
	0x12, 0xd1, 0x53, 0xa8, 0xd2, 0xd0, 0x19, 0x79, 0x17, 0x23, 0xfa, 0x3a, 0x34, 0x03, 0x32, 0x19,
	0x4d, 0xc8, 0x8d, 0x6d, 0xf2, 0x8e, 0xe6, 0xa9, 0xc2, 0x0f, 0x68, 0xe8, 0x74, 0x2f, 0xfa, 0xe2,
	0xb6, 0x15, 0x5f, 0x22, 0x0b, 0x36, 0xc7, 0xa1, 0x75, 0x4d, 0xd8, 0xc8, 0xe3, 0xcd, 0x4e, 0x65,
	0xba, 0x3e, 0x5d, 0x2e, 0x5d, 0x46, 0x93, 0x83, 0x74, 0x05, 0x06, 0xde, 0x18, 0x67, 0x8f, 0xa8,
	0x0b, 0x6b, 0x42, 0x11, 0xef, 0x9b, 0x4f, 0xde, 0x0a, 0x1d, 0xc7, 0x28, 0xb5, 0x1f, 0x14, 0xd8,
	0xb8, 0xf5, 0x45, 0x64, 0x41, 0x89, 0x7c, 0xe3, 0x4f, 0x6d, 0xcb, 0x66, 0xb2, 0xf7, 0xda, 0xff,
	0x26, 0x02, 0xa3, 0x2d, 0xc1, 0x8e, 0x57, 0x70, 0x02, 0x5c, 0xd3, 0xa1, 0x14, 0xeb, 0xd1, 0x43,
	0x28, 0x8e, 0xbd, 0xd0, 0x9d, 0xd0, 0xaa, 0xb2, 0x93, 0xdf, 0x55, 0xb0, 0x3c, 0x35, 0x8b, 0x62,
	0x4d, 0xd7, 0x28, 0x14, 0x05, 0xe2, 0xdf, 0xd4, 0xb0, 0x1f, 0x11, 0x26, 0x8e, 0x3f, 0x35, 0x03,
	0x5e, 0x48, 0xb5, 0xf1, 0x74, 0x49, 0xc2, 0x6d, 0xe9, 0x8e, 0x13, 0xa0, 0xda, 0xb7, 0xb9, 0x88,
	0xa1, 0x38, 0xdc, 0x1e, 0x66, 0x25, 0x1e, 0xe6, 0x5b, 0x53, 0x9a, 0x5b, 0x66, 0x4a, 0xbf, 0x02,
	0xd5, 0x64, 0xcc, 0xb4, 0xae, 0x1c, 0x92, 0xee, 0x9a, 0xe3, 0xb7, 0x24, 0x6d, 0x1c, 0xa4, 0x50,
	0x6d, 0x97, 0x05, 0x33, 0x9c, 0x05, 0xaf, 0xbd, 0x00, 0xed, 0xae, 0xc1, 0x5f, 0xac, 0xee, 0x24,
	0xc2, 0x5c, 0x66, 0x5d, 0x3d, 0xcf, 0x3d, 0x53, 0xf4, 0xdf, 0xf3, 0x50, 0xce, 0xce, 0x1d, 0xda,
	0xcf, 0x16, 0x41, 0x6d, 0x6c, 0xcd, 0x85, 0xdc, 0x49, 0x76, 0x4d, 0x5c, 0x21, 0x23, 0x9d, 0x32,
	0xb5, 0xf1, 0xff, 0x39, 0x87, 0x56, 0xba, 0x78, 0xc4, 0x0c, 0x9e, 0x41, 0x89, 0xba, 0xa6, 0x4f,
	0xaf, 0x3c, 0x26, 0xdf, 0x10, 0x8d, 0x37, 0xde, 0x0b, 0x46, 0x5f, 0x7a, 0xe2, 0x04, 0xa3, 0xf6,
	0x53, 0x0e, 0x4a, 0xb1, 0xfa, 0xbf, 0xe0, 0xff, 0x1a, 0x2a, 0x3e, 0x09, 0x2c, 0xe2, 0x32, 0x3b,
	0x5e, 0xb3, 0x71, 0x95, 0x5b, 0xcb, 0x07, 0x62, 0xf0, 0xe3, 0x01, 0xeb, 0x25, 0x90, 0x58, 0x4b,
	0xe1, 0xc5, 0x9f, 0xab, 0xd6, 0x81, 0xca, 0x9c, 0x19, 0xda, 0x06, 0x48, 0x0d, 0x65, 0xf3, 0x66,
	0x34, 0xb7, 0xab, 0x1e, 0xf7, 0x75, 0xf3, 0x3b, 0x05, 0xb6, 0x6d, 0x6f, 0x11, 0xcf, 0x66, 0x59,
	0x3c, 0x8b, 0x68, 0x2f, 0xba, 0xe8, 0x29, 0xaf, 0x5a, 0x97, 0x36, 0xbb, 0x0a, 0xc7, 0x86, 0xe5,
	0x39, 0x75, 0xe1, 0xb3, 0x67, 0xbb, 0x94, 0x05, 0x61, 0xd4, 0x74, 0x7c, 0x3d, 0xd6, 0x53, 0xb8,
	0x3d, 0xf1, 0xe6, 0xbd, 0x24, 0xee, 0xde, 0x65, 0xf6, 0x0d, 0xfe, 0x5b, 0x6e, 0xab, 0xeb, 0x13,
	0xf7, 0x50, 0x7c, 0x93, 0x43, 0xcb, 0xe7, 0x17, 0x35, 0xce, 0xf7, 0xc7, 0x45, 0xee, 0xf6, 0xf1,
	0x9f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xf8, 0xd0, 0xb4, 0x8d, 0xc7, 0x0b, 0x00, 0x00,
}
