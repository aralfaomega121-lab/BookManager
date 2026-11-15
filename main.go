
package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"

    "github.com/aralfaomega121-lab/BookManager/catalogo"
    "github.com/aralfaomega121-lab/BookManager/libros"
)

func main() {
    lector := bufio.NewScanner(os.Stdin)

    for {
        fmt.Println("\n===== BOOKMANAGER =====")
        fmt.Println("Sistema de Gestion de Libros Electronicos")
        fmt.Println("1. Gestion de usuarios")
        fmt.Println("2. Gestion de libros")
        fmt.Println("3. Busqueda y catalogo")
        fmt.Println("4. Lectura y descarga")
        fmt.Println("5. Administracion")
        fmt.Println("0. Salir")
        fmt.Print("Seleccione una opcion: ")

        if !lector.Scan() {
            return
        }

        opcion := strings.TrimSpace(lector.Text())

        switch opcion {
        case "1":
            fmt.Println("Modulo de usuarios pendiente.")

        case "2":
            fmt.Println("\n--- GESTION DE LIBROS ---")
            fmt.Println("1. Agregar libro")
            fmt.Println("2. Mostrar libros")
            fmt.Print("Seleccione una opcion: ")

            if !lector.Scan() {
                return
            }

            subopcion := strings.TrimSpace(lector.Text())

            switch subopcion {
            case "1":
                fmt.Print("Ingrese el titulo: ")
                if !lector.Scan() {
                    return
                }
                titulo := strings.TrimSpace(lector.Text())

                fmt.Print("Ingrese el autor: ")
                if !lector.Scan() {
                    return
                }
                autor := strings.TrimSpace(lector.Text())

                libros.AgregarLibro(titulo, autor)

            case "2":
                libros.MostrarLibros()

            default:
                fmt.Println("Opcion no valida.")
            }

        
case "3":
    fmt.Println("\n--- BUSQUEDA DE LIBROS ---")
    fmt.Print("Ingrese el titulo o una palabra: ")

    if !lector.Scan() {
        return
    }

    texto := strings.TrimSpace(lector.Text())

    if texto == "" {
        fmt.Println("Debe ingresar una palabra.")
        continue
    }

    resultados := catalogo.BuscarPorTitulo(texto)

    if len(resultados) == 0 {
        fmt.Println("No se encontraron libros.")
    } else {
        fmt.Println("\nLibros encontrados:")

        for i, libro := range resultados {
            fmt.Printf("%d. %s - %s\n",
                i+1, libro.Titulo, libro.Autor)
        }
    }


        case "4":
            fmt.Println("Modulo de lectura pendiente.")

        case "5":
            fmt.Println("Modulo de administracion pendiente.")

        case "0":
            fmt.Println("Gracias por utilizar BookManager.")
            return

        default:
            fmt.Println("Opcion no valida.")
        }
    }
}
