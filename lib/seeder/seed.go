package seeder

import "lukechampine.com/blake3"

type Seeder struct {
	out   [32]byte
	index int
}

func New(seed []byte) *Seeder {
	return &Seeder{
		out: blake3.Sum256(seed),
	}
}

func (s *Seeder) Read(p []byte) (n int, err error) {
	for i := range p {
		p[i] = s.out[s.index]
		s.index++
		if s.index == len(s.out) {
			s.out = blake3.Sum256(s.out[:])
			s.index = 0
		}
	}
	return len(p), nil
}
