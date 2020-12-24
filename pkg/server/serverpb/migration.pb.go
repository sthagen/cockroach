// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: server/serverpb/migration.proto

package serverpb

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import clusterversion "github.com/cockroachdb/cockroach/pkg/clusterversion"

import (
	context "context"
	grpc "google.golang.org/grpc"
)

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

// ValidateTargetClusterVersion is used to verify that the target node is
// running a binary that's able to support the specified cluster version.
type ValidateTargetClusterVersionRequest struct {
	ClusterVersion *clusterversion.ClusterVersion `protobuf:"bytes,1,opt,name=cluster_version,json=clusterVersion,proto3" json:"cluster_version,omitempty"`
}

func (m *ValidateTargetClusterVersionRequest) Reset()         { *m = ValidateTargetClusterVersionRequest{} }
func (m *ValidateTargetClusterVersionRequest) String() string { return proto.CompactTextString(m) }
func (*ValidateTargetClusterVersionRequest) ProtoMessage()    {}
func (*ValidateTargetClusterVersionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_migration_8dfeb6fcf9144e4c, []int{0}
}
func (m *ValidateTargetClusterVersionRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ValidateTargetClusterVersionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalTo(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (dst *ValidateTargetClusterVersionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidateTargetClusterVersionRequest.Merge(dst, src)
}
func (m *ValidateTargetClusterVersionRequest) XXX_Size() int {
	return m.Size()
}
func (m *ValidateTargetClusterVersionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidateTargetClusterVersionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ValidateTargetClusterVersionRequest proto.InternalMessageInfo

// ValidateTargetClusterVersionResponse is the response to a
// ValidateTargetClusterVersionRequest.
type ValidateTargetClusterVersionResponse struct {
}

func (m *ValidateTargetClusterVersionResponse) Reset()         { *m = ValidateTargetClusterVersionResponse{} }
func (m *ValidateTargetClusterVersionResponse) String() string { return proto.CompactTextString(m) }
func (*ValidateTargetClusterVersionResponse) ProtoMessage()    {}
func (*ValidateTargetClusterVersionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_migration_8dfeb6fcf9144e4c, []int{1}
}
func (m *ValidateTargetClusterVersionResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ValidateTargetClusterVersionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalTo(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (dst *ValidateTargetClusterVersionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidateTargetClusterVersionResponse.Merge(dst, src)
}
func (m *ValidateTargetClusterVersionResponse) XXX_Size() int {
	return m.Size()
}
func (m *ValidateTargetClusterVersionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidateTargetClusterVersionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ValidateTargetClusterVersionResponse proto.InternalMessageInfo

// BumpClusterVersionRequest is used to inform a given node of a cluster version
// bump.
type BumpClusterVersionRequest struct {
	ClusterVersion *clusterversion.ClusterVersion `protobuf:"bytes,1,opt,name=cluster_version,json=clusterVersion,proto3" json:"cluster_version,omitempty"`
}

func (m *BumpClusterVersionRequest) Reset()         { *m = BumpClusterVersionRequest{} }
func (m *BumpClusterVersionRequest) String() string { return proto.CompactTextString(m) }
func (*BumpClusterVersionRequest) ProtoMessage()    {}
func (*BumpClusterVersionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_migration_8dfeb6fcf9144e4c, []int{2}
}
func (m *BumpClusterVersionRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BumpClusterVersionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalTo(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (dst *BumpClusterVersionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BumpClusterVersionRequest.Merge(dst, src)
}
func (m *BumpClusterVersionRequest) XXX_Size() int {
	return m.Size()
}
func (m *BumpClusterVersionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_BumpClusterVersionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_BumpClusterVersionRequest proto.InternalMessageInfo

// BumpClusterVersionResponse is the response to an BumpClusterVersionRequest.
type BumpClusterVersionResponse struct {
}

func (m *BumpClusterVersionResponse) Reset()         { *m = BumpClusterVersionResponse{} }
func (m *BumpClusterVersionResponse) String() string { return proto.CompactTextString(m) }
func (*BumpClusterVersionResponse) ProtoMessage()    {}
func (*BumpClusterVersionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_migration_8dfeb6fcf9144e4c, []int{3}
}
func (m *BumpClusterVersionResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BumpClusterVersionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalTo(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (dst *BumpClusterVersionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BumpClusterVersionResponse.Merge(dst, src)
}
func (m *BumpClusterVersionResponse) XXX_Size() int {
	return m.Size()
}
func (m *BumpClusterVersionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_BumpClusterVersionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_BumpClusterVersionResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ValidateTargetClusterVersionRequest)(nil), "cockroach.server.serverpb.ValidateTargetClusterVersionRequest")
	proto.RegisterType((*ValidateTargetClusterVersionResponse)(nil), "cockroach.server.serverpb.ValidateTargetClusterVersionResponse")
	proto.RegisterType((*BumpClusterVersionRequest)(nil), "cockroach.server.serverpb.BumpClusterVersionRequest")
	proto.RegisterType((*BumpClusterVersionResponse)(nil), "cockroach.server.serverpb.BumpClusterVersionResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MigrationClient is the client API for Migration service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MigrationClient interface {
	// ValidateTargetClusterVersion is used to verify that the target node is
	// running a binary that's able to support the specified cluster version.
	// Specifically:
	//
	//   node's minimum supported version <= version <= node's binary version
	ValidateTargetClusterVersion(ctx context.Context, in *ValidateTargetClusterVersionRequest, opts ...grpc.CallOption) (*ValidateTargetClusterVersionResponse, error)
	// BumpClusterVersion is used to inform a given node of a cluster version
	// bump. The node is responsible for durably persisting the message and
	// enabling the corresponding version gates.
	//
	// This RPC is typically used together with ValidateTargetClusterVersion,
	// which checks to see that all nodes in the cluster are running binaries
	// that would be able to support the intended version bump.
	BumpClusterVersion(ctx context.Context, in *BumpClusterVersionRequest, opts ...grpc.CallOption) (*BumpClusterVersionResponse, error)
}

type migrationClient struct {
	cc *grpc.ClientConn
}

func NewMigrationClient(cc *grpc.ClientConn) MigrationClient {
	return &migrationClient{cc}
}

func (c *migrationClient) ValidateTargetClusterVersion(ctx context.Context, in *ValidateTargetClusterVersionRequest, opts ...grpc.CallOption) (*ValidateTargetClusterVersionResponse, error) {
	out := new(ValidateTargetClusterVersionResponse)
	err := c.cc.Invoke(ctx, "/cockroach.server.serverpb.Migration/ValidateTargetClusterVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *migrationClient) BumpClusterVersion(ctx context.Context, in *BumpClusterVersionRequest, opts ...grpc.CallOption) (*BumpClusterVersionResponse, error) {
	out := new(BumpClusterVersionResponse)
	err := c.cc.Invoke(ctx, "/cockroach.server.serverpb.Migration/BumpClusterVersion", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MigrationServer is the server API for Migration service.
type MigrationServer interface {
	// ValidateTargetClusterVersion is used to verify that the target node is
	// running a binary that's able to support the specified cluster version.
	// Specifically:
	//
	//   node's minimum supported version <= version <= node's binary version
	ValidateTargetClusterVersion(context.Context, *ValidateTargetClusterVersionRequest) (*ValidateTargetClusterVersionResponse, error)
	// BumpClusterVersion is used to inform a given node of a cluster version
	// bump. The node is responsible for durably persisting the message and
	// enabling the corresponding version gates.
	//
	// This RPC is typically used together with ValidateTargetClusterVersion,
	// which checks to see that all nodes in the cluster are running binaries
	// that would be able to support the intended version bump.
	BumpClusterVersion(context.Context, *BumpClusterVersionRequest) (*BumpClusterVersionResponse, error)
}

func RegisterMigrationServer(s *grpc.Server, srv MigrationServer) {
	s.RegisterService(&_Migration_serviceDesc, srv)
}

func _Migration_ValidateTargetClusterVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ValidateTargetClusterVersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MigrationServer).ValidateTargetClusterVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cockroach.server.serverpb.Migration/ValidateTargetClusterVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MigrationServer).ValidateTargetClusterVersion(ctx, req.(*ValidateTargetClusterVersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Migration_BumpClusterVersion_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BumpClusterVersionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MigrationServer).BumpClusterVersion(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cockroach.server.serverpb.Migration/BumpClusterVersion",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MigrationServer).BumpClusterVersion(ctx, req.(*BumpClusterVersionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Migration_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cockroach.server.serverpb.Migration",
	HandlerType: (*MigrationServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ValidateTargetClusterVersion",
			Handler:    _Migration_ValidateTargetClusterVersion_Handler,
		},
		{
			MethodName: "BumpClusterVersion",
			Handler:    _Migration_BumpClusterVersion_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "server/serverpb/migration.proto",
}

func (m *ValidateTargetClusterVersionRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ValidateTargetClusterVersionRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.ClusterVersion != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintMigration(dAtA, i, uint64(m.ClusterVersion.Size()))
		n1, err := m.ClusterVersion.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n1
	}
	return i, nil
}

func (m *ValidateTargetClusterVersionResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ValidateTargetClusterVersionResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func (m *BumpClusterVersionRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BumpClusterVersionRequest) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.ClusterVersion != nil {
		dAtA[i] = 0xa
		i++
		i = encodeVarintMigration(dAtA, i, uint64(m.ClusterVersion.Size()))
		n2, err := m.ClusterVersion.MarshalTo(dAtA[i:])
		if err != nil {
			return 0, err
		}
		i += n2
	}
	return i, nil
}

func (m *BumpClusterVersionResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BumpClusterVersionResponse) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	return i, nil
}

func encodeVarintMigration(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *ValidateTargetClusterVersionRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ClusterVersion != nil {
		l = m.ClusterVersion.Size()
		n += 1 + l + sovMigration(uint64(l))
	}
	return n
}

func (m *ValidateTargetClusterVersionResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *BumpClusterVersionRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ClusterVersion != nil {
		l = m.ClusterVersion.Size()
		n += 1 + l + sovMigration(uint64(l))
	}
	return n
}

func (m *BumpClusterVersionResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovMigration(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozMigration(x uint64) (n int) {
	return sovMigration(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ValidateTargetClusterVersionRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMigration
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
			return fmt.Errorf("proto: ValidateTargetClusterVersionRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ValidateTargetClusterVersionRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClusterVersion", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMigration
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
				return ErrInvalidLengthMigration
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ClusterVersion == nil {
				m.ClusterVersion = &clusterversion.ClusterVersion{}
			}
			if err := m.ClusterVersion.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMigration(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMigration
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
func (m *ValidateTargetClusterVersionResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMigration
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
			return fmt.Errorf("proto: ValidateTargetClusterVersionResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ValidateTargetClusterVersionResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipMigration(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMigration
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
func (m *BumpClusterVersionRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMigration
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
			return fmt.Errorf("proto: BumpClusterVersionRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BumpClusterVersionRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClusterVersion", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMigration
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
				return ErrInvalidLengthMigration
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ClusterVersion == nil {
				m.ClusterVersion = &clusterversion.ClusterVersion{}
			}
			if err := m.ClusterVersion.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMigration(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMigration
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
func (m *BumpClusterVersionResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMigration
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
			return fmt.Errorf("proto: BumpClusterVersionResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BumpClusterVersionResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipMigration(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthMigration
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
func skipMigration(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMigration
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
					return 0, ErrIntOverflowMigration
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
					return 0, ErrIntOverflowMigration
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
				return 0, ErrInvalidLengthMigration
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowMigration
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
				next, err := skipMigration(dAtA[start:])
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
	ErrInvalidLengthMigration = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMigration   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("server/serverpb/migration.proto", fileDescriptor_migration_8dfeb6fcf9144e4c)
}

var fileDescriptor_migration_8dfeb6fcf9144e4c = []byte{
	// 278 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2f, 0x4e, 0x2d, 0x2a,
	0x4b, 0x2d, 0xd2, 0x87, 0x50, 0x05, 0x49, 0xfa, 0xb9, 0x99, 0xe9, 0x45, 0x89, 0x25, 0x99, 0xf9,
	0x79, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x92, 0xc9, 0xf9, 0xc9, 0xd9, 0x45, 0xf9, 0x89,
	0xc9, 0x19, 0x7a, 0x10, 0x35, 0x7a, 0x30, 0xa5, 0x52, 0x2a, 0xc9, 0x39, 0xa5, 0xc5, 0x25, 0x60,
	0x5e, 0x71, 0x66, 0x7e, 0x9e, 0x3e, 0x94, 0x1b, 0x0f, 0xe5, 0x43, 0x0c, 0x50, 0xaa, 0xe0, 0x52,
	0x0e, 0x4b, 0xcc, 0xc9, 0x4c, 0x49, 0x2c, 0x49, 0x0d, 0x49, 0x2c, 0x4a, 0x4f, 0x2d, 0x71, 0x86,
	0x28, 0x0b, 0x83, 0xa8, 0x0a, 0x4a, 0x2d, 0x2c, 0x4d, 0x2d, 0x2e, 0x11, 0x0a, 0xe4, 0xe2, 0x47,
	0xd3, 0x2f, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0x6d, 0xa4, 0xa1, 0x87, 0x70, 0x01, 0xaa, 0x85, 0x7a,
	0x68, 0x26, 0xf1, 0x25, 0xa3, 0xf0, 0x95, 0xd4, 0xb8, 0x54, 0xf0, 0xdb, 0x5c, 0x5c, 0x90, 0x9f,
	0x57, 0x9c, 0xaa, 0x94, 0xc7, 0x25, 0xe9, 0x54, 0x9a, 0x5b, 0x40, 0x37, 0x77, 0xc9, 0x70, 0x49,
	0x61, 0xb3, 0x0f, 0xe2, 0x1a, 0xa3, 0xad, 0x4c, 0x5c, 0x9c, 0xbe, 0xb0, 0x48, 0x10, 0x5a, 0xc8,
	0xc8, 0x25, 0x83, 0xcf, 0x13, 0x42, 0x76, 0x7a, 0x38, 0x23, 0x48, 0x8f, 0x88, 0x70, 0x97, 0xb2,
	0x27, 0x5b, 0x3f, 0x34, 0xf4, 0x18, 0x84, 0x9a, 0x19, 0xb9, 0x84, 0x30, 0x3d, 0x24, 0x64, 0x82,
	0xc7, 0x64, 0x9c, 0xe1, 0x2d, 0x65, 0x4a, 0xa2, 0x2e, 0x98, 0x2b, 0x9c, 0xb4, 0x4e, 0x3c, 0x94,
	0x63, 0x38, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x1b, 0x8f, 0xe4, 0x18, 0x1f, 0x3c,
	0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e, 0xe1, 0xc2, 0x63, 0x39, 0x86, 0x1b, 0x8f, 0xe5, 0x18, 0xa2,
	0x38, 0x60, 0x06, 0x25, 0xb1, 0x81, 0x93, 0xa6, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0x08, 0xa5,
	0xf0, 0xc3, 0xfe, 0x02, 0x00, 0x00,
}
