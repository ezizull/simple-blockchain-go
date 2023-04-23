package block

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

func TestNewPrivateKeyFromString(t *testing.T) {
	str := "93f4e0e96a7b3d9dd6a337fc118e2538e7410106bc3072a0bdc5c5176b9b0993"
	privateKey := keyClient.NewPrivateKeyFromString(str)
	assert.Equal(t, blockDomain.PrivateKeyLen, len(privateKey.Bytes()))

	addStr := "ff5bf5ca5153e74eaed0ead9c093fa1dc9eaebc6"
	address := privateKey.ToPublic().Address()
	assert.Equal(t, address.String(), addStr)
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
