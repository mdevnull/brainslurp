// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v5.27.0--rc1
// source: lib/proto/flow.proto

package flow

import (
	issue "github.com/devnull-twitch/brainslurp/lib/proto/issue"
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

type FlowRequirement struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	InCategory       issue.IssueCategory `protobuf:"varint,10,opt,name=in_category,json=inCategory,proto3,enum=brainslurp.IssueCategory" json:"in_category,omitempty"`
	RequiredTagIds   []uint64            `protobuf:"varint,22,rep,packed,name=required_tag_ids,json=requiredTagIds,proto3" json:"required_tag_ids,omitempty"`
	ProhibitedTagIds []uint64            `protobuf:"varint,23,rep,packed,name=prohibited_tag_ids,json=prohibitedTagIds,proto3" json:"prohibited_tag_ids,omitempty"`
}

func (x *FlowRequirement) Reset() {
	*x = FlowRequirement{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lib_proto_flow_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FlowRequirement) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FlowRequirement) ProtoMessage() {}

func (x *FlowRequirement) ProtoReflect() protoreflect.Message {
	mi := &file_lib_proto_flow_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FlowRequirement.ProtoReflect.Descriptor instead.
func (*FlowRequirement) Descriptor() ([]byte, []int) {
	return file_lib_proto_flow_proto_rawDescGZIP(), []int{0}
}

func (x *FlowRequirement) GetInCategory() issue.IssueCategory {
	if x != nil {
		return x.InCategory
	}
	return issue.IssueCategory(0)
}

func (x *FlowRequirement) GetRequiredTagIds() []uint64 {
	if x != nil {
		return x.RequiredTagIds
	}
	return nil
}

func (x *FlowRequirement) GetProhibitedTagIds() []uint64 {
	if x != nil {
		return x.ProhibitedTagIds
	}
	return nil
}

type FlowActions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Title              string   `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty"`
	RemoveAllAssignees bool     `protobuf:"varint,5,opt,name=remove_all_assignees,json=removeAllAssignees,proto3" json:"remove_all_assignees,omitempty"`
	AssignUser         []uint64 `protobuf:"varint,8,rep,packed,name=assign_user,json=assignUser,proto3" json:"assign_user,omitempty"`
	AddTagIds          []uint64 `protobuf:"varint,22,rep,packed,name=add_tag_ids,json=addTagIds,proto3" json:"add_tag_ids,omitempty"`
	RemoveTagIds       []uint64 `protobuf:"varint,23,rep,packed,name=remove_tag_ids,json=removeTagIds,proto3" json:"remove_tag_ids,omitempty"`
}

func (x *FlowActions) Reset() {
	*x = FlowActions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lib_proto_flow_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FlowActions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FlowActions) ProtoMessage() {}

func (x *FlowActions) ProtoReflect() protoreflect.Message {
	mi := &file_lib_proto_flow_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FlowActions.ProtoReflect.Descriptor instead.
func (*FlowActions) Descriptor() ([]byte, []int) {
	return file_lib_proto_flow_proto_rawDescGZIP(), []int{1}
}

func (x *FlowActions) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *FlowActions) GetRemoveAllAssignees() bool {
	if x != nil {
		return x.RemoveAllAssignees
	}
	return false
}

func (x *FlowActions) GetAssignUser() []uint64 {
	if x != nil {
		return x.AssignUser
	}
	return nil
}

func (x *FlowActions) GetAddTagIds() []uint64 {
	if x != nil {
		return x.AddTagIds
	}
	return nil
}

func (x *FlowActions) GetRemoveTagIds() []uint64 {
	if x != nil {
		return x.RemoveTagIds
	}
	return nil
}

type Flow struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Number       uint64             `protobuf:"varint,1,opt,name=number,proto3" json:"number,omitempty"`
	Title        string             `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	CreatedAt    int64              `protobuf:"varint,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	Requirements []*FlowRequirement `protobuf:"bytes,10,rep,name=requirements,proto3" json:"requirements,omitempty"`
	Actions      []*FlowActions     `protobuf:"bytes,15,rep,name=actions,proto3" json:"actions,omitempty"`
}

