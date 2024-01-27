package raft

import (
	"bytes"
	"encoding/gob"
	"errors"
	"fmt"
	"io"
	"reflect"
	"sync"
	"testing"
	"time"

	"github.com/jmsadair/raft/internal/logger"
	"github.com/jmsadair/raft/internal/util"
	"github.com/stretchr/testify/require"
)

const (
	// Max amount of time to elect a leader in seconds.
	maxElectionTime = 5

	// Max amount of time for a configuration change to complete in seconds.
	maxMembershipChangeTime = 10

	// Max amount of time for an operation to be applied in seconds.
	maxSubmissionTime = 10

	// Max amount of time for state machines to match in seconds.
	maxMatchTime = 3

	// Default timeout for futures.
	futureTimeout = 200 * time.Millisecond
)

func checkLogEntry(t *testing.T, expected *LogEntry, actual *LogEntry) {
	require.Equal(t, expected.Index, actual.Index)
	require.Equal(t, expected.Term, actual.Term)
	require.Equal(t, expected.Data, actual.Data)
	require.Equal(t, expected.EntryType, actual.EntryType)
}

func makeOperations(numOperations int) [][]byte {
	operations := make([][]byte, numOperations)
	for i := 1; i <= numOperations; i++ {
		operations[i-1] = []byte(fmt.Sprintf("operation %d", i))
	}
	return operations
}

func makeClusterConfiguration(numServers int) Configuration {
	members := make(map[string]string, numServers)
	isVoter := make(map[string]bool, numServers)
	for i := 0; i < numServers; i++ {
		id := fmt.Sprint(i)
		address := fmt.Sprintf("127.0.0.%d:8080", i)
		members[id] = address
		isVoter[id] = true
	}

	return Configuration{Members: members, IsVoter: isVoter}
}

func makeRaft(
	id string,
	address string,
	dataPath string,
	snapshotting bool,
	snapshotSize int,
) (*Raft, error) {
	logger, err := makeLogger(id)
	if err != nil {
		return nil, err
	}
	fsm := newStateMachineMock(snapshotting, snapshotSize)
	transport, err := newTransportMock(address)
	if err != nil {
		return nil, err
	}
	raft, err := NewRaft(id, address, fsm, dataPath, WithLogger(logger), WithTransport(transport))
	if err != nil {
		return nil, err
	}
	return raft, nil
}

func makeLogger(id string) (Logger, error) {
	prefix := fmt.Sprintf("raft-%s:", id)
	level := logger.Debug
	return logger.NewLogger(logger.WithLevel(level), logger.WithPrefix(prefix))
}

func encodeOperations(operations []Operation) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(operations); err != nil {
		return buf.Bytes(), err
	}
	return buf.Bytes(), nil
}

func decodeOperations(data []byte) ([]Operation, error) {
	var operations []Operation
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)
	if err := dec.Decode(&operations); err != nil {
		return operations, err
	}
	return operations, nil
}

type transportMock struct {
	Transport
	isDisconnected bool
	disconnected   sync.Map
}

func newTransportMock(address string) (*transportMock, error) {
	base, err := NewTransport(address)
	if err != nil {
		return nil, err
	}
	return &transportMock{
		Transport:    base,
		disconnected: sync.Map{},
	}, nil
}

func (t *transportMock) disconnect(address string) {
	t.disconnected.Store(address, true)
}

func (t *transportMock) connect(address string) {
	t.disconnected.Delete(address)
}

func (t *transportMock) SendAppendEntries(
	address string,
	request AppendEntriesRequest,
) (AppendEntriesResponse, error) {
	if _, ok := t.disconnected.Load(address); ok {
		return AppendEntriesResponse{}, errors.New("could not send AppendEntries RPC: disconnected")
	}
	return t.Transport.SendAppendEntries(address, request)
}

func (t *transportMock) SendRequestVote(
	address string,
	request RequestVoteRequest,
) (RequestVoteResponse, error) {
	if _, ok := t.disconnected.Load(address); ok {
		return RequestVoteResponse{}, errors.New("could not send RequestVote RPC: disconnected")
	}
	return t.Transport.SendRequestVote(address, request)
}

func (t *transportMock) SendInstallSnapshot(
	address string,
	request InstallSnapshotRequest,
) (InstallSnapshotResponse, error) {
	if _, ok := t.disconnected.Load(address); ok {
		return InstallSnapshotResponse{}, errors.New(
			"could not send InstallSnapshot RPC: disconnected",
		)
	}
	return t.Transport.SendInstallSnapshot(address, request)
}

