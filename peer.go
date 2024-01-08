package raft

import (
	"context"
	"errors"
	"fmt"
	"net"
	"sync"

	pb "github.com/jmsadair/raft/internal/protobuf"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Peer is an interface representing a component responsible for establishing a connection
// with and making RPCs to another raft node.
type Peer interface {
	// ID returns the ID of the peer.
	ID() string

	// Address returns the network address of the peer.
	Address() string

	// Connect establishes a connection with the peer.
	Connect() error

	// Disconnect tears down a connection with the peer.
	Disconnect() error

	// AppendEntries sends an append entries request to the peer.
	AppendEntries(request AppendEntriesRequest) (AppendEntriesResponse, error)

	// RequestVote sends a request vote request to the peer.
	RequestVote(request RequestVoteRequest) (RequestVoteResponse, error)

	// InstallSnapshot sends a install snapshot request to the peer.
	InstallSnapshot(request InstallSnapshotRequest) (InstallSnapshotResponse, error)
}

// peer is an implementation of the Peer interface.
// This implementation is concurrent safe.
type peer struct {
	// The gRPC client for making Raft protocol calls to the peer.
	client pb.RaftClient

	// The ID of this peer.
	id string

	// The network address of this peer.
	address net.Addr

	// The gRPC client connection to communicate with the peer.
	conn *grpc.ClientConn

	// Prevents a race condition between disconnect/connect and the RPCs.
	mu sync.RWMutex
}

// NewPeer creates a new Peer instance with the provided ID and address.
func NewPeer(id string, address net.Addr) Peer {
	return &peer{id: id, address: address}
}

func (p *peer) ID() string {
	return p.id
}

func (p *peer) Address() string {
	return p.address.String()
}

func (p *peer) Connect() error {
	p.mu.Lock()
	defer p.mu.Unlock()

	if p.client != nil {
		return nil
	}

	conn, err := grpc.Dial(
		p.address.String(),
		[]grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}...)
	if err != nil {
		return fmt.Errorf("could not connect to peer: %w", err)
	}

	p.client = pb.NewRaftClient(conn)
	p.conn = conn

	return nil
}

func (p *peer) Disconnect() error {
	p.mu.Lock()
	defer p.mu.Unlock()

	if p.client == nil {
		return nil
	}

	if err := p.conn.Close(); err != nil {
		return fmt.Errorf("could not disonnect from peer: %w", err)
	}

	p.conn = nil
	p.client = nil

	return nil
}

func (p *peer) AppendEntries(request AppendEntriesRequest) (AppendEntriesResponse, error) {
	p.mu.RLock()
	defer p.mu.RUnlock()

	if p.client == nil {
		return AppendEntriesResponse{}, errors.New(
			"could not make AppendEntries RPC: no connection",
		)
	}

	pbRequest := makeProtoAppendEntriesRequest(request)
	pbResponse, err := p.client.AppendEntries(
		context.Background(),
		pbRequest,
		[]grpc.CallOption{}...)
	if err != nil {
		return AppendEntriesResponse{}, fmt.Errorf("could not make AppendEntries RPC: %w", err)
	}

	return makeAppendEntriesResponse(pbResponse), nil
}

func (p *peer) RequestVote(request RequestVoteRequest) (RequestVoteResponse, error) {
	p.mu.RLock()
	defer p.mu.RUnlock()

	if p.client == nil {
		return RequestVoteResponse{}, errors.New("could not make RequestVote RPC: no connection")
	}

	pbRequest := makeProtoRequestVoteRequest(request)
	pbResponse, err := p.client.RequestVote(context.Background(), pbRequest, []grpc.CallOption{}...)
	if err != nil {
		return RequestVoteResponse{}, fmt.Errorf("could not make RequestVote RPC: %w", err)
	}

	return makeRequestVoteResponse(pbResponse), nil
}

func (p *peer) InstallSnapshot(request InstallSnapshotRequest) (InstallSnapshotResponse, error) {
	p.mu.RLock()
	defer p.mu.RUnlock()

	if p.client == nil {
		return InstallSnapshotResponse{}, errors.New(
			"could not make InstallSnapshot RPC: no connection",
		)
	}

	pbRequest := makeProtoInstallSnapshotRequest(request)
	pbResponse, err := p.client.InstallSnapshot(
		context.Background(),
		pbRequest,
		[]grpc.CallOption{}...)
	if err != nil {
		return InstallSnapshotResponse{}, fmt.Errorf("could not make InstallSnapshot RPC: %w", err)
	}

	return makeInstallSnapshotResponse(pbResponse), nil
}
