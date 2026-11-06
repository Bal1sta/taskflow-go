package main

import (
	"fmt"
	"net/http"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "TaskFlow OK")
}

func main() {

	fs := http.FileServer(http.Dir("../frontend"))
	http.Handle("/", fs)
	
	http.HandleFunc("/health", healthHandler)
	fmt.Println("TaskFlow запущен на порту 8080")



	
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Произошла ошибка:", err)
	}

}