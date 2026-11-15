
package libros

import "fmt"

// Libro guarda la informacion de un libro.
type Libro struct {
    Titulo string
    Autor  string
}

// Lista donde se guardan los libros.
var listaLibros []Libro

// AgregarLibro permite registrar un libro.
func AgregarLibro(titulo string, autor string) {
    if titulo == "" || autor == "" {
        fmt.Println("Debe ingresar titulo y autor.")
        return
    }

    nuevo := Libro{
        Titulo: titulo,
        Autor:  autor,
    }

    listaLibros = append(listaLibros, nuevo)
    fmt.Println("Libro registrado correctamente.")
}

// MostrarLibros presenta los libros registrados.
func MostrarLibros() {
    if len(listaLibros) == 0 {
        fmt.Println("No hay libros registrados.")
        return
    }

    fmt.Println("\n--- LIBROS REGISTRADOS ---")

    for i, libro := range listaLibros {
        fmt.Printf("%d. %s - %s\n",
            i+1, libro.Titulo, libro.Autor)
    }
}

 // ObtenerLibros devuelve una copia de los libros registrados.
func ObtenerLibros() []Libro {
    copia := make([]Libro, len(listaLibros))
    copy(copia, listaLibros)
    return copia
}
