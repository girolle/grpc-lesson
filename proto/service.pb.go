// proto/service.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.5
// 	protoc        v5.29.3
// source: proto/service.proto

// Определяем пакет для наших сообщений и сервисов.

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Сообщение, описывающее партнёра.
type Partner struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int32                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`      // Уникальный идентификатор партнёра
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`   // Имя партнёра
	Email         string                 `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"` // Email партнёра
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Partner) Reset() {
	*x = Partner{}
	mi := &file_proto_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Partner) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Partner) ProtoMessage() {}

func (x *Partner) ProtoReflect() protoreflect.Message {
	mi := &file_proto_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Partner.ProtoReflect.Descriptor instead.
func (*Partner) Descriptor() ([]byte, []int) {
	return file_proto_service_proto_rawDescGZIP(), []int{0}
}

func (x *Partner) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Partner) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Partner) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

// Запрос для создания партнёра.
type CreatePartnerRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`   // Имя нового партнёра
	Email         string                 `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"` // Email нового партнёра
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreatePartnerRequest) Reset() {
	*x = CreatePartnerRequest{}
	mi := &file_proto_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreatePartnerRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePartnerRequest) ProtoMessage() {}

func (x *CreatePartnerRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePartnerRequest.ProtoReflect.Descriptor instead.
func (*CreatePartnerRequest) Descriptor() ([]byte, []int) {
	return file_proto_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreatePartnerRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreatePartnerRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

// Ответ на запрос создания партнёра.
type CreatePartnerResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Partner       *Partner               `protobuf:"bytes,1,opt,name=partner,proto3" json:"partner,omitempty"` // Информация о созданном партнёре
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *CreatePartnerResponse) Reset() {
	*x = CreatePartnerResponse{}
	mi := &file_proto_service_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *CreatePartnerResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePartnerResponse) ProtoMessage() {}

func (x *CreatePartnerResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_service_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePartnerResponse.ProtoReflect.Descriptor instead.
func (*CreatePartnerResponse) Descriptor() ([]byte, []int) {
	return file_proto_service_proto_rawDescGZIP(), []int{2}
}

func (x *CreatePartnerResponse) GetPartner() *Partner {
	if x != nil {
		return x.Partner
	}
	return nil
}

// Запрос для получения статистики по партнёру.
type GetPartnerStatsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	PartnerId     int32                  `protobuf:"varint,1,opt,name=partner_id,json=partnerId,proto3" json:"partner_id,omitempty"` // ID партнёра, статистику которого запрашивают
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetPartnerStatsRequest) Reset() {
	*x = GetPartnerStatsRequest{}
	mi := &file_proto_service_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetPartnerStatsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetPartnerStatsRequest) ProtoMessage() {}

func (x *GetPartnerStatsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_service_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetPartnerStatsRequest.ProtoReflect.Descriptor instead.
func (*GetPartnerStatsRequest) Descriptor() ([]byte, []int) {
	return file_proto_service_proto_rawDescGZIP(), []int{3}
}

func (x *GetPartnerStatsRequest) GetPartnerId() int32 {
	if x != nil {
		return x.PartnerId
	}
	return 0
}

// Сообщение со статистикой партнёра.
type PartnerStats struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	PartnerId     int32                  `protobuf:"varint,1,opt,name=partner_id,json=partnerId,proto3" json:"partner_id,omitempty"` // ID партнёра
	Clicks        int32                  `protobuf:"varint,2,opt,name=clicks,proto3" json:"clicks,omitempty"`                        // Количество кликов по партнерским ссылкам
	Conversions   int32                  `protobuf:"varint,3,opt,name=conversions,proto3" json:"conversions,omitempty"`              // Количество конверсий (например, покупок)
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *PartnerStats) Reset() {
	*x = PartnerStats{}
	mi := &file_proto_service_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PartnerStats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PartnerStats) ProtoMessage() {}

func (x *PartnerStats) ProtoReflect() protoreflect.Message {
	mi := &file_proto_service_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PartnerStats.ProtoReflect.Descriptor instead.
func (*PartnerStats) Descriptor() ([]byte, []int) {
	return file_proto_service_proto_rawDescGZIP(), []int{4}
}

func (x *PartnerStats) GetPartnerId() int32 {
	if x != nil {
		return x.PartnerId
	}
	return 0
}

func (x *PartnerStats) GetClicks() int32 {
	if x != nil {
		return x.Clicks
	}
	return 0
}

func (x *PartnerStats) GetConversions() int32 {
	if x != nil {
		return x.Conversions
	}
	return 0
}

// Запрос на сокращение ссылки.
type ShortenLinkRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	OriginalUrl   string                 `protobuf:"bytes,1,opt,name=original_url,json=originalUrl,proto3" json:"original_url,omitempty"` // Исходная длинная ссылка
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ShortenLinkRequest) Reset() {
	*x = ShortenLinkRequest{}
	mi := &file_proto_service_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ShortenLinkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShortenLinkRequest) ProtoMessage() {}

func (x *ShortenLinkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_service_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShortenLinkRequest.ProtoReflect.Descriptor instead.
func (*ShortenLinkRequest) Descriptor() ([]byte, []int) {
	return file_proto_service_proto_rawDescGZIP(), []int{5}
}

func (x *ShortenLinkRequest) GetOriginalUrl() string {
	if x != nil {
		return x.OriginalUrl
	}
	return ""
}

// Ответ, содержащий сокращённую ссылку.
type ShortenLinkResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ShortUrl      string                 `protobuf:"bytes,1,opt,name=short_url,json=shortUrl,proto3" json:"short_url,omitempty"` // Сокращённая ссылка
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ShortenLinkResponse) Reset() {
	*x = ShortenLinkResponse{}
	mi := &file_proto_service_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ShortenLinkResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ShortenLinkResponse) ProtoMessage() {}

func (x *ShortenLinkResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_service_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ShortenLinkResponse.ProtoReflect.Descriptor instead.
func (*ShortenLinkResponse) Descriptor() ([]byte, []int) {
	return file_proto_service_proto_rawDescGZIP(), []int{6}
}

func (x *ShortenLinkResponse) GetShortUrl() string {
	if x != nil {
		return x.ShortUrl
	}
	return ""
}

// Запрос на получение статистики по сокращённой ссылке.
type GetLinkStatsRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ShortUrl      string                 `protobuf:"bytes,1,opt,name=short_url,json=shortUrl,proto3" json:"short_url,omitempty"` // Сокращённая ссылка, статистику по которой запрашивают
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetLinkStatsRequest) Reset() {
	*x = GetLinkStatsRequest{}
	mi := &file_proto_service_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetLinkStatsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLinkStatsRequest) ProtoMessage() {}

func (x *GetLinkStatsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_service_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLinkStatsRequest.ProtoReflect.Descriptor instead.
func (*GetLinkStatsRequest) Descriptor() ([]byte, []int) {
	return file_proto_service_proto_rawDescGZIP(), []int{7}
}

func (x *GetLinkStatsRequest) GetShortUrl() string {
	if x != nil {
		return x.ShortUrl
	}
	return ""
}

// Сообщение со статистикой сокращённой ссылки.
type LinkStats struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ShortUrl      string                 `protobuf:"bytes,1,opt,name=short_url,json=shortUrl,proto3" json:"short_url,omitempty"` // Сокращённая ссылка
	Clicks        int32                  `protobuf:"varint,2,opt,name=clicks,proto3" json:"clicks,omitempty"`                    // Количество кликов по ней
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *LinkStats) Reset() {
	*x = LinkStats{}
	mi := &file_proto_service_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LinkStats) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LinkStats) ProtoMessage() {}

