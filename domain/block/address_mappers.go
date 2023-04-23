package block

import "encoding/hex"

func (a Address) Bytes() []byte {
	return a.Value
}

func (a Address) String() string {
	return hex.EncodeToString(a.Value)
}
