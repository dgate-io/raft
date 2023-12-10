package raft

import "time"

// OperationTimeoutError represents an error that occurs when an operation submitted to
// raft times out.
type OperationTimeoutError struct {
	// The operation that was submitted to raft.
	Operation []byte
}

// Implements the Error interface for the OperationTimeoutError type.
func (e OperationTimeoutError) Error() string {
	return "The operation timed out while waiting for a response. This could be due to loss of server " +
		"leadership, a partitioned leader, prolonged processing, or a different reason. Try submitting the " +
		"operation to this server again or another server."
}

// Operation is an operation that will be applied to the state machine.
// An operation must be deterministic.
type Operation struct {
	// The operation as bytes. The provided state machine should be capable
	// of decoding these bytes.
	Bytes []byte

	// Indicates whether the operation is read-only. If it is, the log index
	// and log term will not be valid as there is no log entry associated with
	// the operation.
	IsReadOnly bool

	// The log entry index associated with the operation.
	LogIndex uint64

	// The log entry term associated with the operation.
	LogTerm uint64

	// The channel that the result of the operation will be sent over.
	ResponseCh chan OperationResponse
}

// OperationResponse is the response that is generated after applying
// an operation to the state machine.
type OperationResponse struct {
	// The operation applied to the state machine.
	Operation Operation

	// The response returned by the state machine after applying the operation.
	Response interface{}

	// An error encountered during the processing of the response, if any.
	Err error
}

// OperationResponseFuture represents a future response for an operation.
type OperationResponseFuture struct {
	// The operation associated with the future response.
	operation []byte

	// The maximum time to wait for a response.
	timeout time.Duration

	// A buffered channel for receiving the response.
	responseCh chan OperationResponse
}

// NewOperationResponseFuture creates a new OperationResponseFuture instance.
func NewOperationResponseFuture(operation []byte, timeout time.Duration) *OperationResponseFuture {
	return &OperationResponseFuture{
		operation:  operation,
		timeout:    timeout,
		responseCh: make(chan OperationResponse, 1),
	}
}

// Await waits for the response associated with the future operation.
// Note that the returned response may contain an error which should always be
// checked before consuming the content of the response. The content is not valid
// if the error is not nil.
func (o *OperationResponseFuture) Await() OperationResponse {
	for {
		select {
		case response := <-o.responseCh:
			return response
		case <-time.After(o.timeout):
			return OperationResponse{Err: OperationTimeoutError{Operation: o.operation}}
		}
	}
}