func (x *LinkStats) ProtoReflect() protoreflect.Message {
	mi := &file_proto_service_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LinkStats.ProtoReflect.Descriptor instead.
func (*LinkStats) Descriptor() ([]byte, []int) {
	return file_proto_service_proto_rawDescGZIP(), []int{8}
}

func (x *LinkStats) GetShortUrl() string {
	if x != nil {
		return x.ShortUrl
	}
	return ""
}

func (x *LinkStats) GetClicks() int32 {
	if x != nil {
		return x.Clicks
	}
	return 0
}

var File_proto_service_proto protoreflect.FileDescriptor

var file_proto_service_proto_rawDesc = string([]byte{
	0x0a, 0x13, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x74, 0x75, 0x74, 0x6f, 0x72, 0x69, 0x61, 0x6c, 0x22,
	0x43, 0x0a, 0x07, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x22, 0x40, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61,
	0x72, 0x74, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x44, 0x0a, 0x15, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x2b, 0x0a, 0x07, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x11, 0x2e, 0x74, 0x75, 0x74, 0x6f, 0x72, 0x69, 0x61, 0x6c, 0x2e, 0x50, 0x61, 0x72, 0x74,
	0x6e, 0x65, 0x72, 0x52, 0x07, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x22, 0x37, 0x0a, 0x16,
	0x47, 0x65, 0x74, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65,
	0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x70, 0x61, 0x72, 0x74,
	0x6e, 0x65, 0x72, 0x49, 0x64, 0x22, 0x67, 0x0a, 0x0c, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72,
	0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72,
	0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x70, 0x61, 0x72, 0x74, 0x6e,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6c, 0x69, 0x63, 0x6b, 0x73, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x63, 0x6c, 0x69, 0x63, 0x6b, 0x73, 0x12, 0x20, 0x0a, 0x0b,
	0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x37,
	0x0a, 0x12, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x21, 0x0a, 0x0c, 0x6f, 0x72, 0x69, 0x67, 0x69, 0x6e, 0x61, 0x6c,
	0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x6f, 0x72, 0x69, 0x67,
	0x69, 0x6e, 0x61, 0x6c, 0x55, 0x72, 0x6c, 0x22, 0x32, 0x0a, 0x13, 0x53, 0x68, 0x6f, 0x72, 0x74,
	0x65, 0x6e, 0x4c, 0x69, 0x6e, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1b,
	0x0a, 0x09, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x72, 0x6c, 0x22, 0x32, 0x0a, 0x13, 0x47,
	0x65, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x72, 0x6c, 0x22,
	0x40, 0x0a, 0x09, 0x4c, 0x69, 0x6e, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x1b, 0x0a, 0x09,
	0x73, 0x68, 0x6f, 0x72, 0x74, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x73, 0x68, 0x6f, 0x72, 0x74, 0x55, 0x72, 0x6c, 0x12, 0x16, 0x0a, 0x06, 0x63, 0x6c, 0x69,
	0x63, 0x6b, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x63, 0x6c, 0x69, 0x63, 0x6b,
	0x73, 0x32, 0xaf, 0x01, 0x0a, 0x0e, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x50, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61,
	0x72, 0x74, 0x6e, 0x65, 0x72, 0x12, 0x1e, 0x2e, 0x74, 0x75, 0x74, 0x6f, 0x72, 0x69, 0x61, 0x6c,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x74, 0x75, 0x74, 0x6f, 0x72, 0x69, 0x61, 0x6c,
	0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x4b, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x50, 0x61, 0x72,
	0x74, 0x6e, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x20, 0x2e, 0x74, 0x75, 0x74, 0x6f,
	0x72, 0x69, 0x61, 0x6c, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x53,
	0x74, 0x61, 0x74, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x74, 0x75,
	0x74, 0x6f, 0x72, 0x69, 0x61, 0x6c, 0x2e, 0x50, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x53, 0x74,
	0x61, 0x74, 0x73, 0x32, 0xa7, 0x01, 0x0a, 0x15, 0x4c, 0x69, 0x6e, 0x6b, 0x53, 0x68, 0x6f, 0x72,
	0x74, 0x65, 0x6e, 0x69, 0x6e, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4a, 0x0a,
	0x0b, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x4c, 0x69, 0x6e, 0x6b, 0x12, 0x1c, 0x2e, 0x74,
	0x75, 0x74, 0x6f, 0x72, 0x69, 0x61, 0x6c, 0x2e, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x4c,
	0x69, 0x6e, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x74, 0x75, 0x74,
	0x6f, 0x72, 0x69, 0x61, 0x6c, 0x2e, 0x53, 0x68, 0x6f, 0x72, 0x74, 0x65, 0x6e, 0x4c, 0x69, 0x6e,
	0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x42, 0x0a, 0x0c, 0x47, 0x65, 0x74,
	0x4c, 0x69, 0x6e, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x73, 0x12, 0x1d, 0x2e, 0x74, 0x75, 0x74, 0x6f,
	0x72, 0x69, 0x61, 0x6c, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x69, 0x6e, 0x6b, 0x53, 0x74, 0x61, 0x74,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x74, 0x75, 0x74, 0x6f, 0x72,
	0x69, 0x61, 0x6c, 0x2e, 0x4c, 0x69, 0x6e, 0x6b, 0x53, 0x74, 0x61, 0x74, 0x73, 0x42, 0x2c, 0x5a,
	0x2a, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x69, 0x72, 0x6f,
	0x6c, 0x6c, 0x65, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x3b, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
})

var (
	file_proto_service_proto_rawDescOnce sync.Once
	file_proto_service_proto_rawDescData []byte
)

func file_proto_service_proto_rawDescGZIP() []byte {
	file_proto_service_proto_rawDescOnce.Do(func() {
		file_proto_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_proto_service_proto_rawDesc), len(file_proto_service_proto_rawDesc)))
	})
	return file_proto_service_proto_rawDescData
}

