// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow/core/protobuf/master.proto

package tensorflow

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import tensorflow15 "tensorflow/core/framework"
import tensorflow11 "tensorflow/core/framework"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type CreateSessionRequest struct {
	// The initial graph definition.
	GraphDef *tensorflow11.GraphDef `protobuf:"bytes,1,opt,name=graph_def,json=graphDef" json:"graph_def,omitempty"`
	// Configuration options.
	Config *ConfigProto `protobuf:"bytes,2,opt,name=config" json:"config,omitempty"`
	// The target string used from the client's perspective.
	Target string `protobuf:"bytes,3,opt,name=target" json:"target,omitempty"`
}

func (m *CreateSessionRequest) Reset()                    { *m = CreateSessionRequest{} }
func (m *CreateSessionRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateSessionRequest) ProtoMessage()               {}
func (*CreateSessionRequest) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{0} }

func (m *CreateSessionRequest) GetGraphDef() *tensorflow11.GraphDef {
	if m != nil {
		return m.GraphDef
	}
	return nil
}

func (m *CreateSessionRequest) GetConfig() *ConfigProto {
	if m != nil {
		return m.Config
	}
	return nil
}

func (m *CreateSessionRequest) GetTarget() string {
	if m != nil {
		return m.Target
	}
	return ""
}

type CreateSessionResponse struct {
	// The session handle to be used in subsequent calls for the created session.
	//
	// The client must arrange to call CloseSession with this returned
	// session handle to close the session.
	SessionHandle string `protobuf:"bytes,1,opt,name=session_handle,json=sessionHandle" json:"session_handle,omitempty"`
	// The initial version number for the graph, to be used in the next call
	// to ExtendSession.
	GraphVersion int64 `protobuf:"varint,2,opt,name=graph_version,json=graphVersion" json:"graph_version,omitempty"`
}

func (m *CreateSessionResponse) Reset()                    { *m = CreateSessionResponse{} }
func (m *CreateSessionResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateSessionResponse) ProtoMessage()               {}
func (*CreateSessionResponse) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{1} }

func (m *CreateSessionResponse) GetSessionHandle() string {
	if m != nil {
		return m.SessionHandle
	}
	return ""
}

func (m *CreateSessionResponse) GetGraphVersion() int64 {
	if m != nil {
		return m.GraphVersion
	}
	return 0
}

type ExtendSessionRequest struct {
	// REQUIRED: session_handle must be returned by a CreateSession call
	// to the same master service.
	SessionHandle string `protobuf:"bytes,1,opt,name=session_handle,json=sessionHandle" json:"session_handle,omitempty"`
	// REQUIRED: The nodes to be added to the session's graph. If any node has
	// the same name as an existing node, the operation will fail with
	// ILLEGAL_ARGUMENT.
	GraphDef *tensorflow11.GraphDef `protobuf:"bytes,2,opt,name=graph_def,json=graphDef" json:"graph_def,omitempty"`
	// REQUIRED: The version number of the graph to be extended. This will be
	// tested against the current server-side version number, and the operation
	// will fail with FAILED_PRECONDITION if they do not match.
	CurrentGraphVersion int64 `protobuf:"varint,3,opt,name=current_graph_version,json=currentGraphVersion" json:"current_graph_version,omitempty"`
}

func (m *ExtendSessionRequest) Reset()                    { *m = ExtendSessionRequest{} }
func (m *ExtendSessionRequest) String() string            { return proto.CompactTextString(m) }
func (*ExtendSessionRequest) ProtoMessage()               {}
func (*ExtendSessionRequest) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{2} }

func (m *ExtendSessionRequest) GetSessionHandle() string {
	if m != nil {
		return m.SessionHandle
	}
	return ""
}

func (m *ExtendSessionRequest) GetGraphDef() *tensorflow11.GraphDef {
	if m != nil {
		return m.GraphDef
	}
	return nil
}

func (m *ExtendSessionRequest) GetCurrentGraphVersion() int64 {
	if m != nil {
		return m.CurrentGraphVersion
	}
	return 0
}

type ExtendSessionResponse struct {
	// The new version number for the extended graph, to be used in the next call
	// to ExtendSession.
	NewGraphVersion int64 `protobuf:"varint,4,opt,name=new_graph_version,json=newGraphVersion" json:"new_graph_version,omitempty"`
}

func (m *ExtendSessionResponse) Reset()                    { *m = ExtendSessionResponse{} }
func (m *ExtendSessionResponse) String() string            { return proto.CompactTextString(m) }
func (*ExtendSessionResponse) ProtoMessage()               {}
func (*ExtendSessionResponse) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{3} }

