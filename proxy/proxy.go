package main

import "fmt"

// subject
type server interface {
	handleRequest(string, string) (int, string)
}

// real subject
type application struct {
}

func (a *application) handleRequest(url, method string) (int, string) {
	if url == "/app/status" && method == "GET" {
		return 200, "OK"
	}

	if url == "/create/user" && method == "POST" {
		return 201, "User Created"
	}

	return 404, "Not Found"
}

// proxy
type nginx struct {
	application       *application
	maxAllowedRequest int
	rateLimiter       map[string]int
}

func (n *nginx) checkRateLimiting(url string) bool {
	if n.rateLimiter[url] == 0 {
		n.rateLimiter[url] = 1
	}

	if n.rateLimiter[url] > n.maxAllowedRequest {
		return false
	}

	n.rateLimiter[url] = n.rateLimiter[url] + 1
	return true
}

// client will call handleRequest in proxy instead of in real subject
func (n *nginx) handleRequest(url, method string) (int, string) {
	isAllow := n.checkRateLimiting(url)
	if !isAllow {
		return 403, "Not Allowed"
	}

	return n.application.handleRequest(url, method)
}

func newNginxServer() *nginx {
	return &nginx{
		application:       &application{},
		maxAllowedRequest: 2,
		rateLimiter:       make(map[string]int),
	}
}

func main() {
	nginxServer := newNginxServer()
	appStatusUrl := "/app/status"
	createUserUrl := "/create/user"

	httpCode, body := nginxServer.handleRequest(appStatusUrl, "GET")
	fmt.Printf("\nURL: %s\nMethod: %s\nHttpCode: %d\nBody: %s\n", appStatusUrl, "GET", httpCode, body)

	httpCode, body = nginxServer.handleRequest(appStatusUrl, "GET")
	fmt.Printf("\nURL: %s\nMethod: %s\nHttpCode: %d\nBody: %s\n", appStatusUrl, "GET", httpCode, body)

	httpCode, body = nginxServer.handleRequest(appStatusUrl, "GET")
	fmt.Printf("\nURL: %s\nMethod: %s\nHttpCode: %d\nBody: %s\n", appStatusUrl, "GET", httpCode, body)

	httpCode, body = nginxServer.handleRequest(createUserUrl, "POST")
	fmt.Printf("\nURL: %s\nMethod: %s\nHttpCode: %d\nBody: %s\n", createUserUrl, "POST", httpCode, body)

	httpCode, body = nginxServer.handleRequest(createUserUrl, "GET")
	fmt.Printf("\nURL: %s\nMethod: %s\nHttpCode: %d\nBody: %s\n", createUserUrl, "GET", httpCode, body)
}
