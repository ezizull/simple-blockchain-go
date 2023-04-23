package block

func (p *PublicKey) Bytes() []byte {
	return p.Key
}

func (p *PublicKey) Address() Address {
	return Address{
		Value: p.Key[len(p.Key)-AddressLen:],
	}
}