func (x *Flow) Reset() {
	*x = Flow{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lib_proto_flow_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Flow) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Flow) ProtoMessage() {}

func (x *Flow) ProtoReflect() protoreflect.Message {
	mi := &file_lib_proto_flow_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Flow.ProtoReflect.Descriptor instead.
func (*Flow) Descriptor() ([]byte, []int) {
	return file_lib_proto_flow_proto_rawDescGZIP(), []int{2}
}

func (x *Flow) GetNumber() uint64 {
	if x != nil {
		return x.Number
	}
	return 0
}

func (x *Flow) GetTitle() string {
	if x != nil {
		return x.Title
	}
	return ""
}

func (x *Flow) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *Flow) GetRequirements() []*FlowRequirement {
	if x != nil {
		return x.Requirements
	}
	return nil
}

func (x *Flow) GetActions() []*FlowActions {
	if x != nil {
		return x.Actions
	}
	return nil
}

var File_lib_proto_flow_proto protoreflect.FileDescriptor

var file_lib_proto_flow_proto_rawDesc = []byte{
	0x0a, 0x14, 0x6c, 0x69, 0x62, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x66, 0x6c, 0x6f, 0x77,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x62, 0x72, 0x61, 0x69, 0x6e, 0x73, 0x6c, 0x75,
	0x72, 0x70, 0x1a, 0x15, 0x6c, 0x69, 0x62, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x69, 0x73,
	0x73, 0x75, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb1, 0x01, 0x0a, 0x0f, 0x46, 0x6c,
	0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x3a, 0x0a,
	0x0b, 0x69, 0x6e, 0x5f, 0x63, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x18, 0x0a, 0x20, 0x01,
	0x28, 0x0e, 0x32, 0x19, 0x2e, 0x62, 0x72, 0x61, 0x69, 0x6e, 0x73, 0x6c, 0x75, 0x72, 0x70, 0x2e,
	0x49, 0x73, 0x73, 0x75, 0x65, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x52, 0x0a, 0x69,
	0x6e, 0x43, 0x61, 0x74, 0x65, 0x67, 0x6f, 0x72, 0x79, 0x12, 0x28, 0x0a, 0x10, 0x72, 0x65, 0x71,
	0x75, 0x69, 0x72, 0x65, 0x64, 0x5f, 0x74, 0x61, 0x67, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x16, 0x20,
	0x03, 0x28, 0x04, 0x52, 0x0e, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x54, 0x61, 0x67,
	0x49, 0x64, 0x73, 0x12, 0x2c, 0x0a, 0x12, 0x70, 0x72, 0x6f, 0x68, 0x69, 0x62, 0x69, 0x74, 0x65,
	0x64, 0x5f, 0x74, 0x61, 0x67, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x17, 0x20, 0x03, 0x28, 0x04, 0x52,
	0x10, 0x70, 0x72, 0x6f, 0x68, 0x69, 0x62, 0x69, 0x74, 0x65, 0x64, 0x54, 0x61, 0x67, 0x49, 0x64,
	0x73, 0x4a, 0x04, 0x08, 0x14, 0x10, 0x15, 0x4a, 0x04, 0x08, 0x02, 0x10, 0x03, 0x22, 0xc8, 0x01,
	0x0a, 0x0b, 0x46, 0x6c, 0x6f, 0x77, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x14, 0x0a,
	0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69,
	0x74, 0x6c, 0x65, 0x12, 0x30, 0x0a, 0x14, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x5f, 0x61, 0x6c,
	0x6c, 0x5f, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x65, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x12, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x41, 0x6c, 0x6c, 0x41, 0x73, 0x73, 0x69,
	0x67, 0x6e, 0x65, 0x65, 0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x73, 0x73, 0x69, 0x67, 0x6e, 0x5f,
	0x75, 0x73, 0x65, 0x72, 0x18, 0x08, 0x20, 0x03, 0x28, 0x04, 0x52, 0x0a, 0x61, 0x73, 0x73, 0x69,
	0x67, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x12, 0x1e, 0x0a, 0x0b, 0x61, 0x64, 0x64, 0x5f, 0x74, 0x61,
	0x67, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x16, 0x20, 0x03, 0x28, 0x04, 0x52, 0x09, 0x61, 0x64, 0x64,
	0x54, 0x61, 0x67, 0x49, 0x64, 0x73, 0x12, 0x24, 0x0a, 0x0e, 0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65,
	0x5f, 0x74, 0x61, 0x67, 0x5f, 0x69, 0x64, 0x73, 0x18, 0x17, 0x20, 0x03, 0x28, 0x04, 0x52, 0x0c,
	0x72, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x54, 0x61, 0x67, 0x49, 0x64, 0x73, 0x4a, 0x04, 0x08, 0x14,
	0x10, 0x15, 0x4a, 0x04, 0x08, 0x15, 0x10, 0x16, 0x22, 0xc7, 0x01, 0x0a, 0x04, 0x46, 0x6c, 0x6f,
	0x77, 0x12, 0x16, 0x0a, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x06, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x69, 0x74,
	0x6c, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x12,
	0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x3f,
	0x0a, 0x0c, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x0a,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x62, 0x72, 0x61, 0x69, 0x6e, 0x73, 0x6c, 0x75, 0x72,
	0x70, 0x2e, 0x46, 0x6c, 0x6f, 0x77, 0x52, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e,
	0x74, 0x52, 0x0c, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x12,
	0x31, 0x0a, 0x07, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x0f, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x17, 0x2e, 0x62, 0x72, 0x61, 0x69, 0x6e, 0x73, 0x6c, 0x75, 0x72, 0x70, 0x2e, 0x46, 0x6c,
	0x6f, 0x77, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x07, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x42, 0x35, 0x5a, 0x33, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x64, 0x65, 0x76, 0x6e, 0x75, 0x6c, 0x6c, 0x2d, 0x74, 0x77, 0x69, 0x74, 0x63, 0x68, 0x2f,
	0x62, 0x72, 0x61, 0x69, 0x6e, 0x73, 0x6c, 0x75, 0x72, 0x70, 0x2f, 0x6c, 0x69, 0x62, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x66, 0x6c, 0x6f, 0x77, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_lib_proto_flow_proto_rawDescOnce sync.Once
	file_lib_proto_flow_proto_rawDescData = file_lib_proto_flow_proto_rawDesc
)

