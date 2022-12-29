package main
import (
	"net/http"
	"fmt"
	//"strconv"
	// debugging
	//"reflect"
)
func main(){
		http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request){
			if r.Host == "localhost"{
				head:="<html><head><title>"+r.Host+"</title></head><body>"
				nav:="<body><nav><ul><li>"+ r.Host+"</li></ul></nav>"
				fmt.Println(r.Host)
				body:= "<main><h1>Testing</h1></main>"
				footer:="</body></html>"
				fmt.Fprintf(w, head+ nav + body + footer)
			}
			if r.Host == "dcescorts.tk"{
				head:="<html><head><title></title></head><body>"
				nav:="<body>"
				body:= r.Host
				footer:="</body></html>"
				fmt.Fprintf(w, head+ nav + body + footer)
			}
			if r.Host == "vaescorts.tk"{
				head:="<html><head><title></tile></head><body>"
				nav:="<body>"
				fmt.Println(string(r.Host))
				body:= string(r.Host)
				footer:="</body></html>"
				fmt.Fprintf(w, head+ nav + body + footer)
			}
		})
	http.ListenAndServe(":80", nil)
}
