package ssl

import (
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"time"
)

type CertificateInfo struct {
	Version      int
	DNSNames     []string
	Issuer       string
	NotAfter     time.Time
	NotBefore    time.Time
	Subject      string
	SerialNumber string
}

func sslCertificateObj(sslCertificateContent []byte) (*x509.Certificate, error) {
	
	sslCertificatePemDecode, _ := pem.Decode(sslCertificateContent)
	
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("except: ", r)
		}
	}()
	
	certificateObj, err := x509.ParseCertificate(sslCertificatePemDecode.Bytes)
	
	return certificateObj, err
}

/*
解析SSL证书信息
*/

func ParseSSLCertificate(sslCertificateContent []byte) (c *CertificateInfo, err error) {
	var certificateInfo CertificateInfo
	
	sslCertificateX509Obj, err := sslCertificateObj(sslCertificateContent)
	
	if err != nil {
		return nil, err
	}
	
	if sslCertificateX509Obj == nil {
		return nil, errors.New("格式异常")
	}
	
	certificateInfo.Version = sslCertificateX509Obj.Version
	certificateInfo.NotBefore = sslCertificateX509Obj.NotBefore //.Format("2006-01-02 15:04")
	certificateInfo.NotAfter = sslCertificateX509Obj.NotAfter   //.Format("2006-01-02 15:04")
	certificateInfo.Issuer = sslCertificateX509Obj.Issuer.String()
	certificateInfo.Subject = sslCertificateX509Obj.Subject.String()
	certificateInfo.SerialNumber = sslCertificateX509Obj.SerialNumber.String()
	certificateInfo.DNSNames = sslCertificateX509Obj.DNSNames
	
	return &certificateInfo, nil
}

/*
校验域名是否匹配SSL证书
*/

func VerifyHostName(sslCertificateContent []byte, domain string) error {
	sslCertificateX509Obj, err := sslCertificateObj(sslCertificateContent)
	
	if err != nil {
		return err
	}
	
	return sslCertificateX509Obj.VerifyHostname(domain)
}
