package token

const (
	// DefaultStreamBufferSize is the default channel buffer size of the
	// underlying channel in a token stream.
	DefaultStreamBufferSize = 5
)

// Stream is a token stream that allows tokens to be pushed. It uses a buffered
// channel as underlying queue.
//
//	_ = NewStream()           // channel buffer size = 5 (DefaultStreamBufferSize)
//	_ = NewBufferedStream(10) // channel buffer size = 10
type Stream struct {
	ch chan Token
}

// NewStream returns a new token stream with the default buffer size.
func NewStream() *Stream {
	return NewBufferedStream(DefaultStreamBufferSize)
}

// NewBufferedStream returns a new token stream with the given buffer size.
func NewBufferedStream(bufferSize int) *Stream {
	return &Stream{
		ch: make(chan Token, bufferSize),
	}
}

// Tokens returns a channel where all tokens of this stream can be read from
// (once). Multiple calls to Tokens() will return the same channel.
func (s *Stream) Tokens() <-chan Token {
	return s.ch
}

// Next returns the next token that was pushed onto the stream. This is a common
// channel receive operation.
func (s *Stream) Next() Token {
	return <-s.ch
}

// Push will push the given token onto the token stream. This is a common
// channel send operation.
func (s *Stream) Push(t Token) {
	s.ch <- t
}

// Close closes the underlying channel of this token stream.
func (s *Stream) Close() {
	close(s.ch)
}
