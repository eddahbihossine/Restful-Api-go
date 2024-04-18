// create Home Route with GET method
package Routes

import (
	"fmt"
	"net/http"

)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}


