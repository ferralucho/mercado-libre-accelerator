package main

import (
	"fmt"
	"net/http"
	"sync"
)

func miFuncion(wg *sync.WaitGroup) {
	fmt.Println("Dentro de la goroutine")
	wg.Done()
}
/*
func main() {
	fmt.Println("Inicio del programa")
	var wg sync.WaitGroup
	wg.Add(1)
	//cuenta cuantos hilos de ejecucion espera que todos los procesos terminen
	go miFuncion(&wg)
	wg.Wait()
	fmt.Printf("Fin del programa")
}

 */

/*
func main() {
    fmt.Println("Hello World")

    var waitgroup sync.WaitGroup
    waitgroup.Add(1)
    go func() {
        fmt.Println("Inside my goroutine")
        waitgroup.Done()
    }()
    waitgroup.Wait()

    fmt.Println("Finished Execution")
}
 */

/*
go func(url string) {
  fmt.Println(url)
}(url)
 */

var urls = []string {
	"https://www.google.com",
	"https://www.lavoz.com.ar",
	"https://www.mercadolibre.com",
}

func recuperar(url string, wg *sync.WaitGroup) {
	fmt.Println(url)
	res, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	wg.Done()
	fmt.Println(res.Status)
}

func enviarRequest(w http.ResponseWriter, r *http.Request){
	fmt.Println("Enviamos request al endpoint")
	var waitgroup sync.WaitGroup

	for _, url := range urls {
		waitgroup.Add(1)
		go recuperar(url, &waitgroup)
	}
	waitgroup.Wait()
	fmt.Println("Devuelve una respuesta")
	fmt.Println("Proceso terminado")
	fmt.Fprint(w, "Proceso terminado")
}

func handleRequest() {
	http.HandleFunc("/", enviarRequest)
	http.ListenAndServe(":8080", nil)
}

func main() {
	handleRequest()
}
