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
				head:="<html><head><title>"+r.Host+"</title>"
				style:="<style>li a{display:block;color:white;text-align:center;padding:14px 16px;text-decoration:none;}li{float:left;}ul{overflow:hidden;background-color:#333;list-style-type:none;margin:0;padding:0;}</style>"
				nav:="</head><body><nav><ul><li><a href=''>"+r.Host+"</a></li></ul></nav>"
				body:= "<main>"+r.Host+"</main>"
				footer:="</body></html>"
				fmt.Fprintf(w, head+ style+ nav + body + footer)
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
