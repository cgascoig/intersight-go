package main

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"github.com/spf13/viper"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

func getSHA256Digest(data []byte) [sha256.Size]byte {
	ret := sha256.Sum256(data)
	// fmt.Printf("SHA256(%v)=%v\n", data, ret)
	return ret
}

func b64encode(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

func prepareStringToSign(requestTarget string, authHeaders map[string]string, headerList []string) string {
	ss := ""
	ss += "(request-target): " + strings.ToLower(requestTarget)
	for _, k := range headerList {
		ss += "\n" + strings.ToLower(k) + ": " + authHeaders[k]
	}

	return ss
}

func getSignedMessage(data [sha256.Size]byte, key *rsa.PrivateKey) string {
	// fmt.Printf("getSignedMessage -> data: %v\n", data)
	sig, err := rsa.SignPKCS1v15(nil, key, crypto.SHA256, data[:])
	if err != nil {
		panic(err)
	}
	return b64encode(sig)
}

func loadKey(filename string) *rsa.PrivateKey {
	bytes, _ := ioutil.ReadFile(filename)
	block, _ := pem.Decode(bytes)
	key, _ := x509.ParsePKCS1PrivateKey(block.Bytes)
	return key
}

func getAuthHeaderString(authHeaders map[string]string, signedMessage, keyID, algorithm string, headerList []string) string {
	authStr := fmt.Sprintf("Signature keyId=\"%s\",algorithm=\"%s\",headers=\"(request-target)", keyID, algorithm)

	for _, k := range headerList {
		authStr += " " + strings.ToLower(k)
	}
	authStr += "\""

	authStr += ",signature=\"" + signedMessage + "\""

	return authStr
}

func authFunc() runtime.ClientAuthInfoWriter {
	return runtime.ClientAuthInfoWriterFunc(func(r runtime.ClientRequest, _ strfmt.Registry) error {
		date := time.Now().In(time.UTC).Format(time.RFC1123)
		requestTarget := fmt.Sprintf("%s %s", r.GetMethod(), "/api/v1"+r.GetPath()) //TODO: Add query params!!
		bodyDigest := getSHA256Digest(r.GetBody())
		b64BodyDigest := b64encode(bodyDigest[:])

		headerList := []string{
			"Date",
			"Host",
			"Digest",
		}

		authHeaders := map[string]string{
			"Date":   date,
			"Host":   "intersight.com", //r.GetHeaderParams().Get("Host"),
			"Digest": fmt.Sprintf("SHA-256=%s", b64BodyDigest),
		}

		key := loadKey(viper.GetString("keyFile"))
		stringToSign := prepareStringToSign(requestTarget, authHeaders, headerList)
		authDigest := getSHA256Digest([]byte(stringToSign))
		b64SignedMessage := getSignedMessage(authDigest, key)
		authHeaderString := getAuthHeaderString(authHeaders, b64SignedMessage, viper.GetString("keyID"), "rsa-sha256", headerList)

		// fmt.Printf("requestTarget: %v\n", requestTarget)
		// fmt.Printf("authHeader: %v\n", authHeaders)
		// fmt.Printf("stringToSign: %v\n", stringToSign)
		// fmt.Printf("authDigest: %v\n", authDigest)
		// fmt.Printf("b64SignedMessage: %v\n", b64SignedMessage)
		// fmt.Printf("authHeaderString: %v\n", authHeaderString)

		r.SetHeaderParam("Date", date)
		r.SetHeaderParam("Digest", fmt.Sprintf("SHA-256=%s", b64BodyDigest))
		r.SetHeaderParam("Authorization", authHeaderString)

		return nil
	})
}
