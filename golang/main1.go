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

func main2() {
	cert := "key.key"
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Header().Add("Strict-Transport-Security", "max-age=63072000; includeSubDomains")
		w.Write([]byte("It works !!\n"))
	})

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



	cfg := &tls.Config{
		MinVersion:               tls.VersionTLS12,
		CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
		PreferServerCipherSuites: true,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
			tls.TLS_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_RSA_WITH_AES_256_CBC_SHA,
		},
		Certificates: []tls.Certificate{c},
	}
	srv := &http.Server{
		Addr:         ":9000",
		Handler:      mux,
		TLSConfig:    cfg,
		TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
	}

	log.Fatal(srv.ListenAndServeTLS("", ""))
}

// package main

// import (
// 	"crypto/tls"
// 	"crypto/x509"
// 	"encoding/pem"
// 	"fmt"
// 	"io/ioutil"
// 	"log"
// 	"net/http"
// )

// func main() {
// 	cert := "key.key"
// 	b, err := ioutil.ReadFile(cert)
// 	// d,_:= ioutil.ReadFile("cert.crt")
// 	if err != nil {
// 		fmt.Print("could not read", err)
// 	}
// 	var pemBlocks []*pem.Block
// 	var v *pem.Block
// 	var pkey []byte

// 	for {
// 		v, b = pem.Decode(b)
// 		if v == nil {
// 			break
// 		}
// 		if v.Type == "RSA PRIVATE KEY" {
// 			fmt.Print("got rsa", v.Headers)
// 			if x509.IsEncryptedPEMBlock(v) {
// 				pkey, err = x509.DecryptPEMBlock(v, []byte("foobar"))
// 				if err != nil {
// 					fmt.Println("decrypting error", err)
// 				} else {
// 					fmt.Println("no decrypting error", pkey)
// 				}
// 				pkey = pem.EncodeToMemory(&pem.Block{
// 					Type:  v.Type,
// 					Bytes: pkey,
// 				})
// 			} else {
// 				pkey = pem.EncodeToMemory(v)
// 			}
// 		} else {
// 			pemBlocks = append(pemBlocks, v)
// 		}
// 	}
// 	fmt.Println("reached 48")
// 	// c, _ := tls.X509KeyPair(pem.EncodeToMemory(pemBlocks[0]), pkey)
// 	c, _ := tls.X509KeyPair(pem.EncodeToMemory(pemBlocks[0]), pkey)
// 	// create a custom server with `TLSConfig`
// 	s := &http.Server{
// 		Addr:    ":9000",
// 		Handler: nil, // use `http.DefaultServeMux`
// 		TLSConfig: &tls.Config{
// 			Certificates: []tls.Certificate{c},
// 		},
// 	}

// 	// handle `/` route
// 	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
// 		fmt.Fprint(res, "Hello Custom World!")
// 	})

// 	// run server on port "9000"
// 	log.Fatal(s.ListenAndServeTLS("", ""))

// 	// // handle `/` route
// 	// http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
// 	// 	fmt.Fprint(res, "Hello World!")
// 	// })

// 	// // run server on port "9000"
// 	// log.Fatal(http.ListenAndServeTLS(":9000", "cert.crt", "key.key", nil))
// }

// // package main

// // import (
// // 	"crypto/x509"
// // 	"encoding/pem"
// // 	"fmt"
// // )

// // func ExampleCertificate_Verify() {
// // 	// Verifying with a custom list of root certificates.