func (m *ExtendSessionResponse) GetNewGraphVersion() int64 {
	if m != nil {
		return m.NewGraphVersion
	}
	return 0
}

type RunStepRequest struct {
	// REQUIRED: session_handle must be returned by a CreateSession call
	// to the same master service.
	SessionHandle string `protobuf:"bytes,1,opt,name=session_handle,json=sessionHandle" json:"session_handle,omitempty"`
	// Tensors to be fed in the step. Each feed is a named tensor.
	Feed []*NamedTensorProto `protobuf:"bytes,2,rep,name=feed" json:"feed,omitempty"`
	// Fetches. A list of tensor names. The caller expects a tensor to
	// be returned for each fetch[i] (see RunStepResponse.tensor). The
	// order of specified fetches does not change the execution order.
	Fetch []string `protobuf:"bytes,3,rep,name=fetch" json:"fetch,omitempty"`
	// Target Nodes. A list of node names. The named nodes will be run
	// to but their outputs will not be fetched.
	Target []string `protobuf:"bytes,4,rep,name=target" json:"target,omitempty"`
	// Options for the run call.
	Options *RunOptions `protobuf:"bytes,5,opt,name=options" json:"options,omitempty"`
	// Partial run handle (optional). If specified, this will be a partial run
	// execution, run up to the specified fetches.
	PartialRunHandle string `protobuf:"bytes,6,opt,name=partial_run_handle,json=partialRunHandle" json:"partial_run_handle,omitempty"`
}

func (m *RunStepRequest) Reset()                    { *m = RunStepRequest{} }
func (m *RunStepRequest) String() string            { return proto.CompactTextString(m) }
func (*RunStepRequest) ProtoMessage()               {}
func (*RunStepRequest) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{4} }

func (m *RunStepRequest) GetSessionHandle() string {
	if m != nil {
		return m.SessionHandle
	}
	return ""
}

func (m *RunStepRequest) GetFeed() []*NamedTensorProto {
	if m != nil {
		return m.Feed
	}
	return nil
}

func (m *RunStepRequest) GetFetch() []string {
	if m != nil {
		return m.Fetch
	}
	return nil
}

func (m *RunStepRequest) GetTarget() []string {
	if m != nil {
		return m.Target
	}
	return nil
}

func (m *RunStepRequest) GetOptions() *RunOptions {
	if m != nil {
		return m.Options
	}
	return nil
}

func (m *RunStepRequest) GetPartialRunHandle() string {
	if m != nil {
		return m.PartialRunHandle
	}
	return ""
}

type RunStepResponse struct {
	// NOTE: The order of the returned tensors may or may not match
	// the fetch order specified in RunStepRequest.
	Tensor []*NamedTensorProto `protobuf:"bytes,1,rep,name=tensor" json:"tensor,omitempty"`
	// Returned metadata if requested in the options.
	Metadata *RunMetadata `protobuf:"bytes,2,opt,name=metadata" json:"metadata,omitempty"`
}

func (m *RunStepResponse) Reset()                    { *m = RunStepResponse{} }
func (m *RunStepResponse) String() string            { return proto.CompactTextString(m) }
func (*RunStepResponse) ProtoMessage()               {}
func (*RunStepResponse) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{5} }

func (m *RunStepResponse) GetTensor() []*NamedTensorProto {
	if m != nil {
		return m.Tensor
	}
	return nil
}

func (m *RunStepResponse) GetMetadata() *RunMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

type PartialRunSetupRequest struct {
	// REQUIRED: session_handle must be returned by a CreateSession call
	// to the same master service.
	SessionHandle string `protobuf:"bytes,1,opt,name=session_handle,json=sessionHandle" json:"session_handle,omitempty"`
	// Tensors to be fed in future steps.
	Feed []string `protobuf:"bytes,2,rep,name=feed" json:"feed,omitempty"`
	// Fetches. A list of tensor names. The caller expects a tensor to be returned
	// for each fetch[i] (see RunStepResponse.tensor), for corresponding partial
	// RunStepRequests. The order of specified fetches does not change the
	// execution order.
	Fetch []string `protobuf:"bytes,3,rep,name=fetch" json:"fetch,omitempty"`
	// Target Nodes. A list of node names. The named nodes will be run in future
	// steps, but their outputs will not be fetched.
	Target []string `protobuf:"bytes,4,rep,name=target" json:"target,omitempty"`
}