type stateMachineMock struct {
	operations   []Operation
	snapshotting bool
	snapshotSize int
	mu           sync.Mutex
}

func newStateMachineMock(snapshotting bool, snapshotSize int) *stateMachineMock {
	gob.Register(Operation{})
	return &stateMachineMock{
		operations:   make([]Operation, 0),
		snapshotting: snapshotting,
		snapshotSize: snapshotSize,
	}
}

func (s *stateMachineMock) Apply(operation *Operation) interface{} {
	s.mu.Lock()
	defer s.mu.Unlock()
	if operation.OperationType == LeaseBasedReadOnly ||
		operation.OperationType == LinearizableReadOnly {
		return len(s.operations)
	}
	s.operations = append(s.operations, *operation)
	return len(s.operations)
}

func (s *stateMachineMock) Snapshot(snapshotWriter io.Writer) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	snapshotBytes, err := encodeOperations(s.operations)
	if err != nil {
		return fmt.Errorf("error taking snapshot of state machine: error = %v", err)
	}

	if _, err := snapshotWriter.Write(snapshotBytes); err != nil {
		return fmt.Errorf("error taking snapshot of state machine: error = %v", err)
	}

	return nil
}

func (s *stateMachineMock) Restore(snapshotReader io.Reader) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	var buf bytes.Buffer
	_, err := io.Copy(&buf, snapshotReader)
	if err != nil {
		return fmt.Errorf("error restoring state machine: error = %v", err)
	}
	bytes := buf.Bytes()

	entries, err := decodeOperations(bytes)
	if err != nil {
		return fmt.Errorf("error restoring state machine: error = %v", err)
	}

	s.operations = entries

	return nil
}

func (s *stateMachineMock) NeedSnapshot(logSize int) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.snapshotting && logSize%s.snapshotSize == 0
}

func (s *stateMachineMock) appliedOperations() []Operation {
	s.mu.Lock()
	defer s.mu.Unlock()

	operationsCopy := make([]Operation, len(s.operations))
	copy(operationsCopy, s.operations)
	return operationsCopy
}

type testCluster struct {
	// The testing instance associated with the cluster.
	t *testing.T

	// The nodes that make up the cluster.
	nodes map[string]*Raft

	// The ID, address, and voting status of all cluster members.
	configuration Configuration

	// The directories containing the persisted state for each node.
	dirs map[string]string

	// The transport for each node.
	transports map[string]*transportMock

	// The state machine associated with each node.
	stateMachines map[string]*stateMachineMock

	// Indicates whether auto snapshotting will be used.
	snapshotting bool

	// The maximum number of log entries per snapshot if snapshotting is enabled.
	snapshotSize int

	mu sync.RWMutex
}

func newCluster(t *testing.T, numServers int, snapshotting bool, snapshotSize int) *testCluster {
	nodes := make(map[string]*Raft, numServers)
	dirs := make(map[string]string, numServers)
	stateMachines := make(map[string]*stateMachineMock, numServers)
	transports := make(map[string]*transportMock, numServers)
	configuration := makeClusterConfiguration(numServers)

	// Create the nodes.
	for id, address := range configuration.Members {
		tmpDir := t.TempDir()
		node, err := makeRaft(id, address, tmpDir, snapshotting, snapshotSize)
		if err != nil {
			t.Fatalf("failed to create node: error = %v", err)
		}
		stateMachines[id] = node.fsm.(*stateMachineMock)
		transports[id] = node.transport.(*transportMock)
		dirs[id] = tmpDir
		nodes[id] = node
	}

	return &testCluster{
		t:             t,
		nodes:         nodes,
		transports:    transports,
		configuration: configuration,
		stateMachines: stateMachines,
		dirs:          dirs,
		snapshotting:  snapshotting,
		snapshotSize:  snapshotSize,
	}
}

// This starts the test cluster and should always be called before any operations
// are submitted or any failures are inflicted. This is not concurrent safe and should
// only be called once.
func (tc *testCluster) startCluster() {
	for _, node := range tc.nodes {
		if err := node.Bootstrap(tc.configuration.Members); err != nil {
			tc.t.Fatalf("failed to bootstrap node: error = %v", err)
		}
		if err := node.Start(); err != nil {
			tc.t.Fatalf("failed to start node: error = %v", err)
		}
	}
}

