package proxy

// import (
// 	"crypto/tls"
// 	"fmt"
// 	"log"
// 	"net"
// )

// func handleConnection(conn net.Conn) {
// 	defer conn.Close()

// 	// Check if the connection is TLS (Secure)
// 	isTLS := false
// 	_, isTLS = conn.(*tls.Conn)
// 	req := []byte{}
// 	conn.Read(req)
// 	// br := bufio.NewReader(conn)
// 	// firstByte, err := br.Peek(1)
// 	// if err == nil && len(firstByte) > 0 && firstByte[0] == byte(tls.RecordTypeHandshake) {
// 	// 	isTLS = true
// 	// }
// 	println("the request message is : ", string(req))
// 	fmt.Printf("type of the connefction: %T", conn)
// 	if isTLS {
// 		println("tls enabled")
// 		// Handle TLS connection
// 		// Example: Use TLS server to handle secure traffic
// 		tlsConfig := &tls.Config{
// 			InsecureSkipVerify: true, // Change to false if you want to verify the server certificate
// 		}
// 		tlsConn := tls.Server(conn, tlsConfig)
// 		defer tlsConn.Close()

// 		// Perform TLS-specific operations or forwarding
// 		// ...

// 	} else {
// 		println("tls disabled")
// 		// Handle non-TLS connection
// 		// Example: Forwarding the traffic to a destination
// 		// destinationAddr := "targethost:targetport"
// 		// destinationConn, err := net.Dial("tcp", destinationAddr)
// 		// if err != nil {
// 		// 	log.Printf("Failed to connect to destination: %v", err)
// 		// 	return
// 		// }
// 		// defer destinationConn.Close()

// 		// // Copy traffic between the client and the destination
// 		// go io.Copy(destinationConn, conn)
// 		// io.Copy(conn, destinationConn)
// 	}
// }

// func StartProxy() {
// 	proxyAddr := ":8080" // Proxy listens on this address

// 	listener, err := net.Listen("tcp", proxyAddr)
// 	listener = tls.NewListener(listener, &tls.Config{})
// 	if err != nil {
// 		log.Fatalf("Failed to start proxy: %v", err)
// 	}
// 	defer listener.Close()

// 	log.Printf("Proxy started and listening on %s", proxyAddr)

// 	for {
// 		conn, err := listener.Accept()
// 		if err != nil {
// 			log.Printf("Failed to accept incoming connection: %v", err)
// 			continue
// 		}

// 		// Handle each connection in a separate goroutine
// 		go handleConnection(conn)
// 	}
// }

// -------------------------

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net"
	"time"
)

func StartProxy() {
	// Create a TCP listener on a specific port
	listener, err := net.Listen("tcp", ":8080")
	// listener = tls.NewListener(listener, &tls.Config{: })

	if err != nil {
		log.Fatal("Error creating listener:", err)
	}
	defer listener.Close()

	fmt.Println("Server listening on port 8080...")

	// Accept incoming connections
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error accepting connection:", err)
			continue
		}

		// Handle the connection in a separate goroutine
		go handleConnection(conn)
	}
}

type Conn struct {
	net.Conn
	r bufio.Reader
	// buffer       []byte
	// ReadComplete bool
	// pointer      int
}

func readBytes(reader io.Reader) ([]byte, error) {
	var buffer []byte

	for {
		// Create a temporary buffer to hold the incoming bytes
		buf := make([]byte, 1024)
		rand.Seed(time.Now().UnixNano())

		// Read bytes from the Reader
		n, err := reader.Read(buf)
		if err != nil && err != io.EOF {
			return nil, err
		}

		// Append the bytes to the buffer
		buffer = append(buffer, buf[:n]...)

		// If we've reached the end of the input stream, break out of the loop
		if err == io.EOF || n != 1024 {
			break
		}
	}

	return buffer, nil
}