// // 	const rootPEM = `
// // -----BEGIN CERTIFICATE-----
// // MIIEBDCCAuygAwIBAgIDAjppMA0GCSqGSIb3DQEBBQUAMEIxCzAJBgNVBAYTAlVT
// // MRYwFAYDVQQKEw1HZW9UcnVzdCBJbmMuMRswGQYDVQQDExJHZW9UcnVzdCBHbG9i
// // YWwgQ0EwHhcNMTMwNDA1MTUxNTU1WhcNMTUwNDA0MTUxNTU1WjBJMQswCQYDVQQG
// // EwJVUzETMBEGA1UEChMKR29vZ2xlIEluYzElMCMGA1UEAxMcR29vZ2xlIEludGVy
// // bmV0IEF1dGhvcml0eSBHMjCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEB
// // AJwqBHdc2FCROgajguDYUEi8iT/xGXAaiEZ+4I/F8YnOIe5a/mENtzJEiaB0C1NP
// // VaTOgmKV7utZX8bhBYASxF6UP7xbSDj0U/ck5vuR6RXEz/RTDfRK/J9U3n2+oGtv
// // h8DQUB8oMANA2ghzUWx//zo8pzcGjr1LEQTrfSTe5vn8MXH7lNVg8y5Kr0LSy+rE
// // ahqyzFPdFUuLH8gZYR/Nnag+YyuENWllhMgZxUYi+FOVvuOAShDGKuy6lyARxzmZ
// // EASg8GF6lSWMTlJ14rbtCMoU/M4iarNOz0YDl5cDfsCx3nuvRTPPuj5xt970JSXC
// // DTWJnZ37DhF5iR43xa+OcmkCAwEAAaOB+zCB+DAfBgNVHSMEGDAWgBTAephojYn7
// // qwVkDBF9qn1luMrMTjAdBgNVHQ4EFgQUSt0GFhu89mi1dvWBtrtiGrpagS8wEgYD
// // VR0TAQH/BAgwBgEB/wIBADAOBgNVHQ8BAf8EBAMCAQYwOgYDVR0fBDMwMTAvoC2g
// // K4YpaHR0cDovL2NybC5nZW90cnVzdC5jb20vY3Jscy9ndGdsb2JhbC5jcmwwPQYI
// // KwYBBQUHAQEEMTAvMC0GCCsGAQUFBzABhiFodHRwOi8vZ3RnbG9iYWwtb2NzcC5n
// // ZW90cnVzdC5jb20wFwYDVR0gBBAwDjAMBgorBgEEAdZ5AgUBMA0GCSqGSIb3DQEB
// // BQUAA4IBAQA21waAESetKhSbOHezI6B1WLuxfoNCunLaHtiONgaX4PCVOzf9G0JY
// // /iLIa704XtE7JW4S615ndkZAkNoUyHgN7ZVm2o6Gb4ChulYylYbc3GrKBIxbf/a/
// // zG+FA1jDaFETzf3I93k9mTXwVqO94FntT0QJo544evZG0R0SnU++0ED8Vf4GXjza
// // HFa9llF7b1cq26KqltyMdMKVvvBulRP/F/A8rLIQjcxz++iPAsbw+zOzlTvjwsto
// // WHPbqCRiOwY1nQ2pM714A5AuTHhdUDqB1O6gyHA43LL5Z/qHQF1hwFGPa4NrzQU6
// // yuGnBXj8ytqU0CwIPX4WecigUCAkVDNx
// // -----END CERTIFICATE-----`

