// Code generated by protoc-gen-go.
// source: peer/chaincode.proto
// DO NOT EDIT!

package peer

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf1 "github.com/golang/protobuf/ptypes/timestamp"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Confidentiality Levels
type ConfidentialityLevel int32

const (
	ConfidentialityLevel_PUBLIC       ConfidentialityLevel = 0
	ConfidentialityLevel_CONFIDENTIAL ConfidentialityLevel = 1
)

var ConfidentialityLevel_name = map[int32]string{
	0: "PUBLIC",
	1: "CONFIDENTIAL",
}
var ConfidentialityLevel_value = map[string]int32{
	"PUBLIC":       0,
	"CONFIDENTIAL": 1,
}

func (x ConfidentialityLevel) String() string {
	return proto.EnumName(ConfidentialityLevel_name, int32(x))
}
func (ConfidentialityLevel) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

type ChaincodeSpec_Type int32

const (
	ChaincodeSpec_UNDEFINED ChaincodeSpec_Type = 0
	ChaincodeSpec_GOLANG    ChaincodeSpec_Type = 1
	ChaincodeSpec_NODE      ChaincodeSpec_Type = 2
	ChaincodeSpec_CAR       ChaincodeSpec_Type = 3
	ChaincodeSpec_JAVA      ChaincodeSpec_Type = 4
)

var ChaincodeSpec_Type_name = map[int32]string{
	0: "UNDEFINED",
	1: "GOLANG",
	2: "NODE",
	3: "CAR",
	4: "JAVA",
}
var ChaincodeSpec_Type_value = map[string]int32{
	"UNDEFINED": 0,
	"GOLANG":    1,
	"NODE":      2,
	"CAR":       3,
	"JAVA":      4,
}

func (x ChaincodeSpec_Type) String() string {
	return proto.EnumName(ChaincodeSpec_Type_name, int32(x))
}
func (ChaincodeSpec_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{2, 0} }

type ChaincodeDeploymentSpec_ExecutionEnvironment int32

const (
	ChaincodeDeploymentSpec_DOCKER ChaincodeDeploymentSpec_ExecutionEnvironment = 0
	ChaincodeDeploymentSpec_SYSTEM ChaincodeDeploymentSpec_ExecutionEnvironment = 1
)

var ChaincodeDeploymentSpec_ExecutionEnvironment_name = map[int32]string{
	0: "DOCKER",
	1: "SYSTEM",
}
var ChaincodeDeploymentSpec_ExecutionEnvironment_value = map[string]int32{
	"DOCKER": 0,
	"SYSTEM": 1,
}

func (x ChaincodeDeploymentSpec_ExecutionEnvironment) String() string {
	return proto.EnumName(ChaincodeDeploymentSpec_ExecutionEnvironment_name, int32(x))
}
func (ChaincodeDeploymentSpec_ExecutionEnvironment) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor1, []int{3, 0}
}

type ChaincodeMessage_Type int32

const (
	ChaincodeMessage_UNDEFINED           ChaincodeMessage_Type = 0
	ChaincodeMessage_REGISTER            ChaincodeMessage_Type = 1
	ChaincodeMessage_REGISTERED          ChaincodeMessage_Type = 2
	ChaincodeMessage_INIT                ChaincodeMessage_Type = 3
	ChaincodeMessage_READY               ChaincodeMessage_Type = 4
	ChaincodeMessage_TRANSACTION         ChaincodeMessage_Type = 5
	ChaincodeMessage_COMPLETED           ChaincodeMessage_Type = 6
	ChaincodeMessage_ERROR               ChaincodeMessage_Type = 7
	ChaincodeMessage_GET_STATE           ChaincodeMessage_Type = 8
	ChaincodeMessage_PUT_STATE           ChaincodeMessage_Type = 9
	ChaincodeMessage_DEL_STATE           ChaincodeMessage_Type = 10
	ChaincodeMessage_INVOKE_CHAINCODE    ChaincodeMessage_Type = 11
	ChaincodeMessage_RESPONSE            ChaincodeMessage_Type = 13
	ChaincodeMessage_GET_STATE_BY_RANGE  ChaincodeMessage_Type = 14
	ChaincodeMessage_GET_QUERY_RESULT    ChaincodeMessage_Type = 15
	ChaincodeMessage_QUERY_STATE_NEXT    ChaincodeMessage_Type = 16
	ChaincodeMessage_QUERY_STATE_CLOSE   ChaincodeMessage_Type = 17
	ChaincodeMessage_KEEPALIVE           ChaincodeMessage_Type = 18
	ChaincodeMessage_GET_HISTORY_FOR_KEY ChaincodeMessage_Type = 19
)