func (m *PartialRunSetupRequest) Reset()                    { *m = PartialRunSetupRequest{} }
func (m *PartialRunSetupRequest) String() string            { return proto.CompactTextString(m) }
func (*PartialRunSetupRequest) ProtoMessage()               {}
func (*PartialRunSetupRequest) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{6} }

func (m *PartialRunSetupRequest) GetSessionHandle() string {
	if m != nil {
		return m.SessionHandle
	}
	return ""
}

func (m *PartialRunSetupRequest) GetFeed() []string {
	if m != nil {
		return m.Feed
	}
	return nil
}

func (m *PartialRunSetupRequest) GetFetch() []string {
	if m != nil {
		return m.Fetch
	}
	return nil
}

func (m *PartialRunSetupRequest) GetTarget() []string {
	if m != nil {
		return m.Target
	}
	return nil
}

type PartialRunSetupResponse struct {
	// The unique handle corresponding to the ongoing partial run call setup by
	// the invocation to PartialRunSetup. This handle may be passed to
	// RunStepRequest to send and receive tensors for this partial run.
	PartialRunHandle string `protobuf:"bytes,1,opt,name=partial_run_handle,json=partialRunHandle" json:"partial_run_handle,omitempty"`
}

func (m *PartialRunSetupResponse) Reset()                    { *m = PartialRunSetupResponse{} }
func (m *PartialRunSetupResponse) String() string            { return proto.CompactTextString(m) }
func (*PartialRunSetupResponse) ProtoMessage()               {}
func (*PartialRunSetupResponse) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{7} }

func (m *PartialRunSetupResponse) GetPartialRunHandle() string {
	if m != nil {
		return m.PartialRunHandle
	}
	return ""
}

type CloseSessionRequest struct {
	// REQUIRED: session_handle must be returned by a CreateSession call
	// to the same master service.
	SessionHandle string `protobuf:"bytes,1,opt,name=session_handle,json=sessionHandle" json:"session_handle,omitempty"`
}

func (m *CloseSessionRequest) Reset()                    { *m = CloseSessionRequest{} }
func (m *CloseSessionRequest) String() string            { return proto.CompactTextString(m) }
func (*CloseSessionRequest) ProtoMessage()               {}
func (*CloseSessionRequest) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{8} }

func (m *CloseSessionRequest) GetSessionHandle() string {
	if m != nil {
		return m.SessionHandle
	}
	return ""
}

type CloseSessionResponse struct {
}

func (m *CloseSessionResponse) Reset()                    { *m = CloseSessionResponse{} }
func (m *CloseSessionResponse) String() string            { return proto.CompactTextString(m) }
func (*CloseSessionResponse) ProtoMessage()               {}
func (*CloseSessionResponse) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{9} }

// Reset() allows misbehaving or slow sessions to be aborted and closed, and
// causes their resources eventually to be released.  Reset() does not wait
// for the computations in old sessions to cease; it merely starts the
// process of tearing them down.  However, if a new session is started after
// a Reset(), the new session is isolated from changes that old sessions
// (started prior to the Reset()) may continue to make to resources, provided
// all those resources are in containers listed in "containers".
//
// Old sessions may continue to have side-effects on resources not in
// containers listed in "containers", and thus may affect future
// sessions' results in ways that are hard to predict.  Thus, if well-defined
// behavior is desired, is it recommended that all containers be listed in
// "containers".  Similarly, if a device_filter is specified, results may be
// hard to predict.
type ResetRequest struct {
	// A list of container names, which may be empty.
	//
	// If 'container' is not empty, releases resoures in the given
	// containers in all devices.
	//
	// If 'container' is empty, releases resources in the default
	// container in all devices.
	Container []string `protobuf:"bytes,1,rep,name=container" json:"container,omitempty"`
	// When any filters are present, only devices that match the filters
	// will be reset. Each filter can be partially specified,
	// e.g. "/job:ps" "/job:worker/replica:3", etc.
	DeviceFilters []string `protobuf:"bytes,2,rep,name=device_filters,json=deviceFilters" json:"device_filters,omitempty"`
}

func (m *ResetRequest) Reset()                    { *m = ResetRequest{} }
func (m *ResetRequest) String() string            { return proto.CompactTextString(m) }
func (*ResetRequest) ProtoMessage()               {}
func (*ResetRequest) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{10} }

func (m *ResetRequest) GetContainer() []string {
	if m != nil {
		return m.Container
	}
	return nil
}