// // 	const certPEM = `
// // -----BEGIN CERTIFICATE-----
// // MIIDujCCAqKgAwIBAgIIE31FZVaPXTUwDQYJKoZIhvcNAQEFBQAwSTELMAkGA1UE
// // BhMCVVMxEzARBgNVBAoTCkdvb2dsZSBJbmMxJTAjBgNVBAMTHEdvb2dsZSBJbnRl
// // cm5ldCBBdXRob3JpdHkgRzIwHhcNMTQwMTI5MTMyNzQzWhcNMTQwNTI5MDAwMDAw
// // WjBpMQswCQYDVQQGEwJVUzETMBEGA1UECAwKQ2FsaWZvcm5pYTEWMBQGA1UEBwwN
// // TW91bnRhaW4gVmlldzETMBEGA1UECgwKR29vZ2xlIEluYzEYMBYGA1UEAwwPbWFp
// // bC5nb29nbGUuY29tMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEfRrObuSW5T7q
// // 5CnSEqefEmtH4CCv6+5EckuriNr1CjfVvqzwfAhopXkLrq45EQm8vkmf7W96XJhC
// // 7ZM0dYi1/qOCAU8wggFLMB0GA1UdJQQWMBQGCCsGAQUFBwMBBggrBgEFBQcDAjAa
// // BgNVHREEEzARgg9tYWlsLmdvb2dsZS5jb20wCwYDVR0PBAQDAgeAMGgGCCsGAQUF
// // BwEBBFwwWjArBggrBgEFBQcwAoYfaHR0cDovL3BraS5nb29nbGUuY29tL0dJQUcy
// // LmNydDArBggrBgEFBQcwAYYfaHR0cDovL2NsaWVudHMxLmdvb2dsZS5jb20vb2Nz
// // cDAdBgNVHQ4EFgQUiJxtimAuTfwb+aUtBn5UYKreKvMwDAYDVR0TAQH/BAIwADAf
// // BgNVHSMEGDAWgBRK3QYWG7z2aLV29YG2u2IaulqBLzAXBgNVHSAEEDAOMAwGCisG
// // AQQB1nkCBQEwMAYDVR0fBCkwJzAloCOgIYYfaHR0cDovL3BraS5nb29nbGUuY29t
// // L0dJQUcyLmNybDANBgkqhkiG9w0BAQUFAAOCAQEAH6RYHxHdcGpMpFE3oxDoFnP+
// // gtuBCHan2yE2GRbJ2Cw8Lw0MmuKqHlf9RSeYfd3BXeKkj1qO6TVKwCh+0HdZk283
// // TZZyzmEOyclm3UGFYe82P/iDFt+CeQ3NpmBg+GoaVCuWAARJN/KfglbLyyYygcQq
// // 0SgeDh8dRKUiaW3HQSoYvTvdTuqzwK4CXsr3b5/dAOY8uMuG/IAR3FgwTbZ1dtoW
// // RvOTa8hYiU6A475WuZKyEHcwnGYe57u2I2KbMgcKjPniocj4QzgYsVAVKW3IwaOh
// // yE+vPxsiUkvQHdO2fojCkY8jg70jxM+gu59tPDNbw3Uh/2Ij310FgTHsnGQMyA==
// // -----END CERTIFICATE-----`

// // 	// First, create the set of root certificates. For this example we only
// // 	// have one. It's also possible to omit this in order to use the
// // 	// default root set of the current operating system.
// // 	roots := x509.NewCertPool()
// // 	ok := roots.AppendCertsFromPEM([]byte(rootPEM))
// // 	if !ok {
// // 		panic("failed to parse root certificate")
// // 	}

// // 	block, _ := pem.Decode([]byte(certPEM))
// // 	if block == nil {
// // 		panic("failed to parse certificate PEM")
// // 	}
// // 	cert, err := x509.ParseCertificate(block.Bytes)
// // 	if err != nil {
// // 		panic("failed to parse certificate: " + err.Error())
// // 	}

// // 	print(cert.PublicKey)
// // 	// opts := x509.VerifyOptions{
// // 	// 	DNSName: "mail.google.com",
// // 	// 	Roots:   roots,
// // 	// }

// // 	// if _, err := cert.Verify(opts); err != nil {
// // 	// 	panic("failed to verify certificate: " + err.Error())
// // 	// }
// // }

