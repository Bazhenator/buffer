package entities

import (
	"time"

	"github.com/Bazhenator/buffer/internal/errors"
)

const (
	initBufSize = 0 // initBufSize - initial request's buffer size
	statusAccepted = 1 // status 1 when request is accepted
)

// Buffer is a main entity which represents request's buffer of cleaning-service (FIFO queue)
type Buffer struct {
	Requests []*Request

	Tail     *Request
	Head     *Request
	Size     uint64
	Capacity uint64
}

// NewBuffer creates new instance of Buffer
func NewBuffer(capacity uint64) *Buffer {
	buffer := make([]*Request, initBufSize, capacity)

	return &Buffer{
		Requests: buffer,

		Tail:     nil,
		Head:     nil,
		Size:     initBufSize,
		Capacity: capacity,
	}
}

// Append appends new request to Buffer.Requests, sets Head, Tail pointers.
// Returns Buffer.Size.
func (b *Buffer) Append(req *Request) uint64 {
	b.Requests = append(b.Requests, req)
	req.Status = statusAccepted
	req.AppendTime = time.Now() // request appends now
	b.Size = uint64(len(b.Requests))

	if b.Size == 1 {
		b.Head = b.Requests[0]
	}

	b.Tail = b.Requests[len(b.Requests)-1]
	return b.Size
}

// PopTop pops out Buffer.Head request, shifts the remaining elements to start from index 0,
// and updates Head and Tail pointers. 
// Returns the popped request or an error if the buffer is empty.
func (b *Buffer) PopTop() (*Request, error) {
	if b.Size == 0 {
		return nil, errors.ErrEmptyBuf
	}

	// Save the b.Head
	removedRequest := b.Head

	// Remove the first element and shift remaining elements
	b.Requests = b.Requests[1:]
	removedRequest.TimeInBuffer = time.Since(removedRequest.AppendTime) // request spent this time in buffer
	b.Size--

	// Update b.Head and b.Tail
	if b.Size == 0 {
		b.Head = nil
		b.Tail = nil
	} else {
		b.Head = b.Requests[0]
	}

	return removedRequest, nil
}


// PopBottom pops out Buffer.Tail request, removes it from the end of the queue,
// and updates Tail pointer. 
// Returns the popped request or an error if the buffer is empty.
func (b *Buffer) PopBottom() (*Request, error) {
	if b.Size == 0 {
		return nil, errors.ErrEmptyBuf
	}

	// Save the b.Tail
	removedRequest := b.Tail

	// Remove the last element
	b.Requests = b.Requests[:len(b.Requests)-1]
	b.Size--

	// Update Tail
	if b.Size == 0 {
		b.Head = nil
		b.Tail = nil
	} else {
		b.Tail = b.Requests[len(b.Requests)-1]
	}

	return removedRequest, nil
}

// GetHead returns Buffer.Head
func (b *Buffer) GetHead() *Request {
	return b.Head
}

// GetTail returns Buffer.Tail
func (b *Buffer) GetTail() *Request {
	return b.Tail
}

// GetCapacity returns Buffer.Capacity
func (b *Buffer) GetCapacity() uint64 {
	return b.Capacity
}

// GetSize returns Buffer.Size
func (b *Buffer) GetSize() uint64 {
	return b.Size
}