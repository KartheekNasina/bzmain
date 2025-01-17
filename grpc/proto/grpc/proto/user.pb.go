// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v4.23.4
// source: grpc/proto/user.proto

package proto

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

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id                          string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	IsOnline                    bool     `protobuf:"varint,2,opt,name=is_online,json=isOnline,proto3" json:"is_online,omitempty"`
	AllowNotifications          bool     `protobuf:"varint,3,opt,name=allow_notifications,json=allowNotifications,proto3" json:"allow_notifications,omitempty"`
	AllowLocation               bool     `protobuf:"varint,4,opt,name=allow_location,json=allowLocation,proto3" json:"allow_location,omitempty"`
	IsLegalAge                  bool     `protobuf:"varint,5,opt,name=is_legal_age,json=isLegalAge,proto3" json:"is_legal_age,omitempty"`
	PhoneNumberVerified         bool     `protobuf:"varint,6,opt,name=phone_number_verified,json=phoneNumberVerified,proto3" json:"phone_number_verified,omitempty"`
	PhoneNumberVerificationDate string   `protobuf:"bytes,7,opt,name=phone_number_verification_date,json=phoneNumberVerificationDate,proto3" json:"phone_number_verification_date,omitempty"` // Consider using Google's Timestamp type for better handling
	OnboardingStatus            string   `protobuf:"bytes,8,opt,name=onboarding_status,json=onboardingStatus,proto3" json:"onboarding_status,omitempty"`
	OtpVerifiedAt               string   `protobuf:"bytes,9,opt,name=otp_verified_at,json=otpVerifiedAt,proto3" json:"otp_verified_at,omitempty"`                 // Consider using Google's Timestamp type for better handling
	ProfileCompletedAt          string   `protobuf:"bytes,10,opt,name=profile_completed_at,json=profileCompletedAt,proto3" json:"profile_completed_at,omitempty"` // Consider using Google's Timestamp type for better handling
	CreatedAt                   string   `protobuf:"bytes,11,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`                              // Consider using Google's Timestamp type for better handling
	UpdatedAt                   string   `protobuf:"bytes,12,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`                              // Consider using Google's Timestamp type for better handling
	Dob                         string   `protobuf:"bytes,13,opt,name=dob,proto3" json:"dob,omitempty"`
	UserId                      string   `protobuf:"bytes,14,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Name                        string   `protobuf:"bytes,15,opt,name=name,proto3" json:"name,omitempty"`
	AboutMe                     string   `protobuf:"bytes,16,opt,name=about_me,json=aboutMe,proto3" json:"about_me,omitempty"`
	Gender                      string   `protobuf:"bytes,17,opt,name=gender,proto3" json:"gender,omitempty"`
	Email                       string   `protobuf:"bytes,18,opt,name=email,proto3" json:"email,omitempty"`
	ProfileUrl                  string   `protobuf:"bytes,19,opt,name=profile_url,json=profileUrl,proto3" json:"profile_url,omitempty"`
	Images                      []string `protobuf:"bytes,20,rep,name=images,proto3" json:"images,omitempty"` // Using repeated for arrays
	PhoneNumber                 string   `protobuf:"bytes,21,opt,name=phone_number,json=phoneNumber,proto3" json:"phone_number,omitempty"`
	PersonalInterests           []string `protobuf:"bytes,22,rep,name=personal_interests,json=personalInterests,proto3" json:"personal_interests,omitempty"` // Using repeated for arrays
	Provider                    string   `protobuf:"bytes,23,opt,name=provider,proto3" json:"provider,omitempty"`
	BrewInterests               []string `protobuf:"bytes,24,rep,name=brew_interests,json=brewInterests,proto3" json:"brew_interests,omitempty"` // Using repeated for arrays
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_user_proto_msgTypes[0]
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
	return file_grpc_proto_user_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *User) GetIsOnline() bool {
	if x != nil {
		return x.IsOnline
	}
	return false
}

func (x *User) GetAllowNotifications() bool {
	if x != nil {
		return x.AllowNotifications
	}
	return false
}

func (x *User) GetAllowLocation() bool {
	if x != nil {
		return x.AllowLocation
	}
	return false
}

func (x *User) GetIsLegalAge() bool {
	if x != nil {
		return x.IsLegalAge
	}
	return false
}

func (x *User) GetPhoneNumberVerified() bool {
	if x != nil {
		return x.PhoneNumberVerified
	}
	return false
}

func (x *User) GetPhoneNumberVerificationDate() string {
	if x != nil {
		return x.PhoneNumberVerificationDate
	}
	return ""
}

func (x *User) GetOnboardingStatus() string {
	if x != nil {
		return x.OnboardingStatus
	}
	return ""
}

func (x *User) GetOtpVerifiedAt() string {
	if x != nil {
		return x.OtpVerifiedAt
	}
	return ""
}

func (x *User) GetProfileCompletedAt() string {
	if x != nil {
		return x.ProfileCompletedAt
	}
	return ""
}

func (x *User) GetCreatedAt() string {
	if x != nil {
		return x.CreatedAt
	}
	return ""
}

func (x *User) GetUpdatedAt() string {
	if x != nil {
		return x.UpdatedAt
	}
	return ""
}

func (x *User) GetDob() string {
	if x != nil {
		return x.Dob
	}
	return ""
}

func (x *User) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetAboutMe() string {
	if x != nil {
		return x.AboutMe
	}
	return ""
}

func (x *User) GetGender() string {
	if x != nil {
		return x.Gender
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *User) GetProfileUrl() string {
	if x != nil {
		return x.ProfileUrl
	}
	return ""
}

func (x *User) GetImages() []string {
	if x != nil {
		return x.Images
	}
	return nil
}

func (x *User) GetPhoneNumber() string {
	if x != nil {
		return x.PhoneNumber
	}
	return ""
}

func (x *User) GetPersonalInterests() []string {
	if x != nil {
		return x.PersonalInterests
	}
	return nil
}

func (x *User) GetProvider() string {
	if x != nil {
		return x.Provider
	}
	return ""
}

func (x *User) GetBrewInterests() []string {
	if x != nil {
		return x.BrewInterests
	}
	return nil
}

type GetAllUsersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetAllUsersRequest) Reset() {
	*x = GetAllUsersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllUsersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllUsersRequest) ProtoMessage() {}

func (x *GetAllUsersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllUsersRequest.ProtoReflect.Descriptor instead.
func (*GetAllUsersRequest) Descriptor() ([]byte, []int) {
	return file_grpc_proto_user_proto_rawDescGZIP(), []int{1}
}

type GetAllUsersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users []*User `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
}

