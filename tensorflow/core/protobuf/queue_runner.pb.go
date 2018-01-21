// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tensorflow/core/protobuf/queue_runner.proto

package tensorflow

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import tensorflow_error "tensorflow/core/lib/core"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Protocol buffer representing a QueueRunner.
type QueueRunnerDef struct {
	// Queue name.
	QueueName string `protobuf:"bytes,1,opt,name=queue_name,json=queueName" json:"queue_name,omitempty"`
	// A list of enqueue operations.
	EnqueueOpName []string `protobuf:"bytes,2,rep,name=enqueue_op_name,json=enqueueOpName" json:"enqueue_op_name,omitempty"`
	// The operation to run to close the queue.
	CloseOpName string `protobuf:"bytes,3,opt,name=close_op_name,json=closeOpName" json:"close_op_name,omitempty"`
	// The operation to run to cancel the queue.
	CancelOpName string `protobuf:"bytes,4,opt,name=cancel_op_name,json=cancelOpName" json:"cancel_op_name,omitempty"`
	// A list of exception types considered to signal a safely closed queue
	// if raised during enqueue operations.
	QueueClosedExceptionTypes []tensorflow_error.Code `protobuf:"varint,5,rep,packed,name=queue_closed_exception_types,json=queueClosedExceptionTypes,enum=tensorflow.error.Code" json:"queue_closed_exception_types,omitempty"`
}

func (m *QueueRunnerDef) Reset()                    { *m = QueueRunnerDef{} }
func (m *QueueRunnerDef) String() string            { return proto.CompactTextString(m) }
func (*QueueRunnerDef) ProtoMessage()               {}
func (*QueueRunnerDef) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{0} }

func (m *QueueRunnerDef) GetQueueName() string {
	if m != nil {
		return m.QueueName
	}
	return ""
}

func (m *QueueRunnerDef) GetEnqueueOpName() []string {
	if m != nil {
		return m.EnqueueOpName
	}
	return nil
}

func (m *QueueRunnerDef) GetCloseOpName() string {
	if m != nil {
		return m.CloseOpName
	}
	return ""
}

func (m *QueueRunnerDef) GetCancelOpName() string {
	if m != nil {
		return m.CancelOpName
	}
	return ""
}

func (m *QueueRunnerDef) GetQueueClosedExceptionTypes() []tensorflow_error.Code {
	if m != nil {
		return m.QueueClosedExceptionTypes
	}
	return nil
}

func init() {
	proto.RegisterType((*QueueRunnerDef)(nil), "tensorflow.QueueRunnerDef")
}

func init() { proto.RegisterFile("tensorflow/core/protobuf/queue_runner.proto", fileDescriptor8) }

var fileDescriptor8 = []byte{
	// 270 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0x51, 0x4b, 0x84, 0x40,
	0x14, 0x85, 0x99, 0xb6, 0x02, 0xa7, 0xd6, 0xc8, 0x87, 0xb0, 0x28, 0x90, 0x25, 0x42, 0x0a, 0x14,
	0xb6, 0x7f, 0xb0, 0x5b, 0xaf, 0xb5, 0x49, 0xd0, 0xa3, 0xe8, 0x78, 0x8d, 0x25, 0x9d, 0x6b, 0x77,
	0x94, 0xad, 0x7f, 0x1e, 0x3d, 0x85, 0x77, 0x2c, 0x97, 0xde, 0x86, 0xc3, 0xf7, 0x9d, 0xe1, 0x1e,
	0x79, 0xd3, 0x82, 0x36, 0x48, 0x65, 0x85, 0x9b, 0x58, 0x21, 0x41, 0xdc, 0x10, 0xb6, 0x98, 0x77,
	0x65, 0xfc, 0xde, 0x41, 0x07, 0x29, 0x75, 0x5a, 0x03, 0x45, 0x9c, 0x7a, 0x72, 0x84, 0xcf, 0xae,
	0xff, 0x8b, 0xd5, 0x3a, 0xb7, 0x0f, 0x20, 0x42, 0x4a, 0x15, 0x16, 0x60, 0xac, 0x37, 0xfb, 0x16,
	0xd2, 0x7d, 0xea, 0xeb, 0x12, 0x6e, 0xbb, 0x83, 0xd2, 0xbb, 0x90, 0xd2, 0x7e, 0xa0, 0xb3, 0x1a,
	0x7c, 0x11, 0x88, 0xd0, 0x49, 0x1c, 0x4e, 0x1e, 0xb2, 0x1a, 0xbc, 0x2b, 0x79, 0x04, 0xda, 0x02,
	0xd8, 0x58, 0x66, 0x27, 0x98, 0x84, 0x4e, 0x32, 0x1d, 0xe2, 0xc7, 0x86, 0xb9, 0x99, 0x9c, 0xaa,
	0x0a, 0xcd, 0x48, 0x4d, 0xb8, 0xe9, 0x80, 0xc3, 0x81, 0xb9, 0x94, 0xae, 0xca, 0xb4, 0x82, 0xea,
	0x0f, 0xda, 0x65, 0xe8, 0xd0, 0xa6, 0x03, 0xf5, 0x22, 0xcf, 0xed, 0x7f, 0xac, 0x16, 0x29, 0x7c,
	0x28, 0x68, 0xda, 0x35, 0xea, 0xb4, 0xfd, 0x6c, 0xc0, 0xf8, 0x7b, 0xc1, 0x24, 0x74, 0xe7, 0x27,
	0xd1, 0x78, 0x76, 0xc4, 0x87, 0x46, 0x4b, 0x2c, 0x20, 0x39, 0x65, 0x77, 0xc9, 0xea, 0xfd, 0xaf,
	0xf9, 0xdc, 0x8b, 0x8b, 0xb9, 0xf4, 0x91, 0x5e, 0xb7, 0xbd, 0x92, 0xb2, 0x1a, 0x36, 0x48, 0x6f,
	0x8b, 0xe3, 0xad, 0x55, 0x56, 0xfd, 0x54, 0x66, 0x25, 0xbe, 0x84, 0xc8, 0xf7, 0x79, 0xb7, 0xdb,
	0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd7, 0xe1, 0xd0, 0x28, 0x9e, 0x01, 0x00, 0x00,
}