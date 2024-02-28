
package main

import (
	"fmt"
	"net/http"
)

func main(){
	http.HandleFunc("/", 
		func(w http.ResponseWriter, r *http.Request){
			fmt.Fprintf(w, "Aloo world on %s\n",r.URL.Path);	
		});
		
	fileServer := http.FileServer(http.Dir("files/"));
	http.Handle("files/", http.StripPrefix("files/", fileServer));

	http.ListenAndServe(":80", nil);
}
