
package administracion

import (
    "fmt"

    "github.com/aralfaomega121-lab/BookManager/libros"
    "github.com/aralfaomega121-lab/BookManager/usuarios"
)

// MostrarEstadisticas presenta un resumen del sistema.
func MostrarEstadisticas() {
    totalLibros := libros.ContarLibros()
    totalUsuarios := usuarios.ContarUsuarios()

    fmt.Println("\n--- ESTADISTICAS DE BOOKMANAGER ---")
    fmt.Println("Libros registrados:", totalLibros)
    fmt.Println("Usuarios registrados:", totalUsuarios)
}
