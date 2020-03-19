// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: sql/execinfrapb/processors_changefeeds.proto

package execinfrapb

/*
	Beware! This package name must not be changed, even though it doesn't match
	the Go package name, because it defines the Protobuf message names which
	can't be changed without breaking backward compatibility.
*/

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import jobspb "github.com/cockroachdb/cockroach/pkg/jobs/jobspb"
import roachpb "github.com/cockroachdb/cockroach/pkg/roachpb"
import hlc "github.com/cockroachdb/cockroach/pkg/util/hlc"

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

// ChangeAggregatorSpec is the specification for a processor that watches for
// changes in a set of spans. Each span may cross multiple ranges.
type ChangeAggregatorSpec struct {
	Watches []ChangeAggregatorSpec_Watch `protobuf:"bytes,1,rep,name=watches" json:"watches"`
	// Feed is the specification for this changefeed.
	Feed jobspb.ChangefeedDetails `protobuf:"bytes,2,opt,name=feed" json:"feed"`
}

func (m *ChangeAggregatorSpec) Reset()         { *m = ChangeAggregatorSpec{} }
func (m *ChangeAggregatorSpec) String() string { return proto.CompactTextString(m) }
func (*ChangeAggregatorSpec) ProtoMessage()    {}
func (*ChangeAggregatorSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_processors_changefeeds_e29e553eb9bf2f3d, []int{0}
}
func (m *ChangeAggregatorSpec) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ChangeAggregatorSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalTo(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (dst *ChangeAggregatorSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChangeAggregatorSpec.Merge(dst, src)
}
func (m *ChangeAggregatorSpec) XXX_Size() int {
	return m.Size()
}
func (m *ChangeAggregatorSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_ChangeAggregatorSpec.DiscardUnknown(m)
}

var xxx_messageInfo_ChangeAggregatorSpec proto.InternalMessageInfo

type ChangeAggregatorSpec_Watch struct {
	InitialResolved hlc.Timestamp `protobuf:"bytes,1,opt,name=initial_resolved,json=initialResolved" json:"initial_resolved"`
	Span            roachpb.Span  `protobuf:"bytes,2,opt,name=span" json:"span"`
}

func (m *ChangeAggregatorSpec_Watch) Reset()         { *m = ChangeAggregatorSpec_Watch{} }
func (m *ChangeAggregatorSpec_Watch) String() string { return proto.CompactTextString(m) }
func (*ChangeAggregatorSpec_Watch) ProtoMessage()    {}
func (*ChangeAggregatorSpec_Watch) Descriptor() ([]byte, []int) {
	return fileDescriptor_processors_changefeeds_e29e553eb9bf2f3d, []int{0, 0}
}
func (m *ChangeAggregatorSpec_Watch) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ChangeAggregatorSpec_Watch) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalTo(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (dst *ChangeAggregatorSpec_Watch) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChangeAggregatorSpec_Watch.Merge(dst, src)
}
func (m *ChangeAggregatorSpec_Watch) XXX_Size() int {
	return m.Size()
}
func (m *ChangeAggregatorSpec_Watch) XXX_DiscardUnknown() {
	xxx_messageInfo_ChangeAggregatorSpec_Watch.DiscardUnknown(m)
}

var xxx_messageInfo_ChangeAggregatorSpec_Watch proto.InternalMessageInfo

// ChangeFrontierSpec is the specification for a processor that receives
// span-level resolved timestamps, track them, and emits the changefeed-level
// resolved timestamp whenever it changes.
type ChangeFrontierSpec struct {
	// TrackedSpans is the entire span set being watched. Once all these spans
	// have been resolved at a certain timestamp, then it's safe to resolve the
	// changefeed at that timestamp.
	TrackedSpans []roachpb.Span `protobuf:"bytes,1,rep,name=tracked_spans,json=trackedSpans" json:"tracked_spans"`
	// Feed is the specification for this changefeed.
	Feed jobspb.ChangefeedDetails `protobuf:"bytes,2,opt,name=feed" json:"feed"`
	// JobID is the id of this changefeed in the system jobs.
	JobID int64 `protobuf:"varint,3,opt,name=job_id,json=jobId" json:"job_id"`
}

