package main

import
(
	"net/http"
	"./Routes"
	"fmt"
)
// typecasting in go
// typecasting is the process of converting one data type to another data type
// you go int(a)



func main() {
	mux := Routes.NewRouter()
	mux.HandleFunc("/", Routes.Home)

	http.ListenAndServe(":8080", mux)
	fmt.Println("Server is running on port 8080")

}
