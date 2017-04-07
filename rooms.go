package main

import (
  "fmt"
)

func goForward(salas int) {
  salas++
  generateRoom()
}



func generateRoom() {
  var logicAnswer bool
  var randomTreasureNumber, room int
  //var respuesta string
  room = randomNumber(2)
  switch room {
  case 0:
    fmt.Println("Te encuentras en un pasillo con 2 ventanas de cada lado, al final de este hay un cofre.\n")
    logicAnswer = logicQuestion("¿Quieres abrirlo?")
    if logicAnswer {
      randomTreasureNumber = randomNumber(100)
      if randomTreasureNumber < 50 {
        fmt.Println("Has encontrado una pocion")
      } else {
        fmt.Println("¡ERA UNA TRAMPA! Has perdido 5 HP")
      }
    }
  case 1:
    fmt.Println("Entras una sala donde un orco te espera armado y listo para luchar.\n")
    answer = question("¿Que haras?", "luchar", "escapar")
  }
}