// // func ExampleParsePKIXPublicKey() {
// // 	const pubPEM = `
// // -----BEGIN PUBLIC KEY-----
// // MIICIjANBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEAlRuRnThUjU8/prwYxbty
// // WPT9pURI3lbsKMiB6Fn/VHOKE13p4D8xgOCADpdRagdT6n4etr9atzDKUSvpMtR3
// // CP5noNc97WiNCggBjVWhs7szEe8ugyqF23XwpHQ6uV1LKH50m92MbOWfCtjU9p/x
// // qhNpQQ1AZhqNy5Gevap5k8XzRmjSldNAFZMY7Yv3Gi+nyCwGwpVtBUwhuLzgNFK/
// // yDtw2WcWmUU7NuC8Q6MWvPebxVtCfVp/iQU6q60yyt6aGOBkhAX0LpKAEhKidixY
// // nP9PNVBvxgu3XZ4P36gZV6+ummKdBVnc3NqwBLu5+CcdRdusmHPHd5pHf4/38Z3/
// // 6qU2a/fPvWzceVTEgZ47QjFMTCTmCwNt29cvi7zZeQzjtwQgn4ipN9NibRH/Ax/q
// // TbIzHfrJ1xa2RteWSdFjwtxi9C20HUkjXSeI4YlzQMH0fPX6KCE7aVePTOnB69I/
// // a9/q96DiXZajwlpq3wFctrs1oXqBp5DVrCIj8hU2wNgB7LtQ1mCtsYz//heai0K9
// // PhE4X6hiE0YmeAZjR0uHl8M/5aW9xCoJ72+12kKpWAa0SFRWLy6FejNYCYpkupVJ
// // yecLk/4L1W0l6jQQZnWErXZYe0PNFcmwGXy1Rep83kfBRNKRy5tvocalLlwXLdUk
// // AIU+2GKjyT3iMuzZxxFxPFMCAwEAAQ==
// // -----END PUBLIC KEY-----

// // ----BEGIN PRIVATE KEY-----
// // MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQDKlvdvjS7EEkYm
// // RcMYBemMNvIbMRoJPBuwtRHph8pehi/BUXmTUsGUSDelFAl6tH2eMpHD6FpMJkiu
// // PAm875yJYo9nxgnH8SNzS1FSR05qoFkJGSvkQwLd++V4Z17kqg9LyzQf/XGSiGGm
// // VbvPS8rkJUdrFxgTjqhw8KIfNXGG4zLNcB6i0YgW2lZFgaFB16J3/O4KDm21bgdi
// // qI6F5xKel/YSQl+AfF2NlFt5xI3wuZ4YhD7wrp3WWgiS/JIxu0ESDmJTYExdb3K0
// // Dh6bmQl+wYJbkBpO2csNRDehtM2YEU83N8mmlSbIbSSinZYT3KwT9JQbd8QqJZCP
// // /mOPFacrAgMBAAECggEAMqGWR3fWd0RF6ezHfGqF2vgke+1Cn4o5NWmbh2zbg9Iv
// // fzYYl1w4axG9bnFaiSMwvefPjFG2t49d3MW+fUy5J5DNXFcfPKwkev0Y3uJZU8at
// // WdvDn3Gr9sSsrfHPwoBKAFxRs6kIyGFzXjnRDVbY5zn15mrIJqMhr9BEBF6798TA
// // qTG0FYTqkGK0D+FVfaWXvQ0u9Jw0KootS4kKHNwDmbZK2xYI2Ilt1ikeN9MMt8ZD
// // tXV4shTnQaYPty5Atr9Dzh052FTnlwsclVo33XHF2N2dfe7TaJYTaf5uXuh7Vkj3
// // 99bKuvA8iLmEVe+i86L9K9LD5QEywO/sNcqQkNU8sQKBgQDkycMSv1dfSAgSZywU
// // hcJqFWJ3TpAP5aoWqw8Svill3Qs6/2zatU4XC4tzRX9KW437M4ORNee/XrCi9z4L
// // VOeOtR+gp/zn9DebBzaEdTfMlof+znPfx9fVIpBkpjezDgiFNkeeMnKTQrT/kULl
// // Zso29pCfgO/57L9Vi3pjiqKt0wKBgQDir4EDVesZTGcBstSRLIGrUQFmvxICPaMm
// // 0PogvQhUv8UFMxx6nBl/ZB7AZWQH0TRlpVL7iqS2clO1dQOgrtgwgUV4m3Ml9Ivr
// // vI2fgCWzFudPst3oSu9Udc9Pq995Pl2nnwV612p3P6p3wgCaFjbouyyJZxNmHQj/
// // JPtv65nSSQKBgQDVcRnJorLLlHLbYF9yYfunZn3vWl7yRcvxy/KLBNewTZENoIAY
// // Zm8M9ttJVjvTziheg4ep8EVdduSJlOnQPoysyXNROYeril5aBlepKYY+Gu2THV5j
// // FpjYIZ/eFmf+Zwgx5xrXjq7vjZs4lnd3dvcOYec4t1yqqGE0WKR8uzjbuwKBgAYL
// // 0FEadY7TLtwovOqyWTMMkhD/f6d3pWZfpIxC/nnkM4kT9+p9R2DSds+C5MwglFkx
// // s6jp5cLIAduRJ2udvj5s9EFnRAb7ItBC0zQx4s+ICNtjVe/gL8n86m6hkvBU7YKP
// // B0JjhH9xv0Y6cnGprgU/GM0BZs8ObzL+9YXirtOhAoGACk1TGT6I5oMHLDB4yOOe
// // a7dT+kAY6a2W2Sq2e0VN70EkVmXGC6ODpRzPcH7nojcMN5jk8QHHosWNK8DECwAj
// // uCGYn8G/0yAlhkddzE1+y1f5nVm+GCQTrdMMuqOwJnosifdoNDbWg4oGiBRt1uwI
// // aoxbWonlDAZaLC+8Bxe1Hss=
// // -----END PRIVATE KEY-----

