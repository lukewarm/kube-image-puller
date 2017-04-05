// Code generated by protoc-gen-gogo.
// source: snapshot.proto
// DO NOT EDIT!

package api

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// skipping weak import gogoproto "github.com/gogo/protobuf/gogoproto"

import github_com_docker_swarmkit_api_deepcopy "github.com/docker/swarmkit/api/deepcopy"

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Snapshot_Version int32

const (
	// V0 is the initial version of the StoreSnapshot message.
	Snapshot_V0 Snapshot_Version = 0
)

var Snapshot_Version_name = map[int32]string{
	0: "V0",
}
var Snapshot_Version_value = map[string]int32{
	"V0": 0,
}

func (x Snapshot_Version) String() string {
	return proto.EnumName(Snapshot_Version_name, int32(x))
}
func (Snapshot_Version) EnumDescriptor() ([]byte, []int) { return fileDescriptorSnapshot, []int{2, 0} }

// StoreSnapshot is used to store snapshots of the store.
type StoreSnapshot struct {
	Nodes      []*Node      `protobuf:"bytes,1,rep,name=nodes" json:"nodes,omitempty"`
	Services   []*Service   `protobuf:"bytes,2,rep,name=services" json:"services,omitempty"`
	Networks   []*Network   `protobuf:"bytes,3,rep,name=networks" json:"networks,omitempty"`
	Tasks      []*Task      `protobuf:"bytes,4,rep,name=tasks" json:"tasks,omitempty"`
	Clusters   []*Cluster   `protobuf:"bytes,5,rep,name=clusters" json:"clusters,omitempty"`
	Secrets    []*Secret    `protobuf:"bytes,6,rep,name=secrets" json:"secrets,omitempty"`
	Resources  []*Resource  `protobuf:"bytes,7,rep,name=resources" json:"resources,omitempty"`
	Extensions []*Extension `protobuf:"bytes,8,rep,name=extensions" json:"extensions,omitempty"`
}

func (m *StoreSnapshot) Reset()                    { *m = StoreSnapshot{} }
func (*StoreSnapshot) ProtoMessage()               {}
func (*StoreSnapshot) Descriptor() ([]byte, []int) { return fileDescriptorSnapshot, []int{0} }

// ClusterSnapshot stores cluster membership information in snapshots.
type ClusterSnapshot struct {
	Members []*RaftMember `protobuf:"bytes,1,rep,name=members" json:"members,omitempty"`
	Removed []uint64      `protobuf:"varint,2,rep,packed,name=removed" json:"removed,omitempty"`
}

func (m *ClusterSnapshot) Reset()                    { *m = ClusterSnapshot{} }
func (*ClusterSnapshot) ProtoMessage()               {}
func (*ClusterSnapshot) Descriptor() ([]byte, []int) { return fileDescriptorSnapshot, []int{1} }

type Snapshot struct {
	Version    Snapshot_Version `protobuf:"varint,1,opt,name=version,proto3,enum=docker.swarmkit.v1.Snapshot_Version" json:"version,omitempty"`
	Membership ClusterSnapshot  `protobuf:"bytes,2,opt,name=membership" json:"membership"`
	Store      StoreSnapshot    `protobuf:"bytes,3,opt,name=store" json:"store"`
}

func (m *Snapshot) Reset()                    { *m = Snapshot{} }
func (*Snapshot) ProtoMessage()               {}
func (*Snapshot) Descriptor() ([]byte, []int) { return fileDescriptorSnapshot, []int{2} }

func init() {
	proto.RegisterType((*StoreSnapshot)(nil), "docker.swarmkit.v1.StoreSnapshot")
	proto.RegisterType((*ClusterSnapshot)(nil), "docker.swarmkit.v1.ClusterSnapshot")
	proto.RegisterType((*Snapshot)(nil), "docker.swarmkit.v1.Snapshot")
	proto.RegisterEnum("docker.swarmkit.v1.Snapshot_Version", Snapshot_Version_name, Snapshot_Version_value)
}

func (m *StoreSnapshot) Copy() *StoreSnapshot {
	if m == nil {
		return nil
	}
	o := &StoreSnapshot{}
	o.CopyFrom(m)
	return o
}