func file_lib_proto_flow_proto_rawDescGZIP() []byte {
	file_lib_proto_flow_proto_rawDescOnce.Do(func() {
		file_lib_proto_flow_proto_rawDescData = protoimpl.X.CompressGZIP(file_lib_proto_flow_proto_rawDescData)
	})
	return file_lib_proto_flow_proto_rawDescData
}

var file_lib_proto_flow_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_lib_proto_flow_proto_goTypes = []interface{}{
	(*FlowRequirement)(nil),  // 0: brainslurp.FlowRequirement
	(*FlowActions)(nil),      // 1: brainslurp.FlowActions
	(*Flow)(nil),             // 2: brainslurp.Flow
	(issue.IssueCategory)(0), // 3: brainslurp.IssueCategory
}
var file_lib_proto_flow_proto_depIdxs = []int32{
	3, // 0: brainslurp.FlowRequirement.in_category:type_name -> brainslurp.IssueCategory
	0, // 1: brainslurp.Flow.requirements:type_name -> brainslurp.FlowRequirement
	1, // 2: brainslurp.Flow.actions:type_name -> brainslurp.FlowActions
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_lib_proto_flow_proto_init() }
func file_lib_proto_flow_proto_init() {
	if File_lib_proto_flow_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_lib_proto_flow_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FlowRequirement); i {
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
		file_lib_proto_flow_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FlowActions); i {
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
		file_lib_proto_flow_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Flow); i {
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
			RawDescriptor: file_lib_proto_flow_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_lib_proto_flow_proto_goTypes,
		DependencyIndexes: file_lib_proto_flow_proto_depIdxs,
		MessageInfos:      file_lib_proto_flow_proto_msgTypes,
	}.Build()
	File_lib_proto_flow_proto = out.File
	file_lib_proto_flow_proto_rawDesc = nil
	file_lib_proto_flow_proto_goTypes = nil
	file_lib_proto_flow_proto_depIdxs = nil
}
