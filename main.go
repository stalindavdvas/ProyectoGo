package main

import (
    "fmt"
    "html/template"
    "net/http"
    "strings"
)

func main() {
    http.HandleFunc("/", handleIndex)
    http.HandleFunc("/count", handleCount)
    fmt.Println("Servidor escuchando en http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}

func handleIndex(w http.ResponseWriter, r *http.Request) {
    tmpl := template.Must(template.ParseFiles("templates/index.html"))
    tmpl.Execute(w, nil)
}

func handleCount(w http.ResponseWriter, r *http.Request) {
    if r.Method == http.MethodPost {
        err := r.ParseForm()
        if err != nil {
            http.Error(w, "Error al procesar el formulario", http.StatusBadRequest)
            return
        }
        text := r.FormValue("text")
        wordCount := len(strings.Fields(text))
        fmt.Fprintf(w, "El número de palabras ingresadas es: %d", wordCount)
    } else {
        http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
    }
}
