// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/home.proto

/*
	Package shop_srv_home is a generated protocol buffer package.

	It is generated from these files:
		proto/home.proto

	It has these top-level messages:
		HomeHeaderReq
		HomeNav
		HomeCourse
		HomeHeadersResp
		HomeContentsReq
		HomeContentsResp
*/
package shop_srv_home

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import shop_srv_shopproto "shop-micro/shopproto/news"
import shop_srv_shopproto1 "shop-micro/shopproto/product"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type HomeHeaderReq struct {
}

func (m *HomeHeaderReq) Reset()                    { *m = HomeHeaderReq{} }
func (m *HomeHeaderReq) String() string            { return proto.CompactTextString(m) }
func (*HomeHeaderReq) ProtoMessage()               {}
func (*HomeHeaderReq) Descriptor() ([]byte, []int) { return fileDescriptorHome, []int{0} }

type HomeNav struct {
	Id      int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title   string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	EnTitle string `protobuf:"bytes,3,opt,name=enTitle,proto3" json:"enTitle,omitempty"`
	ImgUrl  string `protobuf:"bytes,4,opt,name=imgUrl,proto3" json:"imgUrl,omitempty"`
	Types   int32  `protobuf:"varint,5,opt,name=types,proto3" json:"types,omitempty"`
	ClassId int32  `protobuf:"varint,6,opt,name=classId,proto3" json:"classId,omitempty"`
}

func (m *HomeNav) Reset()                    { *m = HomeNav{} }
func (m *HomeNav) String() string            { return proto.CompactTextString(m) }
func (*HomeNav) ProtoMessage()               {}
func (*HomeNav) Descriptor() ([]byte, []int) { return fileDescriptorHome, []int{1} }

func (m *HomeNav) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *HomeNav) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *HomeNav) GetEnTitle() string {
	if m != nil {
		return m.EnTitle
	}
	return ""
}

func (m *HomeNav) GetImgUrl() string {
	if m != nil {
		return m.ImgUrl
	}
	return ""
}

func (m *HomeNav) GetTypes() int32 {
	if m != nil {
		return m.Types
	}
	return 0
}

func (m *HomeNav) GetClassId() int32 {
	if m != nil {
		return m.ClassId
	}
	return 0
}

type HomeCourse struct {
	Id      int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Title   string `protobuf:"bytes,2,opt,name=title,proto3" json:"title,omitempty"`
	EnTitle string `protobuf:"bytes,3,opt,name=enTitle,proto3" json:"enTitle,omitempty"`
	ImgUrl  string `protobuf:"bytes,4,opt,name=imgUrl,proto3" json:"imgUrl,omitempty"`
	Url     string `protobuf:"bytes,5,opt,name=url,proto3" json:"url,omitempty"`
}

func (m *HomeCourse) Reset()                    { *m = HomeCourse{} }
func (m *HomeCourse) String() string            { return proto.CompactTextString(m) }
func (*HomeCourse) ProtoMessage()               {}
func (*HomeCourse) Descriptor() ([]byte, []int) { return fileDescriptorHome, []int{2} }

func (m *HomeCourse) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *HomeCourse) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *HomeCourse) GetEnTitle() string {
	if m != nil {
		return m.EnTitle
	}
	return ""
}

func (m *HomeCourse) GetImgUrl() string {
	if m != nil {
		return m.ImgUrl
	}
	return ""
}

func (m *HomeCourse) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

type HomeHeadersResp struct {
	HomeNavList    []*HomeNav    `protobuf:"bytes,1,rep,name=homeNavList" json:"homeNavList,omitempty"`
	HomeCourseList []*HomeCourse `protobuf:"bytes,2,rep,name=homeCourseList" json:"homeCourseList,omitempty"`
}

func (m *HomeHeadersResp) Reset()                    { *m = HomeHeadersResp{} }
func (m *HomeHeadersResp) String() string            { return proto.CompactTextString(m) }
func (*HomeHeadersResp) ProtoMessage()               {}
func (*HomeHeadersResp) Descriptor() ([]byte, []int) { return fileDescriptorHome, []int{3} }

