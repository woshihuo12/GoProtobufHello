// Code generated by protoc-gen-go. DO NOT EDIT.
// source: test.proto

/*
Package example is a generated protocol buffer package.

It is generated from these files:
	test.proto

It has these top-level messages:
	Test
*/
package example

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type FOO int32

const (
	FOO_Null FOO = 0
	FOO_X    FOO = 17
)

var FOO_name = map[int32]string{
	0:  "Null",
	17: "X",
}
var FOO_value = map[string]int32{
	"Null": 0,
	"X":    17,
}

func (x FOO) String() string {
	return proto.EnumName(FOO_name, int32(x))
}
func (FOO) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Test struct {
	Id    int32   `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Label string  `protobuf:"bytes,2,opt,name=label" json:"label,omitempty"`
	Type  int32   `protobuf:"varint,3,opt,name=type" json:"type,omitempty"`
	Reps  []int64 `protobuf:"varint,4,rep,packed,name=reps" json:"reps,omitempty"`
}

func (m *Test) Reset()                    { *m = Test{} }
func (m *Test) String() string            { return proto.CompactTextString(m) }
func (*Test) ProtoMessage()               {}
func (*Test) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Test) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Test) GetLabel() string {
	if m != nil {
		return m.Label
	}
	return ""
}

func (m *Test) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *Test) GetReps() []int64 {
	if m != nil {
		return m.Reps
	}
	return nil
}

func init() {
	proto.RegisterType((*Test)(nil), "example.Test")
	proto.RegisterEnum("example.FOO", FOO_name, FOO_value)
}

func init() { proto.RegisterFile("test.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 144 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2a, 0x49, 0x2d, 0x2e,
	0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4f, 0xad, 0x48, 0xcc, 0x2d, 0xc8, 0x49, 0x55,
	0x0a, 0xe1, 0x62, 0x09, 0x49, 0x2d, 0x2e, 0x11, 0xe2, 0xe3, 0x62, 0xca, 0x4c, 0x91, 0x60, 0x54,
	0x60, 0xd4, 0x60, 0x0d, 0x62, 0xca, 0x4c, 0x11, 0x12, 0xe1, 0x62, 0xcd, 0x49, 0x4c, 0x4a, 0xcd,
	0x91, 0x60, 0x52, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0x70, 0x84, 0x84, 0xb8, 0x58, 0x4a, 0x2a, 0x0b,
	0x52, 0x25, 0x98, 0xc1, 0xea, 0xc0, 0x6c, 0x90, 0x58, 0x51, 0x6a, 0x41, 0xb1, 0x04, 0x8b, 0x02,
	0xb3, 0x06, 0x73, 0x10, 0x98, 0xad, 0x25, 0xc6, 0xc5, 0xec, 0xe6, 0xef, 0x2f, 0xc4, 0xc1, 0xc5,
	0xe2, 0x57, 0x9a, 0x93, 0x23, 0xc0, 0x20, 0xc4, 0xca, 0xc5, 0x18, 0x21, 0x20, 0x98, 0xc4, 0x06,
	0xb6, 0xdd, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xc9, 0x07, 0x5d, 0xe8, 0x8b, 0x00, 0x00, 0x00,
}