var ChaincodeMessage_Type_name = map[int32]string{
	0:  "UNDEFINED",
	1:  "REGISTER",
	2:  "REGISTERED",
	3:  "INIT",
	4:  "READY",
	5:  "TRANSACTION",
	6:  "COMPLETED",
	7:  "ERROR",
	8:  "GET_STATE",
	9:  "PUT_STATE",
	10: "DEL_STATE",
	11: "INVOKE_CHAINCODE",
	13: "RESPONSE",
	14: "GET_STATE_BY_RANGE",
	15: "GET_QUERY_RESULT",
	16: "QUERY_STATE_NEXT",
	17: "QUERY_STATE_CLOSE",
	18: "KEEPALIVE",
	19: "GET_HISTORY_FOR_KEY",
}
var ChaincodeMessage_Type_value = map[string]int32{
	"UNDEFINED":           0,
	"REGISTER":            1,
	"REGISTERED":          2,
	"INIT":                3,
	"READY":               4,
	"TRANSACTION":         5,
	"COMPLETED":           6,
	"ERROR":               7,
	"GET_STATE":           8,
	"PUT_STATE":           9,
	"DEL_STATE":           10,
	"INVOKE_CHAINCODE":    11,
	"RESPONSE":            13,
	"GET_STATE_BY_RANGE":  14,
	"GET_QUERY_RESULT":    15,
	"QUERY_STATE_NEXT":    16,
	"QUERY_STATE_CLOSE":   17,
	"KEEPALIVE":           18,
	"GET_HISTORY_FOR_KEY": 19,
}

func (x ChaincodeMessage_Type) String() string {
	return proto.EnumName(ChaincodeMessage_Type_name, int32(x))
}
func (ChaincodeMessage_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor1, []int{6, 0} }

// ChaincodeID contains the path as specified by the deploy transaction
// that created it as well as the hashCode that is generated by the
// system for the path. From the user level (ie, CLI, REST API and so on)
// deploy transaction is expected to provide the path and other requests
// are expected to provide the hashCode. The other value will be ignored.
// Internally, the structure could contain both values. For instance, the
// hashCode will be set when first generated using the path
type ChaincodeID struct {
	// deploy transaction will use the path
	Path string `protobuf:"bytes,1,opt,name=path" json:"path,omitempty"`
	// all other requests will use the name (really a hashcode) generated by
	// the deploy transaction
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	// user friendly version name for the chaincode
	Version string `protobuf:"bytes,3,opt,name=version" json:"version,omitempty"`
}

func (m *ChaincodeID) Reset()                    { *m = ChaincodeID{} }
func (m *ChaincodeID) String() string            { return proto.CompactTextString(m) }
func (*ChaincodeID) ProtoMessage()               {}
func (*ChaincodeID) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

// Carries the chaincode function and its arguments.
// UnmarshalJSON in transaction.go converts the string-based REST/JSON input to
// the []byte-based current ChaincodeInput structure.
type ChaincodeInput struct {
	Args [][]byte `protobuf:"bytes,1,rep,name=args,proto3" json:"args,omitempty"`
}

func (m *ChaincodeInput) Reset()                    { *m = ChaincodeInput{} }
func (m *ChaincodeInput) String() string            { return proto.CompactTextString(m) }
func (*ChaincodeInput) ProtoMessage()               {}
func (*ChaincodeInput) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

// Carries the chaincode specification. This is the actual metadata required for
// defining a chaincode.
type ChaincodeSpec struct {
	Type        ChaincodeSpec_Type `protobuf:"varint,1,opt,name=type,enum=protos.ChaincodeSpec_Type" json:"type,omitempty"`
	ChaincodeId *ChaincodeID       `protobuf:"bytes,2,opt,name=chaincode_id,json=chaincodeId" json:"chaincode_id,omitempty"`
	Input       *ChaincodeInput    `protobuf:"bytes,3,opt,name=input" json:"input,omitempty"`
	Timeout     int32              `protobuf:"varint,4,opt,name=timeout" json:"timeout,omitempty"`
}

func (m *ChaincodeSpec) Reset()                    { *m = ChaincodeSpec{} }
func (m *ChaincodeSpec) String() string            { return proto.CompactTextString(m) }
func (*ChaincodeSpec) ProtoMessage()               {}
func (*ChaincodeSpec) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *ChaincodeSpec) GetChaincodeId() *ChaincodeID {
	if m != nil {
		return m.ChaincodeId
	}
	return nil
}

func (m *ChaincodeSpec) GetInput() *ChaincodeInput {
	if m != nil {
		return m.Input
	}
	return nil
}

// Specify the deployment of a chaincode.
// TODO: Define `codePackage`.
type ChaincodeDeploymentSpec struct {
	ChaincodeSpec *ChaincodeSpec `protobuf:"bytes,1,opt,name=chaincode_spec,json=chaincodeSpec" json:"chaincode_spec,omitempty"`
	// Controls when the chaincode becomes executable.
	EffectiveDate *google_protobuf1.Timestamp                  `protobuf:"bytes,2,opt,name=effective_date,json=effectiveDate" json:"effective_date,omitempty"`
	CodePackage   []byte                                       `protobuf:"bytes,3,opt,name=code_package,json=codePackage,proto3" json:"code_package,omitempty"`
	ExecEnv       ChaincodeDeploymentSpec_ExecutionEnvironment `protobuf:"varint,4,opt,name=exec_env,json=execEnv,enum=protos.ChaincodeDeploymentSpec_ExecutionEnvironment" json:"exec_env,omitempty"`
}

