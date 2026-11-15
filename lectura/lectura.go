
package lectura

import (
    "fmt"
    "os"
    "path/filepath"
)

// LeerEjemplo muestra el contenido del libro digital.
func LeerEjemplo() {
    contenido, err := os.ReadFile("lectura/ejemplo.txt")

    if err != nil {
        fmt.Println("No se pudo abrir el libro:", err)
        return
    }

    fmt.Println("\n--- LECTURA DEL LIBRO ---")
    fmt.Println(string(contenido))
}

// DescargarEjemplo guarda una copia local del libro.
func DescargarEjemplo() {
    contenido, err := os.ReadFile("lectura/ejemplo.txt")

    if err != nil {
        fmt.Println("No se pudo encontrar el libro:", err)
        return
    }

    // Crear la carpeta de descargas.
    err = os.MkdirAll("descargas", 0755)

    if err != nil {
        fmt.Println("No se pudo crear la carpeta:", err)
        return
    }

    // Guardar una copia del archivo.
    destino := filepath.Join("descargas", "ejemplo.txt")

    err = os.WriteFile(destino, contenido, 0644)

    if err != nil {
        fmt.Println("No se pudo guardar el libro:", err)
        return
    }

    fmt.Println("Libro guardado correctamente en:", destino)
}