func (c *Conn) Read(b []byte) (n int, err error) {
	return c.r.Read(b)

	// buffer, err := readBytes(c.conn)
	// if err != nil {
	// 	return 0, err
	// }
	// b = buffer
	// if len(buffer) == 0 && len(c.buffer) != 0 {
	// 	b = c.buffer
	// } else {
	// 	c.buffer = buffer
	// }
	// if c.ReadComplete {
	// 	b = c.buffer[c.pointer:(c.pointer + len(b))]

	// 	return 257, nil
	// }
	// n, err = c.Conn.Read(b)
	// if n > 0 {
	// 	c.buffer = append(c.buffer, b...)
	// }
	// if err != nil {
	// 	return n, err
	// }

	// return n, nil
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	// TODO: Since, net.Conn is an interface, then we can wrap the net.conn to overwrite the read function. This will help in reading from connection multiple times.

	reader := bufio.NewReader(conn)
	// reader2 := bufio.NewReader(conn)

	// Create a buffer to store the peeked data
	// buffer := make([]byte, 5)

	// // Peek at the incoming data without consuming it
	// n, err := reader.Peek(len(buffer))
	// if err != nil {
	// 	log.Fatal("Error peeking:", err)
	// }

	// Inspect the initial data received to determine if it's a TLS handshake
	initialData := make([]byte, 5)
	buffer, err := reader.Peek(len(initialData))
	if err != nil {
		log.Println("Error reading initial data:", err)
		return
	}

	// initialData2 := make([]byte, 5)
	// buffer2, err := reader2.Peek(len(initialData2))
	// print(buffer2)
	// if err != nil {
	// 	log.Panic(err)
	// }
	// reader.Reset(conn)

	// _, err := conn.Read(buffer)
	// if err != nil {
	// 	log.Panic("failed to read the request message", err)
	// 	return
	// }
	// connWrap := Conn{Conn: conn, buffer: []byte{}, pointer: 0, ReadComplete: false}
	// buffer, err := readBytes(&connWrap)
	// if err != nil {
	// 	return
	// }
	// connWrap.ReadComplete = true
	// fmt.Println("buffer in the isTLSHandshake: ", string(buffer), len(buffer))

	// Check if the initial data indicates a TLS handshake
	isTLS := isTLSHandshake(buffer)

	connWrapped := Conn{r: *reader, Conn: conn}
	if isTLS {
		// Handle TLS connection
		handleTLSConnection(&connWrapped)
	} else {
		// Handle plain TCP connection
		handlePlainTCPConnection(conn)
	}
}

func isTLSHandshake(data []byte) bool {
	if len(data) < 5 {
		return false
	}
	fmt.Println("buffer in the isTLSHandshake: ", string(data))

	return data[0] == 0x16 && data[1] == 0x03 && (data[2] == 0x00 || data[2] == 0x01 || data[2] == 0x02 || data[2] == 0x03)
}

func handleTLSConnection(conn net.Conn) {
	fmt.Println("Handling TLS connection from", conn.RemoteAddr().String())

	// Load TLS certificates
	cert, err := tls.LoadX509KeyPair("server.crt", "server.key")
	if err != nil {
		log.Fatal("Error loading TLS certificates:", err)
	}

	// Create a TLS configuration
	config := &tls.Config{
		// Set up your TLS configuration options here
		// For example, you can load certificates from files or use Let's Encrypt
		// Refer to the documentation for more details:
		// https://pkg.go.dev/crypto/tls#Config
		// InsecureSkipVerify: true,
		Certificates: []tls.Certificate{cert},
	}

	// Wrap the TCP connection with TLS
	tlsConn := tls.Server(conn, config)

	req := make([]byte, 1024)
	fmt.Println("before the parsed req: ", string(req))

	_, err = tlsConn.Read(req)
	if err != nil {
		log.Panic("failed reading the request message with error: ", err)
	}
	fmt.Println("after the parsed req: ", string(req))
	// Perform the TLS handshake
	// err = tlsConn.Handshake()
	// if err != nil {
	// 	log.Println("Error performing TLS handshake:", err)
	// 	return
	// }

	// Use the tlsConn for further communication
	// For example, you can read and write data using tlsConn.Read() and tlsConn.Write()

	// Here, we simply close the connection
	tlsConn.Close()
}

func handlePlainTCPConnection(conn net.Conn) {
	fmt.Println("Handling plain TCP connection from", conn.RemoteAddr().String())

	// Use the conn for plain TCP communication
	// For example, you can read and write data using conn.Read() and conn.Write()

	// Here, we simply close the connection
	conn.Close()
}
