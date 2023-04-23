package block

import "crypto/ed25519"

type PrivateKey struct {
	Key ed25519.PrivateKey
}
