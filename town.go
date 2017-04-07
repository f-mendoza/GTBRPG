package main

import (
  "fmt"
  "regexp"
  "strings"
)

func goToTown() {

  var input string
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
        fmt.Println("Info mazmorra")
        goToTown()
        //dungeonInfo()
      } else {
        matched, err = regexp.MatchString("(tienda|shop)", input)
        comprobe(err)
        if matched {
          fmt.Println("Info tienda")
          goToTown()
          //shopInfo()
        } else {
          matched, err = regexp.MatchString("(motel|hotel)", input)
          comprobe(err)
          if matched {
            fmt.Println("Info motel")
            goToTown()
            //motelInfo()
          }
        }
      }
    }
  }
  if !matched {
    goToTown()
  }
}