func (m *ChaincodeDeploymentSpec) Reset()                    { *m = ChaincodeDeploymentSpec{} }
func (m *ChaincodeDeploymentSpec) String() string            { return proto.CompactTextString(m) }
func (*ChaincodeDeploymentSpec) ProtoMessage()               {}
func (*ChaincodeDeploymentSpec) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *ChaincodeDeploymentSpec) GetChaincodeSpec() *ChaincodeSpec {
	if m != nil {
		return m.ChaincodeSpec
	}
	return nil
}

func (m *ChaincodeDeploymentSpec) GetEffectiveDate() *google_protobuf1.Timestamp {
	if m != nil {
		return m.EffectiveDate
	}
	return nil
}

// Carries the chaincode function and its arguments.
type ChaincodeInvocationSpec struct {
	ChaincodeSpec *ChaincodeSpec `protobuf:"bytes,1,opt,name=chaincode_spec,json=chaincodeSpec" json:"chaincode_spec,omitempty"`
	// This field can contain a user-specified ID generation algorithm
	// If supplied, this will be used to generate a ID
	// If not supplied (left empty), sha256base64 will be used
	// The algorithm consists of two parts:
	//  1, a hash function
	//  2, a decoding used to decode user (string) input to bytes
	// Currently, SHA256 with BASE64 is supported (e.g. idGenerationAlg='sha256base64')
	IdGenerationAlg string `protobuf:"bytes,2,opt,name=id_generation_alg,json=idGenerationAlg" json:"id_generation_alg,omitempty"`
}

func (m *ChaincodeInvocationSpec) Reset()                    { *m = ChaincodeInvocationSpec{} }
func (m *ChaincodeInvocationSpec) String() string            { return proto.CompactTextString(m) }
func (*ChaincodeInvocationSpec) ProtoMessage()               {}
func (*ChaincodeInvocationSpec) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{4} }

func (m *ChaincodeInvocationSpec) GetChaincodeSpec() *ChaincodeSpec {
	if m != nil {
		return m.ChaincodeSpec
	}
	return nil
}

// ChaincodeProposalContext contains proposal data that we send to the chaincode
// container shim and allow the chaincode to access through the shim interface.
type ChaincodeProposalContext struct {
	// Creator corresponds to SignatureHeader.Creator
	Creator []byte `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	// Transient corresponds to ChaincodeProposalPayload.Transient
	// TODO: The transient field is supposed to carry application-specific
	// data. They might be realted to access-control, encryption and so on.
	// To simply access to this data, replacing bytes with a map
	// is the next step to be carried.
	Transient []byte `protobuf:"bytes,2,opt,name=transient,proto3" json:"transient,omitempty"`
}

func (m *ChaincodeProposalContext) Reset()                    { *m = ChaincodeProposalContext{} }
func (m *ChaincodeProposalContext) String() string            { return proto.CompactTextString(m) }
func (*ChaincodeProposalContext) ProtoMessage()               {}
func (*ChaincodeProposalContext) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{5} }

type ChaincodeMessage struct {
	Type            ChaincodeMessage_Type       `protobuf:"varint,1,opt,name=type,enum=protos.ChaincodeMessage_Type" json:"type,omitempty"`
	Timestamp       *google_protobuf1.Timestamp `protobuf:"bytes,2,opt,name=timestamp" json:"timestamp,omitempty"`
	Payload         []byte                      `protobuf:"bytes,3,opt,name=payload,proto3" json:"payload,omitempty"`
	Txid            string                      `protobuf:"bytes,4,opt,name=txid" json:"txid,omitempty"`
	ProposalContext *ChaincodeProposalContext   `protobuf:"bytes,5,opt,name=proposal_context,json=proposalContext" json:"proposal_context,omitempty"`
	// event emmited by chaincode. Used only with Init or Invoke.
	// This event is then stored (currently)
	// with Block.NonHashData.TransactionResult
	ChaincodeEvent *ChaincodeEvent `protobuf:"bytes,6,opt,name=chaincode_event,json=chaincodeEvent" json:"chaincode_event,omitempty"`
}

func (m *ChaincodeMessage) Reset()                    { *m = ChaincodeMessage{} }
func (m *ChaincodeMessage) String() string            { return proto.CompactTextString(m) }
func (*ChaincodeMessage) ProtoMessage()               {}
func (*ChaincodeMessage) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{6} }

func (m *ChaincodeMessage) GetTimestamp() *google_protobuf1.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

func (m *ChaincodeMessage) GetProposalContext() *ChaincodeProposalContext {
	if m != nil {
		return m.ProposalContext
	}
	return nil
}

func (m *ChaincodeMessage) GetChaincodeEvent() *ChaincodeEvent {
	if m != nil {
		return m.ChaincodeEvent
	}
	return nil
}

type PutStateInfo struct {
	Key   string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *PutStateInfo) Reset()                    { *m = PutStateInfo{} }
func (m *PutStateInfo) String() string            { return proto.CompactTextString(m) }
func (*PutStateInfo) ProtoMessage()               {}
func (*PutStateInfo) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{7} }

type GetStateByRange struct {
	StartKey string `protobuf:"bytes,1,opt,name=startKey" json:"startKey,omitempty"`
	EndKey   string `protobuf:"bytes,2,opt,name=endKey" json:"endKey,omitempty"`
}

func (m *GetStateByRange) Reset()                    { *m = GetStateByRange{} }
func (m *GetStateByRange) String() string            { return proto.CompactTextString(m) }
func (*GetStateByRange) ProtoMessage()               {}
func (*GetStateByRange) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{8} }

type GetQueryResult struct {
	Query string `protobuf:"bytes,1,opt,name=query" json:"query,omitempty"`
}

func (m *GetQueryResult) Reset()                    { *m = GetQueryResult{} }
func (m *GetQueryResult) String() string            { return proto.CompactTextString(m) }
func (*GetQueryResult) ProtoMessage()               {}
func (*GetQueryResult) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{9} }

type GetHistoryForKey struct {
	Key string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
}

func (m *GetHistoryForKey) Reset()                    { *m = GetHistoryForKey{} }
func (m *GetHistoryForKey) String() string            { return proto.CompactTextString(m) }
func (*GetHistoryForKey) ProtoMessage()               {}
func (*GetHistoryForKey) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{10} }

type QueryStateNext struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *QueryStateNext) Reset()                    { *m = QueryStateNext{} }
func (m *QueryStateNext) String() string            { return proto.CompactTextString(m) }
func (*QueryStateNext) ProtoMessage()               {}
func (*QueryStateNext) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{11} }

type QueryStateClose struct {
	Id string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
}

func (m *QueryStateClose) Reset()                    { *m = QueryStateClose{} }
func (m *QueryStateClose) String() string            { return proto.CompactTextString(m) }
func (*QueryStateClose) ProtoMessage()               {}
func (*QueryStateClose) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{12} }

type QueryStateKeyValue struct {
	Key   string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
	Value []byte `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *QueryStateKeyValue) Reset()                    { *m = QueryStateKeyValue{} }
