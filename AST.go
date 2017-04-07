package main

import (
  "fmt"
  "os"
  //"time"
)

func AST(){
  cls()
  fmt.Println("Cargando archivos...")
  checkFiles()
}

func checkFiles() {
  //var err error
  //if _, err = os.Stat("./bin"); err != nil {
    os.Mkdir("bin", 755)
  //}
  //if _, err = os.Stat("./bin/enemies.xml"); err != nil {
    f, err := os.Create("./bin/enemies.xml")
    if err != nil {
      f.Close()
    } else {
      fmt.Println(err)
    }
    //comprobe(err)
  //}
  //if _, err = os.Stat("./bin/player.xml"); err != nil {
    f, err = os.Create("./bin/player.xml")
    if err != nil {
      f.Close()
    } else {
      fmt.Println(err)
    }
    //comprobe(err)
    //}
  //if _, err = os.Stat("./bin/objects.xml"); err != nil {
    f, err = os.Create("./bin/objects.xml")
    if err != nil {
      f.Close()
    } else {
      fmt.Println(err)
    }
  //}
}
