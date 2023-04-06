package main

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main1() {

	cert := "key.key"
	// mux := http.NewServeMux()
	// mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
	// 	w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
	// 	w.Write([]byte("It works !!\n"))
	// })

	b, _ := ioutil.ReadFile(cert)
	var pemBlocks []*pem.Block
	var v *pem.Block
	var pkey []byte

	for {
		v, b = pem.Decode(b)
		fmt.Println("v=", v, "b=", b)
		if v == nil {
			break
		}
		if v.Type == "RSA PRIVATE KEY" {
			fmt.Println("yes key")
			if x509.IsEncryptedPEMBlock(v) {
				pkey, _ = x509.DecryptPEMBlock(v, []byte("foobar"))
				pkey = pem.EncodeToMemory(&pem.Block{

					Type:  v.Type,
					Bytes: pkey,
				})
				fmt.Println("yes encrypted key", pkey)
			} else {
				pkey = pem.EncodeToMemory(v)
				fmt.Println("yes not encrypted key", pkey)
			}
		} else {
			pemBlocks = append(pemBlocks, v)
			fmt.Println("yes certificate", pemBlocks[0])
		}
	}
	fmt.Println("cer=", pem.EncodeToMemory(pemBlocks[0]), "\n\nkey=", pkey)
	c, _ := tls.X509KeyPair(pem.EncodeToMemory(pemBlocks[0]), pkey)

	// generate a `Certificate` struct
	// cert1, _ := tls.LoadX509KeyPair("cert.crt", "key.key")

	// create a custom server with `TLSConfig`
	s := &http.Server{
		Addr:    ":9000",
		Handler: nil, // use `http.DefaultServeMux`
		TLSConfig: &tls.Config{
			Certificates: []tls.Certificate{c},
		},
	}

	// handle `/` route
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, "Hello Custom World!")
	})

	// run server on port "9000"
	log.Fatal(s.ListenAndServeTLS("", ""))

}