func (m *StoreSnapshot) CopyFrom(src interface{}) {

	o := src.(*StoreSnapshot)
	*m = *o
	if o.Nodes != nil {
		m.Nodes = make([]*Node, len(o.Nodes))
		for i := range m.Nodes {
			m.Nodes[i] = &Node{}
			github_com_docker_swarmkit_api_deepcopy.Copy(m.Nodes[i], o.Nodes[i])
		}
	}

	if o.Services != nil {
		m.Services = make([]*Service, len(o.Services))
		for i := range m.Services {
			m.Services[i] = &Service{}
			github_com_docker_swarmkit_api_deepcopy.Copy(m.Services[i], o.Services[i])
		}
	}

	if o.Networks != nil {
		m.Networks = make([]*Network, len(o.Networks))
		for i := range m.Networks {
			m.Networks[i] = &Network{}
			github_com_docker_swarmkit_api_deepcopy.Copy(m.Networks[i], o.Networks[i])
		}
	}

	if o.Tasks != nil {
		m.Tasks = make([]*Task, len(o.Tasks))
		for i := range m.Tasks {
			m.Tasks[i] = &Task{}
			github_com_docker_swarmkit_api_deepcopy.Copy(m.Tasks[i], o.Tasks[i])
		}
	}

	if o.Clusters != nil {
		m.Clusters = make([]*Cluster, len(o.Clusters))
		for i := range m.Clusters {
			m.Clusters[i] = &Cluster{}
			github_com_docker_swarmkit_api_deepcopy.Copy(m.Clusters[i], o.Clusters[i])
		}
	}

	if o.Secrets != nil {
		m.Secrets = make([]*Secret, len(o.Secrets))
		for i := range m.Secrets {
			m.Secrets[i] = &Secret{}
			github_com_docker_swarmkit_api_deepcopy.Copy(m.Secrets[i], o.Secrets[i])
		}
	}

	if o.Resources != nil {
		m.Resources = make([]*Resource, len(o.Resources))
		for i := range m.Resources {
			m.Resources[i] = &Resource{}
			github_com_docker_swarmkit_api_deepcopy.Copy(m.Resources[i], o.Resources[i])
		}
	}

	if o.Extensions != nil {
		m.Extensions = make([]*Extension, len(o.Extensions))
		for i := range m.Extensions {
			m.Extensions[i] = &Extension{}
			github_com_docker_swarmkit_api_deepcopy.Copy(m.Extensions[i], o.Extensions[i])
		}
	}

}

func (m *ClusterSnapshot) Copy() *ClusterSnapshot {
	if m == nil {
		return nil
	}
	o := &ClusterSnapshot{}
	o.CopyFrom(m)
	return o
}

func (m *ClusterSnapshot) CopyFrom(src interface{}) {

	o := src.(*ClusterSnapshot)
	*m = *o
	if o.Members != nil {
		m.Members = make([]*RaftMember, len(o.Members))
		for i := range m.Members {
			m.Members[i] = &RaftMember{}
			github_com_docker_swarmkit_api_deepcopy.Copy(m.Members[i], o.Members[i])
		}
	}

	if o.Removed != nil {
		m.Removed = make([]uint64, len(o.Removed))
		copy(m.Removed, o.Removed)
	}

}

func (m *Snapshot) Copy() *Snapshot {
	if m == nil {
		return nil
	}
	o := &Snapshot{}
	o.CopyFrom(m)
	return o
}

func (m *Snapshot) CopyFrom(src interface{}) {

	o := src.(*Snapshot)
	*m = *o
	github_com_docker_swarmkit_api_deepcopy.Copy(&m.Membership, &o.Membership)
	github_com_docker_swarmkit_api_deepcopy.Copy(&m.Store, &o.Store)
}

