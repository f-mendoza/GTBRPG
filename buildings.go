package main

import (
  "fmt"
  "regexp"
)

func goToShop() {
   msg("Vendedor", "Bienvenido, ¿Que deseas comprar?")
}

func goToMotel() {
  var input string
  msg("Haley", "Hola, ¿Quieres alquilar una habitacion por 50 de oro o solamente guardar la partida?\n")
  fmt.Printf("Comando: ")
  fmt.Scanf(input)
  matched, err := regexp.MatchString("(alquilar (una )?habitacion|rent (a )?room)", input)
  comprobe(err)
  if matched /*|| player.gold >= 50 */{
    //player.gold -= 50
    fmt.Println("Has dormido y restaurado tus puntos de HP.")
    //player.health = player.maxHealth
  } else {
    matched, err = regexp.MatchString("(guardar( partida)?|save( game)?|savegame)", input)
    comprobe(err)
    if matched {
      //saveGame()
    }
  }
}
