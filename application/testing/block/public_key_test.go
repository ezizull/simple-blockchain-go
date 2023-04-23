package block

import (
	"fmt"
	blockDomain "simple-blockchain/domain/block"
	keyClient "simple-blockchain/infrastructure/blockchain/client/key"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPublicKeyAddress(t *testing.T) {
	privateKey := keyClient.GeneratePrivateKey()
	publicKey := privateKey.ToPublic()
	address := publicKey.Address()

	assert.Equal(t, blockDomain.AddressLen, len(address.Bytes()))
	fmt.Println(address)
}
