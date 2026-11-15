
package catalogo

import (
    "strings"

    "github.com/aralfaomega121-lab/BookManager/libros"
)

// BuscarPorTitulo permite encontrar libros
// utilizando una palabra del titulo.
func BuscarPorTitulo(texto string) []libros.Libro {
    var resultados []libros.Libro

    // Obtener todos los libros registrados.
    lista := libros.ObtenerLibros()

    // Revisar cada libro.
    for _, libro := range lista {
        titulo := strings.ToLower(libro.Titulo)
        busqueda := strings.ToLower(texto)

        if strings.Contains(titulo, busqueda) {
            resultados = append(resultados, libro)
        }
    }

    return resultados
}
