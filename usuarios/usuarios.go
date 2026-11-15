
package usuarios

import (
    "fmt"
    "strings"
)

// Usuario representa a una persona registrada.
type Usuario struct {
    Nombre string
    Correo string
}

// Lista de usuarios registrados.
var listaUsuarios []Usuario

// RegistrarUsuario agrega un usuario al sistema.
func RegistrarUsuario(nombre string, correo string) {
    nombre = strings.TrimSpace(nombre)
    correo = strings.TrimSpace(correo)

    if nombre == "" || correo == "" {
        fmt.Println("Debe ingresar nombre y correo.")
        return
    }

    for _, usuario := range listaUsuarios {
        if strings.EqualFold(usuario.Correo, correo) {
            fmt.Println("Este correo ya esta registrado.")
            return
        }
    }

    nuevo := Usuario{
        Nombre: nombre,
        Correo: correo,
    }

    listaUsuarios = append(listaUsuarios, nuevo)
    fmt.Println("Usuario registrado correctamente.")
}

// MostrarUsuarios presenta los usuarios registrados.
func MostrarUsuarios() {
    if len(listaUsuarios) == 0 {
        fmt.Println("No hay usuarios registrados.")
        return
    }

    fmt.Println("\n--- USUARIOS REGISTRADOS ---")

    for i, usuario := range listaUsuarios {
        fmt.Printf("%d. %s - %s\n",
            i+1, usuario.Nombre, usuario.Correo)
    }
}

 // ContarUsuarios devuelve la cantidad de usuarios.
func ContarUsuarios() int {
    return len(listaUsuarios)
}
