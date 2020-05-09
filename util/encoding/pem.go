/*
-----BEGIN Type-----
Headers
base64-encoded Bytes
-----END Type-----
*/
package encoding

import "encoding/pem"

// PEMDecode
// pem 数据解密
func PEMDecode(data string) (*pem.Block, []byte) {
	block, rest := pem.Decode([]byte(data))
	if block == nil {
		return nil, nil
	}
	return block, rest
}
