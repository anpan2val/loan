// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: loan/loan/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type MsgRequestLoan struct {
	Creator    string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Amount     string `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount,omitempty"`
	Fee        string `protobuf:"bytes,3,opt,name=fee,proto3" json:"fee,omitempty"`
	Collateral string `protobuf:"bytes,4,opt,name=collateral,proto3" json:"collateral,omitempty"`
	Deadline   string `protobuf:"bytes,5,opt,name=deadline,proto3" json:"deadline,omitempty"`
}

func (m *MsgRequestLoan) Reset()         { *m = MsgRequestLoan{} }
func (m *MsgRequestLoan) String() string { return proto.CompactTextString(m) }
func (*MsgRequestLoan) ProtoMessage()    {}
func (*MsgRequestLoan) Descriptor() ([]byte, []int) {
	return fileDescriptor_a914cc78de7d0204, []int{0}
}
func (m *MsgRequestLoan) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRequestLoan) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRequestLoan.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRequestLoan) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRequestLoan.Merge(m, src)
}
func (m *MsgRequestLoan) XXX_Size() int {
	return m.Size()
}
func (m *MsgRequestLoan) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRequestLoan.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRequestLoan proto.InternalMessageInfo

func (m *MsgRequestLoan) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgRequestLoan) GetAmount() string {
	if m != nil {
		return m.Amount
	}
	return ""
}

func (m *MsgRequestLoan) GetFee() string {
	if m != nil {
		return m.Fee
	}
	return ""
}

func (m *MsgRequestLoan) GetCollateral() string {
	if m != nil {
		return m.Collateral
	}
	return ""
}

func (m *MsgRequestLoan) GetDeadline() string {
	if m != nil {
		return m.Deadline
	}
	return ""
}

type MsgRequestLoanResponse struct {
}

func (m *MsgRequestLoanResponse) Reset()         { *m = MsgRequestLoanResponse{} }
func (m *MsgRequestLoanResponse) String() string { return proto.CompactTextString(m) }
func (*MsgRequestLoanResponse) ProtoMessage()    {}
func (*MsgRequestLoanResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a914cc78de7d0204, []int{1}
}
func (m *MsgRequestLoanResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgRequestLoanResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgRequestLoanResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgRequestLoanResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgRequestLoanResponse.Merge(m, src)
}
func (m *MsgRequestLoanResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgRequestLoanResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgRequestLoanResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgRequestLoanResponse proto.InternalMessageInfo

type MsgApproveLoan struct {
	Creator string `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Id      uint64 `protobuf:"varint,2,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *MsgApproveLoan) Reset()         { *m = MsgApproveLoan{} }
func (m *MsgApproveLoan) String() string { return proto.CompactTextString(m) }
func (*MsgApproveLoan) ProtoMessage()    {}
func (*MsgApproveLoan) Descriptor() ([]byte, []int) {
	return fileDescriptor_a914cc78de7d0204, []int{2}
}
func (m *MsgApproveLoan) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgApproveLoan) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgApproveLoan.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgApproveLoan) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgApproveLoan.Merge(m, src)
}
func (m *MsgApproveLoan) XXX_Size() int {
	return m.Size()
}
func (m *MsgApproveLoan) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgApproveLoan.DiscardUnknown(m)
}

var xxx_messageInfo_MsgApproveLoan proto.InternalMessageInfo

func (m *MsgApproveLoan) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgApproveLoan) GetId() uint64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type MsgApproveLoanResponse struct {
}

func (m *MsgApproveLoanResponse) Reset()         { *m = MsgApproveLoanResponse{} }
func (m *MsgApproveLoanResponse) String() string { return proto.CompactTextString(m) }
func (*MsgApproveLoanResponse) ProtoMessage()    {}
func (*MsgApproveLoanResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_a914cc78de7d0204, []int{3}
}
func (m *MsgApproveLoanResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgApproveLoanResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgApproveLoanResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgApproveLoanResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgApproveLoanResponse.Merge(m, src)
}
func (m *MsgApproveLoanResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgApproveLoanResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgApproveLoanResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgApproveLoanResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgRequestLoan)(nil), "anpan2val.loan.loan.MsgRequestLoan")
	proto.RegisterType((*MsgRequestLoanResponse)(nil), "anpan2val.loan.loan.MsgRequestLoanResponse")
	proto.RegisterType((*MsgApproveLoan)(nil), "anpan2val.loan.loan.MsgApproveLoan")
	proto.RegisterType((*MsgApproveLoanResponse)(nil), "anpan2val.loan.loan.MsgApproveLoanResponse")
}

