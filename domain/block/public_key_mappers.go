package block

func (p *PublicKey) Bytes() []byte {
	return p.Key
}
