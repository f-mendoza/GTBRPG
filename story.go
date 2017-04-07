package main

import(
  "fmt"
)

func newGame(){
  name := chooseName()
  introduction(name)
}

func chooseName() string {
  var name string
  msg("Guardia de la puerta", "Hola, ermitaño. ¿Cual es tu nombre?\n")
  for name == "" {
    fmt.Printf("Nombre: ")
    fmt.Scanln(&name)
  }
  return name
}
func introduction(name string) {
  cls()
  welcoming := "Hola " + name + ", bienvenido a Sidón.\n"
  msg("Guardia de la puerta", welcoming)
  msg("Guardia de la puerta", "Si eres valiente, puedes aventurarte directamente en la mazmorra en busca de riquezas.\n")
  msg("Guardia de la puerta", "Puedes visitar la tienda, donde podras comprar objetos y equipo para tu travesia en la mazmorra.\n")
  msg("Guardia de la puerta", "Si estas agotado luego de pasar un rato en la mazmorra, puedes alquilar una habitacion en el motel de Haley por 50 de oro, donde tus puntos de HP seran restaurados completamente.\n")
  goToTown()
}
