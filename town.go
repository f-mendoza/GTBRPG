package main

import (
  "fmt"
  "regexp"
  "strings"
)

var input string

func goToTown() {
  fmt.Println("¿Que deseas hacer?")
  for input == "" {
    fmt.Printf("Comando: ")
    fmt.Scanln(&input)
  }
  input = strings.ToLower(input)
  commandAnalize(input)
}

func commandAnalize(input string) {
  var (
    matched bool
    err error
  )
  matched, err = regexp.MatchString("(ir( a)?|go( to)?)", input)
  comprobe(err)
  if matched {
    matched, err = regexp.MatchString("(mazmorra|masmorra|dungeon)", input)
    comprobe(err)
    if matched {
      fmt.Println("Elegiste ir a la mazmorra")
      goToTown()
      //goToDungeon()
    } else {
      matched, err = regexp.MatchString("(tienda|shop)", input)
      comprobe(err)
      if matched {
        fmt.Println("Elegiste ir a la tienda")
        goToTown()
        //goToShop()
      } else {
        matched, err = regexp.MatchString("(motel|hotel)", input)
        comprobe(err)
        if matched {
          fmt.Println("Elegiste ir al motel")
          goToTown()
          //goToMotel()
        }
      }
    }
  } else {
    matched, err = regexp.MatchString("info", input)
    comprobe(err)
    if matched {
      matched, err = regexp.MatchString("(mazmorra|masmorra|dungeon)", input)
      comprobe(err)
      if matched {
        msg("Guardia de la puerta", "En la mazmorra, te enfrentaras con criaturas corruptas que una vez fueron parte de una civilizacion, ademas te encontraras oro y diferentes tesoros a medida que avanzas.")
        goToTown()
      } else {
        matched, err = regexp.MatchString("(tienda|shop)", input)
        comprobe(err)
        if matched {
          msg("Guardia de la puerta", "En la tienda, puedes comprar objetos, armaduras, armas y escudos para tu travesia en la mazmorra.")
          goToTown()
        } else {
          matched, err = regexp.MatchString("(motel|hotel)", input)
          comprobe(err)
          if matched {
            msg("Guardia de la puerta","En el motel de Haley, puedes alquilar una habitacion por 50 de oro, recuperaras todo tu HP y guardaras la partida.\nSi no tienes el oro suficiente para alquilar una habitacion, puedes simplemente guardar la partida.")
            goToTown()
          }
        }
      }
    }
  }
  if !matched {
    goToTown()
  }
}