// // -----BEGIN CERTIFICATE-----
// // MIIEBDCCAuygAwIBAgIDAjppMA0GCSqGSIb3DQEBBQUAMEIxCzAJBgNVBAYTAlVT
// // MRYwFAYDVQQKEw1HZW9UcnVzdCBJbmMuMRswGQYDVQQDExJHZW9UcnVzdCBHbG9i
// // YWwgQ0EwHhcNMTMwNDA1MTUxNTU1WhcNMTUwNDA0MTUxNTU1WjBJMQswCQYDVQQG
// // EwJVUzETMBEGA1UEChMKR29vZ2xlIEluYzElMCMGA1UEAxMcR29vZ2xlIEludGVy
// // bmV0IEF1dGhvcml0eSBHMjCCASIwDQYJKoZIhvcNAQEBBQADggEPADCCAQoCggEB
// // AJwqBHdc2FCROgajguDYUEi8iT/xGXAaiEZ+4I/F8YnOIe5a/mENtzJEiaB0C1NP
// // VaTOgmKV7utZX8bhBYASxF6UP7xbSDj0U/ck5vuR6RXEz/RTDfRK/J9U3n2+oGtv
// // h8DQUB8oMANA2ghzUWx//zo8pzcGjr1LEQTrfSTe5vn8MXH7lNVg8y5Kr0LSy+rE
// // ahqyzFPdFUuLH8gZYR/Nnag+YyuENWllhMgZxUYi+FOVvuOAShDGKuy6lyARxzmZ
// // EASg8GF6lSWMTlJ14rbtCMoU/M4iarNOz0YDl5cDfsCx3nuvRTPPuj5xt970JSXC
// // DTWJnZ37DhF5iR43xa+OcmkCAwEAAaOB+zCB+DAfBgNVHSMEGDAWgBTAephojYn7
// // qwVkDBF9qn1luMrMTjAdBgNVHQ4EFgQUSt0GFhu89mi1dvWBtrtiGrpagS8wEgYD
// // VR0TAQH/BAgwBgEB/wIBADAOBgNVHQ8BAf8EBAMCAQYwOgYDVR0fBDMwMTAvoC2g
// // K4YpaHR0cDovL2NybC5nZW90cnVzdC5jb20vY3Jscy9ndGdsb2JhbC5jcmwwPQYI
// // KwYBBQUHAQEEMTAvMC0GCCsGAQUFBzABhiFodHRwOi8vZ3RnbG9iYWwtb2NzcC5n
// // ZW90cnVzdC5jb20wFwYDVR0gBBAwDjAMBgorBgEEAdZ5AgUBMA0GCSqGSIb3DQEB
// // BQUAA4IBAQA21waAESetKhSbOHezI6B1WLuxfoNCunLaHtiONgaX4PCVOzf9G0JY
// // /iLIa704XtE7JW4S615ndkZAkNoUyHgN7ZVm2o6Gb4ChulYylYbc3GrKBIxbf/a/
// // zG+FA1jDaFETzf3I93k9mTXwVqO94FntT0QJo544evZG0R0SnU++0ED8Vf4GXjza
// // HFa9llF7b1cq26KqltyMdMKVvvBulRP/F/A8rLIQjcxz++iPAsbw+zOzlTvjwsto
// // WHPbqCRiOwY1nQ2pM714A5AuTHhdUDqB1O6gyHA43LL5Z/qHQF1hwFGPa4NrzQU6
// // yuGnBXj8ytqU0CwIPX4WecigUCAkVDNx
// // -----END CERTIFICATE-----
// // `

