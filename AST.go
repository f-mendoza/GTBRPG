package main

import (
  "fmt"
  "os"
)

func AST(){
  cls()
  fmt.Println("Cargando archivos...")
  checkFiles()
}

func checkFiles() {
    os.Mkdir("bin", 755)
    f, err := os.Create("./bin/enemies.xml")
    if err != nil {
      f.Close()
    } else {
      fmt.Println(err)
    }
    f, err = os.Create("./bin/player.xml")
    if err != nil {
      f.Close()
    } else {
      fmt.Println(err)
    }
    f, err = os.Create("./bin/objects.xml")
    if err != nil {
      f.Close()
    } else {
      fmt.Println(err)
    }
}
