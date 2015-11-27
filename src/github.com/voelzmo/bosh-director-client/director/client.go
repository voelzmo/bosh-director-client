package director

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func NewClient(rootCAPath string) *http.Client {
	pemData, err := ioutil.ReadFile(rootCAPath)
	if err != nil {
		log.Fatal(fmt.Sprintf("Parsing the rootCA cert failed: %s", err))
	}
	pool := x509.NewCertPool()
	pool.AppendCertsFromPEM(pemData)
	tr := &http.Transport{
		TLSClientConfig:    &tls.Config{RootCAs: pool},
		DisableCompression: true,
	}
	client := &http.Client{Transport: tr}
	return client
}
