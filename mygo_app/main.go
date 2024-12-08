package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) > 0 {
		fmt.Printf("Servidor iniciado na porta :%s", args[0])
		http.ListenAndServe(":"+args[0], nil)
	}
	
	fmt.Println("Servidor iniciado na porta :8080")
	http.ListenAndServe(":8080", nil)
}
