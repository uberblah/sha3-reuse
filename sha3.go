package sha3r

import (
	"encoding/binary"
	"fmt"

	"golang.org/x/crypto/sha3"
)

// SHA3r is a reusable sponge hash based on golang's sha3
//   SHA3r's state is not modified when reading
//   to modify state when reading, use SHA3rm
type SHA3r struct {
	state sha3.ShakeHash
}

// NewSHA3r returns a reusable version of the sha3 ShakeHash
func NewSHA3r() sha3.ShakeHash {
	return &SHA3r{
		state: sha3.NewShake256(),
	}
}

func (s *SHA3r) Read(data []byte) (n int, e error) {
	clone := s.state.Clone()
	n, e = s.state.Read(data)
	s.state = clone
	return
}

func (s *SHA3r) Write(data []byte) (n int, e error) {
	n, e = s.state.Write(data)
	return
}

// Reset resets the hash state
func (s *SHA3r) Reset() {
	s.state.Reset()
}

// Clone returns a copy of the hash
func (s *SHA3r) Clone() sha3.ShakeHash {
	return &SHA3r{
		state: s.state,
	}
}

// SHA3rm is a version of SHA3r whose state is modified during reads
type SHA3rm struct {
	state sha3.ShakeHash
}

// NewSHA3rm returns a reusable version of the sha3 ShakeHash
func NewSHA3rm() sha3.ShakeHash {
	return &SHA3rm{
		state: sha3.NewShake256(),
	}
}

// Read reads data, but also performs a write to the internal hash,
//   so that reads will mutate state. n may be non-zero even if an
//   error has occurred, so always check the error status
func (s *SHA3rm) Read(data []byte) (n int, e error) {
	// update the state based on how much we're about to read
	bs := make([]byte, 4)
	binary.LittleEndian.PutUint32(bs, uint32(len(data)))
	var nw int
	nw, e = s.state.Write(bs)
	if e != nil {
		return
	}
	if nw < len(bs) {
		e = fmt.Errorf("SHA3rm.Read: only %d/%d bytes written in state mutation",
			nw, len(bs))
		return
	}

	// get the requested data without invalidating our state
	clone := s.state.Clone()
	n, e = s.state.Read(data)
	s.state = clone
	if e != nil {
		return
	}

	return
}

func (s *SHA3rm) Write(data []byte) (n int, e error) {
	n, e = s.state.Write(data)
	return
}

// Reset resets the hash state
func (s *SHA3rm) Reset() {
	s.state.Reset()
}

// Clone returns a copy of the hash
func (s *SHA3rm) Clone() sha3.ShakeHash {
	return &SHA3r{
		state: s.state,
	}
}
