package key

import (
	blockDomain "simple-blockchain/domain/block"
	keyClient "simple-blockchain/infrastructure/blockchain/client/key"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGeneratePrivateKey(t *testing.T) {
	privateKey := keyClient.GeneratePrivateKey()
	assert.Equal(t, len(privateKey.Bytes()), blockDomain.PrivateKeyLen)

	publicKey := privateKey.ToPublic()
	assert.Equal(t, len(publicKey.Bytes()), blockDomain.PubKeyLen)
}

func TestPrivateKeySign(t *testing.T) {
	privateKey := keyClient.GeneratePrivateKey()
	publicKey := privateKey.ToPublic()
	msg := []byte("foo bar baz")

	signature := privateKey.Sing(msg)
	assert.True(t, signature.Verify(publicKey, msg))

	// test with invalid msg
	assert.False(t, signature.Verify(publicKey, []byte("foo")))

	// test with invalid publicKey
	invalidPrivateKey := keyClient.GeneratePrivateKey()
	invalidPublicKey := invalidPrivateKey.ToPublic()
	assert.False(t, signature.Verify(invalidPublicKey, msg))
}
