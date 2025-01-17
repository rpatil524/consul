package pbpeering

import (
	"time"

	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/hashicorp/consul/agent/structs"
	"github.com/hashicorp/consul/api"
	"github.com/hashicorp/consul/lib"
)

// RequestDatacenter implements structs.RPCInfo
func (req *GenerateTokenRequest) RequestDatacenter() string {
	// Cross-datacenter requests are not allowed for peering actions because
	// they rely on WAN-federation.
	return ""
}

// RequestDatacenter implements structs.RPCInfo
func (req *EstablishRequest) RequestDatacenter() string {
	// Cross-datacenter requests are not allowed for peering actions because
	// they rely on WAN-federation.
	return ""
}

// RequestDatacenter implements structs.RPCInfo
func (req *PeeringReadRequest) RequestDatacenter() string {
	// Cross-datacenter requests are not allowed for peering actions because
	// they rely on WAN-federation.
	return ""
}

// RequestDatacenter implements structs.RPCInfo
func (req *PeeringListRequest) RequestDatacenter() string {
	// Cross-datacenter requests are not allowed for peering actions because
	// they rely on WAN-federation.
	return ""
}

// RequestDatacenter implements structs.RPCInfo
func (req *PeeringWriteRequest) RequestDatacenter() string {
	// Cross-datacenter requests are not allowed for peering actions because
	// they rely on WAN-federation.
	return ""
}

// RequestDatacenter implements structs.RPCInfo
func (req *PeeringDeleteRequest) RequestDatacenter() string {
	// Cross-datacenter requests are not allowed for peering actions because
	// they rely on WAN-federation.
	return ""
}

// RequestDatacenter implements structs.RPCInfo
func (req *TrustBundleReadRequest) RequestDatacenter() string {
	// Cross-datacenter requests are not allowed for peering actions because
	// they rely on WAN-federation.
	return ""
}

// RequestDatacenter implements structs.RPCInfo
func (req *TrustBundleListByServiceRequest) RequestDatacenter() string {
	// Cross-datacenter requests are not allowed for peering actions because
	// they rely on WAN-federation.
	return ""
}

// ShouldDial returns true when the peering was stored via the peering initiation endpoint,
// AND the peering is not marked as terminated by our peer.
// If we generated a token for this peer we did not store our server addresses under PeerServerAddresses.
// These server addresses are for dialing, and only the peer initiating the peering will do the dialing.
func (p *Peering) ShouldDial() bool {
	return len(p.PeerServerAddresses) > 0
}

func (x PeeringState) GoString() string {
	return x.String()
}

// ConcatenatedRootPEMs concatenates and returns all PEM-encoded public certificates
// in a peer's trust bundle.
func (b *PeeringTrustBundle) ConcatenatedRootPEMs() string {
	if b == nil {
		return ""
	}

	var rootPEMs string
	for _, pem := range b.RootPEMs {
		rootPEMs += lib.EnsureTrailingNewline(pem)
	}
	return rootPEMs
}

// enumcover:PeeringState
func PeeringStateToAPI(s PeeringState) api.PeeringState {
	switch s {
	case PeeringState_PENDING:
		return api.PeeringStatePending
	case PeeringState_ESTABLISHING:
		return api.PeeringStateEstablishing
	case PeeringState_ACTIVE:
		return api.PeeringStateActive
	case PeeringState_FAILING:
		return api.PeeringStateFailing
	case PeeringState_DELETING:
		return api.PeeringStateDeleting
	case PeeringState_TERMINATED:
		return api.PeeringStateTerminated
	case PeeringState_UNDEFINED:
		fallthrough
	default:
		return api.PeeringStateUndefined
	}
}

// enumcover:api.PeeringState
func PeeringStateFromAPI(t api.PeeringState) PeeringState {
	switch t {
	case api.PeeringStatePending:
		return PeeringState_PENDING
	case api.PeeringStateEstablishing:
		return PeeringState_ESTABLISHING
	case api.PeeringStateActive:
		return PeeringState_ACTIVE
	case api.PeeringStateFailing:
		return PeeringState_FAILING
	case api.PeeringStateDeleting:
		return PeeringState_DELETING
	case api.PeeringStateTerminated:
		return PeeringState_TERMINATED
	case api.PeeringStateUndefined:
		fallthrough
	default:
		return PeeringState_UNDEFINED
	}
}

func (p *Peering) IsActive() bool {
	if p != nil && p.State == PeeringState_TERMINATED {
		return false
	}
	if p == nil || p.DeletedAt == nil {
		return true
	}

	// The minimum protobuf timestamp is the Unix epoch rather than go's zero.
	return structs.IsZeroProtoTime(p.DeletedAt)
}

func (p *Peering) ToAPI() *api.Peering {
	var t api.Peering
	PeeringToAPI(p, &t)
	return &t
}

// TODO consider using mog for this
func (resp *PeeringListResponse) ToAPI() []*api.Peering {
	list := make([]*api.Peering, len(resp.Peerings))
	for i, p := range resp.Peerings {
		list[i] = p.ToAPI()
	}
	return list
}

// TODO consider using mog for this
func (resp *GenerateTokenResponse) ToAPI() *api.PeeringGenerateTokenResponse {
	var t api.PeeringGenerateTokenResponse
	GenerateTokenResponseToAPI(resp, &t)
	return &t
}

// TODO consider using mog for this
func (resp *EstablishResponse) ToAPI() *api.PeeringEstablishResponse {
	var t api.PeeringEstablishResponse
	EstablishResponseToAPI(resp, &t)
	return &t
}

// convenience
func NewGenerateTokenRequestFromAPI(req *api.PeeringGenerateTokenRequest) *GenerateTokenRequest {
	if req == nil {
		return nil
	}
	t := &GenerateTokenRequest{}
	GenerateTokenRequestFromAPI(req, t)
	return t
}

// convenience
func NewEstablishRequestFromAPI(req *api.PeeringEstablishRequest) *EstablishRequest {
	if req == nil {
		return nil
	}
	t := &EstablishRequest{}
	EstablishRequestFromAPI(req, t)
	return t
}

func TimePtrFromProto(s *timestamp.Timestamp) *time.Time {
	if s == nil {
		return nil
	}
	t := structs.TimeFromProto(s)
	return &t
}

func TimePtrToProto(s *time.Time) *timestamp.Timestamp {
	if s == nil {
		return nil
	}
	return structs.TimeToProto(*s)
}
