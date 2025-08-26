package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"math/big"
	"os"
	"time"
)

var (
	CA_PrivateKey    *rsa.PrivateKey
	CA_Certficate    *x509.Certificate
	TEMP_Certificate []byte
	TEMP_PrivateKey  *rsa.PrivateKey
)

func Init(caCertPath, caKeyPath string) error {
	var err error
	certBytes, err := os.ReadFile(caCertPath)
	if err != nil {
		return err
	}
	block, _ := pem.Decode(certBytes)
	if block == nil {
		return fmt.Errorf("bad pem format of ca crt")
	}
	CA_Certficate, err = x509.ParseCertificate(block.Bytes)
	if err != nil {
		return err
	}

	keyBytes, err := os.ReadFile(caKeyPath)
	if err != nil {
		return err
	}
	block, _ = pem.Decode(keyBytes)
	if block == nil {
		return fmt.Errorf("bad pem format of ca key")
	}
	generalKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return err
	}
	var ok bool
	CA_PrivateKey, ok = generalKey.(*rsa.PrivateKey)
	if !ok {
		return fmt.Errorf("invalid rsa key")
	}
	return nil
}

func GenerateTempCertificateFromCSR(pemedCsr string) ([]byte, error) {
	block, _ := pem.Decode([]byte(pemedCsr))
	if block == nil {
		return nil, fmt.Errorf("empty pem string")
	}
	fmt.Println("block type", block.Type)
	csr, err := x509.ParseCertificateRequest(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("cannot parse block bytes: %v", err)
	}
	fmt.Println("subject", csr.Subject)
	template := &x509.Certificate{
		SerialNumber:          new(big.Int).SetInt64(time.Now().UnixNano()),
		Subject:               csr.Subject,
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(10, 0, 0),
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageKeyEncipherment,
		ExtKeyUsage:           []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth},
		BasicConstraintsValid: true,
	}

	certBytes, err := x509.CreateCertificate(
		rand.Reader,
		template,
		CA_Certficate,
		csr.PublicKey,
		CA_PrivateKey,
	)
	if err != nil {
		return nil, err
	}
	cert := pem.EncodeToMemory(&pem.Block{
		Type:  "CERTIFICATE",
		Bytes: certBytes,
	})
	return cert, nil

}

func GenerateTempCertificate(serialNumber string) error {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return err
	}
	template := &x509.Certificate{
		SerialNumber: new(big.Int).SetInt64(time.Now().UnixNano()),
		Subject: pkix.Name{
			Country:      []string{"CN"},
			Organization: []string{"Test Inc."},
			Locality:     []string{"NB"},
			Province:     []string{"ZJ"},
			CommonName:   "RSA File Certificate",
			SerialNumber: serialNumber,
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(10, 0, 0),
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageKeyEncipherment | x509.KeyUsageCertSign | x509.KeyUsageCRLSign,
		BasicConstraintsValid: true,
		IsCA:                  true,
	}

	certBytes, err := x509.CreateCertificate(
		rand.Reader,
		template,
		CA_Certficate,
		&privateKey.PublicKey,
		CA_PrivateKey,
	)
	if err != nil {
		return  err
	}
	TEMP_Certificate = pem.EncodeToMemory(&pem.Block{
		Type:  "CERTIFICATE",
		Bytes: certBytes,
	})
	TEMP_PrivateKey = privateKey
	return  nil
}

func GenerateTempCertificate2Local(serialNumber string) error {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return err
	}

	template := &x509.Certificate{
		SerialNumber: new(big.Int).SetInt64(time.Now().UnixNano()),
		Subject: pkix.Name{
			Country:      []string{"CN"},
			Organization: []string{"Test Inc."},
			Locality:     []string{"NB"},
			Province:     []string{"ZJ"},
			CommonName:   "RSA File Certificate",
			SerialNumber: serialNumber,
		},
		NotBefore:             time.Now(),
		NotAfter:              time.Now().AddDate(10, 0, 0),
		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageKeyEncipherment | x509.KeyUsageCertSign | x509.KeyUsageCRLSign,
		BasicConstraintsValid: true,
		IsCA:                  false,
	}

	certBytes, err := x509.CreateCertificate(
		rand.Reader,
		template,
		CA_Certficate,
		&privateKey.PublicKey,
		CA_PrivateKey,
	)
	if err != nil {
		return  err
	}

	certFile, err := os.Create("server.crt")
    if err != nil {
        return err
    }
    defer certFile.Close()

	err = pem.Encode(certFile, &pem.Block{
        Type:  "CERTIFICATE",
        Bytes: certBytes,
    })
    if err != nil {
        return err
    }

	keyFile, err := os.Create("server.key")
    if err != nil {
        return err
    }
    defer keyFile.Close()

    privKeyBytes, err := x509.MarshalPKCS8PrivateKey(privateKey)
    if err != nil {
        return err
    }

    err = pem.Encode(keyFile, &pem.Block{
        Type:  "PRIVATE KEY",
        Bytes: privKeyBytes,
    })
    if err != nil {
        return err
    }
    return nil
}


// func GenerateTempCertificate(caCert *x509.Certificate, caPrivateKey rsa.PrivateKey, serialNumber string) ([]byte, *rsa.PrivateKey, error) {
// 	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	template := &x509.Certificate{
// 		SerialNumber: new(big.Int).SetInt64(time.Now().UnixNano()),
// 		Subject: pkix.Name{
// 			Country:      []string{"CN"},
// 			Organization: []string{"Test Inc."},
// 			Locality:     []string{"NB"},
// 			Province:     []string{"ZJ"},
// 			CommonName:   "RSA File Certificate",
// 			SerialNumber: serialNumber,
// 		},
// 		NotBefore:             time.Now(),
// 		NotAfter:              time.Now().AddDate(10, 0, 0),
// 		KeyUsage:              x509.KeyUsageDigitalSignature | x509.KeyUsageKeyEncipherment | x509.KeyUsageCertSign | x509.KeyUsageCRLSign,
// 		BasicConstraintsValid: true,
// 		IsCA:                  true,
// 	}

// 	certBytes, err := x509.CreateCertificate(
// 		rand.Reader,
// 		template,
// 		caCert,
// 		&privateKey.PublicKey,
// 		caPrivateKey,
// 	)
// 	if err != nil {
// 		return nil, nil, err
// 	}
// 	cert := pem.EncodeToMemory(&pem.Block{
// 		Type:  "CERTIFICATE",
// 		Bytes: certBytes,
// 	})
// 	return cert, privateKey, nil
// }
