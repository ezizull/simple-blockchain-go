package block

import "crypto/ed25519"

type PublicKey struct {
	Key ed25519.PublicKey
}