// // 	block, _ := pem.Decode([]byte(pubPEM))
// // 	if block == nil {
// // 		panic("failed to parse PEM block containing the public key")
// // 	}

// // 	// pub, err := x509.ParsePKIXPublicKey(block.Bytes)
// // 	// if err != nil {
// // 	// 	panic("failed to parse DER encoded public key: " + err.Error())
// // 	// }
// // 	// pub, err := x509.ParseCertificate(block.Bytes)
// // 	// if err != nil {
// // 	// 	panic("failed to parse DER encoded public key: " + err.Error())
// // 	// }
// // 	pub, err := x509.ParseECPrivateKey(block.Bytes)
// // 	if err != nil {
// // 		panic("failed to parse DER encoded public key: " + err.Error())
// // 	}

// // 	fmt.Print(pub)

// // 	// switch pub := pub.(type) {
// // 	// case *rsa.PublicKey:
// // 	// 	fmt.Println("pub is of type RSA:", pub)
// // 	// case *dsa.PublicKey:
// // 	// 	fmt.Println("pub is of type DSA:", pub)
// // 	// case *ecdsa.PublicKey:
// // 	// 	fmt.Println("pub is of type ECDSA:", pub)
// // 	// case ed25519.PublicKey:
// // 	// 	fmt.Println("pub is of type Ed25519:", pub)
// // 	// default:
// // 	// 	panic("unknown type of public key")
// // 	// }
// // }

// // func main() {
// // 	ExampleParsePKIXPublicKey()
// // 	// ExampleCertificate_Verify()
// // }

// // // package main

// // // import (
// // // 	"encoding/pem"
// // // 	"io/ioutil"
// // // 	"log"
// // // 	"strings"
// // // )