func (m *QueryStateKeyValue) String() string            { return proto.CompactTextString(m) }
func (*QueryStateKeyValue) ProtoMessage()               {}
func (*QueryStateKeyValue) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{13} }

type QueryStateResponse struct {
	KeysAndValues []*QueryStateKeyValue `protobuf:"bytes,1,rep,name=keys_and_values,json=keysAndValues" json:"keys_and_values,omitempty"`
	HasMore       bool                  `protobuf:"varint,2,opt,name=has_more,json=hasMore" json:"has_more,omitempty"`
	Id            string                `protobuf:"bytes,3,opt,name=id" json:"id,omitempty"`
}

func (m *QueryStateResponse) Reset()                    { *m = QueryStateResponse{} }
func (m *QueryStateResponse) String() string            { return proto.CompactTextString(m) }
func (*QueryStateResponse) ProtoMessage()               {}
func (*QueryStateResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{14} }

func (m *QueryStateResponse) GetKeysAndValues() []*QueryStateKeyValue {
	if m != nil {
		return m.KeysAndValues
	}
	return nil
}

func init() {
	proto.RegisterType((*ChaincodeID)(nil), "protos.ChaincodeID")
	proto.RegisterType((*ChaincodeInput)(nil), "protos.ChaincodeInput")
	proto.RegisterType((*ChaincodeSpec)(nil), "protos.ChaincodeSpec")
	proto.RegisterType((*ChaincodeDeploymentSpec)(nil), "protos.ChaincodeDeploymentSpec")
	proto.RegisterType((*ChaincodeInvocationSpec)(nil), "protos.ChaincodeInvocationSpec")
	proto.RegisterType((*ChaincodeProposalContext)(nil), "protos.ChaincodeProposalContext")
	proto.RegisterType((*ChaincodeMessage)(nil), "protos.ChaincodeMessage")
	proto.RegisterType((*PutStateInfo)(nil), "protos.PutStateInfo")
	proto.RegisterType((*GetStateByRange)(nil), "protos.GetStateByRange")
	proto.RegisterType((*GetQueryResult)(nil), "protos.GetQueryResult")
	proto.RegisterType((*GetHistoryForKey)(nil), "protos.GetHistoryForKey")
	proto.RegisterType((*QueryStateNext)(nil), "protos.QueryStateNext")
	proto.RegisterType((*QueryStateClose)(nil), "protos.QueryStateClose")
	proto.RegisterType((*QueryStateKeyValue)(nil), "protos.QueryStateKeyValue")
	proto.RegisterType((*QueryStateResponse)(nil), "protos.QueryStateResponse")
	proto.RegisterEnum("protos.ConfidentialityLevel", ConfidentialityLevel_name, ConfidentialityLevel_value)
	proto.RegisterEnum("protos.ChaincodeSpec_Type", ChaincodeSpec_Type_name, ChaincodeSpec_Type_value)
	proto.RegisterEnum("protos.ChaincodeDeploymentSpec_ExecutionEnvironment", ChaincodeDeploymentSpec_ExecutionEnvironment_name, ChaincodeDeploymentSpec_ExecutionEnvironment_value)
	proto.RegisterEnum("protos.ChaincodeMessage_Type", ChaincodeMessage_Type_name, ChaincodeMessage_Type_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion3

// Client API for ChaincodeSupport service

type ChaincodeSupportClient interface {
	Register(ctx context.Context, opts ...grpc.CallOption) (ChaincodeSupport_RegisterClient, error)
}

type chaincodeSupportClient struct {
	cc *grpc.ClientConn
}

func NewChaincodeSupportClient(cc *grpc.ClientConn) ChaincodeSupportClient {
	return &chaincodeSupportClient{cc}
}

func (c *chaincodeSupportClient) Register(ctx context.Context, opts ...grpc.CallOption) (ChaincodeSupport_RegisterClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_ChaincodeSupport_serviceDesc.Streams[0], c.cc, "/protos.ChaincodeSupport/Register", opts...)
	if err != nil {
		return nil, err
	}
	x := &chaincodeSupportRegisterClient{stream}
	return x, nil
}

type ChaincodeSupport_RegisterClient interface {
	Send(*ChaincodeMessage) error
	Recv() (*ChaincodeMessage, error)
	grpc.ClientStream
}

type chaincodeSupportRegisterClient struct {
	grpc.ClientStream
}

func (x *chaincodeSupportRegisterClient) Send(m *ChaincodeMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *chaincodeSupportRegisterClient) Recv() (*ChaincodeMessage, error) {
	m := new(ChaincodeMessage)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// Server API for ChaincodeSupport service

type ChaincodeSupportServer interface {
	Register(ChaincodeSupport_RegisterServer) error
}

func RegisterChaincodeSupportServer(s *grpc.Server, srv ChaincodeSupportServer) {
	s.RegisterService(&_ChaincodeSupport_serviceDesc, srv)
}

func _ChaincodeSupport_Register_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ChaincodeSupportServer).Register(&chaincodeSupportRegisterServer{stream})
}

