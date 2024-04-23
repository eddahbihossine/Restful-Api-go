// create Home Route with GET method
package routes
package main

import (
    "net/http" // require http package for the http server
)

func HelloHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Hello, world!"))
}

func SetupRoutes() {
    fmt.Println("Server is running on port 8080")
    http.HandleFunc("/", HelloHandler)
}

func main() {
    SetupRoutes()
    http.ListenAndServe(":8080", nil)
}