func (m *ResetRequest) GetDeviceFilters() []string {
	if m != nil {
		return m.DeviceFilters
	}
	return nil
}

type ResetResponse struct {
}

func (m *ResetResponse) Reset()                    { *m = ResetResponse{} }
func (m *ResetResponse) String() string            { return proto.CompactTextString(m) }
func (*ResetResponse) ProtoMessage()               {}
func (*ResetResponse) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{11} }

type ListDevicesRequest struct {
	// Optional: session_handle must be returned by a CreateSession call to the
	// same master service.
	//
	// When session_handle is empty, the ClusterSpec provided when the master was
	// started is used to compute the available devices. If the session_handle is
	// provided but not recognized, an error is returned. Finally, if a valid
	// session_handle is provided, the cluster configuration for that session is
	// used when computing the response.
	SessionHandle string `protobuf:"bytes,1,opt,name=session_handle,json=sessionHandle" json:"session_handle,omitempty"`
}

func (m *ListDevicesRequest) Reset()                    { *m = ListDevicesRequest{} }
func (m *ListDevicesRequest) String() string            { return proto.CompactTextString(m) }
func (*ListDevicesRequest) ProtoMessage()               {}
func (*ListDevicesRequest) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{12} }

func (m *ListDevicesRequest) GetSessionHandle() string {
	if m != nil {
		return m.SessionHandle
	}
	return ""
}

type ListDevicesResponse struct {
	LocalDevice  []*tensorflow15.DeviceAttributes `protobuf:"bytes,1,rep,name=local_device,json=localDevice" json:"local_device,omitempty"`
	RemoteDevice []*tensorflow15.DeviceAttributes `protobuf:"bytes,2,rep,name=remote_device,json=remoteDevice" json:"remote_device,omitempty"`
}

func (m *ListDevicesResponse) Reset()                    { *m = ListDevicesResponse{} }
func (m *ListDevicesResponse) String() string            { return proto.CompactTextString(m) }
func (*ListDevicesResponse) ProtoMessage()               {}
func (*ListDevicesResponse) Descriptor() ([]byte, []int) { return fileDescriptor5, []int{13} }

func (m *ListDevicesResponse) GetLocalDevice() []*tensorflow15.DeviceAttributes {
	if m != nil {
		return m.LocalDevice
	}
	return nil
}

func (m *ListDevicesResponse) GetRemoteDevice() []*tensorflow15.DeviceAttributes {
	if m != nil {
		return m.RemoteDevice
	}
	return nil
}

func init() {
	proto.RegisterType((*CreateSessionRequest)(nil), "tensorflow.CreateSessionRequest")
	proto.RegisterType((*CreateSessionResponse)(nil), "tensorflow.CreateSessionResponse")
	proto.RegisterType((*ExtendSessionRequest)(nil), "tensorflow.ExtendSessionRequest")
	proto.RegisterType((*ExtendSessionResponse)(nil), "tensorflow.ExtendSessionResponse")
	proto.RegisterType((*RunStepRequest)(nil), "tensorflow.RunStepRequest")
	proto.RegisterType((*RunStepResponse)(nil), "tensorflow.RunStepResponse")
	proto.RegisterType((*PartialRunSetupRequest)(nil), "tensorflow.PartialRunSetupRequest")
	proto.RegisterType((*PartialRunSetupResponse)(nil), "tensorflow.PartialRunSetupResponse")
	proto.RegisterType((*CloseSessionRequest)(nil), "tensorflow.CloseSessionRequest")
	proto.RegisterType((*CloseSessionResponse)(nil), "tensorflow.CloseSessionResponse")
	proto.RegisterType((*ResetRequest)(nil), "tensorflow.ResetRequest")
	proto.RegisterType((*ResetResponse)(nil), "tensorflow.ResetResponse")
	proto.RegisterType((*ListDevicesRequest)(nil), "tensorflow.ListDevicesRequest")
	proto.RegisterType((*ListDevicesResponse)(nil), "tensorflow.ListDevicesResponse")
}

func init() { proto.RegisterFile("tensorflow/core/protobuf/master.proto", fileDescriptor5) }