// // // func pain() {
// // // 	var privk = "-----BEGIN PRIVATE KEY-----\n" +
// // // 		"MIIEvQIBADANBgkqhkiG9w0BAQEFAASCBKcwggSjAgEAAoIBAQDKlvdvjS7EEkYm" +
// // // 		"RcMYBemMNvIbMRoJPBuwtRHph8pehi/BUXmTUsGUSDelFAl6tH2eMpHD6FpMJkiu" +
// // // 		"PAm875yJYo9nxgnH8SNzS1FSR05qoFkJGSvkQwLd++V4Z17kqg9LyzQf/XGSiGGm" +
// // // 		"VbvPS8rkJUdrFxgTjqhw8KIfNXGG4zLNcB6i0YgW2lZFgaFB16J3/O4KDm21bgdi" +
// // // 		"qI6F5xKel/YSQl+AfF2NlFt5xI3wuZ4YhD7wrp3WWgiS/JIxu0ESDmJTYExdb3K0" +
// // // 		"Dh6bmQl+wYJbkBpO2csNRDehtM2YEU83N8mmlSbIbSSinZYT3KwT9JQbd8QqJZCP" +
// // // 		"/mOPFacrAgMBAAECggEAMqGWR3fWd0RF6ezHfGqF2vgke+1Cn4o5NWmbh2zbg9Iv" +
// // // 		"fzYYl1w4axG9bnFaiSMwvefPjFG2t49d3MW+fUy5J5DNXFcfPKwkev0Y3uJZU8at" +
// // // 		"WdvDn3Gr9sSsrfHPwoBKAFxRs6kIyGFzXjnRDVbY5zn15mrIJqMhr9BEBF6798TA" +
// // // 		"qTG0FYTqkGK0D+FVfaWXvQ0u9Jw0KootS4kKHNwDmbZK2xYI2Ilt1ikeN9MMt8ZD" +
// // // 		"tXV4shTnQaYPty5Atr9Dzh052FTnlwsclVo33XHF2N2dfe7TaJYTaf5uXuh7Vkj3" +
// // // 		"99bKuvA8iLmEVe+i86L9K9LD5QEywO/sNcqQkNU8sQKBgQDkycMSv1dfSAgSZywU" +
// // // 		"hcJqFWJ3TpAP5aoWqw8Svill3Qs6/2zatU4XC4tzRX9KW437M4ORNee/XrCi9z4L" +
// // // 		"VOeOtR+gp/zn9DebBzaEdTfMlof+znPfx9fVIpBkpjezDgiFNkeeMnKTQrT/kULl" +
// // // 		"Zso29pCfgO/57L9Vi3pjiqKt0wKBgQDir4EDVesZTGcBstSRLIGrUQFmvxICPaMm" +
// // // 		"0PogvQhUv8UFMxx6nBl/ZB7AZWQH0TRlpVL7iqS2clO1dQOgrtgwgUV4m3Ml9Ivr" +
// // // 		"vI2fgCWzFudPst3oSu9Udc9Pq995Pl2nnwV612p3P6p3wgCaFjbouyyJZxNmHQj/" +
// // // 		"JPtv65nSSQKBgQDVcRnJorLLlHLbYF9yYfunZn3vWl7yRcvxy/KLBNewTZENoIAY" +
// // // 		"Zm8M9ttJVjvTziheg4ep8EVdduSJlOnQPoysyXNROYeril5aBlepKYY+Gu2THV5j" +
// // // 		"FpjYIZ/eFmf+Zwgx5xrXjq7vjZs4lnd3dvcOYec4t1yqqGE0WKR8uzjbuwKBgAYL" +
// // // 		"0FEadY7TLtwovOqyWTMMkhD/f6d3pWZfpIxC/nnkM4kT9+p9R2DSds+C5MwglFkx" +
// // // 		"s6jp5cLIAduRJ2udvj5s9EFnRAb7ItBC0zQx4s+ICNtjVe/gL8n86m6hkvBU7YKP" +
// // // 		"B0JjhH9xv0Y6cnGprgU/GM0BZs8ObzL+9YXirtOhAoGACk1TGT6I5oMHLDB4yOOe" +
// // // 		"a7dT+kAY6a2W2Sq2e0VN70EkVmXGC6ODpRzPcH7nojcMN5jk8QHHosWNK8DECwAj" +
// // // 		"uCGYn8G/0yAlhkddzE1+y1f5nVm+GCQTrdMMuqOwJnosifdoNDbWg4oGiBRt1uwI" +
// // // 		"aoxbWonlDAZaLC+8Bxe1Hss=" +
// // // 		"\n-----END PRIVATE KEY-----"
// // // 	r := strings.NewReader(privk)
// // // 	pemBytes, err := ioutil.ReadAll(r)
// // // 	if err != nil {
// // // 		log.Fatal(err)
// // // 	}

// // // 	block, _ := pem.Decode(pemBytes)
// // // 	log.Println(block)
// // // 	if block == nil {
// // // 		log.Println(block)
// // // 	}
// // // }
