package main
import (
         "io"
         "net/http"
)
func hello(w http.ResponseWriter, r *http.Request){
        io.WriteString("<body style='background-color: yellow'>Hello World</body>")
}
func main(){
        http.HandleFunc("/",hello)
        http.ListenAndServe(":80",nil)
}