func (x *GetAllUsersResponse) Reset() {
	*x = GetAllUsersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_grpc_proto_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllUsersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllUsersResponse) ProtoMessage() {}

func (x *GetAllUsersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_grpc_proto_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllUsersResponse.ProtoReflect.Descriptor instead.
func (*GetAllUsersResponse) Descriptor() ([]byte, []int) {
	return file_grpc_proto_user_proto_rawDescGZIP(), []int{2}
}

func (x *GetAllUsersResponse) GetUsers() []*User {
	if x != nil {
		return x.Users
	}
	return nil
}

var File_grpc_proto_user_proto protoreflect.FileDescriptor

var file_grpc_proto_user_proto_rawDesc = []byte{
	0x0a, 0x15, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x73, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x75, 0x73, 0x65, 0x72, 0x22, 0xc1, 0x06,
	0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73, 0x5f, 0x6f, 0x6e, 0x6c,
	0x69, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69, 0x73, 0x4f, 0x6e, 0x6c,
	0x69, 0x6e, 0x65, 0x12, 0x2f, 0x0a, 0x13, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x6e, 0x6f, 0x74,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08,
	0x52, 0x12, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x4e, 0x6f, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x61, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x6c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0d, 0x61, 0x6c,
	0x6c, 0x6f, 0x77, 0x4c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x0a, 0x0c, 0x69,
	0x73, 0x5f, 0x6c, 0x65, 0x67, 0x61, 0x6c, 0x5f, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x0a, 0x69, 0x73, 0x4c, 0x65, 0x67, 0x61, 0x6c, 0x41, 0x67, 0x65, 0x12, 0x32, 0x0a,
	0x15, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x5f, 0x76, 0x65,
	0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x08, 0x52, 0x13, 0x70, 0x68,
	0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65,
	0x64, 0x12, 0x43, 0x0a, 0x1e, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64,
	0x61, 0x74, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x1b, 0x70, 0x68, 0x6f, 0x6e, 0x65,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x12, 0x2b, 0x0a, 0x11, 0x6f, 0x6e, 0x62, 0x6f, 0x61, 0x72,
	0x64, 0x69, 0x6e, 0x67, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x10, 0x6f, 0x6e, 0x62, 0x6f, 0x61, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x26, 0x0a, 0x0f, 0x6f, 0x74, 0x70, 0x5f, 0x76, 0x65, 0x72, 0x69, 0x66,
	0x69, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x6f, 0x74,
	0x70, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x65, 0x64, 0x41, 0x74, 0x12, 0x30, 0x0a, 0x14, 0x70,
	0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x63, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64,
	0x5f, 0x61, 0x74, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x12, 0x70, 0x72, 0x6f, 0x66, 0x69,
	0x6c, 0x65, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x1d, 0x0a, 0x0a,
	0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x64,
	0x6f, 0x62, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x64, 0x6f, 0x62, 0x12, 0x17, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x0f,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x61, 0x62,
	0x6f, 0x75, 0x74, 0x5f, 0x6d, 0x65, 0x18, 0x10, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x62,
	0x6f, 0x75, 0x74, 0x4d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x18,
	0x11, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x67, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x12, 0x14, 0x0a,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x12, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x12, 0x1f, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c, 0x65, 0x5f, 0x75,
	0x72, 0x6c, 0x18, 0x13, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x66, 0x69, 0x6c,
	0x65, 0x55, 0x72, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x18, 0x14,
	0x20, 0x03, 0x28, 0x09, 0x52, 0x06, 0x69, 0x6d, 0x61, 0x67, 0x65, 0x73, 0x12, 0x21, 0x0a, 0x0c,
	0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x15, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0b, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12,
	0x2d, 0x0a, 0x12, 0x70, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x5f, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x65, 0x73, 0x74, 0x73, 0x18, 0x16, 0x20, 0x03, 0x28, 0x09, 0x52, 0x11, 0x70, 0x65, 0x72,
	0x73, 0x6f, 0x6e, 0x61, 0x6c, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x73, 0x74, 0x73, 0x12, 0x1a,
	0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x17, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x25, 0x0a, 0x0e, 0x62, 0x72,
	0x65, 0x77, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x73, 0x74, 0x73, 0x18, 0x18, 0x20, 0x03,
	0x28, 0x09, 0x52, 0x0d, 0x62, 0x72, 0x65, 0x77, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x65, 0x73, 0x74,
	0x73, 0x22, 0x14, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x37, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x41, 0x6c,
	0x6c, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20,
	0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0a, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73,
	0x32, 0x51, 0x0a, 0x0b, 0x55, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x42, 0x0a, 0x0b, 0x67, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x18,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x76, 0x69, 0x76, 0x65, 0x6b, 0x62, 0x6e, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x62, 0x7a,
	0x2d, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2f, 0x62, 0x7a, 0x2d, 0x6d, 0x61, 0x69, 0x6e,
	0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_grpc_proto_user_proto_rawDescOnce sync.Once
	file_grpc_proto_user_proto_rawDescData = file_grpc_proto_user_proto_rawDesc
)

func file_grpc_proto_user_proto_rawDescGZIP() []byte {
	file_grpc_proto_user_proto_rawDescOnce.Do(func() {
		file_grpc_proto_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_grpc_proto_user_proto_rawDescData)
	})
	return file_grpc_proto_user_proto_rawDescData
}