func (m *HomeHeadersResp) GetHomeNavList() []*HomeNav {
	if m != nil {
		return m.HomeNavList
	}
	return nil
}

func (m *HomeHeadersResp) GetHomeCourseList() []*HomeCourse {
	if m != nil {
		return m.HomeCourseList
	}
	return nil
}

type HomeContentsReq struct {
}

func (m *HomeContentsReq) Reset()                    { *m = HomeContentsReq{} }
func (m *HomeContentsReq) String() string            { return proto.CompactTextString(m) }
func (*HomeContentsReq) ProtoMessage()               {}
func (*HomeContentsReq) Descriptor() ([]byte, []int) { return fileDescriptorHome, []int{4} }

type HomeContentsResp struct {
	HomeGoodsList []*shop_srv_shopproto1.Product `protobuf:"bytes,1,rep,name=HomeGoodsList" json:"HomeGoodsList,omitempty"`
	HomeNewsList  []*shop_srv_shopproto.News     `protobuf:"bytes,2,rep,name=HomeNewsList" json:"HomeNewsList,omitempty"`
}

func (m *HomeContentsResp) Reset()                    { *m = HomeContentsResp{} }
func (m *HomeContentsResp) String() string            { return proto.CompactTextString(m) }
func (*HomeContentsResp) ProtoMessage()               {}
func (*HomeContentsResp) Descriptor() ([]byte, []int) { return fileDescriptorHome, []int{5} }

func (m *HomeContentsResp) GetHomeGoodsList() []*shop_srv_shopproto1.Product {
	if m != nil {
		return m.HomeGoodsList
	}
	return nil
}

func (m *HomeContentsResp) GetHomeNewsList() []*shop_srv_shopproto.News {
	if m != nil {
		return m.HomeNewsList
	}
	return nil
}

func init() {
	proto.RegisterType((*HomeHeaderReq)(nil), "shop.srv.home.HomeHeaderReq")
	proto.RegisterType((*HomeNav)(nil), "shop.srv.home.HomeNav")
	proto.RegisterType((*HomeCourse)(nil), "shop.srv.home.HomeCourse")
	proto.RegisterType((*HomeHeadersResp)(nil), "shop.srv.home.HomeHeadersResp")
	proto.RegisterType((*HomeContentsReq)(nil), "shop.srv.home.HomeContentsReq")
	proto.RegisterType((*HomeContentsResp)(nil), "shop.srv.home.HomeContentsResp")
}
func (m *HomeHeaderReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HomeHeaderReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *HomeNav) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HomeNav) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintHome(dAtA, i, uint64(m.Id))
	}
	if len(m.Title) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintHome(dAtA, i, uint64(len(m.Title)))
		i += copy(dAtA[i:], m.Title)
	}
	if len(m.EnTitle) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintHome(dAtA, i, uint64(len(m.EnTitle)))
		i += copy(dAtA[i:], m.EnTitle)
	}
	if len(m.ImgUrl) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintHome(dAtA, i, uint64(len(m.ImgUrl)))
		i += copy(dAtA[i:], m.ImgUrl)
	}
	if m.Types != 0 {
		dAtA[i] = 0x28
		i++
		i = encodeVarintHome(dAtA, i, uint64(m.Types))
	}
	if m.ClassId != 0 {
		dAtA[i] = 0x30
		i++
		i = encodeVarintHome(dAtA, i, uint64(m.ClassId))
	}
	return i, nil
}

