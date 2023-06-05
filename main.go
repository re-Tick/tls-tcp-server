package main

import (
	"proxy-tls-tcp/client"
	"proxy-tls-tcp/proxy"
	"time"
)

func main() {
	go func() {
		time.Sleep(5 * time.Second)
		client.ClientCall()
	}()
	proxy.StartProxy()

}

// import (
// 	"crypto/tls"
// 	"fmt"
// 	"io/ioutil"
// 	"log"
// 	"net"
// 	"net/http"
// 	"time"
// )

// func main() {

// 	// client
// 	go func() {
// 		time.Sleep(5 * time.Second)
// 		// Create a TLS configuration
// 		config := &tls.Config{
// 			InsecureSkipVerify: true,
// 		}

// 		// Create a transport with the TLS configuration
// 		transport := &http.Transport{
// 			TLSClientConfig: config,
// 		}

// 		// Create an HTTP client with the custom transport
// 		client := &http.Client{
// 			Transport: transport,
// 		}

// 		// Make a GET request to the server
// 		response, err := client.Get("https://localhost:8080")
// 		if err != nil {
// 			log.Fatal("Error making GET request:", err)
// 		}
// 		defer response.Body.Close()

// 		// Read the response body
// 		body, err := ioutil.ReadAll(response.Body)
// 		if err != nil {
// 			log.Fatal("Error reading response body:", err)
// 		}

// 		// Print the response
// 		fmt.Println("Response:", string(body))
// 	}()

// 	// Load TLS certificates
// 	cert, err := tls.LoadX509KeyPair("server.crt", "server.key")
// 	if err != nil {
// 		log.Fatal("Error loading TLS certificates:", err)
// 	}

// 	// Create a TLS configuration
// 	config := &tls.Config{
// 		Certificates: []tls.Certificate{cert},
// 	}

// 	// Create a TCP listener on a specific port
// 	listener, err := net.Listen("tcp", ":8080")
// 	if err != nil {
// 		log.Fatal("Error creating listener:", err)
// 	}

// 	// Wrap the listener with TLS
// 	// tlsListener := tls.NewListener(listener, config)

// 	fmt.Println("Server listening on port 8080...")

// 	// Accept incoming connections
// 	for {
// 		conn, err := listener.Accept()
// 		if err != nil {
// 			log.Println("Error accepting connection:", err)
// 			continue
// 		}
// 		tlsConn := tls.Server(conn, config)

// 		// Handle the connection in a separate goroutine
// 		go handleConnection(tlsConn)
// 	}
// }

// func handleConnection(conn net.Conn) {
// 	defer conn.Close()

// 	// Read data from the client
// 	data := make([]byte, 1024)
// 	n, err := conn.Read(data)
// 	if err != nil {
// 		log.Println("Error reading data:", err)
// 		return
// 	}

// 	// Print the received data
// 	fmt.Println("Received:", string(data[:n]))

// 	// Send a response back to the client
// 	response := []byte("HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: 12\r\n\r\nHello World!")
// 	_, err = conn.Write(response)
// 	if err != nil {
// 		log.Println("Error sending response:", err)
// 		return
// 	}
// }