var file_grpc_proto_user_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_grpc_proto_user_proto_goTypes = []interface{}{
	(*User)(nil),                // 0: user.User
	(*GetAllUsersRequest)(nil),  // 1: user.GetAllUsersRequest
	(*GetAllUsersResponse)(nil), // 2: user.GetAllUsersResponse
}
var file_grpc_proto_user_proto_depIdxs = []int32{
	0, // 0: user.GetAllUsersResponse.users:type_name -> user.User
	1, // 1: user.UserService.getAllUsers:input_type -> user.GetAllUsersRequest
	2, // 2: user.UserService.getAllUsers:output_type -> user.GetAllUsersResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_grpc_proto_user_proto_init() }
func file_grpc_proto_user_proto_init() {
	if File_grpc_proto_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_grpc_proto_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_grpc_proto_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllUsersRequest); i {
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
		file_grpc_proto_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllUsersResponse); i {
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
			RawDescriptor: file_grpc_proto_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_grpc_proto_user_proto_goTypes,
		DependencyIndexes: file_grpc_proto_user_proto_depIdxs,
		MessageInfos:      file_grpc_proto_user_proto_msgTypes,
	}.Build()
	File_grpc_proto_user_proto = out.File
	file_grpc_proto_user_proto_rawDesc = nil
	file_grpc_proto_user_proto_goTypes = nil
	file_grpc_proto_user_proto_depIdxs = nil
}
