/*
Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by protoc-gen-gogo.
// source: k8s.io/kubernetes/pkg/apis/apps/v1alpha1/generated.proto
// DO NOT EDIT!

/*
	Package v1alpha1 is a generated protocol buffer package.

	It is generated from these files:
		k8s.io/kubernetes/pkg/apis/apps/v1alpha1/generated.proto

	It has these top-level messages:
		PetSet
		PetSetList
		PetSetSpec
		PetSetStatus
*/
package v1alpha1

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

import k8s_io_kubernetes_pkg_api_unversioned "k8s.io/client-go/1.4/pkg/api/unversioned"
import k8s_io_kubernetes_pkg_api_v1 "k8s.io/client-go/1.4/pkg/api/v1"

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.GoGoProtoPackageIsVersion1

func (m *PetSet) Reset()                    { *m = PetSet{} }
func (*PetSet) ProtoMessage()               {}
func (*PetSet) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{0} }

func (m *PetSetList) Reset()                    { *m = PetSetList{} }
func (*PetSetList) ProtoMessage()               {}
func (*PetSetList) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{1} }

func (m *PetSetSpec) Reset()                    { *m = PetSetSpec{} }
func (*PetSetSpec) ProtoMessage()               {}
func (*PetSetSpec) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{2} }

func (m *PetSetStatus) Reset()                    { *m = PetSetStatus{} }
func (*PetSetStatus) ProtoMessage()               {}
func (*PetSetStatus) Descriptor() ([]byte, []int) { return fileDescriptorGenerated, []int{3} }

func init() {
	proto.RegisterType((*PetSet)(nil), "k8s.io.client-go.1.4.pkg.apis.apps.v1alpha1.PetSet")
	proto.RegisterType((*PetSetList)(nil), "k8s.io.client-go.1.4.pkg.apis.apps.v1alpha1.PetSetList")
	proto.RegisterType((*PetSetSpec)(nil), "k8s.io.client-go.1.4.pkg.apis.apps.v1alpha1.PetSetSpec")
	proto.RegisterType((*PetSetStatus)(nil), "k8s.io.client-go.1.4.pkg.apis.apps.v1alpha1.PetSetStatus")
}
func (m *PetSet) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *PetSet) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0xa
	i++
	i = encodeVarintGenerated(data, i, uint64(m.ObjectMeta.Size()))
	n1, err := m.ObjectMeta.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n1
	data[i] = 0x12
	i++
	i = encodeVarintGenerated(data, i, uint64(m.Spec.Size()))
	n2, err := m.Spec.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n2
	data[i] = 0x1a
	i++
	i = encodeVarintGenerated(data, i, uint64(m.Status.Size()))
	n3, err := m.Status.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n3
	return i, nil
}

