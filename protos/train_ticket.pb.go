// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v4.24.4
// source: train_ticket.proto

package protos

import (
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

// Message for user details
type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FirstName string `protobuf:"bytes,1,opt,name=first_name,json=firstName,proto3" json:"first_name,omitempty"`
	LastName  string `protobuf:"bytes,2,opt,name=last_name,json=lastName,proto3" json:"last_name,omitempty"`
	Email     string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_train_ticket_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_train_ticket_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_train_ticket_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetFirstName() string {
	if x != nil {
		return x.FirstName
	}
	return ""
}

func (x *User) GetLastName() string {
	if x != nil {
		return x.LastName
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

// Message for ticket purchase request
type PurchaseRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	From      string  `protobuf:"bytes,1,opt,name=from,proto3" json:"from,omitempty"`
	To        string  `protobuf:"bytes,2,opt,name=to,proto3" json:"to,omitempty"`
	User      *User   `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
	PricePaid float64 `protobuf:"fixed64,4,opt,name=price_paid,json=pricePaid,proto3" json:"price_paid,omitempty"`
	Section   string  `protobuf:"bytes,5,opt,name=section,proto3" json:"section,omitempty"` // Section A or B //can modify as Enum
}

func (x *PurchaseRequest) Reset() {
	*x = PurchaseRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_train_ticket_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PurchaseRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PurchaseRequest) ProtoMessage() {}

func (x *PurchaseRequest) ProtoReflect() protoreflect.Message {
	mi := &file_train_ticket_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PurchaseRequest.ProtoReflect.Descriptor instead.
func (*PurchaseRequest) Descriptor() ([]byte, []int) {
	return file_train_ticket_proto_rawDescGZIP(), []int{1}
}

func (x *PurchaseRequest) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *PurchaseRequest) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *PurchaseRequest) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *PurchaseRequest) GetPricePaid() float64 {
	if x != nil {
		return x.PricePaid
	}
	return 0
}

func (x *PurchaseRequest) GetSection() string {
	if x != nil {
		return x.Section
	}
	return ""
}

// Message for ticket purchase response
type PurchaseResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReceiptId  string  `protobuf:"bytes,1,opt,name=receipt_id,json=receiptId,proto3" json:"receipt_id,omitempty"`
	From       string  `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`
	To         string  `protobuf:"bytes,3,opt,name=to,proto3" json:"to,omitempty"`
	User       *User   `protobuf:"bytes,4,opt,name=user,proto3" json:"user,omitempty"`
	PricePaid  float64 `protobuf:"fixed64,5,opt,name=price_paid,json=pricePaid,proto3" json:"price_paid,omitempty"`
	Section    string  `protobuf:"bytes,6,opt,name=section,proto3" json:"section,omitempty"`
	SeatNumber string  `protobuf:"bytes,7,opt,name=seat_number,json=seatNumber,proto3" json:"seat_number,omitempty"`
}

func (x *PurchaseResponse) Reset() {
	*x = PurchaseResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_train_ticket_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PurchaseResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PurchaseResponse) ProtoMessage() {}

func (x *PurchaseResponse) ProtoReflect() protoreflect.Message {
	mi := &file_train_ticket_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PurchaseResponse.ProtoReflect.Descriptor instead.
func (*PurchaseResponse) Descriptor() ([]byte, []int) {
	return file_train_ticket_proto_rawDescGZIP(), []int{2}
}

func (x *PurchaseResponse) GetReceiptId() string {
	if x != nil {
		return x.ReceiptId
	}
	return ""
}

func (x *PurchaseResponse) GetFrom() string {
	if x != nil {
		return x.From
	}
	return ""
}

func (x *PurchaseResponse) GetTo() string {
	if x != nil {
		return x.To
	}
	return ""
}

func (x *PurchaseResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *PurchaseResponse) GetPricePaid() float64 {
	if x != nil {
		return x.PricePaid
	}
	return 0
}

func (x *PurchaseResponse) GetSection() string {
	if x != nil {
		return x.Section
	}
	return ""
}

func (x *PurchaseResponse) GetSeatNumber() string {
	if x != nil {
		return x.SeatNumber
	}
	return ""
}

// Message for seat modification
type ModifySeatRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ReceiptId     string `protobuf:"bytes,1,opt,name=receipt_id,json=receiptId,proto3" json:"receipt_id,omitempty"`
	NewSeatNumber string `protobuf:"bytes,2,opt,name=new_seat_number,json=newSeatNumber,proto3" json:"new_seat_number,omitempty"`
}

func (x *ModifySeatRequest) Reset() {
	*x = ModifySeatRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_train_ticket_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModifySeatRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModifySeatRequest) ProtoMessage() {}

func (x *ModifySeatRequest) ProtoReflect() protoreflect.Message {
	mi := &file_train_ticket_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModifySeatRequest.ProtoReflect.Descriptor instead.
func (*ModifySeatRequest) Descriptor() ([]byte, []int) {
	return file_train_ticket_proto_rawDescGZIP(), []int{3}
}

func (x *ModifySeatRequest) GetReceiptId() string {
	if x != nil {
		return x.ReceiptId
	}
	return ""
}

func (x *ModifySeatRequest) GetNewSeatNumber() string {
	if x != nil {
		return x.NewSeatNumber
	}
	return ""
}

// Message for seat modification response
type ModifySeatResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *ModifySeatResponse) Reset() {
	*x = ModifySeatResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_train_ticket_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModifySeatResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModifySeatResponse) ProtoMessage() {}

func (x *ModifySeatResponse) ProtoReflect() protoreflect.Message {
	mi := &file_train_ticket_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModifySeatResponse.ProtoReflect.Descriptor instead.
func (*ModifySeatResponse) Descriptor() ([]byte, []int) {
	return file_train_ticket_proto_rawDescGZIP(), []int{4}
}

func (x *ModifySeatResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

// Message for getting user details by section
type UsersInSectionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Section string `protobuf:"bytes,1,opt,name=section,proto3" json:"section,omitempty"`
}

func (x *UsersInSectionRequest) Reset() {
	*x = UsersInSectionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_train_ticket_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UsersInSectionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UsersInSectionRequest) ProtoMessage() {}

func (x *UsersInSectionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_train_ticket_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UsersInSectionRequest.ProtoReflect.Descriptor instead.
func (*UsersInSectionRequest) Descriptor() ([]byte, []int) {
	return file_train_ticket_proto_rawDescGZIP(), []int{5}
}

func (x *UsersInSectionRequest) GetSection() string {
	if x != nil {
		return x.Section
	}
	return ""
}

// Message for user details response
type UserDetails struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId     string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	SeatNumber string `protobuf:"bytes,2,opt,name=seat_number,json=seatNumber,proto3" json:"seat_number,omitempty"`
	User       *User  `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *UserDetails) Reset() {
	*x = UserDetails{}
	if protoimpl.UnsafeEnabled {
		mi := &file_train_ticket_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserDetails) ProtoMessage() {}

func (x *UserDetails) ProtoReflect() protoreflect.Message {
	mi := &file_train_ticket_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserDetails.ProtoReflect.Descriptor instead.
func (*UserDetails) Descriptor() ([]byte, []int) {
	return file_train_ticket_proto_rawDescGZIP(), []int{6}
}

func (x *UserDetails) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UserDetails) GetSeatNumber() string {
	if x != nil {
		return x.SeatNumber
	}
	return ""
}

func (x *UserDetails) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

// Response for getting users by section
type UsersInSectionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users []*UserDetails `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
}

func (x *UsersInSectionResponse) Reset() {
	*x = UsersInSectionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_train_ticket_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UsersInSectionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UsersInSectionResponse) ProtoMessage() {}

func (x *UsersInSectionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_train_ticket_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UsersInSectionResponse.ProtoReflect.Descriptor instead.
func (*UsersInSectionResponse) Descriptor() ([]byte, []int) {
	return file_train_ticket_proto_rawDescGZIP(), []int{7}
}

func (x *UsersInSectionResponse) GetUsers() []*UserDetails {
	if x != nil {
		return x.Users
	}
	return nil
}

// Message for removing a user
type RemoveUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
}

func (x *RemoveUserRequest) Reset() {
	*x = RemoveUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_train_ticket_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveUserRequest) ProtoMessage() {}

func (x *RemoveUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_train_ticket_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveUserRequest.ProtoReflect.Descriptor instead.
func (*RemoveUserRequest) Descriptor() ([]byte, []int) {
	return file_train_ticket_proto_rawDescGZIP(), []int{8}
}

func (x *RemoveUserRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

// Response for removing a user
type RemoveUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
}

func (x *RemoveUserResponse) Reset() {
	*x = RemoveUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_train_ticket_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveUserResponse) ProtoMessage() {}

func (x *RemoveUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_train_ticket_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveUserResponse.ProtoReflect.Descriptor instead.
func (*RemoveUserResponse) Descriptor() ([]byte, []int) {
	return file_train_ticket_proto_rawDescGZIP(), []int{9}
}

func (x *RemoveUserResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

var File_train_ticket_proto protoreflect.FileDescriptor

var file_train_ticket_proto_rawDesc = []byte{
	0x0a, 0x12, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x5f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x74, 0x72, 0x61,
	0x69, 0x6e, 0x5f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x22, 0x58, 0x0a, 0x04, 0x55, 0x73, 0x65,
	0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x66, 0x69, 0x72, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x66, 0x69, 0x72, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x61, 0x73, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x22, 0x9d, 0x01, 0x0a, 0x0f, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74,
	0x6f, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x2d, 0x0a, 0x04, 0x75,
	0x73, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x5f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x5f, 0x70, 0x61, 0x69, 0x64, 0x18, 0x04, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x50, 0x61, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x22, 0xde, 0x01, 0x0a, 0x10, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x63, 0x65,
	0x69, 0x70, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72, 0x65,
	0x63, 0x65, 0x69, 0x70, 0x74, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x66, 0x72, 0x6f, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x74,
	0x6f, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x74, 0x6f, 0x12, 0x2d, 0x0a, 0x04, 0x75,
	0x73, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x5f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x5f, 0x70, 0x61, 0x69, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x01, 0x52, 0x09,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x50, 0x61, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x65, 0x61, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x62,
	0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x73, 0x65, 0x61, 0x74, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x22, 0x5a, 0x0a, 0x11, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x53, 0x65,
	0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x72, 0x65, 0x63,
	0x65, 0x69, 0x70, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x72,
	0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x0f, 0x6e, 0x65, 0x77, 0x5f,
	0x73, 0x65, 0x61, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x6e, 0x65, 0x77, 0x53, 0x65, 0x61, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x22, 0x2e, 0x0a, 0x12, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x53, 0x65, 0x61, 0x74, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x22, 0x31, 0x0a, 0x15, 0x55, 0x73, 0x65, 0x72, 0x73, 0x49, 0x6e, 0x53, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x65, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x65, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x22, 0x76, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69,
	0x6c, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x73,
	0x65, 0x61, 0x74, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0a, 0x73, 0x65, 0x61, 0x74, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x2d, 0x0a, 0x04,
	0x75, 0x73, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x5f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74,
	0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0x50, 0x0a, 0x16, 0x55,
	0x73, 0x65, 0x72, 0x73, 0x49, 0x6e, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x74, 0x72,
	0x61, 0x69, 0x6e, 0x5f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x22, 0x2c, 0x0a,
	0x11, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x22, 0x2e, 0x0a, 0x12, 0x52,
	0x65, 0x6d, 0x6f, 0x76, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x32, 0xfa, 0x03, 0x0a, 0x12,
	0x54, 0x72, 0x61, 0x69, 0x6e, 0x54, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x12, 0x5d, 0x0a, 0x0e, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x54, 0x69,
	0x63, 0x6b, 0x65, 0x74, 0x12, 0x24, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x74, 0x72,
	0x61, 0x69, 0x6e, 0x5f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x50, 0x75, 0x72, 0x63, 0x68,
	0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x5f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74,
	0x2e, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x59, 0x0a, 0x0a, 0x47, 0x65, 0x74, 0x52, 0x65, 0x63, 0x65, 0x69, 0x70, 0x74, 0x12,
	0x24, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x5f, 0x74,
	0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x50, 0x75, 0x72, 0x63, 0x68, 0x61, 0x73, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x74,
	0x72, 0x61, 0x69, 0x6e, 0x5f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x50, 0x75, 0x72, 0x63,
	0x68, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x6c, 0x0a, 0x11,
	0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x49, 0x6e, 0x53, 0x65, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x12, 0x2a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e,
	0x5f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x73, 0x49, 0x6e, 0x53,
	0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x2b, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x5f, 0x74, 0x69, 0x63,
	0x6b, 0x65, 0x74, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x73, 0x49, 0x6e, 0x53, 0x65, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5d, 0x0a, 0x0a, 0x52, 0x65,
	0x6d, 0x6f, 0x76, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x26, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x73, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x5f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x52,
	0x65, 0x6d, 0x6f, 0x76, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x27, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x5f,
	0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x55, 0x73, 0x65,
	0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5d, 0x0a, 0x0a, 0x4d, 0x6f, 0x64,
	0x69, 0x66, 0x79, 0x53, 0x65, 0x61, 0x74, 0x12, 0x26, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x5f, 0x74, 0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x4d, 0x6f,
	0x64, 0x69, 0x66, 0x79, 0x53, 0x65, 0x61, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x27, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x74, 0x72, 0x61, 0x69, 0x6e, 0x5f, 0x74,
	0x69, 0x63, 0x6b, 0x65, 0x74, 0x2e, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x53, 0x65, 0x61, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x0b, 0x5a, 0x09, 0x2e, 0x2e, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_train_ticket_proto_rawDescOnce sync.Once
	file_train_ticket_proto_rawDescData = file_train_ticket_proto_rawDesc
)

func file_train_ticket_proto_rawDescGZIP() []byte {
	file_train_ticket_proto_rawDescOnce.Do(func() {
		file_train_ticket_proto_rawDescData = protoimpl.X.CompressGZIP(file_train_ticket_proto_rawDescData)
	})
	return file_train_ticket_proto_rawDescData
}

var file_train_ticket_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_train_ticket_proto_goTypes = []any{
	(*User)(nil),                   // 0: protos.train_ticket.User
	(*PurchaseRequest)(nil),        // 1: protos.train_ticket.PurchaseRequest
	(*PurchaseResponse)(nil),       // 2: protos.train_ticket.PurchaseResponse
	(*ModifySeatRequest)(nil),      // 3: protos.train_ticket.ModifySeatRequest
	(*ModifySeatResponse)(nil),     // 4: protos.train_ticket.ModifySeatResponse
	(*UsersInSectionRequest)(nil),  // 5: protos.train_ticket.UsersInSectionRequest
	(*UserDetails)(nil),            // 6: protos.train_ticket.UserDetails
	(*UsersInSectionResponse)(nil), // 7: protos.train_ticket.UsersInSectionResponse
	(*RemoveUserRequest)(nil),      // 8: protos.train_ticket.RemoveUserRequest
	(*RemoveUserResponse)(nil),     // 9: protos.train_ticket.RemoveUserResponse
}
var file_train_ticket_proto_depIdxs = []int32{
	0, // 0: protos.train_ticket.PurchaseRequest.user:type_name -> protos.train_ticket.User
	0, // 1: protos.train_ticket.PurchaseResponse.user:type_name -> protos.train_ticket.User
	0, // 2: protos.train_ticket.UserDetails.user:type_name -> protos.train_ticket.User
	6, // 3: protos.train_ticket.UsersInSectionResponse.users:type_name -> protos.train_ticket.UserDetails
	1, // 4: protos.train_ticket.TrainTicketService.PurchaseTicket:input_type -> protos.train_ticket.PurchaseRequest
	1, // 5: protos.train_ticket.TrainTicketService.GetReceipt:input_type -> protos.train_ticket.PurchaseRequest
	5, // 6: protos.train_ticket.TrainTicketService.GetUsersInSection:input_type -> protos.train_ticket.UsersInSectionRequest
	8, // 7: protos.train_ticket.TrainTicketService.RemoveUser:input_type -> protos.train_ticket.RemoveUserRequest
	3, // 8: protos.train_ticket.TrainTicketService.ModifySeat:input_type -> protos.train_ticket.ModifySeatRequest
	2, // 9: protos.train_ticket.TrainTicketService.PurchaseTicket:output_type -> protos.train_ticket.PurchaseResponse
	2, // 10: protos.train_ticket.TrainTicketService.GetReceipt:output_type -> protos.train_ticket.PurchaseResponse
	7, // 11: protos.train_ticket.TrainTicketService.GetUsersInSection:output_type -> protos.train_ticket.UsersInSectionResponse
	9, // 12: protos.train_ticket.TrainTicketService.RemoveUser:output_type -> protos.train_ticket.RemoveUserResponse
	4, // 13: protos.train_ticket.TrainTicketService.ModifySeat:output_type -> protos.train_ticket.ModifySeatResponse
	9, // [9:14] is the sub-list for method output_type
	4, // [4:9] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_train_ticket_proto_init() }
func file_train_ticket_proto_init() {
	if File_train_ticket_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_train_ticket_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*User); i {
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
		file_train_ticket_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*PurchaseRequest); i {
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
		file_train_ticket_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*PurchaseResponse); i {
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
		file_train_ticket_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*ModifySeatRequest); i {
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
		file_train_ticket_proto_msgTypes[4].Exporter = func(v any, i int) any {
			switch v := v.(*ModifySeatResponse); i {
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
		file_train_ticket_proto_msgTypes[5].Exporter = func(v any, i int) any {
			switch v := v.(*UsersInSectionRequest); i {
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
		file_train_ticket_proto_msgTypes[6].Exporter = func(v any, i int) any {
			switch v := v.(*UserDetails); i {
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
		file_train_ticket_proto_msgTypes[7].Exporter = func(v any, i int) any {
			switch v := v.(*UsersInSectionResponse); i {
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
		file_train_ticket_proto_msgTypes[8].Exporter = func(v any, i int) any {
			switch v := v.(*RemoveUserRequest); i {
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
		file_train_ticket_proto_msgTypes[9].Exporter = func(v any, i int) any {
			switch v := v.(*RemoveUserResponse); i {
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
			RawDescriptor: file_train_ticket_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_train_ticket_proto_goTypes,
		DependencyIndexes: file_train_ticket_proto_depIdxs,
		MessageInfos:      file_train_ticket_proto_msgTypes,
	}.Build()
	File_train_ticket_proto = out.File
	file_train_ticket_proto_rawDesc = nil
	file_train_ticket_proto_goTypes = nil
	file_train_ticket_proto_depIdxs = nil
}
