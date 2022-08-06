package main
import(
	"fmt"
	"os"
	"net/http"
)
//underlinig the http.Get function
type hotdog int

// type handler
func(c hotdog) serveHTTP(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello World")
}

func main(){
var Hello hotdog
	http.ListenAndServe(":8080", Hello)

}

//underlining the http.HandleFunc function