// This stops the test cluster. This is not concurrent safe and should only
// be called once.
func (tc *testCluster) stopCluster() {
	for _, node := range tc.nodes {
		node.Stop()
	}
}

func (tc *testCluster) submit(
	expectFail bool,
	operationType OperationType,
	operations ...[]byte,
) {
	for _, operation := range operations {
		tc.mu.RLock()

		// Attempt to submit the operations.
		start := time.Now()
		success := false
		for time.Since(start).Seconds() < maxSubmissionTime {
			// Try to submit to this node. It might be the leader.
			for _, node := range tc.nodes {
				operationFuture := node.SubmitOperation(operation, operationType, futureTimeout)
				response := operationFuture.Await()
				if err := response.Error(); err == nil {
					if expectFail {
						tc.t.Fatal("expected the operation to fail, but it was successful")
					}
					result := response.Success()
					if string(result.Operation.Bytes) != string(operation) {
						tc.t.Fatal("operation response does not match submitted operation")
					}
					success = true
					break
				}
			}

			if success {
				break
			}

			// Sleep a little bit in case the cluster needs to stabilize.
			tc.mu.RUnlock()
			time.Sleep(defaultElectionTimeout)
			tc.mu.RLock()
		}

		if !success && !expectFail {
			tc.mu.RUnlock()
			tc.t.Fatalf(
				"cluster timed out trying to apply operation: operation = %s",
				string(operation),
			)
		}

		tc.mu.RUnlock()
	}
}

func (tc *testCluster) addServer(id string, address string, isVoter bool) {
	// Create a new node if necessary.
	tc.mu.Lock()
	if _, ok := tc.nodes[id]; !ok {
		// This is currently not supported and it isn't really useful to create a new node
		// as a voting member if it is going to be dynamically added to the cluster.
		if isVoter {
			tc.t.Fatalf(
				"cannot add a node as a voter that does not already exist as a non-voting member",
			)
		}

		// Create the node
		tmpDir := tc.t.TempDir()
		node, err := makeRaft(id, address, tmpDir, tc.snapshotting, tc.snapshotSize)
		if err != nil {
			tc.t.Fatalf("failed to make node: error = %v", err)
		}
		tc.nodes[id] = node
		tc.dirs[id] = tmpDir
		tc.stateMachines[id] = node.fsm.(*stateMachineMock)
		tc.transports[id] = node.transport.(*transportMock)

		// Start the node as a non-voting member with no configuration.
		if err := node.Start(); err != nil {
			tc.t.Fatalf("failed to start node: error = %v", err)
		}
	}
	tc.mu.Unlock()

	// Attempt to add the node to the cluster.
	tc.mu.RLock()
	defer tc.mu.RUnlock()
	start := time.Now()
	for time.Since(start).Seconds() < maxMembershipChangeTime {
		for _, node := range tc.nodes {
			// Submit the request to add a  member to the cluster.
			// This node might be the leader.
			future := node.AddServer(id, address, isVoter, futureTimeout)
			response := future.Await()
			if err := response.Error(); err != nil {
				continue
			}

			// Make sure the configuration contains the node.
			configuration := response.Success()
			actualAddress, ok := configuration.Members[id]
			actualIsVoter := configuration.IsVoter[id]
			if !ok {
				tc.t.Fatalf(
					"membership change returned success, but node is missing from configuration: ID = %s",
					id,
				)
			}
			if actualAddress != address {
				tc.t.Fatalf(
					"membership change returned success, but node has incorrect address: ID = %s, actualAddress = %s, expectedAddress = %s",
					id,
					actualAddress,
					address,
				)
			}
			if actualIsVoter != isVoter {
				tc.t.Fatalf(
					"membership change returned success, but node has incorrect voting status: ID = %s, actualIsVoter = %t, expectedIsVoter = %t",
					id,
					actualIsVoter,
					isVoter,
				)
			}
			return
		}

		// Sleep a little in case the cluster needs to stabilize.
		tc.mu.RUnlock()
		time.Sleep(defaultElectionTimeout)
		tc.mu.RLock()
	}

	tc.t.Fatalf("timed out trying to add a node: ID = %s, address = %s", id, address)
}