func (m *HomeCourse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HomeCourse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintHome(dAtA, i, uint64(m.Id))
	}
	if len(m.Title) > 0 {
		dAtA[i] = 0x12
		i++
		i = encodeVarintHome(dAtA, i, uint64(len(m.Title)))
		i += copy(dAtA[i:], m.Title)
	}
	if len(m.EnTitle) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintHome(dAtA, i, uint64(len(m.EnTitle)))
		i += copy(dAtA[i:], m.EnTitle)
	}
	if len(m.ImgUrl) > 0 {
		dAtA[i] = 0x22
		i++
		i = encodeVarintHome(dAtA, i, uint64(len(m.ImgUrl)))
		i += copy(dAtA[i:], m.ImgUrl)
	}
	if len(m.Url) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintHome(dAtA, i, uint64(len(m.Url)))
		i += copy(dAtA[i:], m.Url)
	}
	return i, nil
}

func (m *HomeHeadersResp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HomeHeadersResp) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.HomeNavList) > 0 {
		for _, msg := range m.HomeNavList {
			dAtA[i] = 0xa
			i++
			i = encodeVarintHome(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.HomeCourseList) > 0 {
		for _, msg := range m.HomeCourseList {
			dAtA[i] = 0x12
			i++
			i = encodeVarintHome(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *HomeContentsReq) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HomeContentsReq) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *HomeContentsResp) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *HomeContentsResp) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.HomeGoodsList) > 0 {
		for _, msg := range m.HomeGoodsList {
			dAtA[i] = 0xa
			i++
			i = encodeVarintHome(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.HomeNewsList) > 0 {
		for _, msg := range m.HomeNewsList {
			dAtA[i] = 0x12
			i++
			i = encodeVarintHome(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func encodeVarintHome(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *HomeHeaderReq) Size() (n int) {
	var l int
	_ = l
	return n
}

func (m *HomeNav) Size() (n int) {
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovHome(uint64(m.Id))
	}
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovHome(uint64(l))
	}
	l = len(m.EnTitle)
	if l > 0 {
		n += 1 + l + sovHome(uint64(l))
	}
	l = len(m.ImgUrl)
	if l > 0 {
		n += 1 + l + sovHome(uint64(l))
	}
	if m.Types != 0 {
		n += 1 + sovHome(uint64(m.Types))
	}
	if m.ClassId != 0 {
		n += 1 + sovHome(uint64(m.ClassId))
	}
	return n
}

func (m *HomeCourse) Size() (n int) {
	var l int
	_ = l
	if m.Id != 0 {
		n += 1 + sovHome(uint64(m.Id))
	}
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovHome(uint64(l))
	}
	l = len(m.EnTitle)
	if l > 0 {
		n += 1 + l + sovHome(uint64(l))
	}
	l = len(m.ImgUrl)
	if l > 0 {
		n += 1 + l + sovHome(uint64(l))
	}
	l = len(m.Url)
	if l > 0 {
		n += 1 + l + sovHome(uint64(l))
	}
	return n
}

func (m *HomeHeadersResp) Size() (n int) {
	var l int
	_ = l
	if len(m.HomeNavList) > 0 {
		for _, e := range m.HomeNavList {
			l = e.Size()
			n += 1 + l + sovHome(uint64(l))
		}
	}
	if len(m.HomeCourseList) > 0 {
		for _, e := range m.HomeCourseList {
			l = e.Size()
			n += 1 + l + sovHome(uint64(l))
		}
	}
	return n
}

func (m *HomeContentsReq) Size() (n int) {
	var l int
	_ = l
	return n
}

func (m *HomeContentsResp) Size() (n int) {
	var l int
	_ = l
	if len(m.HomeGoodsList) > 0 {
		for _, e := range m.HomeGoodsList {
			l = e.Size()
			n += 1 + l + sovHome(uint64(l))
		}
	}
	if len(m.HomeNewsList) > 0 {
		for _, e := range m.HomeNewsList {
			l = e.Size()
			n += 1 + l + sovHome(uint64(l))
		}
	}
	return n
}

func sovHome(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozHome(x uint64) (n int) {
	return sovHome(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *HomeHeaderReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHome
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: HomeHeaderReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HomeHeaderReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipHome(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHome
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *HomeNav) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHome
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: HomeNav: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HomeNav: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHome
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHome
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthHome
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EnTitle", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHome
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthHome
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EnTitle = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ImgUrl", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHome
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthHome
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ImgUrl = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Types", wireType)
			}
			m.Types = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHome
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Types |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClassId", wireType)
			}
			m.ClassId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHome
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ClassId |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipHome(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHome
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *HomeCourse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHome
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: HomeCourse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HomeCourse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHome
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHome
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthHome
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EnTitle", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHome
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthHome
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.EnTitle = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ImgUrl", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHome
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthHome
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ImgUrl = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Url", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHome
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthHome
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Url = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHome(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHome
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *HomeHeadersResp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHome
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: HomeHeadersResp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HomeHeadersResp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HomeNavList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHome
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthHome
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.HomeNavList = append(m.HomeNavList, &HomeNav{})
			if err := m.HomeNavList[len(m.HomeNavList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HomeCourseList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHome
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthHome
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.HomeCourseList = append(m.HomeCourseList, &HomeCourse{})
			if err := m.HomeCourseList[len(m.HomeCourseList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHome(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHome
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *HomeContentsReq) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHome
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: HomeContentsReq: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HomeContentsReq: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipHome(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHome
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *HomeContentsResp) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowHome
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: HomeContentsResp: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: HomeContentsResp: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HomeGoodsList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHome
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthHome
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.HomeGoodsList = append(m.HomeGoodsList, &shop_srv_shopproto1.Product{})
			if err := m.HomeGoodsList[len(m.HomeGoodsList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field HomeNewsList", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowHome
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthHome
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.HomeNewsList = append(m.HomeNewsList, &shop_srv_shopproto.News{})
			if err := m.HomeNewsList[len(m.HomeNewsList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipHome(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthHome
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipHome(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowHome
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowHome
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowHome
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthHome
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowHome
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipHome(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthHome = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowHome   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("proto/home.proto", fileDescriptorHome) }

var fileDescriptorHome = []byte{
	// 438 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x52, 0xd1, 0x8a, 0xd4, 0x30,
	0x14, 0xdd, 0xb4, 0xce, 0x2c, 0x7b, 0xc7, 0xdd, 0xa9, 0x41, 0x96, 0x58, 0xa5, 0x96, 0xe2, 0x43,
	0x11, 0xec, 0xc2, 0xfa, 0xe2, 0x83, 0x2f, 0xeb, 0x82, 0xae, 0x20, 0xa2, 0xd1, 0xfd, 0x80, 0xb1,
	0x0d, 0x36, 0xd0, 0x36, 0x35, 0xe9, 0x74, 0xf1, 0x27, 0x04, 0xf1, 0xc5, 0x3f, 0xf1, 0x17, 0x7c,
	0xf4, 0x13, 0x64, 0xfc, 0x11, 0x49, 0x32, 0xed, 0xb4, 0x52, 0x7d, 0xf3, 0x65, 0x26, 0xf7, 0xe4,
	0xdc, 0xdb, 0x73, 0x4e, 0x2e, 0x78, 0xb5, 0x14, 0x8d, 0x38, 0xc9, 0x45, 0xc9, 0x12, 0x73, 0xc4,
	0x87, 0x2a, 0x17, 0x75, 0xa2, 0x64, 0x9b, 0x68, 0xd0, 0xbf, 0xa7, 0xcb, 0x07, 0x25, 0x4f, 0xa5,
	0x38, 0xd1, 0x47, 0xcb, 0xaf, 0xd8, 0x95, 0x32, 0x3f, 0xb6, 0xc9, 0xbf, 0x3f, 0xc9, 0xaa, 0xa5,
	0xc8, 0xd6, 0x69, 0xd3, 0xfd, 0x5b, 0x6e, 0xb4, 0x84, 0xc3, 0x0b, 0x51, 0xb2, 0x0b, 0xb6, 0xca,
	0x98, 0xa4, 0xec, 0x43, 0xf4, 0x19, 0xc1, 0xbe, 0x46, 0x5e, 0xae, 0x5a, 0x7c, 0x04, 0x0e, 0xcf,
	0x08, 0x0a, 0x51, 0xec, 0x52, 0x87, 0x67, 0xf8, 0x26, 0xcc, 0x1a, 0xde, 0x14, 0x8c, 0x38, 0x21,
	0x8a, 0x0f, 0xa8, 0x2d, 0x30, 0x81, 0x7d, 0x56, 0xbd, 0x35, 0xb8, 0x6b, 0xf0, 0xae, 0xc4, 0xc7,
	0x30, 0xe7, 0xe5, 0xfb, 0x4b, 0x59, 0x90, 0x6b, 0xe6, 0x62, 0x5b, 0x99, 0x39, 0x1f, 0x6b, 0xa6,
	0xc8, 0x2c, 0x44, 0xf1, 0x8c, 0xda, 0x42, 0xcf, 0x49, 0x8b, 0x95, 0x52, 0xcf, 0x33, 0x32, 0x37,
	0x78, 0x57, 0x46, 0x2d, 0x80, 0x96, 0x74, 0x2e, 0xd6, 0x52, 0xb1, 0xff, 0xa6, 0xca, 0x03, 0x77,
	0x2d, 0x0b, 0xa3, 0xe9, 0x80, 0xea, 0x63, 0xf4, 0x09, 0xc1, 0x72, 0x97, 0x8e, 0xa2, 0x4c, 0xd5,
	0xf8, 0x11, 0x2c, 0x72, 0x1b, 0xcf, 0x0b, 0xae, 0x1a, 0x82, 0x42, 0x37, 0x5e, 0x9c, 0x1e, 0x27,
	0xa3, 0x77, 0x4a, 0xb6, 0x01, 0xd2, 0x21, 0x15, 0x9f, 0xc1, 0x51, 0xde, 0xbb, 0x30, 0xcd, 0x8e,
	0x69, 0xbe, 0x35, 0xd1, 0x6c, 0x49, 0xf4, 0x8f, 0x86, 0xe8, 0x86, 0xd5, 0x73, 0x2e, 0xaa, 0x86,
	0x55, 0x8d, 0xd2, 0xef, 0xf5, 0x05, 0x81, 0x37, 0xc6, 0x54, 0x8d, 0xcf, 0xec, 0xab, 0x3e, 0x13,
	0x22, 0x53, 0x03, 0x99, 0xb7, 0x77, 0x5f, 0xea, 0xf7, 0x22, 0x79, 0x65, 0xf7, 0x81, 0x8e, 0x3b,
	0xf0, 0x63, 0xb8, 0x6e, 0x5c, 0xb0, 0x2b, 0x35, 0xd0, 0x4a, 0xa6, 0x26, 0x68, 0x0e, 0x1d, 0xb1,
	0x4f, 0xbf, 0x21, 0x58, 0x68, 0xe0, 0x0d, 0x93, 0x2d, 0x4f, 0x19, 0x7e, 0x0d, 0xcb, 0xa7, 0xbc,
	0xca, 0x06, 0x61, 0xe2, 0x3b, 0x13, 0xb6, 0xfb, 0x35, 0xf4, 0x83, 0xbf, 0xde, 0x1a, 0x87, 0xd1,
	0x1e, 0xbe, 0x04, 0xaf, 0x1b, 0xd9, 0x79, 0xc7, 0xc1, 0x64, 0x94, 0x7d, 0x58, 0xfe, 0xdd, 0x7f,
	0xde, 0xeb, 0xb1, 0x4f, 0xbc, 0xef, 0x9b, 0x00, 0xfd, 0xd8, 0x04, 0xe8, 0xe7, 0x26, 0x40, 0x5f,
	0x7f, 0x05, 0x7b, 0xef, 0xe6, 0xc6, 0xe5, 0xc3, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x42, 0xd7,
	0x43, 0xe9, 0x9e, 0x03, 0x00, 0x00,
}