func (m *PetSetList) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *PetSetList) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	data[i] = 0xa
	i++
	i = encodeVarintGenerated(data, i, uint64(m.ListMeta.Size()))
	n4, err := m.ListMeta.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n4
	if len(m.Items) > 0 {
		for _, msg := range m.Items {
			data[i] = 0x12
			i++
			i = encodeVarintGenerated(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *PetSetSpec) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *PetSetSpec) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Replicas != nil {
		data[i] = 0x8
		i++
		i = encodeVarintGenerated(data, i, uint64(*m.Replicas))
	}
	if m.Selector != nil {
		data[i] = 0x12
		i++
		i = encodeVarintGenerated(data, i, uint64(m.Selector.Size()))
		n5, err := m.Selector.MarshalTo(data[i:])
		if err != nil {
			return 0, err
		}
		i += n5
	}
	data[i] = 0x1a
	i++
	i = encodeVarintGenerated(data, i, uint64(m.Template.Size()))
	n6, err := m.Template.MarshalTo(data[i:])
	if err != nil {
		return 0, err
	}
	i += n6
	if len(m.VolumeClaimTemplates) > 0 {
		for _, msg := range m.VolumeClaimTemplates {
			data[i] = 0x22
			i++
			i = encodeVarintGenerated(data, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(data[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	data[i] = 0x2a
	i++
	i = encodeVarintGenerated(data, i, uint64(len(m.ServiceName)))
	i += copy(data[i:], m.ServiceName)
	return i, nil
}

func (m *PetSetStatus) Marshal() (data []byte, err error) {
	size := m.Size()
	data = make([]byte, size)
	n, err := m.MarshalTo(data)
	if err != nil {
		return nil, err
	}
	return data[:n], nil
}

func (m *PetSetStatus) MarshalTo(data []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.ObservedGeneration != nil {
		data[i] = 0x8
		i++
		i = encodeVarintGenerated(data, i, uint64(*m.ObservedGeneration))
	}
	data[i] = 0x10
	i++
	i = encodeVarintGenerated(data, i, uint64(m.Replicas))
	return i, nil
}

func encodeFixed64Generated(data []byte, offset int, v uint64) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	data[offset+4] = uint8(v >> 32)
	data[offset+5] = uint8(v >> 40)
	data[offset+6] = uint8(v >> 48)
	data[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Generated(data []byte, offset int, v uint32) int {
	data[offset] = uint8(v)
	data[offset+1] = uint8(v >> 8)
	data[offset+2] = uint8(v >> 16)
	data[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintGenerated(data []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		data[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	data[offset] = uint8(v)
	return offset + 1
}
func (m *PetSet) Size() (n int) {
	var l int
	_ = l
	l = m.ObjectMeta.Size()
	n += 1 + l + sovGenerated(uint64(l))
	l = m.Spec.Size()
	n += 1 + l + sovGenerated(uint64(l))
	l = m.Status.Size()
	n += 1 + l + sovGenerated(uint64(l))
	return n
}

func (m *PetSetList) Size() (n int) {
	var l int
	_ = l
	l = m.ListMeta.Size()
	n += 1 + l + sovGenerated(uint64(l))
	if len(m.Items) > 0 {
		for _, e := range m.Items {
			l = e.Size()
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	return n
}

func (m *PetSetSpec) Size() (n int) {
	var l int
	_ = l
	if m.Replicas != nil {
		n += 1 + sovGenerated(uint64(*m.Replicas))
	}
	if m.Selector != nil {
		l = m.Selector.Size()
		n += 1 + l + sovGenerated(uint64(l))
	}
	l = m.Template.Size()
	n += 1 + l + sovGenerated(uint64(l))
	if len(m.VolumeClaimTemplates) > 0 {
		for _, e := range m.VolumeClaimTemplates {
			l = e.Size()
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	l = len(m.ServiceName)
	n += 1 + l + sovGenerated(uint64(l))
	return n
}

func (m *PetSetStatus) Size() (n int) {
	var l int
	_ = l
	if m.ObservedGeneration != nil {
		n += 1 + sovGenerated(uint64(*m.ObservedGeneration))
	}
	n += 1 + sovGenerated(uint64(m.Replicas))
	return n
}

func sovGenerated(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozGenerated(x uint64) (n int) {
	return sovGenerated(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *PetSet) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&PetSet{`,
		`ObjectMeta:` + strings.Replace(strings.Replace(this.ObjectMeta.String(), "ObjectMeta", "k8s_io_kubernetes_pkg_api_v1.ObjectMeta", 1), `&`, ``, 1) + `,`,
		`Spec:` + strings.Replace(strings.Replace(this.Spec.String(), "PetSetSpec", "PetSetSpec", 1), `&`, ``, 1) + `,`,
		`Status:` + strings.Replace(strings.Replace(this.Status.String(), "PetSetStatus", "PetSetStatus", 1), `&`, ``, 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *PetSetList) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&PetSetList{`,
		`ListMeta:` + strings.Replace(strings.Replace(this.ListMeta.String(), "ListMeta", "k8s_io_kubernetes_pkg_api_unversioned.ListMeta", 1), `&`, ``, 1) + `,`,
		`Items:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.Items), "PetSet", "PetSet", 1), `&`, ``, 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *PetSetSpec) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&PetSetSpec{`,
		`Replicas:` + valueToStringGenerated(this.Replicas) + `,`,
		`Selector:` + strings.Replace(fmt.Sprintf("%v", this.Selector), "LabelSelector", "k8s_io_kubernetes_pkg_api_unversioned.LabelSelector", 1) + `,`,
		`Template:` + strings.Replace(strings.Replace(this.Template.String(), "PodTemplateSpec", "k8s_io_kubernetes_pkg_api_v1.PodTemplateSpec", 1), `&`, ``, 1) + `,`,
		`VolumeClaimTemplates:` + strings.Replace(strings.Replace(fmt.Sprintf("%v", this.VolumeClaimTemplates), "PersistentVolumeClaim", "k8s_io_kubernetes_pkg_api_v1.PersistentVolumeClaim", 1), `&`, ``, 1) + `,`,
		`ServiceName:` + fmt.Sprintf("%v", this.ServiceName) + `,`,
		`}`,
	}, "")
	return s
}
func (this *PetSetStatus) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&PetSetStatus{`,
		`ObservedGeneration:` + valueToStringGenerated(this.ObservedGeneration) + `,`,
		`Replicas:` + fmt.Sprintf("%v", this.Replicas) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringGenerated(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *PetSet) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PetSet: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PetSet: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ObjectMeta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ObjectMeta.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Spec", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Spec.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Status.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
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
func (m *PetSetList) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PetSetList: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PetSetList: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ListMeta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ListMeta.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Items", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Items = append(m.Items, PetSet{})
			if err := m.Items[len(m.Items)-1].Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
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
func (m *PetSetSpec) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PetSetSpec: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PetSetSpec: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Replicas", wireType)
			}
			var v int32
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Replicas = &v
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Selector", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Selector == nil {
				m.Selector = &k8s_io_kubernetes_pkg_api_unversioned.LabelSelector{}
			}
			if err := m.Selector.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Template", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Template.Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VolumeClaimTemplates", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				msglen |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VolumeClaimTemplates = append(m.VolumeClaimTemplates, k8s_io_kubernetes_pkg_api_v1.PersistentVolumeClaim{})
			if err := m.VolumeClaimTemplates[len(m.VolumeClaimTemplates)-1].Unmarshal(data[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServiceName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ServiceName = string(data[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
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
func (m *PetSetStatus) Unmarshal(data []byte) error {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: PetSetStatus: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PetSetStatus: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ObservedGeneration", wireType)
			}
			var v int64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				v |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.ObservedGeneration = &v
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Replicas", wireType)
			}
			m.Replicas = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				m.Replicas |= (int32(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(data[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
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
func skipGenerated(data []byte) (n int, err error) {
	l := len(data)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenerated
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := data[iNdEx]
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
					return 0, ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if data[iNdEx-1] < 0x80 {
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
					return 0, ErrIntOverflowGenerated
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := data[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthGenerated
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowGenerated
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := data[iNdEx]
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
				next, err := skipGenerated(data[start:])
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
	ErrInvalidLengthGenerated = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenerated   = fmt.Errorf("proto: integer overflow")
)

var fileDescriptorGenerated = []byte{
	// 611 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x93, 0x4f, 0x6f, 0xd3, 0x4c,
	0x10, 0xc6, 0xeb, 0xa6, 0xa9, 0xfc, 0x6e, 0xf3, 0x22, 0xb4, 0x54, 0x28, 0x8a, 0x50, 0x8a, 0x72,
	0x8a, 0x50, 0xb3, 0x26, 0x85, 0xa2, 0x9e, 0x8d, 0x04, 0x42, 0x02, 0x5a, 0x39, 0x10, 0x21, 0x10,
	0x48, 0x6b, 0x67, 0x48, 0x97, 0xd8, 0x5e, 0xcb, 0xbb, 0xce, 0x99, 0x03, 0xdc, 0x39, 0xf3, 0x31,
	0xf8, 0x08, 0x9c, 0x72, 0xec, 0x91, 0x53, 0x05, 0xe5, 0x8b, 0xb0, 0x5e, 0xff, 0x49, 0xa8, 0x93,
	0x96, 0x1e, 0x36, 0xca, 0xee, 0xce, 0xf3, 0xdb, 0x99, 0x67, 0xc6, 0xe8, 0x60, 0x72, 0x20, 0x08,
	0xe3, 0xd6, 0x24, 0x71, 0x21, 0x0e, 0x41, 0x82, 0xb0, 0xa2, 0xc9, 0xd8, 0xa2, 0x11, 0x13, 0xea,
	0x27, 0x12, 0xd6, 0xb4, 0x4f, 0xfd, 0xe8, 0x98, 0xf6, 0xad, 0x31, 0x84, 0x10, 0x53, 0x09, 0x23,
	0x12, 0xc5, 0x5c, 0x72, 0xdc, 0xcd, 0x94, 0x64, 0xae, 0x24, 0x4a, 0x49, 0x52, 0x25, 0x49, 0x95,
	0xa4, 0x50, 0xb6, 0x7a, 0x63, 0x26, 0x8f, 0x13, 0x97, 0x78, 0x3c, 0xb0, 0xc6, 0x7c, 0xcc, 0x2d,
	0x0d, 0x70, 0x93, 0xf7, 0x7a, 0xa7, 0x37, 0xfa, 0x5f, 0x06, 0x6e, 0xed, 0xad, 0x4c, 0xc9, 0x8a,
	0x41, 0xf0, 0x24, 0xf6, 0xe0, 0x7c, 0x32, 0xad, 0xfd, 0xd5, 0x9a, 0x24, 0x9c, 0x42, 0x2c, 0x18,
	0x0f, 0x61, 0x54, 0x91, 0xed, 0xae, 0x96, 0x4d, 0x2b, 0x15, 0xb7, 0x7a, 0xcb, 0xa3, 0xe3, 0x24,
	0x94, 0x2c, 0xa8, 0xe6, 0xd4, 0x5f, 0x1e, 0x9e, 0x48, 0xe6, 0x5b, 0x2c, 0x94, 0x42, 0xc6, 0xe7,
	0x25, 0x9d, 0xaf, 0xeb, 0x68, 0xf3, 0x08, 0xe4, 0x00, 0x24, 0x7e, 0x85, 0xcc, 0x00, 0x24, 0x1d,
	0x51, 0x49, 0x9b, 0xc6, 0x6d, 0xa3, 0xbb, 0xb5, 0xd7, 0x25, 0x2b, 0x1d, 0x57, 0x5e, 0x93, 0x43,
	0xf7, 0x03, 0x78, 0xf2, 0x99, 0xd2, 0xd8, 0x78, 0x76, 0xba, 0xb3, 0x76, 0x76, 0xba, 0x83, 0xe6,
	0x67, 0x4e, 0x49, 0xc3, 0x43, 0xb4, 0x21, 0x22, 0xf0, 0x9a, 0xeb, 0x9a, 0x7a, 0x9f, 0xfc, 0x6b,
	0x1f, 0x49, 0x96, 0xd9, 0x40, 0x69, 0xed, 0x46, 0xfe, 0xc2, 0x46, 0xba, 0x73, 0x34, 0x0f, 0xbf,
	0x43, 0x9b, 0x42, 0x52, 0x99, 0x88, 0x66, 0x4d, 0x93, 0x1f, 0x5c, 0x99, 0xac, 0xd5, 0xf6, 0xb5,
	0x9c, 0xbd, 0x99, 0xed, 0x9d, 0x9c, 0xda, 0xf9, 0x6e, 0x20, 0x94, 0x05, 0x3e, 0x65, 0x42, 0xe2,
	0xb7, 0x15, 0x83, 0xac, 0x0b, 0x0c, 0x5a, 0x98, 0x02, 0x92, 0xca, 0xb5, 0x4f, 0xd7, 0xf3, 0x97,
	0xcc, 0xe2, 0x64, 0xc1, 0xa5, 0x97, 0xa8, 0xce, 0x24, 0x04, 0x42, 0xd9, 0x54, 0x53, 0xec, 0xbb,
	0x57, 0x2d, 0xc6, 0xfe, 0x3f, 0x87, 0xd7, 0x9f, 0xa4, 0x18, 0x27, 0xa3, 0x75, 0xbe, 0xd5, 0x8a,
	0x22, 0x52, 0xe7, 0x70, 0x17, 0x99, 0x31, 0x44, 0x3e, 0xf3, 0xa8, 0xd0, 0x45, 0xd4, 0xed, 0x46,
	0x9a, 0x8f, 0x93, 0x9f, 0x39, 0xe5, 0xad, 0x72, 0xd7, 0x14, 0xe0, 0xab, 0x6e, 0xf2, 0xf8, 0xf2,
	0xce, 0xfd, 0x5d, 0x2e, 0x75, 0xc1, 0x1f, 0xe4, 0xda, 0x8c, 0x5f, 0xec, 0x9c, 0x92, 0x89, 0xdf,
	0x20, 0x53, 0x25, 0x18, 0xf9, 0x6a, 0x1a, 0xf3, 0xfe, 0xf5, 0x2e, 0x9e, 0xb7, 0x23, 0x3e, 0x7a,
	0x91, 0x0b, 0xf4, 0x48, 0x94, 0x66, 0x16, 0xa7, 0x4e, 0x09, 0xc4, 0x9f, 0x0d, 0xb4, 0x3d, 0xe5,
	0x7e, 0x12, 0xc0, 0x43, 0x9f, 0xb2, 0xa0, 0x88, 0x10, 0xcd, 0x0d, 0x6d, 0xee, 0xbd, 0x4b, 0x5e,
	0x4a, 0x4b, 0x11, 0x12, 0x42, 0x39, 0x9c, 0x33, 0xec, 0x5b, 0xf9, 0x7b, 0xdb, 0xc3, 0x25, 0x60,
	0x67, 0xe9, 0x73, 0x78, 0x1f, 0x6d, 0x09, 0x88, 0xa7, 0xcc, 0x83, 0xe7, 0x34, 0x80, 0x66, 0x5d,
	0xd5, 0xf9, 0x9f, 0x7d, 0x23, 0x07, 0x6d, 0x0d, 0xe6, 0x57, 0xce, 0x62, 0x5c, 0xe7, 0x93, 0x81,
	0x1a, 0x8b, 0x23, 0x8a, 0x1f, 0x21, 0xcc, 0xdd, 0x34, 0x02, 0x46, 0x8f, 0xb3, 0x4f, 0x58, 0x59,
	0xad, 0x1b, 0x58, 0xb3, 0x6f, 0x2a, 0x14, 0x3e, 0xac, 0xdc, 0x3a, 0x4b, 0x14, 0x78, 0x77, 0xa1,
	0xfd, 0xeb, 0xba, 0xfd, 0xa5, 0x8b, 0xd5, 0x11, 0xb0, 0xef, 0xcc, 0x7e, 0xb5, 0xd7, 0x4e, 0xd4,
	0xfa, 0xa1, 0xd6, 0xc7, 0xb3, 0xb6, 0x31, 0x53, 0xeb, 0x44, 0xad, 0x9f, 0x6a, 0x7d, 0xf9, 0xdd,
	0x5e, 0x7b, 0x6d, 0x16, 0x43, 0xf8, 0x27, 0x00, 0x00, 0xff, 0xff, 0x14, 0xcf, 0x45, 0x01, 0xd8,
	0x05, 0x00, 0x00,
}