func (tc *testCluster) removeServer(id string) {
	// Attempt to remove the node from the cluster.
	tc.mu.RLock()
	start := time.Now()
	for time.Since(start).Seconds() < maxMembershipChangeTime {
		for _, node := range tc.nodes {
			// Submit the request to remove a member to the cluster.
			// This node might be the leader.
			future := node.RemoveServer(id, futureTimeout)
			response := future.Await()
			if err := response.Error(); err != nil {
				continue
			}
			tc.mu.RUnlock()

			// Make sure the configuration foes not contain the removed node.
			configuration := response.Success()
			_, inMembers := configuration.Members[id]
			_, inIsVoter := configuration.IsVoter[id]
			if inMembers || inIsVoter {
				tc.t.Fatalf(
					"membership change returned success, but node is still in configuration: ID = %s",
					id,
				)
			}

			// Stop the node.
			tc.mu.Lock()
			defer tc.mu.Unlock()
			removeNode, ok := tc.nodes[id]
			removeNode.Stop()
			if !ok {
				tc.t.Fatalf("tried to remove a node that does not exist: ID = %s", id)
			}
			delete(tc.nodes, id)
			delete(tc.dirs, id)
			delete(tc.stateMachines, id)
			delete(tc.transports, id)
			return
		}

		// Sleep a little in case the cluster needs to stabilize.
		tc.mu.RUnlock()
		time.Sleep(defaultElectionTimeout)
		tc.mu.RLock()
	}

	tc.mu.RUnlock()
	tc.t.Fatalf("timed out trying to remove a node: ID = %s", id)
}

func (tc *testCluster) checkStateMachines(expectedMatches int, submittedOperations [][]byte) {
	tc.mu.RLock()
	defer tc.mu.RUnlock()

	var matchID string
	startTime := time.Now()
	allAppliedOperations := make(map[string][]Operation, len(tc.nodes))
	for time.Since(startTime) < maxMatchTime {
		// Take the state machine with the most applied operations to be the source of truth.
		for id := range tc.nodes {
			appliedOperations := tc.stateMachines[id].appliedOperations()
			if matchID == "" || len(appliedOperations) > len(allAppliedOperations[matchID]) {
				matchID = id
			}
			allAppliedOperations[id] = appliedOperations
		}

		// Check if the applied operations from the other state machines match source of truth.
		matches := 1
		for id, appliedOperations := range allAppliedOperations {
			if id == matchID {
				continue
			}
			if reflect.DeepEqual(allAppliedOperations[matchID], appliedOperations) {
				matches++
			}
		}

		// Check that we have at least the expected number of matches and that
		// the applied operations are correct.
		if matches >= expectedMatches {
			// Check that applied log indices are monotonically increasing on all
			// state machines, not just those that are supposed to match.
			for _, operations := range allAppliedOperations {
				tc.checkMonotonicity(operations)
			}
			// Check that the matching state machines do contain all submitted
			// operations. Since these match, it's fine to just check one.
			tc.checkContainsAll(matchID, submittedOperations, allAppliedOperations[matchID])
			return
		}

		tc.mu.RUnlock()
		time.Sleep(defaultElectionTimeout)
		tc.mu.RLock()
	}

	// There are not enough matches.
	// The state machines have diverged. Find where two state machines differ.
	expectedAppliedOperation := allAppliedOperations[matchID]
	for actualID, actualAppliedOperations := range allAppliedOperations {
		if actualID == matchID {
			continue
		}
		tc.compareOperations(matchID, expectedAppliedOperation, actualID, actualAppliedOperations)
	}
}

func (tc *testCluster) checkMonotonicity(operations []Operation) {
	lastIndex := uint64(0)
	for _, operation := range operations {
		// The log index of each operation should never decrease.
		// Note that there may be skipped log indices due to no-op entries and configuration entries.
		if operation.LogIndex <= lastIndex {
			tc.t.Fatalf(
				"operations are not monotonic, indices should strictly increase: lastIndex = %d, index = %d",
				lastIndex,
				operation.LogIndex,
			)
		}
		lastIndex = operation.LogIndex
	}
}

func (tc *testCluster) checkContainsAll(
	id string,
	submittedOperations [][]byte,
	appliedOperations []Operation,
) {
	appliedSet := make(map[string]bool)
	for _, operation := range appliedOperations {
		appliedSet[string(operation.Bytes)] = true
	}
	for _, operation := range submittedOperations {
		if _, ok := appliedSet[string(operation)]; !ok {
			tc.t.Fatalf("state machine is missing operations: ID = %s", id)
		}
	}
}