func (m *ChangeFrontierSpec) Reset()         { *m = ChangeFrontierSpec{} }
func (m *ChangeFrontierSpec) String() string { return proto.CompactTextString(m) }
func (*ChangeFrontierSpec) ProtoMessage()    {}
func (*ChangeFrontierSpec) Descriptor() ([]byte, []int) {
	return fileDescriptor_processors_changefeeds_e29e553eb9bf2f3d, []int{1}
}
func (m *ChangeFrontierSpec) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ChangeFrontierSpec) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalTo(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (dst *ChangeFrontierSpec) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ChangeFrontierSpec.Merge(dst, src)
}
func (m *ChangeFrontierSpec) XXX_Size() int {
	return m.Size()
}
func (m *ChangeFrontierSpec) XXX_DiscardUnknown() {
	xxx_messageInfo_ChangeFrontierSpec.DiscardUnknown(m)
}

var xxx_messageInfo_ChangeFrontierSpec proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ChangeAggregatorSpec)(nil), "cockroach.sql.distsqlrun.ChangeAggregatorSpec")
	proto.RegisterType((*ChangeAggregatorSpec_Watch)(nil), "cockroach.sql.distsqlrun.ChangeAggregatorSpec.Watch")
	proto.RegisterType((*ChangeFrontierSpec)(nil), "cockroach.sql.distsqlrun.ChangeFrontierSpec")
}
func (m *ChangeAggregatorSpec) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ChangeAggregatorSpec) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Watches) > 0 {
		for _, msg := range m.Watches {
			dAtA[i] = 0xa
			i++
			i = encodeVarintProcessorsChangefeeds(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	dAtA[i] = 0x12
	i++
	i = encodeVarintProcessorsChangefeeds(dAtA, i, uint64(m.Feed.Size()))
	n1, err := m.Feed.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	return i, nil
}

func (m *ChangeAggregatorSpec_Watch) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ChangeAggregatorSpec_Watch) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	dAtA[i] = 0xa
	i++
	i = encodeVarintProcessorsChangefeeds(dAtA, i, uint64(m.InitialResolved.Size()))
	n2, err := m.InitialResolved.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	dAtA[i] = 0x12
	i++
	i = encodeVarintProcessorsChangefeeds(dAtA, i, uint64(m.Span.Size()))
	n3, err := m.Span.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n3
	return i, nil
}

func (m *ChangeFrontierSpec) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ChangeFrontierSpec) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.TrackedSpans) > 0 {
		for _, msg := range m.TrackedSpans {
			dAtA[i] = 0xa
			i++
			i = encodeVarintProcessorsChangefeeds(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	dAtA[i] = 0x12
	i++
	i = encodeVarintProcessorsChangefeeds(dAtA, i, uint64(m.Feed.Size()))
	n4, err := m.Feed.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n4
	dAtA[i] = 0x18
	i++
	i = encodeVarintProcessorsChangefeeds(dAtA, i, uint64(m.JobID))
	return i, nil
}

func encodeVarintProcessorsChangefeeds(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *ChangeAggregatorSpec) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Watches) > 0 {
		for _, e := range m.Watches {
			l = e.Size()
			n += 1 + l + sovProcessorsChangefeeds(uint64(l))
		}
	}
	l = m.Feed.Size()
	n += 1 + l + sovProcessorsChangefeeds(uint64(l))
	return n
}

func (m *ChangeAggregatorSpec_Watch) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.InitialResolved.Size()
	n += 1 + l + sovProcessorsChangefeeds(uint64(l))
	l = m.Span.Size()
	n += 1 + l + sovProcessorsChangefeeds(uint64(l))
	return n
}

func (m *ChangeFrontierSpec) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.TrackedSpans) > 0 {
		for _, e := range m.TrackedSpans {
			l = e.Size()
			n += 1 + l + sovProcessorsChangefeeds(uint64(l))
		}
	}
	l = m.Feed.Size()
	n += 1 + l + sovProcessorsChangefeeds(uint64(l))
	n += 1 + sovProcessorsChangefeeds(uint64(m.JobID))
	return n
}