func init() { proto.RegisterFile("loan/loan/tx.proto", fileDescriptor_a914cc78de7d0204) }

var fileDescriptor_a914cc78de7d0204 = []byte{
	// 307 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0xbb, 0x6d, 0xad, 0x3a, 0x42, 0x91, 0x15, 0xca, 0xd2, 0xc3, 0x22, 0x15, 0x44, 0x10,
	0x52, 0xa8, 0x37, 0x4f, 0xea, 0xd9, 0x5e, 0x72, 0xf4, 0x22, 0xdb, 0x66, 0xac, 0x81, 0xed, 0xee,
	0x9a, 0xdd, 0x94, 0xfa, 0x16, 0x7d, 0x2b, 0x3d, 0xf6, 0xe8, 0x51, 0x92, 0x17, 0x91, 0xac, 0x49,
	0x49, 0xa1, 0x45, 0x2f, 0x43, 0x66, 0xfe, 0xc9, 0xfc, 0xdf, 0x0e, 0x03, 0x54, 0x6a, 0xa1, 0x86,
	0x3e, 0xb8, 0x65, 0x60, 0x12, 0xed, 0x34, 0x3d, 0x13, 0xca, 0x08, 0x35, 0x5a, 0x08, 0x19, 0x14,
	0x82, 0x0f, 0x83, 0x15, 0x81, 0xee, 0xd8, 0xce, 0x42, 0x7c, 0x4b, 0xd1, 0xba, 0x47, 0x2d, 0x14,
	0x65, 0x70, 0x38, 0x4d, 0x50, 0x38, 0x9d, 0x30, 0x72, 0x4e, 0xae, 0x8e, 0xc3, 0x2a, 0xa5, 0x3d,
	0xe8, 0x88, 0xb9, 0x4e, 0x95, 0x63, 0x4d, 0x2f, 0x94, 0x19, 0x3d, 0x85, 0xd6, 0x0b, 0x22, 0x6b,
	0xf9, 0x62, 0xf1, 0x49, 0x39, 0xc0, 0x54, 0x4b, 0x29, 0x1c, 0x26, 0x42, 0xb2, 0xb6, 0x17, 0x6a,
	0x15, 0xda, 0x87, 0xa3, 0x08, 0x45, 0x24, 0x63, 0x85, 0xec, 0xc0, 0xab, 0x9b, 0x7c, 0xc0, 0xa0,
	0xb7, 0x4d, 0x14, 0xa2, 0x35, 0x5a, 0x59, 0x1c, 0xdc, 0x7a, 0xd6, 0x7b, 0x63, 0x12, 0xbd, 0xc0,
	0x3f, 0x58, 0xbb, 0xd0, 0x8c, 0x23, 0xcf, 0xd9, 0x0e, 0x9b, 0x71, 0x54, 0x4e, 0xad, 0xfd, 0x5b,
	0x4d, 0x1d, 0x7d, 0x10, 0x68, 0x8d, 0xed, 0x8c, 0x3e, 0xc3, 0x49, 0x7d, 0x0d, 0x17, 0xc1, 0x8e,
	0x7d, 0x05, 0xdb, 0x64, 0xfd, 0xeb, 0x7f, 0x34, 0x55, 0x46, 0x85, 0x41, 0x9d, 0x7d, 0xaf, 0x41,
	0xad, 0x69, 0xbf, 0xc1, 0x8e, 0x97, 0x3c, 0xdc, 0x7d, 0x66, 0x9c, 0xac, 0x33, 0x4e, 0xbe, 0x33,
	0x4e, 0x56, 0x39, 0x6f, 0xac, 0x73, 0xde, 0xf8, 0xca, 0x79, 0xe3, 0xe9, 0x72, 0x16, 0xbb, 0xd7,
	0x74, 0x12, 0x4c, 0xf5, 0x7c, 0xb8, 0x19, 0xf8, 0x7b, 0x1f, 0xcb, 0xf2, 0x4c, 0xde, 0x0d, 0xda,
	0x49, 0xc7, 0x9f, 0xca, 0xcd, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xfd, 0xe5, 0x62, 0x04, 0x40,
	0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	RequestLoan(ctx context.Context, in *MsgRequestLoan, opts ...grpc.CallOption) (*MsgRequestLoanResponse, error)
	ApproveLoan(ctx context.Context, in *MsgApproveLoan, opts ...grpc.CallOption) (*MsgApproveLoanResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) RequestLoan(ctx context.Context, in *MsgRequestLoan, opts ...grpc.CallOption) (*MsgRequestLoanResponse, error) {
	out := new(MsgRequestLoanResponse)
	err := c.cc.Invoke(ctx, "/anpan2val.loan.loan.Msg/RequestLoan", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) ApproveLoan(ctx context.Context, in *MsgApproveLoan, opts ...grpc.CallOption) (*MsgApproveLoanResponse, error) {
	out := new(MsgApproveLoanResponse)
	err := c.cc.Invoke(ctx, "/anpan2val.loan.loan.Msg/ApproveLoan", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	RequestLoan(context.Context, *MsgRequestLoan) (*MsgRequestLoanResponse, error)
	ApproveLoan(context.Context, *MsgApproveLoan) (*MsgApproveLoanResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) RequestLoan(ctx context.Context, req *MsgRequestLoan) (*MsgRequestLoanResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RequestLoan not implemented")
}
func (*UnimplementedMsgServer) ApproveLoan(ctx context.Context, req *MsgApproveLoan) (*MsgApproveLoanResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ApproveLoan not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_RequestLoan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgRequestLoan)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).RequestLoan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/anpan2val.loan.loan.Msg/RequestLoan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).RequestLoan(ctx, req.(*MsgRequestLoan))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_ApproveLoan_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgApproveLoan)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).ApproveLoan(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/anpan2val.loan.loan.Msg/ApproveLoan",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).ApproveLoan(ctx, req.(*MsgApproveLoan))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "anpan2val.loan.loan.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RequestLoan",
			Handler:    _Msg_RequestLoan_Handler,
		},
		{
			MethodName: "ApproveLoan",
			Handler:    _Msg_ApproveLoan_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "loan/loan/tx.proto",
}

