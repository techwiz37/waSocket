// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.1
// 	protoc        v3.21.12
// source: waArmadilloICDC/WAArmadilloICDC.proto

package waArmadilloICDC

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"

	_ "embed"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ICDCIdentityList struct {
	state              protoimpl.MessageState `protogen:"open.v1"`
	Seq                *int32                 `protobuf:"varint,1,opt,name=seq" json:"seq,omitempty"`
	Timestamp          *int64                 `protobuf:"varint,2,opt,name=timestamp" json:"timestamp,omitempty"`
	Devices            [][]byte               `protobuf:"bytes,3,rep,name=devices" json:"devices,omitempty"`
	SigningDeviceIndex *int32                 `protobuf:"varint,4,opt,name=signingDeviceIndex" json:"signingDeviceIndex,omitempty"`
	unknownFields      protoimpl.UnknownFields
	sizeCache          protoimpl.SizeCache
}

func (x *ICDCIdentityList) Reset() {
	*x = ICDCIdentityList{}
	mi := &file_waArmadilloICDC_WAArmadilloICDC_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ICDCIdentityList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ICDCIdentityList) ProtoMessage() {}

func (x *ICDCIdentityList) ProtoReflect() protoreflect.Message {
	mi := &file_waArmadilloICDC_WAArmadilloICDC_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ICDCIdentityList.ProtoReflect.Descriptor instead.
func (*ICDCIdentityList) Descriptor() ([]byte, []int) {
	return file_waArmadilloICDC_WAArmadilloICDC_proto_rawDescGZIP(), []int{0}
}

func (x *ICDCIdentityList) GetSeq() int32 {
	if x != nil && x.Seq != nil {
		return *x.Seq
	}
	return 0
}

func (x *ICDCIdentityList) GetTimestamp() int64 {
	if x != nil && x.Timestamp != nil {
		return *x.Timestamp
	}
	return 0
}

func (x *ICDCIdentityList) GetDevices() [][]byte {
	if x != nil {
		return x.Devices
	}
	return nil
}

func (x *ICDCIdentityList) GetSigningDeviceIndex() int32 {
	if x != nil && x.SigningDeviceIndex != nil {
		return *x.SigningDeviceIndex
	}
	return 0
}

type SignedICDCIdentityList struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Details       []byte                 `protobuf:"bytes,1,opt,name=details" json:"details,omitempty"`
	Signature     []byte                 `protobuf:"bytes,2,opt,name=signature" json:"signature,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SignedICDCIdentityList) Reset() {
	*x = SignedICDCIdentityList{}
	mi := &file_waArmadilloICDC_WAArmadilloICDC_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SignedICDCIdentityList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignedICDCIdentityList) ProtoMessage() {}

func (x *SignedICDCIdentityList) ProtoReflect() protoreflect.Message {
	mi := &file_waArmadilloICDC_WAArmadilloICDC_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignedICDCIdentityList.ProtoReflect.Descriptor instead.
func (*SignedICDCIdentityList) Descriptor() ([]byte, []int) {
	return file_waArmadilloICDC_WAArmadilloICDC_proto_rawDescGZIP(), []int{1}
}

func (x *SignedICDCIdentityList) GetDetails() []byte {
	if x != nil {
		return x.Details
	}
	return nil
}

func (x *SignedICDCIdentityList) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

var File_waArmadilloICDC_WAArmadilloICDC_proto protoreflect.FileDescriptor

//go:embed WAArmadilloICDC.pb.raw
var file_waArmadilloICDC_WAArmadilloICDC_proto_rawDesc []byte

var (
	file_waArmadilloICDC_WAArmadilloICDC_proto_rawDescOnce sync.Once
	file_waArmadilloICDC_WAArmadilloICDC_proto_rawDescData = file_waArmadilloICDC_WAArmadilloICDC_proto_rawDesc
)

func file_waArmadilloICDC_WAArmadilloICDC_proto_rawDescGZIP() []byte {
	file_waArmadilloICDC_WAArmadilloICDC_proto_rawDescOnce.Do(func() {
		file_waArmadilloICDC_WAArmadilloICDC_proto_rawDescData = protoimpl.X.CompressGZIP(file_waArmadilloICDC_WAArmadilloICDC_proto_rawDescData)
	})
	return file_waArmadilloICDC_WAArmadilloICDC_proto_rawDescData
}

var file_waArmadilloICDC_WAArmadilloICDC_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_waArmadilloICDC_WAArmadilloICDC_proto_goTypes = []any{
	(*ICDCIdentityList)(nil),       // 0: WAArmadilloICDC.ICDCIdentityList
	(*SignedICDCIdentityList)(nil), // 1: WAArmadilloICDC.SignedICDCIdentityList
}
var file_waArmadilloICDC_WAArmadilloICDC_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_waArmadilloICDC_WAArmadilloICDC_proto_init() }
func file_waArmadilloICDC_WAArmadilloICDC_proto_init() {
	if File_waArmadilloICDC_WAArmadilloICDC_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_waArmadilloICDC_WAArmadilloICDC_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_waArmadilloICDC_WAArmadilloICDC_proto_goTypes,
		DependencyIndexes: file_waArmadilloICDC_WAArmadilloICDC_proto_depIdxs,
		MessageInfos:      file_waArmadilloICDC_WAArmadilloICDC_proto_msgTypes,
	}.Build()
	File_waArmadilloICDC_WAArmadilloICDC_proto = out.File
	file_waArmadilloICDC_WAArmadilloICDC_proto_rawDesc = nil
	file_waArmadilloICDC_WAArmadilloICDC_proto_goTypes = nil
	file_waArmadilloICDC_WAArmadilloICDC_proto_depIdxs = nil
}