func (tc *testCluster) compareOperations(
	expectedID string,
	expectedOperations []Operation,
	actualID string,
	actualOperations []Operation,
) {
	// The arrays of operations match one another.
	if reflect.DeepEqual(expectedOperations, actualOperations) {
		return
	}

	// The arrays of operations do not match.
	// Try to find the first index where they differ for debugging purposes.
	for i := 0; i < util.Min(len(expectedOperations), len(actualOperations)); i++ {
		expectedOperation := expectedOperations[i]
		actualOperation := actualOperations[i]
		if reflect.DeepEqual(expectedOperation, actualOperation) {
			continue
		}
		tc.t.Fatalf(
			"state machines do not match: expectedID = %s, expectedLogIndex = %d, expectedLogTerm = %d, actualID = %s, actualLogIndex = %d, actualLogTerm = %d",
			expectedID,
			expectedOperation.LogIndex,
			expectedOperation.LogTerm,
			actualID,
			actualOperation.LogIndex,
			actualOperation.LogTerm,
		)
	}

	// The prefix of both arrays match, but one is shorter or longer than the other.
	tc.t.Fatal("state machines do not match: incorrect number of operations")
}

func (tc *testCluster) checkLeaders(expectNoLeader bool) string {
	// Any leaders detected.
	leaders := make([]string, 0, 1)

	// Check the nodes to see which, if any, are in the leader state.
	start := time.Now()
	for time.Since(start).Seconds() < maxElectionTime {
		tc.mu.RLock()
		for _, node := range tc.nodes {
			// Get the status of the node, it may be a leader.
			status := node.Status()

			// If the node is a leader, and it is connected, then it is
			// a legitimate leader. Leaders that are disconnected or
			// partitioned are ignored. It is assumed that disconnected
			// nodes are either:
			// 1. Completely disconnected from all other nodes - it
			//    cannot communicate with any other nodes, and no other
			//    nodes can communicate with it.
			// 2. In a minority partition - it may only communicate with
			//    a minority of the cluster. Members of the majority partition
			//    cannot communicate with it.
			if status.State == Leader && !tc.transports[node.id].isDisconnected {
				leaders = append(leaders, status.ID)
			}
		}
		tc.mu.RUnlock()

		if len(leaders) > 1 {
			tc.t.Fatalf("cluster has more than one leader: leaders = %v", leaders)
		}

		if len(leaders) == 1 {
			break
		}

		// If no leaders were found, sleep for a sufficient amount of time to allow
		// an election to take place.
		time.Sleep(defaultElectionTimeout)
	}

	if len(leaders) == 0 && !expectNoLeader {
		tc.t.Fatal("cluster failed to elect a leader in a reasonable amount of time")
	}

	if len(leaders) != 0 && expectNoLeader {
		tc.t.Fatalf("cluster elected leader without quorum: leaders = %v", leaders)
	}

	if expectNoLeader {
		return ""
	}

	return leaders[0]
}

func (tc *testCluster) crashServer(id string) {
	tc.mu.RLock()
	defer tc.mu.RUnlock()

	node, ok := tc.nodes[id]
	if !ok {
		tc.t.Fatalf("attempted to crash node that does not exist: ID = %s", id)
	}
	status := node.Status()
	if status.State == Shutdown {
		tc.t.Fatalf("attempted to crash server that was already crashed: ID = %s", id)
	}

	node.Stop()
}

func (tc *testCluster) crashRandom() string {
	tc.mu.RLock()
	defer tc.mu.RUnlock()

	notCrashed := make([]*Raft, 0, len(tc.nodes))
	for _, node := range tc.nodes {
		status := node.Status()
		if status.State != Shutdown {
			notCrashed = append(notCrashed, node)
		}
	}
	i := util.RandomInt(0, len(notCrashed))
	notCrashed[i].Stop()

	return notCrashed[i].id
}

func (tc *testCluster) restartServers() {
	tc.mu.Lock()
	defer tc.mu.Unlock()

	for id, node := range tc.nodes {
		status := node.Status()
		if status.State == Shutdown {
			tc.restart(id)
		}
	}
}

func (tc *testCluster) restartServer(id string) {
	tc.mu.Lock()
	defer tc.mu.Unlock()

	tc.restart(id)
}

