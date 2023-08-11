package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {

	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
		}

		go handle(conn)
	}
}

func handle(conn net.Conn) {
	request(conn)
	defer conn.Close()
}

func request(conn net.Conn) {

	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			mux(conn, ln)
		}
		if ln == "" {
			break
		}
		i++
	}
}

func mux(conn net.Conn, ln string) {

	method := strings.Fields(ln)[0]
	fmt.Println("method: ", method)
	uri := strings.Fields(ln)[1]
	fmt.Println("uri: ", uri)

	if method == "GET" {
		switch uri {
		case "/":
			index(conn)
		case "/about":
			about(conn)
		case "/contact":
			contact(conn)
		case "/apply":
			apply(conn)
		case "/faq":
			faq(conn)
		default:
			notFound(conn)
		}
	}
}

func index(conn net.Conn) {

	body := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title></title></head><body><strong>Hello World <br/></strong>
	<a href="/about">About</a><br/>
	<a href="/contact">Contact</a><br/>
	<a href="/apply">Apply</a><br/>
	<a href="/faq">FAQ</a><br/>
	</body></html>`
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func about(conn net.Conn) {
	body := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title></title></head><body><strong>About <br/></strong>
	<a href="/">Index</a><br/>
	<a href="/contact">Contact</a><br/>
	<a href="/apply">Apply</a><br/>
	<a href="/faq">FAQ</a><br/>
	</body></html>`
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func contact(conn net.Conn) {
	body := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title></title></head><body><strong>Contact <br/></strong>
	<a href="/">Index</a><br/>
	<a href="/about">About</a><br/>
	<a href="/apply">Apply</a><br/>
	<a href="/faq">FAQ</a><br/>
	</body></html>`
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func apply(conn net.Conn) {
	body := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title></title></head><body><strong>Apply <br/></strong>
	<a href="/">Index</a><br/>
	<a href="/about">About</a><br/>
	<a href="/contact">Contact</a><br/>
	<a href="/faq">FAQ</a><br/>
	</body></html>`
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func faq(conn net.Conn) {

	body := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title></title></head><body><strong>FAQ <br/></strong>
	<a href="/">Index</a><br/>
	<a href="/about">About</a><br/>
	<a href="/contact">Contact</a><br/>
	<a href="/apply">Apply</a><br/>
	</body></html>`
	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func notFound(conn net.Conn) {
	body := `<!DOCTYPE html><html lang="en"><head><meta charset="UTF-8"><title></title></head><body><strong>Not Found <br/></strong>
	<a href="/">Index</a><br/>
	<a href="/about">About</a><br/>
	<a href="/contact">Contact</a><br/>
	<a href="/apply">Apply</a><br/>
	<a href="/faq">FAQ</a><br/>
	</body></html>`
	fmt.Fprint(conn, "HTTP/1.1 404 Not Found\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}
