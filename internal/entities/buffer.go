package entities

import (
	"github.com/Bazhenator/buffer/internal/errors"
)

const (
	initBufSize = 0 // initBufSize - initial request's buffer size
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

// Append appends new request to Buffer.Requests, sets Head, Tail requests and returns Buffer.Size
func (b *Buffer) Append(req *Request) uint64 {
	b.Requests = append(b.Requests, req)
	b.Size = uint64(len(b.Requests))

	if b.Size == 1 {
		b.Head = b.Requests[0]
	}

	b.Tail = b.Requests[len(b.Requests)-1]
	return b.Size
}

// PopTop pops out Buffer.Head request and returns it if Buffer is not empty, else returns error
func (b *Buffer) PopTop() (*Request, error) {
	if b.Size == 0 {
		return nil, errors.ErrEmptyBuf
	}

	return b.Head, nil
}

// PopBottom pops out Buffer.Tail and returns it if buffer is not empty, else returns error
func (b *Buffer) PopBottom() (*Request, error) {
	if b.Size == 0 {
		return nil, errors.ErrEmptyBuf
	}

	return b.Tail, nil
}