func (m *MsgRequestLoan) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRequestLoan) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRequestLoan) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Deadline) > 0 {
		i -= len(m.Deadline)
		copy(dAtA[i:], m.Deadline)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Deadline)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Collateral) > 0 {
		i -= len(m.Collateral)
		copy(dAtA[i:], m.Collateral)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Collateral)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Fee) > 0 {
		i -= len(m.Fee)
		copy(dAtA[i:], m.Fee)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Fee)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Amount) > 0 {
		i -= len(m.Amount)
		copy(dAtA[i:], m.Amount)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Amount)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgRequestLoanResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgRequestLoanResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgRequestLoanResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgApproveLoan) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgApproveLoan) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgApproveLoan) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Id != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Id))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgApproveLoanResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgApproveLoanResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgApproveLoanResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgRequestLoan) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Amount)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Fee)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Collateral)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Deadline)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgRequestLoanResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgApproveLoan) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Id != 0 {
		n += 1 + sovTx(uint64(m.Id))
	}
	return n
}

func (m *MsgApproveLoanResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgRequestLoan) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgRequestLoan: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRequestLoan: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Amount = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fee", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Fee = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Collateral", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Collateral = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Deadline", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Deadline = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgRequestLoanResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgRequestLoanResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgRequestLoanResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgApproveLoan) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgApproveLoan: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgApproveLoan: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			m.Id = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Id |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgApproveLoanResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: MsgApproveLoanResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgApproveLoanResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowTx
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
			if length < 0 {
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
