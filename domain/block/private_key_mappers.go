package block

import (
	"crypto/ed25519"
)

func (p PrivateKey) Bytes() []byte {
	return p.Key
}

func (p PrivateKey) Sing(msg []byte) *Signature {
	return &Signature{
		Value: ed25519.Sign(p.Key, msg),
	}
}

func (p PrivateKey) ToPublic() *PublicKey {
	b := make([]byte, PubKeyLen)
	copy(b, p.Key[32:])

	return &PublicKey{
		Key: b,
	}
}
