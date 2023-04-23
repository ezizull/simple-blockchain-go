package key

import (
	"crypto/ed25519"
	"crypto/rand"
	"encoding/hex"
	"io"
	blockDomain "simple-blockchain/domain/block"
)

func GeneratePrivateKey() *blockDomain.PrivateKey {
	seed := make([]byte, blockDomain.SeedLen)
	_, err := io.ReadFull(rand.Reader, seed)
	if err != nil {
		panic(err)
	}

	return &blockDomain.PrivateKey{
		Key: ed25519.NewKeyFromSeed(seed),
	}
}

func NewPrivateKeyFromSeed(seed []byte) *blockDomain.PrivateKey {
	if len(seed) != blockDomain.SeedLen {
		panic("invalid seed length, must be 32")
	}

	return &blockDomain.PrivateKey{
		Key: ed25519.NewKeyFromSeed(seed),
	}
}

func NewPrivateKeyFromString(s string) *blockDomain.PrivateKey {
	b, err := hex.DecodeString(s)
	if err != nil {
		panic(err)
	}

	return NewPrivateKeyFromSeed(b)
}