func (m *StoreSnapshot) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *StoreSnapshot) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Nodes) > 0 {
		for _, msg := range m.Nodes {
			dAtA[i] = 0xa
			i++
			i = encodeVarintSnapshot(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.Services) > 0 {
		for _, msg := range m.Services {
			dAtA[i] = 0x12
			i++
			i = encodeVarintSnapshot(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.Networks) > 0 {
		for _, msg := range m.Networks {
			dAtA[i] = 0x1a
			i++
			i = encodeVarintSnapshot(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.Tasks) > 0 {
		for _, msg := range m.Tasks {
			dAtA[i] = 0x22
			i++
			i = encodeVarintSnapshot(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.Clusters) > 0 {
		for _, msg := range m.Clusters {
			dAtA[i] = 0x2a
			i++
			i = encodeVarintSnapshot(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.Secrets) > 0 {
		for _, msg := range m.Secrets {
			dAtA[i] = 0x32
			i++
			i = encodeVarintSnapshot(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.Resources) > 0 {
		for _, msg := range m.Resources {
			dAtA[i] = 0x3a
			i++
			i = encodeVarintSnapshot(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.Extensions) > 0 {
		for _, msg := range m.Extensions {
			dAtA[i] = 0x42
			i++
			i = encodeVarintSnapshot(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	return i, nil
}

func (m *ClusterSnapshot) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClusterSnapshot) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Members) > 0 {
		for _, msg := range m.Members {
			dAtA[i] = 0xa
			i++
			i = encodeVarintSnapshot(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.Removed) > 0 {
		dAtA2 := make([]byte, len(m.Removed)*10)
		var j1 int
		for _, num := range m.Removed {
			for num >= 1<<7 {
				dAtA2[j1] = uint8(uint64(num)&0x7f | 0x80)
				num >>= 7
				j1++
			}
			dAtA2[j1] = uint8(num)
			j1++
		}
		dAtA[i] = 0x12
		i++
		i = encodeVarintSnapshot(dAtA, i, uint64(j1))
		i += copy(dAtA[i:], dAtA2[:j1])
	}
	return i, nil
}

func (m *Snapshot) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Snapshot) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Version != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintSnapshot(dAtA, i, uint64(m.Version))
	}
	dAtA[i] = 0x12
	i++
	i = encodeVarintSnapshot(dAtA, i, uint64(m.Membership.Size()))
	n3, err := m.Membership.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n3
	dAtA[i] = 0x1a
	i++
	i = encodeVarintSnapshot(dAtA, i, uint64(m.Store.Size()))
	n4, err := m.Store.MarshalTo(dAtA[i:])
	if err != nil {
		return 0, err
	}
	i += n4
	return i, nil
}

func encodeFixed64Snapshot(dAtA []byte, offset int, v uint64) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	dAtA[offset+4] = uint8(v >> 32)
	dAtA[offset+5] = uint8(v >> 40)
	dAtA[offset+6] = uint8(v >> 48)
	dAtA[offset+7] = uint8(v >> 56)
	return offset + 8
}
func encodeFixed32Snapshot(dAtA []byte, offset int, v uint32) int {
	dAtA[offset] = uint8(v)
	dAtA[offset+1] = uint8(v >> 8)
	dAtA[offset+2] = uint8(v >> 16)
	dAtA[offset+3] = uint8(v >> 24)
	return offset + 4
}
func encodeVarintSnapshot(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}

func (m *StoreSnapshot) Size() (n int) {
	var l int
	_ = l
	if len(m.Nodes) > 0 {
		for _, e := range m.Nodes {
			l = e.Size()
			n += 1 + l + sovSnapshot(uint64(l))
		}
	}
	if len(m.Services) > 0 {
		for _, e := range m.Services {
			l = e.Size()
			n += 1 + l + sovSnapshot(uint64(l))
		}
	}
	if len(m.Networks) > 0 {
		for _, e := range m.Networks {
			l = e.Size()
			n += 1 + l + sovSnapshot(uint64(l))
		}
	}
	if len(m.Tasks) > 0 {
		for _, e := range m.Tasks {
			l = e.Size()
			n += 1 + l + sovSnapshot(uint64(l))
		}
	}
	if len(m.Clusters) > 0 {
		for _, e := range m.Clusters {
			l = e.Size()
			n += 1 + l + sovSnapshot(uint64(l))
		}
	}
	if len(m.Secrets) > 0 {
		for _, e := range m.Secrets {
			l = e.Size()
			n += 1 + l + sovSnapshot(uint64(l))
		}
	}
	if len(m.Resources) > 0 {
		for _, e := range m.Resources {
			l = e.Size()
			n += 1 + l + sovSnapshot(uint64(l))
		}
	}
	if len(m.Extensions) > 0 {
		for _, e := range m.Extensions {
			l = e.Size()
			n += 1 + l + sovSnapshot(uint64(l))
		}
	}
	return n
}

func (m *ClusterSnapshot) Size() (n int) {
	var l int
	_ = l
	if len(m.Members) > 0 {
		for _, e := range m.Members {
			l = e.Size()
			n += 1 + l + sovSnapshot(uint64(l))
		}
	}
	if len(m.Removed) > 0 {
		l = 0
		for _, e := range m.Removed {
			l += sovSnapshot(uint64(e))
		}
		n += 1 + sovSnapshot(uint64(l)) + l
	}
	return n
}

func (m *Snapshot) Size() (n int) {
	var l int
	_ = l
	if m.Version != 0 {
		n += 1 + sovSnapshot(uint64(m.Version))
	}
	l = m.Membership.Size()
	n += 1 + l + sovSnapshot(uint64(l))
	l = m.Store.Size()
	n += 1 + l + sovSnapshot(uint64(l))
	return n
}

func sovSnapshot(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozSnapshot(x uint64) (n int) {
	return sovSnapshot(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *StoreSnapshot) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&StoreSnapshot{`,
		`Nodes:` + strings.Replace(fmt.Sprintf("%v", this.Nodes), "Node", "Node", 1) + `,`,
		`Services:` + strings.Replace(fmt.Sprintf("%v", this.Services), "Service", "Service", 1) + `,`,
		`Networks:` + strings.Replace(fmt.Sprintf("%v", this.Networks), "Network", "Network", 1) + `,`,
		`Tasks:` + strings.Replace(fmt.Sprintf("%v", this.Tasks), "Task", "Task", 1) + `,`,
		`Clusters:` + strings.Replace(fmt.Sprintf("%v", this.Clusters), "Cluster", "Cluster", 1) + `,`,
		`Secrets:` + strings.Replace(fmt.Sprintf("%v", this.Secrets), "Secret", "Secret", 1) + `,`,
		`Resources:` + strings.Replace(fmt.Sprintf("%v", this.Resources), "Resource", "Resource", 1) + `,`,
		`Extensions:` + strings.Replace(fmt.Sprintf("%v", this.Extensions), "Extension", "Extension", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *ClusterSnapshot) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&ClusterSnapshot{`,
		`Members:` + strings.Replace(fmt.Sprintf("%v", this.Members), "RaftMember", "RaftMember", 1) + `,`,
		`Removed:` + fmt.Sprintf("%v", this.Removed) + `,`,
		`}`,
	}, "")
	return s
}
func (this *Snapshot) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&Snapshot{`,
		`Version:` + fmt.Sprintf("%v", this.Version) + `,`,
		`Membership:` + strings.Replace(strings.Replace(this.Membership.String(), "ClusterSnapshot", "ClusterSnapshot", 1), `&`, ``, 1) + `,`,
		`Store:` + strings.Replace(strings.Replace(this.Store.String(), "StoreSnapshot", "StoreSnapshot", 1), `&`, ``, 1) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringSnapshot(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *StoreSnapshot) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSnapshot
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
			return fmt.Errorf("proto: StoreSnapshot: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: StoreSnapshot: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Nodes", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSnapshot
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
				return ErrInvalidLengthSnapshot
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Nodes = append(m.Nodes, &Node{})
			if err := m.Nodes[len(m.Nodes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Services", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSnapshot
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
				return ErrInvalidLengthSnapshot
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Services = append(m.Services, &Service{})
			if err := m.Services[len(m.Services)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Networks", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSnapshot
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
				return ErrInvalidLengthSnapshot
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Networks = append(m.Networks, &Network{})
			if err := m.Networks[len(m.Networks)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tasks", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSnapshot
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
				return ErrInvalidLengthSnapshot
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Tasks = append(m.Tasks, &Task{})
			if err := m.Tasks[len(m.Tasks)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Clusters", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSnapshot
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
				return ErrInvalidLengthSnapshot
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Clusters = append(m.Clusters, &Cluster{})
			if err := m.Clusters[len(m.Clusters)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Secrets", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSnapshot
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
				return ErrInvalidLengthSnapshot
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Secrets = append(m.Secrets, &Secret{})
			if err := m.Secrets[len(m.Secrets)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Resources", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSnapshot
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
				return ErrInvalidLengthSnapshot
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Resources = append(m.Resources, &Resource{})
			if err := m.Resources[len(m.Resources)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Extensions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSnapshot
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
				return ErrInvalidLengthSnapshot
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Extensions = append(m.Extensions, &Extension{})
			if err := m.Extensions[len(m.Extensions)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSnapshot(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSnapshot
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
func (m *ClusterSnapshot) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSnapshot
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
			return fmt.Errorf("proto: ClusterSnapshot: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClusterSnapshot: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Members", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSnapshot
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
				return ErrInvalidLengthSnapshot
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Members = append(m.Members, &RaftMember{})
			if err := m.Members[len(m.Members)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType == 2 {
				var packedLen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowSnapshot
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					packedLen |= (int(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if packedLen < 0 {
					return ErrInvalidLengthSnapshot
				}
				postIndex := iNdEx + packedLen
				if postIndex > l {
					return io.ErrUnexpectedEOF
				}
				for iNdEx < postIndex {
					var v uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowSnapshot
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						v |= (uint64(b) & 0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					m.Removed = append(m.Removed, v)
				}
			} else if wireType == 0 {
				var v uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowSnapshot
					}
					if iNdEx >= l {
						return io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					v |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				m.Removed = append(m.Removed, v)
			} else {
				return fmt.Errorf("proto: wrong wireType = %d for field Removed", wireType)
			}
		default:
			iNdEx = preIndex
			skippy, err := skipSnapshot(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSnapshot
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
func (m *Snapshot) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowSnapshot
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
			return fmt.Errorf("proto: Snapshot: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Snapshot: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Version", wireType)
			}
			m.Version = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSnapshot
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Version |= (Snapshot_Version(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Membership", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSnapshot
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
				return ErrInvalidLengthSnapshot
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Membership.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Store", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowSnapshot
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
				return ErrInvalidLengthSnapshot
			}
			postIndex := iNdEx + msglen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Store.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipSnapshot(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthSnapshot
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
func skipSnapshot(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowSnapshot
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
					return 0, ErrIntOverflowSnapshot
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
					return 0, ErrIntOverflowSnapshot
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
				return 0, ErrInvalidLengthSnapshot
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowSnapshot
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
				next, err := skipSnapshot(dAtA[start:])
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
	ErrInvalidLengthSnapshot = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowSnapshot   = fmt.Errorf("proto: integer overflow")
)

func init() { proto.RegisterFile("snapshot.proto", fileDescriptorSnapshot) }

var fileDescriptorSnapshot = []byte{
	// 449 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x93, 0xbf, 0x6f, 0xd3, 0x40,
	0x14, 0xc7, 0x73, 0xf9, 0xe5, 0xf0, 0xaa, 0x16, 0x38, 0x31, 0x9c, 0x02, 0x98, 0x10, 0x18, 0x32,
	0x19, 0x08, 0x48, 0x20, 0xa4, 0x32, 0x14, 0x31, 0x30, 0xd0, 0xe1, 0x82, 0x2a, 0x56, 0xc7, 0x79,
	0x6d, 0x83, 0x89, 0x2f, 0xba, 0x77, 0x75, 0x19, 0xf9, 0xf3, 0x22, 0xb1, 0x30, 0x32, 0x21, 0xea,
	0x85, 0x7f, 0x03, 0xdd, 0x9d, 0x6d, 0x22, 0xe1, 0x74, 0xbb, 0x44, 0x9f, 0xcf, 0xfb, 0xbe, 0xbb,
	0xf7, 0x0c, 0x07, 0x94, 0xc5, 0x6b, 0x3a, 0x57, 0x26, 0x5a, 0x6b, 0x65, 0x14, 0xe7, 0x0b, 0x95,
	0xa4, 0xa8, 0x23, 0xba, 0x8c, 0xf5, 0x2a, 0x5d, 0x9a, 0x28, 0x7f, 0x36, 0xdc, 0x57, 0xf3, 0xcf,
	0x98, 0x18, 0xf2, 0xc8, 0x10, 0x74, 0x7c, 0x5a, 0xe2, 0xc3, 0x3b, 0x67, 0xea, 0x4c, 0xb9, 0xe3,
	0x13, 0x7b, 0xf2, 0xff, 0x8e, 0xbf, 0x77, 0x60, 0x7f, 0x66, 0x94, 0xc6, 0x59, 0x59, 0x9c, 0x47,
	0xd0, 0xcb, 0xd4, 0x02, 0x49, 0xb0, 0x51, 0x67, 0xb2, 0x37, 0x15, 0xd1, 0xff, 0x31, 0xd1, 0xb1,
	0x5a, 0xa0, 0xf4, 0x18, 0x7f, 0x09, 0x03, 0x42, 0x9d, 0x2f, 0x13, 0x24, 0xd1, 0x76, 0xca, 0xdd,
	0x26, 0x65, 0xe6, 0x19, 0x59, 0xc3, 0x56, 0xcc, 0xd0, 0x5c, 0x2a, 0x9d, 0x92, 0xe8, 0xec, 0x16,
	0x8f, 0x3d, 0x23, 0x6b, 0xd8, 0x76, 0x68, 0x62, 0x4a, 0x49, 0x74, 0x77, 0x77, 0xf8, 0x31, 0xa6,
	0x54, 0x7a, 0xcc, 0x06, 0x25, 0x5f, 0x2e, 0xc8, 0xa0, 0x26, 0xd1, 0xdb, 0x1d, 0xf4, 0xd6, 0x33,
	0xb2, 0x86, 0xf9, 0x0b, 0x08, 0x08, 0x13, 0x8d, 0x86, 0x44, 0xdf, 0x79, 0xc3, 0xe6, 0x9b, 0x59,
	0x44, 0x56, 0x28, 0x7f, 0x0d, 0x37, 0x34, 0x92, 0xba, 0xd0, 0xf6, 0x45, 0x02, 0xe7, 0xdd, 0x6b,
	0xf2, 0x64, 0x09, 0xc9, 0x7f, 0x38, 0x3f, 0x04, 0xc0, 0xaf, 0x06, 0x33, 0x5a, 0xaa, 0x8c, 0xc4,
	0xc0, 0xc9, 0xf7, 0x9b, 0xe4, 0x77, 0x15, 0x25, 0xb7, 0x84, 0x31, 0xc2, 0xcd, 0xf2, 0x16, 0xf5,
	0x38, 0x5f, 0x41, 0xb0, 0xc2, 0xd5, 0xdc, 0xde, 0xdd, 0x0f, 0x34, 0x6c, 0xec, 0x25, 0x3e, 0x35,
	0x1f, 0x1c, 0x26, 0x2b, 0x9c, 0x0b, 0x08, 0x34, 0xae, 0x54, 0x8e, 0x0b, 0x37, 0xd7, 0xae, 0xac,
	0x7e, 0x8e, 0xff, 0x30, 0x18, 0xd4, 0x01, 0x6f, 0x20, 0xc8, 0x51, 0xdb, 0x7c, 0xc1, 0x46, 0x6c,
	0x72, 0x30, 0x7d, 0xdc, 0xf8, 0x48, 0xd5, 0xee, 0x9e, 0x78, 0x56, 0x56, 0x12, 0x7f, 0x0f, 0x50,
	0x26, 0x9e, 0x2f, 0xd7, 0xa2, 0x3d, 0x62, 0x93, 0xbd, 0xe9, 0xa3, 0x6b, 0xe6, 0x53, 0x55, 0x3a,
	0xea, 0x6e, 0x7e, 0x3d, 0x68, 0xc9, 0x2d, 0x99, 0x1f, 0x42, 0x8f, 0xec, 0x2e, 0x8b, 0x8e, 0xab,
	0xf2, 0xb0, 0xb1, 0x91, 0xed, 0x65, 0x2f, 0x6b, 0x78, 0x6b, 0x7c, 0x1b, 0x82, 0xb2, 0x3b, 0xde,
	0x87, 0xf6, 0xc9, 0xd3, 0x5b, 0xad, 0x23, 0xb1, 0xb9, 0x0a, 0x5b, 0x3f, 0xaf, 0xc2, 0xd6, 0xb7,
	0x22, 0x64, 0x9b, 0x22, 0x64, 0x3f, 0x8a, 0x90, 0xfd, 0x2e, 0x42, 0xf6, 0xa9, 0x3d, 0xef, 0xbb,
	0x2f, 0xe8, 0xf9, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0x91, 0xfe, 0xed, 0x18, 0x98, 0x03, 0x00,
	0x00,
}