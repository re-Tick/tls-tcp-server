package client

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// func ClientCall() {
// 	// _, err := http.Get("https://localhost:8080/ritik")
// 	// if err != nil {
// 	// 	log.Panic("failed to simulate a request. error: ", err)
// 	// }
// 	println("client called")

// 	const rootPEM = `
// -- GlobalSign Root R2, valid until Dec 15, 2021
// -----BEGIN CERTIFICATE-----
// MIIDujCCAqKgAwIBAgILBAAAAAABD4Ym5g0wDQYJKoZIhvcNAQEFBQAwTDEgMB4G
// A1UECxMXR2xvYmFsU2lnbiBSb290IENBIC0gUjIxEzARBgNVBAoTCkdsb2JhbFNp
// Z24xEzARBgNVBAMTCkdsb2JhbFNpZ24wHhcNMDYxMjE1MDgwMDAwWhcNMjExMjE1
// MDgwMDAwWjBMMSAwHgYDVQQLExdHbG9iYWxTaWduIFJvb3QgQ0EgLSBSMjETMBEG
// A1UEChMKR2xvYmFsU2lnbjETMBEGA1UEAxMKR2xvYmFsU2lnbjCCASIwDQYJKoZI
// hvcNAQEBBQADggEPADCCAQoCggEBAKbPJA6+Lm8omUVCxKs+IVSbC9N/hHD6ErPL
// v4dfxn+G07IwXNb9rfF73OX4YJYJkhD10FPe+3t+c4isUoh7SqbKSaZeqKeMWhG8
// eoLrvozps6yWJQeXSpkqBy+0Hne/ig+1AnwblrjFuTosvNYSuetZfeLQBoZfXklq
// tTleiDTsvHgMCJiEbKjNS7SgfQx5TfC4LcshytVsW33hoCmEofnTlEnLJGKRILzd
// C9XZzPnqJworc5HGnRusyMvo4KD0L5CLTfuwNhv2GXqF4G3yYROIXJ/gkwpRl4pa
// zq+r1feqCapgvdzZX99yqWATXgAByUr6P6TqBwMhAo6CygPCm48CAwEAAaOBnDCB
// mTAOBgNVHQ8BAf8EBAMCAQYwDwYDVR0TAQH/BAUwAwEB/zAdBgNVHQ4EFgQUm+IH
// V2ccHsBqBt5ZtJot39wZhi4wNgYDVR0fBC8wLTAroCmgJ4YlaHR0cDovL2NybC5n
// bG9iYWxzaWduLm5ldC9yb290LXIyLmNybDAfBgNVHSMEGDAWgBSb4gdXZxwewGoG
// 3lm0mi3f3BmGLjANBgkqhkiG9w0BAQUFAAOCAQEAmYFThxxol4aR7OBKuEQLq4Gs
// J0/WwbgcQ3izDJr86iw8bmEbTUsp9Z8FHSbBuOmDAGJFtqkIk7mpM0sYmsL4h4hO
// 291xNBrBVNpGP+DTKqttVCL1OmLNIG+6KYnX3ZHu01yiPqFbQfXf5WRDLenVOavS
// ot+3i9DAgBkcRcAtjOj4LaR0VknFBbVPFd5uRHg5h6h+u/N5GJG79G+dwfCMNYxd
// AfvDbbnvRG15RjF+Cv6pgsH/76tuIMRQyV+dTZsXjAzlAcmgQWpzU/qlULRuJQ/7
// TBj0/VLZjmmx6BEP3ojY+x1J96relc8geMJgEtslQIxq/H5COEBkEveegeGTLg==
// -----END CERTIFICATE-----`

// 	// First, create the set of root certificates. For this example we only
// 	// have one. It's also possible to omit this in order to use the
// 	// default root set of the current operating system.
// 	roots := x509.NewCertPool()
// 	ok := roots.AppendCertsFromPEM([]byte(rootPEM))
// 	if !ok {
// 		panic("failed to parse root certificate")
// 	}

// 	conn, err := tls.Dial("tcp", "localhost:8080", &tls.Config{
// 		// RootCAs: roots,
// 		InsecureSkipVerify: true,
// 		ServerName:         "ritik.com",
// 	})
// 	if err != nil {
// 		panic("failed to connect: " + err.Error())
// 	}
// 	defer conn.Close()
// 	conn.Write([]byte("Hii ritik"))
// }

// ----------2 ---------------------

func ClientCall() {
	config := &tls.Config{
		InsecureSkipVerify: true,
	}

	// Create a transport with the TLS configuration
	transport := &http.Transport{
		TLSClientConfig: config,
	}

	// Create an HTTP client with the custom transport
	client := &http.Client{
		Transport: transport,
	}

	// Make a GET request to the server
	response, err := client.Get("https://localhost:8080")
	if err != nil {
		log.Fatal("Error making GET request:", err)
	}
	defer response.Body.Close()

	// Read the response body
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal("Error reading response body:", err)
	}

	// Print the response
	fmt.Println("Response:", string(body))
}