func (tc *testCluster) restart(id string) {
	crashedNode, ok := tc.nodes[id]
	if !ok {
		tc.t.Fatalf("attempted to restart node which does not exist: ID = %s", id)
	}
	node, err := makeRaft(id, crashedNode.address, tc.dirs[id], tc.snapshotting, tc.snapshotSize)
	if err != nil {
		tc.t.Fatalf("failed to create node: error = %v", err)
	}
	tc.nodes[id] = node
	tc.transports[id] = node.transport.(*transportMock)
	tc.stateMachines[id] = node.fsm.(*stateMachineMock)
	node.Start()
}

func (tc *testCluster) createPartition() {
	tc.mu.Lock()
	defer tc.mu.Unlock()

	// The number of nodes in the partition.
	partitionSize := (len(tc.nodes) - 1) / 2

	disconnected := make(map[string]bool)

	// Choose random nodes to partition.
	for id := range tc.nodes {
		tc.transports[id].isDisconnected = true
		disconnected[id] = true
		if len(disconnected) == partitionSize {
			break
		}
	}

	// Disconnect all nodes in the partition set from those
	// that are not, but maintain connections between the nodes
	// that are in the partition set.
	for id1, node1 := range tc.nodes {
		if _, ok := disconnected[id1]; ok {
			for id2, node2 := range tc.nodes {
				if _, ok := disconnected[id2]; ok {
					continue
				}
				tc.transports[node1.id].disconnect(node2.address)
				tc.transports[node2.id].disconnect(node1.address)
			}
		}
	}
}

func (tc *testCluster) reconnectServer(id string) {
	tc.mu.Lock()
	defer tc.mu.Unlock()

	node1, ok := tc.nodes[id]
	if !ok {
		tc.t.Fatalf("attempted to reconnect node that does not exist: ID = %s", id)
	}

	for _, node2 := range tc.nodes {
		if node1 == node2 {
			continue
		}
		tc.transports[node1.id].connect(node2.address)
		tc.transports[node2.id].connect(node1.address)
	}

	tc.transports[id].isDisconnected = false
}

func (tc *testCluster) reconnectAllServers() {
	tc.mu.Lock()
	defer tc.mu.Unlock()

	for _, node1 := range tc.nodes {
		for _, node2 := range tc.nodes {
			if node1 == node2 {
				continue
			}
			tc.transports[node1.id].connect(node2.address)
			tc.transports[node1.id].isDisconnected = false
		}
	}
}

func (tc *testCluster) disconnectRandom() string {
	tc.mu.Lock()
	defer tc.mu.Unlock()

	notDisconnected := make([]string, 0, len(tc.nodes))

	for id := range tc.nodes {
		if !tc.transports[id].isDisconnected {
			notDisconnected = append(notDisconnected, id)
		}
	}

	i := util.RandomInt(0, len(notDisconnected))
	tc.disconnect(notDisconnected[i])

	return notDisconnected[i]
}

func (tc *testCluster) disconnectServer(id string) {
	tc.mu.Lock()
	defer tc.mu.Unlock()

	tc.disconnect(id)
}

func (tc *testCluster) disconnect(id string) {
	node1, ok := tc.nodes[id]
	if !ok {
		tc.t.Fatalf("attempted to disconnected node that does not exist: ID = %s", id)
	}

	for _, node2 := range tc.nodes {
		if node1 == node2 {
			continue
		}
		tc.transports[node1.id].disconnect(node2.address)
		tc.transports[node2.id].disconnect(node1.address)
	}

	tc.transports[node1.id].isDisconnected = true
}

func (tc *testCluster) unusedIDandAddress() (string, string) {
	tc.mu.RLock()
	defer tc.mu.RUnlock()

	// Collect the addresses being used.
	addresses := make(map[string]bool, len(tc.nodes))
	for _, node := range tc.nodes {
		addresses[node.address] = true
	}

	// Find an unused ID.
	i := 0
	var id string
	for {
		id = fmt.Sprint(i)
		if _, ok := tc.nodes[id]; !ok {
			break
		}
		i++
	}

	// Find an unused address.
	i = 0
	var address string
	for {
		address = fmt.Sprintf("127.0.0.%d:8080", i)
		if _, ok := addresses[address]; !ok {
			break
		}
		i++
	}

	return id, address
}

func (tc *testCluster) nodeIDs() []string {
	tc.mu.RLock()
	defer tc.mu.RUnlock()

	// Collect all node IDs.
	nodeIDs := make([]string, 0, len(tc.nodes))
	for id := range tc.nodes {
		nodeIDs = append(nodeIDs, id)
	}

	return nodeIDs
}
