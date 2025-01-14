package ciphersuite

import (
	"github.com/pion/dtls/v2/pkg/crypto/ciphersuite"
	"github.com/pion/dtls/v2/pkg/crypto/clientcertificate"
)

// Aes128Ccm is a base class used by multiple AES-CCM Ciphers
type Aes128Ccm struct {
	AesCcm
}

func newAes128Ccm(clientCertificateType clientcertificate.Type, id ID, psk bool, cryptoCCMTagLen ciphersuite.CCMTagLen) *Aes128Ccm {
	return &Aes128Ccm{
		AesCcm: AesCcm{
			clientCertificateType: clientCertificateType,
			id:                    id,
			psk:                   psk,
			cryptoCCMTagLen:       cryptoCCMTagLen,
		},
	}
}

// Init initializes the internal Cipher with keying material
func (c *Aes128Ccm) Init(masterSecret, clientRandom, serverRandom []byte, isClient bool) error {
	const prfKeyLen = 16
	return c.AesCcm.Init(masterSecret, clientRandom, serverRandom, isClient, prfKeyLen)
}