var file_proto_service_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_proto_service_proto_goTypes = []any{
	(*Partner)(nil),                // 0: tutorial.Partner
	(*CreatePartnerRequest)(nil),   // 1: tutorial.CreatePartnerRequest
	(*CreatePartnerResponse)(nil),  // 2: tutorial.CreatePartnerResponse
	(*GetPartnerStatsRequest)(nil), // 3: tutorial.GetPartnerStatsRequest
	(*PartnerStats)(nil),           // 4: tutorial.PartnerStats
	(*ShortenLinkRequest)(nil),     // 5: tutorial.ShortenLinkRequest
	(*ShortenLinkResponse)(nil),    // 6: tutorial.ShortenLinkResponse
	(*GetLinkStatsRequest)(nil),    // 7: tutorial.GetLinkStatsRequest
	(*LinkStats)(nil),              // 8: tutorial.LinkStats
}
var file_proto_service_proto_depIdxs = []int32{
	0, // 0: tutorial.CreatePartnerResponse.partner:type_name -> tutorial.Partner
	1, // 1: tutorial.PartnerService.CreatePartner:input_type -> tutorial.CreatePartnerRequest
	3, // 2: tutorial.PartnerService.GetPartnerStats:input_type -> tutorial.GetPartnerStatsRequest
	5, // 3: tutorial.LinkShorteningService.ShortenLink:input_type -> tutorial.ShortenLinkRequest
	7, // 4: tutorial.LinkShorteningService.GetLinkStats:input_type -> tutorial.GetLinkStatsRequest
	2, // 5: tutorial.PartnerService.CreatePartner:output_type -> tutorial.CreatePartnerResponse
	4, // 6: tutorial.PartnerService.GetPartnerStats:output_type -> tutorial.PartnerStats
	6, // 7: tutorial.LinkShorteningService.ShortenLink:output_type -> tutorial.ShortenLinkResponse
	8, // 8: tutorial.LinkShorteningService.GetLinkStats:output_type -> tutorial.LinkStats
	5, // [5:9] is the sub-list for method output_type
	1, // [1:5] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_service_proto_init() }
func file_proto_service_proto_init() {
	if File_proto_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_proto_service_proto_rawDesc), len(file_proto_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   2,
		},
		GoTypes:           file_proto_service_proto_goTypes,
		DependencyIndexes: file_proto_service_proto_depIdxs,
		MessageInfos:      file_proto_service_proto_msgTypes,
	}.Build()
	File_proto_service_proto = out.File
	file_proto_service_proto_goTypes = nil
	file_proto_service_proto_depIdxs = nil
}
