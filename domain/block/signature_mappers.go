package block

import (
	"crypto/ed25519"
)

func (s *Signature) Bytes() []byte {
	return s.Value
}

func (s *Signature) Verify(publicKey *PublicKey, msg []byte) bool {
	return ed25519.Verify(publicKey.Key, msg, s.Value)
}
