package main

import (
    "net/http"
	"./Routes/Home"
	"fmt" // Import your routes package
)

func main() {
    // Setup routes from the routes package
    Routes.SetupRoutes()
    // // Start the HTTP server
    http.ListenAndServe(":8080", nil)
	fmt.Println("Server is running on port 8080")
    fmt.Println(ss)

}
