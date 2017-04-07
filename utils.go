package main

import (
  "fmt"
  "strings"
  "math/rand"
  "time"
  "os"
  "os/exec"
)

var answer string

func question(question string, options ...string) string {
  for {
    fmt.Printf(question)
    fmt.Scanln(&answer)
    answer = strings.ToLower(answer)
    for _, check := range options {
      if answer == check {
        return answer
      } else {
        fmt.Println("Respuesta incorrecta intentalo de nuevo.")
        break
      }
    }
  }
}

func logicQuestion(question string) bool {
  for {
    fmt.Printf(question)
    fmt.Scanln(&answer)
    answer = strings.ToLower(answer)
    if answer == "si" {
      return true
    } else if answer == "no" {
      return false
    } else {
      fmt.Println("Respuesta incorrecta, intentalo de nuevo")
    }
  }
}

func randomNumber(max int) int {
  seed := rand.NewSource(time.Now().UnixNano())
  randomNumber := rand.New(seed)
  return randomNumber.Intn(max)
}

func msg(name string, input string) {
  fmt.Printf(name + ": ")
  for _, c := range input {
    fmt.Printf("%c", c)
    time.Sleep(50 * time.Millisecond)
  }
}

func cls() {
  cmd := exec.Command("cmd", "/c", "cls")
  cmd.Stdout = os.Stdout
  cmd.Run()
}

func comprobe(err error) bool{
  if err != nil {
    fmt.Printf("Ha ocurrido un error.")
    panic(err)
  }
  return true
}
