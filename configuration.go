package raft

import (
	"fmt"
	"strings"

	pb "github.com/dgate-io/raft/internal/protobuf"
	"google.golang.org/protobuf/proto"
)

// Configuration represents a cluster of nodes.
type Configuration struct {
	// All members of the cluster. Maps node ID to address.
	Members map[string]string

	// Maps node ID to a boolean that indicates whether the node
	// is a voting member or not. Voting members are those that
	// have their vote counted in elections and their match index
	// considered when the leader is advancing the commit index.
	// Non-voting members merely receive log entries. They are
	// not considered for election or commitment purposes.
	IsVoter map[string]bool

	// The log index of the configuration.
	Index uint64
}

// NewConfiguration creates a new configuration with the provided
// members and index. By default, all members in the returned configuration
// will have voter status.
func NewConfiguration(index uint64, members map[string]string) *Configuration {
	configuration := &Configuration{
		Index:   index,
		Members: members,
		IsVoter: make(map[string]bool, len(members)),
	}
	for id := range members {
		configuration.IsVoter[id] = true
	}
	return configuration
}

// Clone creates a deep-copy of the configuration.
func (c *Configuration) Clone() Configuration {
	configuration := Configuration{
		Index:   c.Index,
		IsVoter: make(map[string]bool, len(c.Members)),
		Members: make(map[string]string, len(c.Members)),
	}

	for id := range c.Members {
		configuration.IsVoter[id] = c.IsVoter[id]
		configuration.Members[id] = c.Members[id]
	}

	return configuration
}

// String returns a string representation of the configuration.
func (c *Configuration) String() string {
	var builder strings.Builder

	builder.WriteString(fmt.Sprintf("logIndex:%d, members:", c.Index))
	index := 0
	for nodeID, address := range c.Members {
		if c.IsVoter[nodeID] {
			builder.WriteString(fmt.Sprintf("(%s, %s, voter)", nodeID, address))
		} else {
			builder.WriteString(fmt.Sprintf("(%s, %s, non-voter)", nodeID, address))
		}
		index++
		if index < len(c.Members) {
			builder.WriteString(",")
		}
	}

	return fmt.Sprintf("{%s}", strings.TrimSuffix(builder.String(), ","))
}

func encodeConfiguration(configuration *Configuration) ([]byte, error) {
	pbConfiguration := &pb.Configuration{
		Members: configuration.Members,
		IsVoter: configuration.IsVoter,
		Index:   configuration.Index,
	}
	data, err := proto.Marshal(pbConfiguration)
	if err != nil {
		return nil, fmt.Errorf("could not marshal protobuf message: %w", err)
	}
	return data, nil
}

func decodeConfiguration(data []byte) (Configuration, error) {
	pbConfiguration := &pb.Configuration{}
	if err := proto.Unmarshal(data, pbConfiguration); err != nil {
		return Configuration{}, fmt.Errorf("could not unmarshal protobuf message: %w", err)
	}
	configuration := Configuration{
		Members: pbConfiguration.GetMembers(),
		IsVoter: pbConfiguration.GetIsVoter(),
		Index:   pbConfiguration.GetIndex(),
	}
	return configuration, nil
}
