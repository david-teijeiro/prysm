package beaconv1

import (
	"context"
	"errors"

	ptypes "github.com/gogo/protobuf/types"
	ethpb "github.com/prysmaticlabs/ethereumapis/eth/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

// ListPoolAttestations retrieves attestations known by the node but
// not necessarily incorporated into any block.
func (bs *Server) ListPoolAttestations(ctx context.Context, req *ethpb.AttestationsPoolRequest) (*ethpb.AttestationsPoolResponse, error) {
	atts := bs.AttestationsPool.AggregatedAttestations()
	numAtts := len(atts)
	if numAtts == 0 {
		return &ethpb.AttestationsPoolResponse{
			Data: make([]*ethpb.Attestation, 0),
		}, nil
	}
	v1Atts := make([]*ethpb.Attestation, len(atts))
	for i, att := range atts {
		marshaledAtt, err := att.Marshal()
		if err != nil {
			return nil, status.Errorf(codes.Internal, "Could not marshal attestation: %v", err)
		}
		v1Att := &ethpb.Attestation{}
		if err := proto.Unmarshal(marshaledAtt, v1Att); err != nil {
			return nil, status.Errorf(codes.Internal, "Could not unmarshal attestation: %v", err)
		}
		v1Atts[i] = v1Att
	}
	return &ethpb.AttestationsPoolResponse{
		Data: v1Atts,
	}, nil
}

// SubmitAttestation submits Attestation object to node. If attestation passes all validation
// constraints, node MUST publish attestation on appropriate subnet.
func (bs *Server) SubmitAttestation(ctx context.Context, req *ethpb.Attestation) (*ptypes.Empty, error) {
	return nil, errors.New("unimplemented")
}

// ListPoolAttesterSlashings retrieves attester slashings known by the node but
// not necessarily incorporated into any block.
func (bs *Server) ListPoolAttesterSlashings(ctx context.Context, req *ptypes.Empty) (*ethpb.AttesterSlashingsPoolResponse, error) {
	return nil, errors.New("unimplemented")
}

// SubmitAttesterSlashing submits AttesterSlashing object to node's pool and
// if passes validation node MUST broadcast it to network.
func (bs *Server) SubmitAttesterSlashing(ctx context.Context, req *ethpb.AttesterSlashing) (*ptypes.Empty, error) {
	return nil, errors.New("unimplemented")
}

// ListPoolProposerSlashings retrieves proposer slashings known by the node
// but not necessarily incorporated into any block.
func (bs *Server) ListPoolProposerSlashings(ctx context.Context, req *ptypes.Empty) (*ethpb.ProposerSlashingPoolResponse, error) {
	return nil, errors.New("unimplemented")
}

// SubmitProposerSlashing submits AttesterSlashing object to node's pool and if
// passes validation node MUST broadcast it to network.
func (bs *Server) SubmitProposerSlashing(ctx context.Context, req *ethpb.ProposerSlashing) (*ptypes.Empty, error) {
	return nil, errors.New("unimplemented")
}

// ListPoolVoluntaryExits retrieves voluntary exits known by the node but
// not necessarily incorporated into any block.
func (bs *Server) ListPoolVoluntaryExits(ctx context.Context, req *ptypes.Empty) (*ethpb.VoluntaryExitsPoolResponse, error) {
	return nil, errors.New("unimplemented")
}

// SubmitVoluntaryExit submits SignedVoluntaryExit object to node's pool
// and if passes validation node MUST broadcast it to network.
func (bs *Server) SubmitVoluntaryExit(ctx context.Context, req *ethpb.SignedVoluntaryExit) (*ptypes.Empty, error) {
	return nil, errors.New("unimplemented")
}