func sovProcessorsChangefeeds(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozProcessorsChangefeeds(x uint64) (n int) {
	return sovProcessorsChangefeeds(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ChangeAggregatorSpec) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProcessorsChangefeeds
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
			return fmt.Errorf("proto: ChangeAggregatorSpec: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ChangeAggregatorSpec: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Watches", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProcessorsChangefeeds
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
				return ErrInvalidLengthProcessorsChangefeeds
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Watches = append(m.Watches, ChangeAggregatorSpec_Watch{})
			if err := m.Watches[len(m.Watches)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Feed", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProcessorsChangefeeds
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
				return ErrInvalidLengthProcessorsChangefeeds
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Feed.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProcessorsChangefeeds(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthProcessorsChangefeeds
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
func (m *ChangeAggregatorSpec_Watch) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProcessorsChangefeeds
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
			return fmt.Errorf("proto: Watch: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Watch: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InitialResolved", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProcessorsChangefeeds
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
				return ErrInvalidLengthProcessorsChangefeeds
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.InitialResolved.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Span", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProcessorsChangefeeds
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
				return ErrInvalidLengthProcessorsChangefeeds
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Span.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProcessorsChangefeeds(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthProcessorsChangefeeds
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
func (m *ChangeFrontierSpec) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProcessorsChangefeeds
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
			return fmt.Errorf("proto: ChangeFrontierSpec: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ChangeFrontierSpec: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TrackedSpans", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProcessorsChangefeeds
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
				return ErrInvalidLengthProcessorsChangefeeds
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TrackedSpans = append(m.TrackedSpans, roachpb.Span{})
			if err := m.TrackedSpans[len(m.TrackedSpans)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Feed", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProcessorsChangefeeds
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
				return ErrInvalidLengthProcessorsChangefeeds
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Feed.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field JobID", wireType)
			}
			m.JobID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProcessorsChangefeeds
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.JobID |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipProcessorsChangefeeds(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthProcessorsChangefeeds
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
func skipProcessorsChangefeeds(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProcessorsChangefeeds
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
					return 0, ErrIntOverflowProcessorsChangefeeds
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
					return 0, ErrIntOverflowProcessorsChangefeeds
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
				return 0, ErrInvalidLengthProcessorsChangefeeds
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowProcessorsChangefeeds
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
				next, err := skipProcessorsChangefeeds(dAtA[start:])
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
	ErrInvalidLengthProcessorsChangefeeds = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProcessorsChangefeeds   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("sql/execinfrapb/processors_changefeeds.proto", fileDescriptor_processors_changefeeds_e29e553eb9bf2f3d)
}

var fileDescriptor_processors_changefeeds_e29e553eb9bf2f3d = []byte{
	// 430 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x92, 0xc1, 0x6a, 0x14, 0x31,
	0x1c, 0xc6, 0x27, 0xdd, 0x5d, 0x85, 0xac, 0x45, 0x09, 0x45, 0x87, 0x05, 0xd3, 0xa5, 0x78, 0xd8,
	0x43, 0xcd, 0x60, 0xf1, 0x05, 0x5c, 0x4b, 0xa1, 0x1e, 0x3c, 0x6c, 0x0b, 0x82, 0x97, 0x21, 0x93,
	0x49, 0x67, 0xb2, 0x4d, 0x27, 0xd9, 0x24, 0x55, 0x9f, 0xc1, 0x93, 0xaf, 0xe4, 0x45, 0xf6, 0xd8,
	0x63, 0x4f, 0x45, 0x67, 0x5f, 0x44, 0x92, 0xc9, 0xc0, 0x22, 0x8a, 0x97, 0x5e, 0x66, 0xc2, 0x3f,
	0xdf, 0xf7, 0x9b, 0xef, 0xff, 0x31, 0xf0, 0xd0, 0xae, 0x64, 0xc6, 0xbf, 0x70, 0x26, 0x9a, 0x0b,
	0x43, 0x75, 0x91, 0x69, 0xa3, 0x18, 0xb7, 0x56, 0x19, 0x9b, 0xb3, 0x9a, 0x36, 0x15, 0xbf, 0xe0,
	0xbc, 0xb4, 0x44, 0x1b, 0xe5, 0x14, 0x4a, 0x99, 0x62, 0x97, 0x46, 0x51, 0x56, 0x13, 0xbb, 0x92,
	0xa4, 0x14, 0xd6, 0xd9, 0x95, 0x34, 0xd7, 0xcd, 0xe4, 0xe9, 0x52, 0x15, 0x36, 0xf3, 0x0f, 0x5d,
	0x84, 0x57, 0xe7, 0x98, 0xa0, 0xa0, 0xd6, 0x45, 0x56, 0x52, 0x47, 0xe3, 0x2c, 0xbd, 0x76, 0x42,
	0x66, 0xb5, 0x64, 0x99, 0x13, 0x57, 0xdc, 0x3a, 0x7a, 0xa5, 0xe3, 0xcd, 0x5e, 0xa5, 0x2a, 0x15,
	0x8e, 0x99, 0x3f, 0x75, 0xd3, 0x83, 0xef, 0x3b, 0x70, 0xef, 0x6d, 0xc8, 0xf2, 0xa6, 0xaa, 0x0c,
	0xaf, 0xa8, 0x53, 0xe6, 0x4c, 0x73, 0x86, 0xce, 0xe1, 0xc3, 0xcf, 0xd4, 0xb1, 0x9a, 0xdb, 0x14,
	0x4c, 0x07, 0xb3, 0xf1, 0xd1, 0x6b, 0xf2, 0xaf, 0x80, 0xe4, 0x6f, 0x00, 0xf2, 0xc1, 0xbb, 0xe7,
	0xc3, 0xf5, 0xdd, 0x7e, 0xb2, 0xe8, 0x51, 0xe8, 0x04, 0x0e, 0xfd, 0xce, 0xe9, 0xce, 0x14, 0xcc,
	0xc6, 0x47, 0x87, 0x7f, 0x20, 0xc3, 0x6e, 0xdd, 0x9e, 0x91, 0xe9, 0xc5, 0xc7, 0xdc, 0x51, 0x21,
	0x6d, 0x44, 0x05, 0xff, 0xe4, 0x2b, 0x80, 0xa3, 0xf0, 0x01, 0xf4, 0x1e, 0x3e, 0x11, 0x8d, 0x70,
	0x82, 0xca, 0xdc, 0x70, 0xab, 0xe4, 0x27, 0x5e, 0xa6, 0x20, 0xd0, 0x9f, 0x6f, 0xd1, 0x7d, 0x2b,
	0xa4, 0x96, 0x8c, 0x9c, 0xf7, 0xad, 0x44, 0xdc, 0xe3, 0x68, 0x5e, 0x44, 0x2f, 0x7a, 0x05, 0x87,
	0x56, 0xd3, 0x26, 0x26, 0x7c, 0xb6, 0xc5, 0x88, 0x6d, 0x93, 0x33, 0x4d, 0x9b, 0x3e, 0x8c, 0x97,
	0x1e, 0xfc, 0x00, 0x10, 0x75, 0x71, 0x4f, 0x8c, 0x6a, 0x9c, 0xe0, 0x5d, 0x83, 0x73, 0xb8, 0xeb,
	0x0c, 0x65, 0x97, 0xbc, 0xcc, 0xbd, 0xac, 0xef, 0xf1, 0x3f, 0xc8, 0x47, 0xd1, 0xe3, 0x47, 0xf7,
	0xd6, 0x17, 0x7a, 0x01, 0x1f, 0x2c, 0x55, 0x91, 0x8b, 0x32, 0x1d, 0x4c, 0xc1, 0x6c, 0x30, 0xdf,
	0xf5, 0x77, 0xed, 0xdd, 0xfe, 0xe8, 0x9d, 0x2a, 0x4e, 0x8f, 0x17, 0xa3, 0xa5, 0x2a, 0x4e, 0xcb,
	0xf9, 0xcb, 0xf5, 0x2f, 0x9c, 0xac, 0x5b, 0x0c, 0x6e, 0x5a, 0x0c, 0x6e, 0x5b, 0x0c, 0x7e, 0xb6,
	0x18, 0x7c, 0xdb, 0xe0, 0xe4, 0x66, 0x83, 0x93, 0xdb, 0x0d, 0x4e, 0x3e, 0x8e, 0xb7, 0x7e, 0xe7,
	0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x4e, 0x41, 0x27, 0xac, 0xe0, 0x02, 0x00, 0x00,
}
