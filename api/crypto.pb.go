// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: api/crypto.proto

package api

import (
	empty "github.com/golang/protobuf/ptypes/empty"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// The request message containing the data of crypto.
type Crypto struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id          int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name        string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Code        string `protobuf:"bytes,3,opt,name=code,proto3" json:"code,omitempty"`
	Upvote      int64  `protobuf:"varint,4,opt,name=upvote,proto3" json:"upvote,omitempty"`
	Downvote    int64  `protobuf:"varint,5,opt,name=downvote,proto3" json:"downvote,omitempty"`
	Description string `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
}

func (x *Crypto) Reset() {
	*x = Crypto{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_crypto_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Crypto) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Crypto) ProtoMessage() {}

func (x *Crypto) ProtoReflect() protoreflect.Message {
	mi := &file_api_crypto_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Crypto.ProtoReflect.Descriptor instead.
func (*Crypto) Descriptor() ([]byte, []int) {
	return file_api_crypto_proto_rawDescGZIP(), []int{0}
}

func (x *Crypto) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Crypto) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Crypto) GetCode() string {
	if x != nil {
		return x.Code
	}
	return ""
}

func (x *Crypto) GetUpvote() int64 {
	if x != nil {
		return x.Upvote
	}
	return 0
}

func (x *Crypto) GetDownvote() int64 {
	if x != nil {
		return x.Downvote
	}
	return 0
}

func (x *Crypto) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

type CryptoId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CryptoId) Reset() {
	*x = CryptoId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_crypto_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CryptoId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CryptoId) ProtoMessage() {}

func (x *CryptoId) ProtoReflect() protoreflect.Message {
	mi := &file_api_crypto_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CryptoId.ProtoReflect.Descriptor instead.
func (*CryptoId) Descriptor() ([]byte, []int) {
	return file_api_crypto_proto_rawDescGZIP(), []int{1}
}

func (x *CryptoId) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type Upvotes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Total int64 `protobuf:"varint,1,opt,name=total,proto3" json:"total,omitempty"`
}

func (x *Upvotes) Reset() {
	*x = Upvotes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_api_crypto_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Upvotes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Upvotes) ProtoMessage() {}

func (x *Upvotes) ProtoReflect() protoreflect.Message {
	mi := &file_api_crypto_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Upvotes.ProtoReflect.Descriptor instead.
func (*Upvotes) Descriptor() ([]byte, []int) {
	return file_api_crypto_proto_rawDescGZIP(), []int{2}
}

func (x *Upvotes) GetTotal() int64 {
	if x != nil {
		return x.Total
	}
	return 0
}

var File_api_crypto_proto protoreflect.FileDescriptor

var file_api_crypto_proto_rawDesc = []byte{
	0x0a, 0x10, 0x61, 0x70, 0x69, 0x2f, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x06, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74,
	0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x96, 0x01, 0x0a, 0x06, 0x43, 0x72, 0x79, 0x70,
	0x74, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x75, 0x70,
	0x76, 0x6f, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x70, 0x76, 0x6f,
	0x74, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x6f, 0x77, 0x6e, 0x76, 0x6f, 0x74, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x64, 0x6f, 0x77, 0x6e, 0x76, 0x6f, 0x74, 0x65, 0x12, 0x20,
	0x0a, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0x1a, 0x0a, 0x08, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x22, 0x1f, 0x0a, 0x07,
	0x55, 0x70, 0x76, 0x6f, 0x74, 0x65, 0x73, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x32, 0xbf, 0x02,
	0x0a, 0x07, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x73, 0x12, 0x29, 0x0a, 0x03, 0x47, 0x65, 0x74,
	0x12, 0x10, 0x2e, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f,
	0x49, 0x64, 0x1a, 0x0e, 0x2e, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x79, 0x70,
	0x74, 0x6f, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x06, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x12, 0x16,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x0e, 0x2e, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e,
	0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x22, 0x00, 0x30, 0x01, 0x12, 0x2f, 0x0a, 0x09, 0x41, 0x64,
	0x64, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x12, 0x0e, 0x2e, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f,
	0x2e, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x1a, 0x10, 0x2e, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f,
	0x2e, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x49, 0x64, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x06, 0x55,
	0x70, 0x76, 0x6f, 0x74, 0x65, 0x12, 0x10, 0x2e, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x43,
	0x72, 0x79, 0x70, 0x74, 0x6f, 0x49, 0x64, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x00, 0x12, 0x36, 0x0a, 0x08, 0x44, 0x6f, 0x77, 0x6e, 0x76, 0x6f, 0x74, 0x65, 0x12, 0x10, 0x2e,
	0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x2e, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x49, 0x64, 0x1a,
	0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x34, 0x0a, 0x0b, 0x4c, 0x69, 0x76,
	0x65, 0x55, 0x70, 0x56, 0x6f, 0x74, 0x65, 0x73, 0x12, 0x10, 0x2e, 0x43, 0x72, 0x79, 0x70, 0x74,
	0x6f, 0x2e, 0x43, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x49, 0x64, 0x1a, 0x0f, 0x2e, 0x43, 0x72, 0x79,
	0x70, 0x74, 0x6f, 0x2e, 0x55, 0x70, 0x76, 0x6f, 0x74, 0x65, 0x73, 0x22, 0x00, 0x30, 0x01, 0x42,
	0x44, 0x0a, 0x17, 0x69, 0x6f, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x73, 0x2e, 0x63, 0x72, 0x79, 0x70, 0x74, 0x6f, 0x42, 0x0b, 0x43, 0x72, 0x79, 0x70,
	0x74, 0x6f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x1a, 0x67, 0x6f, 0x2d, 0x6d, 0x69,
	0x63, 0x72, 0x6f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2d, 0x73, 0x61, 0x6d, 0x70, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_api_crypto_proto_rawDescOnce sync.Once
	file_api_crypto_proto_rawDescData = file_api_crypto_proto_rawDesc
)

func file_api_crypto_proto_rawDescGZIP() []byte {
	file_api_crypto_proto_rawDescOnce.Do(func() {
		file_api_crypto_proto_rawDescData = protoimpl.X.CompressGZIP(file_api_crypto_proto_rawDescData)
	})
	return file_api_crypto_proto_rawDescData
}

var file_api_crypto_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_api_crypto_proto_goTypes = []interface{}{
	(*Crypto)(nil),      // 0: Crypto.Crypto
	(*CryptoId)(nil),    // 1: Crypto.CryptoId
	(*Upvotes)(nil),     // 2: Crypto.Upvotes
	(*empty.Empty)(nil), // 3: google.protobuf.Empty
}
var file_api_crypto_proto_depIdxs = []int32{
	1, // 0: Crypto.Cryptos.Get:input_type -> Crypto.CryptoId
	3, // 1: Crypto.Cryptos.GetAll:input_type -> google.protobuf.Empty
	0, // 2: Crypto.Cryptos.AddCrypto:input_type -> Crypto.Crypto
	1, // 3: Crypto.Cryptos.Upvote:input_type -> Crypto.CryptoId
	1, // 4: Crypto.Cryptos.Downvote:input_type -> Crypto.CryptoId
	1, // 5: Crypto.Cryptos.LiveUpVotes:input_type -> Crypto.CryptoId
	0, // 6: Crypto.Cryptos.Get:output_type -> Crypto.Crypto
	0, // 7: Crypto.Cryptos.GetAll:output_type -> Crypto.Crypto
	1, // 8: Crypto.Cryptos.AddCrypto:output_type -> Crypto.CryptoId
	3, // 9: Crypto.Cryptos.Upvote:output_type -> google.protobuf.Empty
	3, // 10: Crypto.Cryptos.Downvote:output_type -> google.protobuf.Empty
	2, // 11: Crypto.Cryptos.LiveUpVotes:output_type -> Crypto.Upvotes
	6, // [6:12] is the sub-list for method output_type
	0, // [0:6] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_api_crypto_proto_init() }
func file_api_crypto_proto_init() {
	if File_api_crypto_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_api_crypto_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Crypto); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_crypto_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CryptoId); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_api_crypto_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Upvotes); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_api_crypto_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_crypto_proto_goTypes,
		DependencyIndexes: file_api_crypto_proto_depIdxs,
		MessageInfos:      file_api_crypto_proto_msgTypes,
	}.Build()
	File_api_crypto_proto = out.File
	file_api_crypto_proto_rawDesc = nil
	file_api_crypto_proto_goTypes = nil
	file_api_crypto_proto_depIdxs = nil
}
