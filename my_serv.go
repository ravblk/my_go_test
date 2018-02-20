package main

import (
"net/http"
"os"
"mime"
"path"
"log"
)


func requestHandler(w http.ResponseWriter, r *http.Request) {

	file := r.URL.Path
	base :=  "D:/my_serv/www"
	if file == "/" {
		file = "/index.html"
	}

	reqFile, err := os.OpenFile(base + file, os.O_RDONLY, 0)
	if err != nil {
	        log.Printf("https://127.0.0.1:8080 ERROR ",err,"file ",file)
		file = "/404.html"
		reqFile, err = os.OpenFile(base + file, os.O_RDONLY, 0)
		w.WriteHeader(http.StatusNotFound)

	} else {
		w.WriteHeader(http.StatusOK)	
                log.Printf("request file ",file)	
	}
	fi, err := reqFile.Stat()
	contentType := mime.TypeByExtension(path.Ext(file));
	bytes := make([]byte, fi.Size())
	reqFile.Read(bytes)

	w.Header().Set("Content-Type", contentType)
	w.Header().Set("Server", "my_serv/1.0")
	w.Write(bytes)
}
 
func main(){
	http.HandleFunc("/", requestHandler)
	http.ListenAndServe(":8080", nil)
}