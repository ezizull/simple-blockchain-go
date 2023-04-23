package key

import (
	"crypto/ed25519"
	"crypto/rand"
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
