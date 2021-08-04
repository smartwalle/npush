package aps

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"golang.org/x/crypto/pkcs12"
	"io/ioutil"
)

var (
	ErrExpired = errors.New("certificate has expired or is not yet valid")
)

func Load(filename, password string) (tls.Certificate, error) {
	p12, err := ioutil.ReadFile(filename)
	if err != nil {
		return tls.Certificate{}, fmt.Errorf("unable to load %s: %v", filename, err)
	}
	return Decode(p12, password)
}

func Decode(p12 []byte, password string) (tls.Certificate, error) {
	privateKey, cert, err := pkcs12.Decode(p12, password)
	if err != nil {
		return tls.Certificate{}, err
	}
	if err = verify(cert); err != nil {
		return tls.Certificate{}, err
	}
	return tls.Certificate{
		Certificate: [][]byte{cert.Raw},
		PrivateKey:  privateKey,
		Leaf:        cert,
	}, nil
}

func verify(cert *x509.Certificate) error {
	_, err := cert.Verify(x509.VerifyOptions{})
	if err == nil {
		return nil
	}

	switch e := err.(type) {
	case x509.CertificateInvalidError:
		switch e.Reason {
		case x509.Expired:
			return ErrExpired
		case x509.IncompatibleUsage:
			// Apple cert fail on go 1.10
			return nil
		default:
			return err
		}
	case x509.UnknownAuthorityError:
		// Apple cert isn't in the cert pool
		// ignoring this error
		return nil
	default:
		return err
	}
}
