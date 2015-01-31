// Code generated by protoc-gen-go.
// source: disk/client.proto
// DO NOT EDIT!

/*
Package disk is a generated protocol buffer package.

It is generated from these files:
	disk/client.proto

It has these top-level messages:
	Header
	Contact
	RatchetState
	Inbox
	Outbox
	Draft
	State
*/
package disk

import proto "github.com/golang/protobuf/proto"
import math "math"
import protos "github.com/agl/pond/protos"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

// Header is placed at the beginning of the state file and is *unencrypted and
// unauthenticated*. Its purpose is only to describe how to decrypt the
// remainder of the file.
type Header struct {
	// nonce_smear_copies contains the number of copies of the nonce that
	// follow the header. Each copy of the nonce is different and, XORed
	// together, they result in the real nonce. The intent is that this may
	// make recovery of old state files more difficult on HDDs.
	NonceSmearCopies *int32 `protobuf:"varint,1,opt,name=nonce_smear_copies,def=1365" json:"nonce_smear_copies,omitempty"`
	// kdf_salt contains the salt for the KDF function.
	KdfSalt  []byte         `protobuf:"bytes,2,opt,name=kdf_salt" json:"kdf_salt,omitempty"`
	Scrypt   *Header_SCrypt `protobuf:"bytes,3,opt,name=scrypt" json:"scrypt,omitempty"`
	TpmNvram *Header_TPM    `protobuf:"bytes,4,opt,name=tpm_nvram" json:"tpm_nvram,omitempty"`
	// no_erasure_storage exists to signal that there is no erasure storage
	// for this state file, as opposed to the state file using a method
	// that isn't recognised by the client.
	NoErasureStorage *bool  `protobuf:"varint,5,opt,name=no_erasure_storage" json:"no_erasure_storage,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Header) Reset()         { *m = Header{} }
func (m *Header) String() string { return proto.CompactTextString(m) }
func (*Header) ProtoMessage()    {}

const Default_Header_NonceSmearCopies int32 = 1365

func (m *Header) GetNonceSmearCopies() int32 {
	if m != nil && m.NonceSmearCopies != nil {
		return *m.NonceSmearCopies
	}
	return Default_Header_NonceSmearCopies
}

func (m *Header) GetKdfSalt() []byte {
	if m != nil {
		return m.KdfSalt
	}
	return nil
}

func (m *Header) GetScrypt() *Header_SCrypt {
	if m != nil {
		return m.Scrypt
	}
	return nil
}

func (m *Header) GetTpmNvram() *Header_TPM {
	if m != nil {
		return m.TpmNvram
	}
	return nil
}

func (m *Header) GetNoErasureStorage() bool {
	if m != nil && m.NoErasureStorage != nil {
		return *m.NoErasureStorage
	}
	return false
}

type Header_SCrypt struct {
	N                *int32 `protobuf:"varint,2,opt,def=32768" json:"N,omitempty"`
	R                *int32 `protobuf:"varint,3,opt,name=r,def=16" json:"r,omitempty"`
	P                *int32 `protobuf:"varint,4,opt,name=p,def=1" json:"p,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Header_SCrypt) Reset()         { *m = Header_SCrypt{} }
func (m *Header_SCrypt) String() string { return proto.CompactTextString(m) }
func (*Header_SCrypt) ProtoMessage()    {}

const Default_Header_SCrypt_N int32 = 32768
const Default_Header_SCrypt_R int32 = 16
const Default_Header_SCrypt_P int32 = 1

func (m *Header_SCrypt) GetN() int32 {
	if m != nil && m.N != nil {
		return *m.N
	}
	return Default_Header_SCrypt_N
}

func (m *Header_SCrypt) GetR() int32 {
	if m != nil && m.R != nil {
		return *m.R
	}
	return Default_Header_SCrypt_R
}

func (m *Header_SCrypt) GetP() int32 {
	if m != nil && m.P != nil {
		return *m.P
	}
	return Default_Header_SCrypt_P
}

// TPM contains information about an erasure key stored in TPM NVRAM.
type Header_TPM struct {
	Index            *uint32 `protobuf:"varint,1,req,name=index" json:"index,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Header_TPM) Reset()         { *m = Header_TPM{} }
func (m *Header_TPM) String() string { return proto.CompactTextString(m) }
func (*Header_TPM) ProtoMessage()    {}

func (m *Header_TPM) GetIndex() uint32 {
	if m != nil && m.Index != nil {
		return *m.Index
	}
	return 0
}

type Contact struct {
	Id                  *uint64                `protobuf:"fixed64,1,req,name=id" json:"id,omitempty"`
	Name                *string                `protobuf:"bytes,2,req,name=name" json:"name,omitempty"`
	GroupKey            []byte                 `protobuf:"bytes,3,req,name=group_key" json:"group_key,omitempty"`
	SupportedVersion    *int32                 `protobuf:"varint,16,opt,name=supported_version" json:"supported_version,omitempty"`
	KeyExchangeBytes    []byte                 `protobuf:"bytes,4,opt,name=key_exchange_bytes" json:"key_exchange_bytes,omitempty"`
	PandaKeyExchange    []byte                 `protobuf:"bytes,18,opt,name=panda_key_exchange" json:"panda_key_exchange,omitempty"`
	PandaError          *string                `protobuf:"bytes,19,opt,name=panda_error" json:"panda_error,omitempty"`
	TheirGroup          []byte                 `protobuf:"bytes,5,opt,name=their_group" json:"their_group,omitempty"`
	MyGroupKey          []byte                 `protobuf:"bytes,6,opt,name=my_group_key" json:"my_group_key,omitempty"`
	Generation          *uint32                `protobuf:"varint,7,opt,name=generation" json:"generation,omitempty"`
	TheirServer         *string                `protobuf:"bytes,8,opt,name=their_server" json:"their_server,omitempty"`
	TheirPub            []byte                 `protobuf:"bytes,9,opt,name=their_pub" json:"their_pub,omitempty"`
	TheirIdentityPublic []byte                 `protobuf:"bytes,10,opt,name=their_identity_public" json:"their_identity_public,omitempty"`
	RevokedUs           *bool                  `protobuf:"varint,21,opt,name=revoked_us" json:"revoked_us,omitempty"`
	LastPrivate         []byte                 `protobuf:"bytes,11,opt,name=last_private" json:"last_private,omitempty"`
	CurrentPrivate      []byte                 `protobuf:"bytes,12,opt,name=current_private" json:"current_private,omitempty"`
	TheirLastPublic     []byte                 `protobuf:"bytes,13,opt,name=their_last_public" json:"their_last_public,omitempty"`
	TheirCurrentPublic  []byte                 `protobuf:"bytes,14,opt,name=their_current_public" json:"their_current_public,omitempty"`
	Ratchet             *RatchetState          `protobuf:"bytes,20,opt,name=ratchet" json:"ratchet,omitempty"`
	PreviousTags        []*Contact_PreviousTag `protobuf:"bytes,17,rep,name=previous_tags" json:"previous_tags,omitempty"`
	Events              []*Contact_Event       `protobuf:"bytes,22,rep,name=events" json:"events,omitempty"`
	IsPending           *bool                  `protobuf:"varint,15,opt,name=is_pending,def=0" json:"is_pending,omitempty"`
	IntroducedBy        *uint64                `protobuf:"fixed64,23,opt,name=introduced_by" json:"introduced_by,omitempty"`
	ReintroducedBy          []uint64               `protobuf:"fixed64,24,rep,name=reintroduced_by" json:"reintroduced_by,omitempty"`
	IntroducedTo        []uint64               `protobuf:"fixed64,25,rep,name=introduced_to" json:"introduced_to,omitempty"`
	XXX_unrecognized    []byte                 `json:"-"`
}

func (m *Contact) Reset()         { *m = Contact{} }
func (m *Contact) String() string { return proto.CompactTextString(m) }
func (*Contact) ProtoMessage()    {}

const Default_Contact_IsPending bool = false

func (m *Contact) GetId() uint64 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *Contact) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Contact) GetGroupKey() []byte {
	if m != nil {
		return m.GroupKey
	}
	return nil
}

func (m *Contact) GetSupportedVersion() int32 {
	if m != nil && m.SupportedVersion != nil {
		return *m.SupportedVersion
	}
	return 0
}

func (m *Contact) GetKeyExchangeBytes() []byte {
	if m != nil {
		return m.KeyExchangeBytes
	}
	return nil
}

func (m *Contact) GetPandaKeyExchange() []byte {
	if m != nil {
		return m.PandaKeyExchange
	}
	return nil
}

func (m *Contact) GetPandaError() string {
	if m != nil && m.PandaError != nil {
		return *m.PandaError
	}
	return ""
}

func (m *Contact) GetTheirGroup() []byte {
	if m != nil {
		return m.TheirGroup
	}
	return nil
}

func (m *Contact) GetMyGroupKey() []byte {
	if m != nil {
		return m.MyGroupKey
	}
	return nil
}

func (m *Contact) GetGeneration() uint32 {
	if m != nil && m.Generation != nil {
		return *m.Generation
	}
	return 0
}

func (m *Contact) GetTheirServer() string {
	if m != nil && m.TheirServer != nil {
		return *m.TheirServer
	}
	return ""
}

func (m *Contact) GetTheirPub() []byte {
	if m != nil {
		return m.TheirPub
	}
	return nil
}

func (m *Contact) GetTheirIdentityPublic() []byte {
	if m != nil {
		return m.TheirIdentityPublic
	}
	return nil
}

func (m *Contact) GetRevokedUs() bool {
	if m != nil && m.RevokedUs != nil {
		return *m.RevokedUs
	}
	return false
}

func (m *Contact) GetLastPrivate() []byte {
	if m != nil {
		return m.LastPrivate
	}
	return nil
}

func (m *Contact) GetCurrentPrivate() []byte {
	if m != nil {
		return m.CurrentPrivate
	}
	return nil
}

func (m *Contact) GetTheirLastPublic() []byte {
	if m != nil {
		return m.TheirLastPublic
	}
	return nil
}

func (m *Contact) GetTheirCurrentPublic() []byte {
	if m != nil {
		return m.TheirCurrentPublic
	}
	return nil
}

func (m *Contact) GetRatchet() *RatchetState {
	if m != nil {
		return m.Ratchet
	}
	return nil
}

func (m *Contact) GetPreviousTags() []*Contact_PreviousTag {
	if m != nil {
		return m.PreviousTags
	}
	return nil
}

func (m *Contact) GetEvents() []*Contact_Event {
	if m != nil {
		return m.Events
	}
	return nil
}

func (m *Contact) GetIsPending() bool {
	if m != nil && m.IsPending != nil {
		return *m.IsPending
	}
	return Default_Contact_IsPending
}

func (m *Contact) GetIntroducedBy() uint64 {
	if m != nil && m.IntroducedBy != nil {
		return *m.IntroducedBy
	}
	return 0
}

func (m *Contact) GetReintroducedBy() []uint64 {
	if m != nil {
		return m.ReintroducedBy
	}
	return nil
}

func (m *Contact) GetIntroducedTo() []uint64 {
	if m != nil {
		return m.IntroducedTo
	}
	return nil
}

type Contact_PreviousTag struct {
	Tag              []byte `protobuf:"bytes,1,req,name=tag" json:"tag,omitempty"`
	Expired          *int64 `protobuf:"varint,2,req,name=expired" json:"expired,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *Contact_PreviousTag) Reset()         { *m = Contact_PreviousTag{} }
func (m *Contact_PreviousTag) String() string { return proto.CompactTextString(m) }
func (*Contact_PreviousTag) ProtoMessage()    {}

func (m *Contact_PreviousTag) GetTag() []byte {
	if m != nil {
		return m.Tag
	}
	return nil
}

func (m *Contact_PreviousTag) GetExpired() int64 {
	if m != nil && m.Expired != nil {
		return *m.Expired
	}
	return 0
}

type Contact_Event struct {
	Time             *int64  `protobuf:"varint,1,req,name=time" json:"time,omitempty"`
	Message          *string `protobuf:"bytes,2,req,name=message" json:"message,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Contact_Event) Reset()         { *m = Contact_Event{} }
func (m *Contact_Event) String() string { return proto.CompactTextString(m) }
func (*Contact_Event) ProtoMessage()    {}

func (m *Contact_Event) GetTime() int64 {
	if m != nil && m.Time != nil {
		return *m.Time
	}
	return 0
}

func (m *Contact_Event) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

type RatchetState struct {
	RootKey            []byte                    `protobuf:"bytes,1,req,name=root_key" json:"root_key,omitempty"`
	SendHeaderKey      []byte                    `protobuf:"bytes,2,req,name=send_header_key" json:"send_header_key,omitempty"`
	RecvHeaderKey      []byte                    `protobuf:"bytes,3,req,name=recv_header_key" json:"recv_header_key,omitempty"`
	NextSendHeaderKey  []byte                    `protobuf:"bytes,4,req,name=next_send_header_key" json:"next_send_header_key,omitempty"`
	NextRecvHeaderKey  []byte                    `protobuf:"bytes,5,req,name=next_recv_header_key" json:"next_recv_header_key,omitempty"`
	SendChainKey       []byte                    `protobuf:"bytes,6,req,name=send_chain_key" json:"send_chain_key,omitempty"`
	RecvChainKey       []byte                    `protobuf:"bytes,7,req,name=recv_chain_key" json:"recv_chain_key,omitempty"`
	SendRatchetPrivate []byte                    `protobuf:"bytes,8,req,name=send_ratchet_private" json:"send_ratchet_private,omitempty"`
	RecvRatchetPublic  []byte                    `protobuf:"bytes,9,req,name=recv_ratchet_public" json:"recv_ratchet_public,omitempty"`
	SendCount          *uint32                   `protobuf:"varint,10,req,name=send_count" json:"send_count,omitempty"`
	RecvCount          *uint32                   `protobuf:"varint,11,req,name=recv_count" json:"recv_count,omitempty"`
	PrevSendCount      *uint32                   `protobuf:"varint,12,req,name=prev_send_count" json:"prev_send_count,omitempty"`
	Ratchet            *bool                     `protobuf:"varint,13,req,name=ratchet" json:"ratchet,omitempty"`
	V2                 *bool                     `protobuf:"varint,17,opt,name=v2" json:"v2,omitempty"`
	Private0           []byte                    `protobuf:"bytes,14,opt,name=private0" json:"private0,omitempty"`
	Private1           []byte                    `protobuf:"bytes,15,opt,name=private1" json:"private1,omitempty"`
	SavedKeys          []*RatchetState_SavedKeys `protobuf:"bytes,16,rep,name=saved_keys" json:"saved_keys,omitempty"`
	XXX_unrecognized   []byte                    `json:"-"`
}

func (m *RatchetState) Reset()         { *m = RatchetState{} }
func (m *RatchetState) String() string { return proto.CompactTextString(m) }
func (*RatchetState) ProtoMessage()    {}

func (m *RatchetState) GetRootKey() []byte {
	if m != nil {
		return m.RootKey
	}
	return nil
}

func (m *RatchetState) GetSendHeaderKey() []byte {
	if m != nil {
		return m.SendHeaderKey
	}
	return nil
}

func (m *RatchetState) GetRecvHeaderKey() []byte {
	if m != nil {
		return m.RecvHeaderKey
	}
	return nil
}

func (m *RatchetState) GetNextSendHeaderKey() []byte {
	if m != nil {
		return m.NextSendHeaderKey
	}
	return nil
}

func (m *RatchetState) GetNextRecvHeaderKey() []byte {
	if m != nil {
		return m.NextRecvHeaderKey
	}
	return nil
}

func (m *RatchetState) GetSendChainKey() []byte {
	if m != nil {
		return m.SendChainKey
	}
	return nil
}

func (m *RatchetState) GetRecvChainKey() []byte {
	if m != nil {
		return m.RecvChainKey
	}
	return nil
}

func (m *RatchetState) GetSendRatchetPrivate() []byte {
	if m != nil {
		return m.SendRatchetPrivate
	}
	return nil
}

func (m *RatchetState) GetRecvRatchetPublic() []byte {
	if m != nil {
		return m.RecvRatchetPublic
	}
	return nil
}

func (m *RatchetState) GetSendCount() uint32 {
	if m != nil && m.SendCount != nil {
		return *m.SendCount
	}
	return 0
}

func (m *RatchetState) GetRecvCount() uint32 {
	if m != nil && m.RecvCount != nil {
		return *m.RecvCount
	}
	return 0
}

func (m *RatchetState) GetPrevSendCount() uint32 {
	if m != nil && m.PrevSendCount != nil {
		return *m.PrevSendCount
	}
	return 0
}

func (m *RatchetState) GetRatchet() bool {
	if m != nil && m.Ratchet != nil {
		return *m.Ratchet
	}
	return false
}

func (m *RatchetState) GetV2() bool {
	if m != nil && m.V2 != nil {
		return *m.V2
	}
	return false
}

func (m *RatchetState) GetPrivate0() []byte {
	if m != nil {
		return m.Private0
	}
	return nil
}

func (m *RatchetState) GetPrivate1() []byte {
	if m != nil {
		return m.Private1
	}
	return nil
}

func (m *RatchetState) GetSavedKeys() []*RatchetState_SavedKeys {
	if m != nil {
		return m.SavedKeys
	}
	return nil
}

type RatchetState_SavedKeys struct {
	HeaderKey        []byte                               `protobuf:"bytes,1,req,name=header_key" json:"header_key,omitempty"`
	MessageKeys      []*RatchetState_SavedKeys_MessageKey `protobuf:"bytes,2,rep,name=message_keys" json:"message_keys,omitempty"`
	XXX_unrecognized []byte                               `json:"-"`
}

func (m *RatchetState_SavedKeys) Reset()         { *m = RatchetState_SavedKeys{} }
func (m *RatchetState_SavedKeys) String() string { return proto.CompactTextString(m) }
func (*RatchetState_SavedKeys) ProtoMessage()    {}

func (m *RatchetState_SavedKeys) GetHeaderKey() []byte {
	if m != nil {
		return m.HeaderKey
	}
	return nil
}

func (m *RatchetState_SavedKeys) GetMessageKeys() []*RatchetState_SavedKeys_MessageKey {
	if m != nil {
		return m.MessageKeys
	}
	return nil
}

type RatchetState_SavedKeys_MessageKey struct {
	Num              *uint32 `protobuf:"varint,1,req,name=num" json:"num,omitempty"`
	Key              []byte  `protobuf:"bytes,2,req,name=key" json:"key,omitempty"`
	CreationTime     *int64  `protobuf:"varint,3,req,name=creation_time" json:"creation_time,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *RatchetState_SavedKeys_MessageKey) Reset()         { *m = RatchetState_SavedKeys_MessageKey{} }
func (m *RatchetState_SavedKeys_MessageKey) String() string { return proto.CompactTextString(m) }
func (*RatchetState_SavedKeys_MessageKey) ProtoMessage()    {}

func (m *RatchetState_SavedKeys_MessageKey) GetNum() uint32 {
	if m != nil && m.Num != nil {
		return *m.Num
	}
	return 0
}

func (m *RatchetState_SavedKeys_MessageKey) GetKey() []byte {
	if m != nil {
		return m.Key
	}
	return nil
}

func (m *RatchetState_SavedKeys_MessageKey) GetCreationTime() int64 {
	if m != nil && m.CreationTime != nil {
		return *m.CreationTime
	}
	return 0
}

type Inbox struct {
	Id               *uint64 `protobuf:"fixed64,1,req,name=id" json:"id,omitempty"`
	From             *uint64 `protobuf:"fixed64,2,req,name=from" json:"from,omitempty"`
	ReceivedTime     *int64  `protobuf:"varint,3,req,name=received_time" json:"received_time,omitempty"`
	Acked            *bool   `protobuf:"varint,4,req,name=acked" json:"acked,omitempty"`
	Message          []byte  `protobuf:"bytes,5,opt,name=message" json:"message,omitempty"`
	Read             *bool   `protobuf:"varint,6,req,name=read" json:"read,omitempty"`
	Sealed           []byte  `protobuf:"bytes,7,opt,name=sealed" json:"sealed,omitempty"`
	Retained         *bool   `protobuf:"varint,8,opt,name=retained,def=0" json:"retained,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Inbox) Reset()         { *m = Inbox{} }
func (m *Inbox) String() string { return proto.CompactTextString(m) }
func (*Inbox) ProtoMessage()    {}

const Default_Inbox_Retained bool = false

func (m *Inbox) GetId() uint64 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *Inbox) GetFrom() uint64 {
	if m != nil && m.From != nil {
		return *m.From
	}
	return 0
}

func (m *Inbox) GetReceivedTime() int64 {
	if m != nil && m.ReceivedTime != nil {
		return *m.ReceivedTime
	}
	return 0
}

func (m *Inbox) GetAcked() bool {
	if m != nil && m.Acked != nil {
		return *m.Acked
	}
	return false
}

func (m *Inbox) GetMessage() []byte {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *Inbox) GetRead() bool {
	if m != nil && m.Read != nil {
		return *m.Read
	}
	return false
}

func (m *Inbox) GetSealed() []byte {
	if m != nil {
		return m.Sealed
	}
	return nil
}

func (m *Inbox) GetRetained() bool {
	if m != nil && m.Retained != nil {
		return *m.Retained
	}
	return Default_Inbox_Retained
}

type Outbox struct {
	Id               *uint64 `protobuf:"fixed64,1,req,name=id" json:"id,omitempty"`
	To               *uint64 `protobuf:"fixed64,2,req,name=to" json:"to,omitempty"`
	Server           *string `protobuf:"bytes,3,req,name=server" json:"server,omitempty"`
	Created          *int64  `protobuf:"varint,4,req,name=created" json:"created,omitempty"`
	Sent             *int64  `protobuf:"varint,5,opt,name=sent" json:"sent,omitempty"`
	Message          []byte  `protobuf:"bytes,6,opt,name=message" json:"message,omitempty"`
	Request          []byte  `protobuf:"bytes,7,opt,name=request" json:"request,omitempty"`
	Acked            *int64  `protobuf:"varint,8,opt,name=acked" json:"acked,omitempty"`
	Revocation       *bool   `protobuf:"varint,9,opt,name=revocation" json:"revocation,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Outbox) Reset()         { *m = Outbox{} }
func (m *Outbox) String() string { return proto.CompactTextString(m) }
func (*Outbox) ProtoMessage()    {}

func (m *Outbox) GetId() uint64 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *Outbox) GetTo() uint64 {
	if m != nil && m.To != nil {
		return *m.To
	}
	return 0
}

func (m *Outbox) GetServer() string {
	if m != nil && m.Server != nil {
		return *m.Server
	}
	return ""
}

func (m *Outbox) GetCreated() int64 {
	if m != nil && m.Created != nil {
		return *m.Created
	}
	return 0
}

func (m *Outbox) GetSent() int64 {
	if m != nil && m.Sent != nil {
		return *m.Sent
	}
	return 0
}

func (m *Outbox) GetMessage() []byte {
	if m != nil {
		return m.Message
	}
	return nil
}

func (m *Outbox) GetRequest() []byte {
	if m != nil {
		return m.Request
	}
	return nil
}

func (m *Outbox) GetAcked() int64 {
	if m != nil && m.Acked != nil {
		return *m.Acked
	}
	return 0
}

func (m *Outbox) GetRevocation() bool {
	if m != nil && m.Revocation != nil {
		return *m.Revocation
	}
	return false
}

type Draft struct {
	Id               *uint64                      `protobuf:"fixed64,1,req,name=id" json:"id,omitempty"`
	Created          *int64                       `protobuf:"varint,2,req,name=created" json:"created,omitempty"`
	ToNormal         []uint64                     `protobuf:"fixed64,3,rep,name=to_normal" json:"to_normal,omitempty"`
	ToIntroduce      []uint64                     `protobuf:"fixed64,8,rep,name=to_introduce" json:"to_introduce,omitempty"`
	Body             *string                      `protobuf:"bytes,4,req,name=body" json:"body,omitempty"`
	InReplyTo        *uint64                      `protobuf:"fixed64,5,opt,name=in_reply_to" json:"in_reply_to,omitempty"`
	Attachments      []*protos.Message_Attachment `protobuf:"bytes,6,rep,name=attachments" json:"attachments,omitempty"`
	Detachments      []*protos.Message_Detachment `protobuf:"bytes,7,rep,name=detachments" json:"detachments,omitempty"`
	XXX_unrecognized []byte                       `json:"-"`
}

func (m *Draft) Reset()         { *m = Draft{} }
func (m *Draft) String() string { return proto.CompactTextString(m) }
func (*Draft) ProtoMessage()    {}

func (m *Draft) GetId() uint64 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *Draft) GetCreated() int64 {
	if m != nil && m.Created != nil {
		return *m.Created
	}
	return 0
}

func (m *Draft) GetToNormal() []uint64 {
	if m != nil {
		return m.ToNormal
	}
	return nil
}

func (m *Draft) GetToIntroduce() []uint64 {
	if m != nil {
		return m.ToIntroduce
	}
	return nil
}

func (m *Draft) GetBody() string {
	if m != nil && m.Body != nil {
		return *m.Body
	}
	return ""
}

func (m *Draft) GetInReplyTo() uint64 {
	if m != nil && m.InReplyTo != nil {
		return *m.InReplyTo
	}
	return 0
}

func (m *Draft) GetAttachments() []*protos.Message_Attachment {
	if m != nil {
		return m.Attachments
	}
	return nil
}

func (m *Draft) GetDetachments() []*protos.Message_Detachment {
	if m != nil {
		return m.Detachments
	}
	return nil
}

type State struct {
	Identity                 []byte                 `protobuf:"bytes,1,req,name=identity" json:"identity,omitempty"`
	Public                   []byte                 `protobuf:"bytes,2,req,name=public" json:"public,omitempty"`
	Private                  []byte                 `protobuf:"bytes,3,req,name=private" json:"private,omitempty"`
	Server                   *string                `protobuf:"bytes,4,req,name=server" json:"server,omitempty"`
	Group                    []byte                 `protobuf:"bytes,5,req,name=group" json:"group,omitempty"`
	GroupPrivate             []byte                 `protobuf:"bytes,6,req,name=group_private" json:"group_private,omitempty"`
	PreviousGroupPrivateKeys []*State_PreviousGroup `protobuf:"bytes,12,rep,name=previous_group_private_keys" json:"previous_group_private_keys,omitempty"`
	Generation               *uint32                `protobuf:"varint,7,req,name=generation" json:"generation,omitempty"`
	LastErasureStorageTime   *int64                 `protobuf:"varint,13,opt,name=last_erasure_storage_time" json:"last_erasure_storage_time,omitempty"`
	Contacts                 []*Contact             `protobuf:"bytes,8,rep,name=contacts" json:"contacts,omitempty"`
	Inbox                    []*Inbox               `protobuf:"bytes,9,rep,name=inbox" json:"inbox,omitempty"`
	Outbox                   []*Outbox              `protobuf:"bytes,10,rep,name=outbox" json:"outbox,omitempty"`
	Drafts                   []*Draft               `protobuf:"bytes,11,rep,name=drafts" json:"drafts,omitempty"`
	XXX_unrecognized         []byte                 `json:"-"`
}

func (m *State) Reset()         { *m = State{} }
func (m *State) String() string { return proto.CompactTextString(m) }
func (*State) ProtoMessage()    {}

func (m *State) GetIdentity() []byte {
	if m != nil {
		return m.Identity
	}
	return nil
}

func (m *State) GetPublic() []byte {
	if m != nil {
		return m.Public
	}
	return nil
}

func (m *State) GetPrivate() []byte {
	if m != nil {
		return m.Private
	}
	return nil
}

func (m *State) GetServer() string {
	if m != nil && m.Server != nil {
		return *m.Server
	}
	return ""
}

func (m *State) GetGroup() []byte {
	if m != nil {
		return m.Group
	}
	return nil
}

func (m *State) GetGroupPrivate() []byte {
	if m != nil {
		return m.GroupPrivate
	}
	return nil
}

func (m *State) GetPreviousGroupPrivateKeys() []*State_PreviousGroup {
	if m != nil {
		return m.PreviousGroupPrivateKeys
	}
	return nil
}

func (m *State) GetGeneration() uint32 {
	if m != nil && m.Generation != nil {
		return *m.Generation
	}
	return 0
}

func (m *State) GetLastErasureStorageTime() int64 {
	if m != nil && m.LastErasureStorageTime != nil {
		return *m.LastErasureStorageTime
	}
	return 0
}

func (m *State) GetContacts() []*Contact {
	if m != nil {
		return m.Contacts
	}
	return nil
}

func (m *State) GetInbox() []*Inbox {
	if m != nil {
		return m.Inbox
	}
	return nil
}

func (m *State) GetOutbox() []*Outbox {
	if m != nil {
		return m.Outbox
	}
	return nil
}

func (m *State) GetDrafts() []*Draft {
	if m != nil {
		return m.Drafts
	}
	return nil
}

type State_PreviousGroup struct {
	Group            []byte `protobuf:"bytes,1,req,name=group" json:"group,omitempty"`
	GroupPrivate     []byte `protobuf:"bytes,2,req,name=group_private" json:"group_private,omitempty"`
	Expired          *int64 `protobuf:"varint,3,req,name=expired" json:"expired,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *State_PreviousGroup) Reset()         { *m = State_PreviousGroup{} }
func (m *State_PreviousGroup) String() string { return proto.CompactTextString(m) }
func (*State_PreviousGroup) ProtoMessage()    {}

func (m *State_PreviousGroup) GetGroup() []byte {
	if m != nil {
		return m.Group
	}
	return nil
}

func (m *State_PreviousGroup) GetGroupPrivate() []byte {
	if m != nil {
		return m.GroupPrivate
	}
	return nil
}

func (m *State_PreviousGroup) GetExpired() int64 {
	if m != nil && m.Expired != nil {
		return *m.Expired
	}
	return 0
}

func init() {
}