type ChaincodeSupport_RegisterServer interface {
	Send(*ChaincodeMessage) error
	Recv() (*ChaincodeMessage, error)
	grpc.ServerStream
}

type chaincodeSupportRegisterServer struct {
	grpc.ServerStream
}

func (x *chaincodeSupportRegisterServer) Send(m *ChaincodeMessage) error {
	return x.ServerStream.SendMsg(m)
}

func (x *chaincodeSupportRegisterServer) Recv() (*ChaincodeMessage, error) {
	m := new(ChaincodeMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _ChaincodeSupport_serviceDesc = grpc.ServiceDesc{
	ServiceName: "protos.ChaincodeSupport",
	HandlerType: (*ChaincodeSupportServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Register",
			Handler:       _ChaincodeSupport_Register_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: fileDescriptor1,
}

func init() { proto.RegisterFile("peer/chaincode.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 1184 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xac, 0x56, 0xcf, 0x6f, 0xdb, 0xc6,
	0x12, 0x8e, 0x6c, 0xc9, 0x96, 0x46, 0xb2, 0xb4, 0xd9, 0x38, 0x8e, 0x62, 0xbc, 0x87, 0xe7, 0x10,
	0xc1, 0x83, 0x5f, 0xf0, 0x20, 0xb7, 0x6e, 0x10, 0xf4, 0x10, 0xb4, 0xa0, 0xc9, 0xb5, 0xc2, 0x4a,
	0x26, 0x95, 0x15, 0x6d, 0xc4, 0xbd, 0x10, 0x34, 0x39, 0x96, 0x89, 0xc8, 0x24, 0x4b, 0xae, 0x04,
	0xeb, 0x1c, 0xf4, 0xdf, 0xea, 0xb5, 0x7f, 0x56, 0x8b, 0x5d, 0xea, 0x87, 0x1d, 0x25, 0x40, 0x0e,
	0x3d, 0x69, 0xbf, 0xd9, 0x6f, 0x66, 0x67, 0xbe, 0x9d, 0x59, 0x11, 0x76, 0x53, 0xc4, 0xec, 0x28,
	0xb8, 0xf1, 0xa3, 0x38, 0x48, 0x42, 0xec, 0xa4, 0x59, 0x22, 0x12, 0xba, 0xa5, 0x7e, 0xf2, 0xfd,
	0xe7, 0x0f, 0x77, 0x71, 0x8a, 0xb1, 0x28, 0x28, 0xfb, 0xff, 0x19, 0x25, 0xc9, 0x68, 0x8c, 0x47,
	0x0a, 0x5d, 0x4d, 0xae, 0x8f, 0x44, 0x74, 0x8b, 0xb9, 0xf0, 0x6f, 0xd3, 0x82, 0xa0, 0x39, 0x50,
	0x37, 0x16, 0x8e, 0x96, 0x49, 0x29, 0x94, 0x53, 0x5f, 0xdc, 0xb4, 0x4b, 0x07, 0xa5, 0xc3, 0x1a,
	0x57, 0x6b, 0x69, 0x8b, 0xfd, 0x5b, 0x6c, 0x6f, 0x14, 0x36, 0xb9, 0xa6, 0x6d, 0xd8, 0x9e, 0x62,
	0x96, 0x47, 0x49, 0xdc, 0xde, 0x54, 0xe6, 0x05, 0xd4, 0x5e, 0x42, 0x73, 0x15, 0x30, 0x4e, 0x27,
	0x42, 0xfa, 0xfb, 0xd9, 0x28, 0x6f, 0x97, 0x0e, 0x36, 0x0f, 0x1b, 0x5c, 0xad, 0xb5, 0xbf, 0x4a,
	0xb0, 0xb3, 0xa4, 0x0d, 0x53, 0x0c, 0x68, 0x07, 0xca, 0x62, 0x96, 0xa2, 0x3a, 0xb9, 0x79, 0xbc,
	0x5f, 0xa4, 0x97, 0x77, 0x1e, 0x90, 0x3a, 0xee, 0x2c, 0x45, 0xae, 0x78, 0xf4, 0x0d, 0x34, 0x96,
	0x15, 0x7b, 0x51, 0xa8, 0xb2, 0xab, 0x1f, 0x3f, 0x59, 0xf3, 0xb3, 0x4c, 0x5e, 0x5f, 0x12, 0xad,
	0x90, 0xfe, 0x1f, 0x2a, 0x91, 0x4c, 0x4b, 0xe5, 0x5d, 0x3f, 0xde, 0x5b, 0x77, 0x90, 0xbb, 0xbc,
	0x20, 0xc9, 0x3a, 0xa5, 0x62, 0xc9, 0x44, 0xb4, 0xcb, 0x07, 0xa5, 0xc3, 0x0a, 0x5f, 0x40, 0xed,
	0x27, 0x28, 0xcb, 0x6c, 0xe8, 0x0e, 0xd4, 0xce, 0x6d, 0x93, 0x9d, 0x5a, 0x36, 0x33, 0xc9, 0x23,
	0x0a, 0xb0, 0xd5, 0x75, 0xfa, 0xba, 0xdd, 0x25, 0x25, 0x5a, 0x85, 0xb2, 0xed, 0x98, 0x8c, 0x6c,
	0xd0, 0x6d, 0xd8, 0x34, 0x74, 0x4e, 0x36, 0xa5, 0xe9, 0x17, 0xfd, 0x42, 0x27, 0x65, 0xed, 0x8f,
	0x0d, 0x78, 0xb6, 0x3c, 0xd3, 0xc4, 0x74, 0x9c, 0xcc, 0x6e, 0x31, 0x16, 0x4a, 0x8b, 0xb7, 0xd0,
	0x5c, 0xd5, 0x96, 0xa7, 0x18, 0x28, 0x55, 0xea, 0xc7, 0x4f, 0xbf, 0xa8, 0x0a, 0xdf, 0x09, 0x1e,
	0x28, 0xa9, 0x43, 0x13, 0xaf, 0xaf, 0x31, 0x10, 0xd1, 0x14, 0xbd, 0xd0, 0x17, 0x38, 0xd7, 0x66,
	0xbf, 0x53, 0x34, 0x43, 0x67, 0xd1, 0x0c, 0x1d, 0x77, 0xd1, 0x0c, 0x7c, 0x67, 0xe9, 0x61, 0xfa,
	0x02, 0xe9, 0x0b, 0x68, 0xa8, 0xb3, 0x53, 0x3f, 0xf8, 0xe8, 0x8f, 0x50, 0x69, 0xd5, 0xe0, 0x75,
	0x69, 0x1b, 0x14, 0x26, 0xea, 0x40, 0x15, 0xef, 0x30, 0xf0, 0x30, 0x9e, 0x2a, 0x69, 0x9a, 0xc7,
	0xaf, 0xd7, 0xb2, 0x7b, 0x58, 0x56, 0x87, 0xdd, 0x61, 0x30, 0x11, 0x51, 0x12, 0xb3, 0x78, 0x1a,
	0x65, 0x49, 0x2c, 0x37, 0xf8, 0xb6, 0x8c, 0xc2, 0xe2, 0xa9, 0xd6, 0x81, 0xdd, 0x2f, 0x11, 0xa4,
	0xa2, 0xa6, 0x63, 0xf4, 0x18, 0x2f, 0xd4, 0x1d, 0x5e, 0x0e, 0x5d, 0x76, 0x46, 0x4a, 0xda, 0xa7,
	0xd2, 0x3d, 0x01, 0xad, 0x78, 0x9a, 0x04, 0xbe, 0x74, 0xfd, 0x07, 0x04, 0x7c, 0x05, 0x8f, 0xa3,
	0xd0, 0x1b, 0x61, 0x8c, 0x99, 0x0a, 0xe9, 0xf9, 0xe3, 0xd1, 0xbc, 0xfb, 0x5b, 0x51, 0xd8, 0x5d,
	0xda, 0xf5, 0xf1, 0x48, 0xe3, 0xd0, 0x5e, 0xc6, 0x1a, 0x64, 0x49, 0x9a, 0xe4, 0xfe, 0xd8, 0x48,
	0x62, 0x81, 0x77, 0xaa, 0x79, 0x82, 0x0c, 0x7d, 0x91, 0x64, 0xea, 0xf8, 0x06, 0x5f, 0x40, 0xfa,
	0x2f, 0xa8, 0x89, 0xcc, 0x8f, 0xf3, 0x08, 0x63, 0xa1, 0x22, 0x37, 0xf8, 0xca, 0xa0, 0xfd, 0x5e,
	0x01, 0xb2, 0x0c, 0x7a, 0x86, 0x79, 0x2e, 0xf5, 0xfe, 0xfe, 0xc1, 0x7c, 0xfc, 0x7b, 0xad, 0x90,
	0x39, 0xef, 0xfe, 0x88, 0xfc, 0x08, 0xb5, 0xe5, 0xb8, 0x7f, 0x43, 0x0f, 0xac, 0xc8, 0x32, 0xf3,
	0xd4, 0x9f, 0x8d, 0x13, 0x3f, 0x9c, 0x5f, 0xfd, 0x02, 0xca, 0x61, 0x16, 0x77, 0x51, 0xa8, 0xae,
	0xbc, 0xc6, 0xd5, 0x9a, 0xf6, 0x80, 0xa4, 0xf3, 0xd2, 0xbd, 0xa0, 0xa8, 0xbd, 0x5d, 0x51, 0xc7,
	0x1d, 0xac, 0xa5, 0xf9, 0x99, 0x46, 0xbc, 0x95, 0x7e, 0x26, 0xda, 0xcf, 0xd0, 0x5a, 0x5d, 0x9d,
	0x7a, 0xca, 0xda, 0x5b, 0x5f, 0x99, 0x54, 0x26, 0x77, 0xf9, 0xea, 0xa6, 0x15, 0xd6, 0xfe, 0xdc,
	0xf8, 0xf2, 0x64, 0x36, 0xa0, 0xca, 0x59, 0xd7, 0x1a, 0xba, 0x8c, 0x93, 0x12, 0x6d, 0x02, 0x2c,
	0x10, 0x33, 0xc9, 0x86, 0x1c, 0x4c, 0xcb, 0xb6, 0x5c, 0xb2, 0x49, 0x6b, 0x50, 0xe1, 0x4c, 0x37,
	0x2f, 0x49, 0x99, 0xb6, 0xa0, 0xee, 0x72, 0xdd, 0x1e, 0xea, 0x86, 0x6b, 0x39, 0x36, 0xa9, 0xc8,
	0x90, 0x86, 0x73, 0x36, 0xe8, 0x33, 0x97, 0x99, 0x64, 0x4b, 0x52, 0x19, 0xe7, 0x0e, 0x27, 0xdb,
	0x72, 0xa7, 0xcb, 0x5c, 0x6f, 0xe8, 0xea, 0x2e, 0x23, 0x55, 0x09, 0x07, 0xe7, 0x0b, 0x58, 0x93,
	0xd0, 0x64, 0xfd, 0x39, 0x04, 0xba, 0x0b, 0xc4, 0xb2, 0x2f, 0x9c, 0x1e, 0xf3, 0x8c, 0x77, 0xba,
	0x65, 0x1b, 0xf2, 0x91, 0xa8, 0x17, 0x09, 0x0e, 0x07, 0x8e, 0x3d, 0x64, 0x64, 0x87, 0xee, 0x01,
	0x5d, 0x06, 0xf4, 0x4e, 0x2e, 0x3d, 0xae, 0xdb, 0x5d, 0x46, 0x9a, 0xd2, 0x57, 0xda, 0xdf, 0x9f,
	0x33, 0x7e, 0xe9, 0x71, 0x36, 0x3c, 0xef, 0xbb, 0xa4, 0x25, 0xad, 0x85, 0xa5, 0xe0, 0xdb, 0xec,
	0x83, 0x4b, 0x08, 0x7d, 0x0a, 0x8f, 0xef, 0x5b, 0x8d, 0xbe, 0x33, 0x64, 0xe4, 0xb1, 0xcc, 0xa6,
	0xc7, 0xd8, 0x40, 0xef, 0x5b, 0x17, 0x8c, 0x50, 0xfa, 0x0c, 0x9e, 0xc8, 0x88, 0xef, 0xac, 0xa1,
	0xeb, 0xf0, 0x4b, 0xef, 0xd4, 0xe1, 0x5e, 0x8f, 0x5d, 0x92, 0x27, 0xda, 0x1b, 0x68, 0x0c, 0x26,
	0x62, 0x28, 0x7c, 0x81, 0x56, 0x7c, 0x9d, 0x50, 0x02, 0x9b, 0x1f, 0x71, 0x36, 0xff, 0x6f, 0x90,
	0x4b, 0xba, 0x0b, 0x95, 0xa9, 0x3f, 0x9e, 0xe0, 0xbc, 0x87, 0x0b, 0xa0, 0x31, 0x68, 0x75, 0xb1,
	0xf0, 0x3b, 0x99, 0x71, 0x3f, 0x1e, 0x21, 0xdd, 0x87, 0x6a, 0x2e, 0xfc, 0x4c, 0xf4, 0x96, 0xfe,
	0x4b, 0x4c, 0xf7, 0x60, 0x0b, 0xe3, 0x50, 0xee, 0x14, 0x33, 0x36, 0x47, 0xda, 0x7f, 0xa1, 0xd9,
	0x45, 0xf1, 0x7e, 0x82, 0xd9, 0x8c, 0x63, 0x3e, 0x19, 0x0b, 0x79, 0xdc, 0x6f, 0x12, 0xce, 0x43,
	0x14, 0x40, 0x7b, 0x09, 0xa4, 0x8b, 0xe2, 0x5d, 0x94, 0x8b, 0x24, 0x9b, 0x9d, 0x26, 0x99, 0x8c,
	0xb9, 0x96, 0xaa, 0x76, 0x00, 0x4d, 0x15, 0x4a, 0xa5, 0x65, 0xcb, 0x4e, 0x6b, 0xc2, 0x46, 0x14,
	0xce, 0x29, 0x1b, 0x51, 0xa8, 0xbd, 0x80, 0xd6, 0x8a, 0x61, 0x8c, 0x93, 0x1c, 0xd7, 0x28, 0x6f,
	0x81, 0xae, 0x28, 0x3d, 0x9c, 0x5d, 0xc8, 0x7a, 0xbf, 0x59, 0x97, 0x4f, 0xa5, 0xfb, 0xee, 0x1c,
	0xf3, 0x34, 0x89, 0x73, 0xa4, 0x27, 0xd0, 0xfa, 0x88, 0xb3, 0xdc, 0xf3, 0xe3, 0xd0, 0x53, 0xc4,
	0xe2, 0xaf, 0xb2, 0xbe, 0xfa, 0x13, 0x5c, 0x3f, 0x93, 0xef, 0x48, 0x17, 0x3d, 0x0e, 0x15, 0xca,
	0xe9, 0x73, 0xa8, 0xde, 0xf8, 0xb9, 0x77, 0x9b, 0x64, 0xc5, 0x99, 0x55, 0xbe, 0x7d, 0xe3, 0xe7,
	0x67, 0x49, 0xb6, 0xa8, 0x61, 0x73, 0x51, 0xc3, 0xab, 0xd7, 0xb0, 0x6b, 0x24, 0xf1, 0x75, 0x14,
	0x62, 0x2c, 0x22, 0x7f, 0x1c, 0x89, 0x59, 0x1f, 0xa7, 0x38, 0x96, 0x6f, 0xeb, 0xe0, 0xfc, 0xa4,
	0x6f, 0x19, 0xe4, 0x11, 0x25, 0xd0, 0x30, 0x1c, 0xfb, 0xd4, 0x32, 0x99, 0xed, 0x5a, 0x7a, 0x9f,
	0x94, 0x8e, 0x3f, 0xdc, 0x7b, 0x92, 0x86, 0x93, 0x34, 0x4d, 0x32, 0x41, 0x4d, 0xa8, 0x72, 0x1c,
	0x45, 0xb9, 0xc0, 0x8c, 0xb6, 0xbf, 0xf6, 0x20, 0xed, 0x7f, 0x75, 0x47, 0x7b, 0x74, 0x58, 0xfa,
	0xae, 0x74, 0x62, 0xc0, 0x5e, 0x92, 0x8d, 0x3a, 0x37, 0xb3, 0x14, 0xb3, 0x31, 0x86, 0x23, 0xcc,
	0xe6, 0x0e, 0xbf, 0xfe, 0x6f, 0x14, 0x89, 0x9b, 0xc9, 0x55, 0x27, 0x48, 0x6e, 0x8f, 0xee, 0x6d,
	0x1f, 0x5d, 0xfb, 0x57, 0x59, 0x14, 0x14, 0xdf, 0x34, 0xf9, 0x91, 0xfc, 0xf8, 0xb9, 0x2a, 0x3e,
	0x85, 0x7e, 0xf8, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x51, 0xfa, 0x66, 0xd7, 0x29, 0x09, 0x00, 0x00,
}
