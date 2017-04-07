package main

import (
  "fmt"
)

func main(){
  AST()
  cls()
  var option uint8
  fmt.Println("Bienvenido a RPGame. v0.01\nElige una de las primeras opciones para comenzar")
  fmt.Println("1) Nuevo juego\n2) Cargar juego\n3) Salir.")
  fmt.Printf("Opcion: ")
  fmt.Scanln(&option)
  switch option {
  case 1:
    newGame()
  case 2:
    //loadGame()
  case 3:
    //exit()
  default:
    fmt.Println("Esa opcion no esta disponible")
  }
}