var fileDescriptor5 = []byte{
	// 673 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x55, 0x4f, 0x4f, 0x1b, 0x3f,
	0x10, 0xd5, 0x26, 0x90, 0x1f, 0x19, 0x12, 0xf8, 0xd5, 0x84, 0xb0, 0x42, 0x1c, 0xa2, 0xad, 0x90,
	0xa2, 0xb6, 0x22, 0xfc, 0xe9, 0x8d, 0x4a, 0x15, 0x84, 0x96, 0x1e, 0xfa, 0x07, 0x39, 0x55, 0xaf,
	0x2b, 0xb3, 0x3b, 0x1b, 0x56, 0xdd, 0xd8, 0xa9, 0xed, 0x25, 0x3d, 0xf4, 0xd2, 0xaf, 0xd0, 0x5b,
	0xef, 0xfd, 0x8e, 0xed, 0xb1, 0x8a, 0x6d, 0xc2, 0x26, 0x25, 0x52, 0xe0, 0xc6, 0xbc, 0x79, 0x6f,
	0xe6, 0xf9, 0xd9, 0x6c, 0x60, 0x57, 0x23, 0x57, 0x42, 0x26, 0x99, 0x18, 0x75, 0x22, 0x21, 0xb1,
	0x33, 0x94, 0x42, 0x8b, 0xcb, 0x3c, 0xe9, 0x0c, 0x98, 0xd2, 0x28, 0xf7, 0x4c, 0x4d, 0xe0, 0x96,
	0xb6, 0x7d, 0x30, 0x2b, 0x49, 0x24, 0x1b, 0xe0, 0x48, 0xc8, 0xcf, 0x9d, 0x18, 0xaf, 0xd3, 0x08,
	0x43, 0xa6, 0xb5, 0x4c, 0x2f, 0x73, 0x8d, 0xca, 0xca, 0xb7, 0x77, 0xe7, 0x4b, 0xfa, 0x92, 0x0d,
	0xaf, 0xe6, 0xd1, 0x26, 0x66, 0x22, 0xc1, 0x93, 0xb4, 0xef, 0x68, 0x4f, 0xe7, 0xd2, 0x38, 0x1b,
	0x60, 0x1c, 0xda, 0xb6, 0x25, 0x07, 0x3f, 0x3c, 0x68, 0x74, 0x25, 0x32, 0x8d, 0x3d, 0x54, 0x2a,
	0x15, 0x9c, 0xe2, 0x97, 0x1c, 0x95, 0x26, 0x07, 0x50, 0x35, 0xbb, 0xc3, 0x18, 0x13, 0xdf, 0x6b,
	0x79, 0xed, 0xd5, 0xc3, 0xc6, 0xde, 0xed, 0xe4, 0xbd, 0xf3, 0x71, 0xf3, 0x0c, 0x13, 0xba, 0xd2,
	0x77, 0x7f, 0x91, 0x0e, 0x54, 0xac, 0x11, 0xbf, 0x64, 0xf8, 0x5b, 0x45, 0x7e, 0xd7, 0x74, 0x2e,
	0xc6, 0x4b, 0xa9, 0xa3, 0x91, 0x26, 0x54, 0x34, 0x93, 0x7d, 0xd4, 0x7e, 0xb9, 0xe5, 0xb5, 0xab,
	0xd4, 0x55, 0x41, 0x04, 0x9b, 0x33, 0x9e, 0xd4, 0x50, 0x70, 0x85, 0x64, 0x17, 0xd6, 0x94, 0x85,
	0xc2, 0x2b, 0xc6, 0xe3, 0x0c, 0x8d, 0xb3, 0x2a, 0xad, 0x3b, 0xf4, 0x8d, 0x01, 0xc9, 0x63, 0xa8,
	0x5b, 0xef, 0xd7, 0x28, 0xc7, 0xb0, 0xf1, 0x53, 0xa6, 0x35, 0x03, 0x7e, 0xb2, 0x58, 0xf0, 0xcb,
	0x83, 0xc6, 0xab, 0xaf, 0x1a, 0x79, 0x3c, 0x73, 0xf2, 0x05, 0x97, 0x4c, 0x05, 0x54, 0x5a, 0x28,
	0xa0, 0x43, 0xd8, 0x8c, 0x72, 0x29, 0x91, 0xeb, 0x70, 0xda, 0x5f, 0xd9, 0xf8, 0xdb, 0x70, 0xcd,
	0xf3, 0xa2, 0xcd, 0x2e, 0x6c, 0xce, 0xb8, 0x74, 0x59, 0x3c, 0x81, 0x47, 0x1c, 0x47, 0x33, 0x83,
	0x96, 0xcc, 0xa0, 0x75, 0x8e, 0xa3, 0xa9, 0x21, 0xbf, 0x3d, 0x58, 0xa3, 0x39, 0xef, 0x69, 0x1c,
	0xde, 0xf3, 0x94, 0xfb, 0xb0, 0x94, 0x20, 0xc6, 0x7e, 0xa9, 0x55, 0x6e, 0xaf, 0x1e, 0xee, 0x14,
	0x0f, 0xf8, 0x7e, 0xfc, 0x9a, 0x3e, 0x9a, 0xda, 0x5e, 0xab, 0x61, 0x92, 0x06, 0x2c, 0x27, 0xa8,
	0xa3, 0x2b, 0xbf, 0xdc, 0x2a, 0xb7, 0xab, 0xd4, 0x16, 0x85, 0xab, 0x5e, 0x32, 0xb0, 0xab, 0xc8,
	0x3e, 0xfc, 0x27, 0x86, 0x3a, 0x15, 0x5c, 0xf9, 0xcb, 0x26, 0xc3, 0x66, 0x71, 0x05, 0xcd, 0xf9,
	0x07, 0xdb, 0xa5, 0x37, 0x34, 0xf2, 0x0c, 0xc8, 0x90, 0x49, 0x9d, 0xb2, 0x2c, 0x94, 0xf9, 0xc4,
	0x7c, 0xc5, 0x98, 0xff, 0xdf, 0x75, 0x68, 0xee, 0xfc, 0x07, 0xdf, 0x60, 0x7d, 0x72, 0x70, 0x17,
	0xdc, 0x73, 0xa8, 0xd8, 0x15, 0xbe, 0xb7, 0xc0, 0xa1, 0x1c, 0x97, 0x1c, 0xc1, 0xca, 0x00, 0x35,
	0x8b, 0x99, 0x66, 0x77, 0x3d, 0x6f, 0x9a, 0xf3, 0x77, 0xae, 0x4d, 0x27, 0xc4, 0xe0, 0xbb, 0x07,
	0xcd, 0x8b, 0x89, 0xa5, 0x1e, 0xea, 0xfc, 0xbe, 0xf9, 0x93, 0x42, 0xfe, 0xd5, 0x87, 0x24, 0x1c,
	0x9c, 0xc3, 0xd6, 0x3f, 0x16, 0x5c, 0x12, 0x77, 0x47, 0xe9, 0xcd, 0x89, 0xf2, 0x05, 0x6c, 0x74,
	0x33, 0xa1, 0xf0, 0x41, 0xff, 0x2e, 0x41, 0x13, 0x1a, 0xd3, 0x6a, 0xeb, 0x21, 0xe8, 0x41, 0x8d,
	0xa2, 0x42, 0x7d, 0x33, 0x6e, 0x07, 0xaa, 0x91, 0xe0, 0x9a, 0xa5, 0x1c, 0xed, 0x05, 0x55, 0xe9,
	0x2d, 0x30, 0x5e, 0xe6, 0x3e, 0xa2, 0x49, 0x9a, 0x69, 0x94, 0xca, 0x05, 0x53, 0xb7, 0xe8, 0x6b,
	0x0b, 0x06, 0xeb, 0x50, 0x77, 0x43, 0xdd, 0x96, 0x63, 0x20, 0x6f, 0x53, 0xa5, 0xcf, 0x0c, 0x4b,
	0xdd, 0xd3, 0xfa, 0x4f, 0x0f, 0x36, 0xa6, 0xd4, 0x2e, 0xbe, 0x97, 0x50, 0xcb, 0x44, 0xc4, 0xb2,
	0xd0, 0x2e, 0xbf, 0xeb, 0x39, 0x59, 0xc9, 0xc9, 0xe4, 0x83, 0x4f, 0x57, 0x8d, 0xc2, 0xc2, 0xe4,
	0x04, 0xea, 0x12, 0x07, 0x42, 0xe3, 0xcd, 0x84, 0xd2, 0x02, 0x13, 0x6a, 0x56, 0x62, 0xf1, 0xd3,
	0x63, 0xd8, 0x16, 0xb2, 0x5f, 0x14, 0xc4, 0xa9, 0xd2, 0x32, 0xe7, 0x3a, 0x1d, 0xe0, 0xa9, 0x7f,
	0x36, 0x2e, 0x8c, 0x30, 0xa6, 0x16, 0x33, 0xaf, 0x5a, 0x5d, 0x78, 0x7f, 0x3c, 0xef, 0xb2, 0x62,
	0x7e, 0x03, 0x8e, 0xfe, 0x06, 0x00, 0x00, 0xff, 0xff, 0x03, 0xea, 0x8d, 0xaa, 0xe6, 0x06, 0x00,
	0x00,
}
