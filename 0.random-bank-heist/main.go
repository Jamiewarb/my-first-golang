package main

import (
  "fmt"
  "math/rand"
  "time"
)

func main() {
  rand.Seed(time.Now().UnixNano())

  isHeistOn := true
  eludedGuards := rand.Intn(100) >= 50
  openedVault := rand.Intn(100) >= 70
  leftSafely := rand.Intn(5)

  if (eludedGuards) {
    fmt.Println("Looks like you've managed to make it past the guards. Good job, but remember, this is the first step.")
  } else {
    isHeistOn = false
    fmt.Println("Plan a better disguise next time?")
  }

  if (isHeistOn && openedVault) {
    fmt.Println("Grab and GO!")
  } else if (isHeistOn) {
    isHeistOn = false
    fmt.Println("You couldn't open the vault and are caught")
  }

  if (isHeistOn) {
    switch leftSafely {
      case 0:
        isHeistOn = false
        fmt.Println("You failed to escape the bank and are caught by police")
      case 1:
        isHeistOn = false
        fmt.Println("You take a wrong turn in your car and are caught by police")
      case 2:
        isHeistOn = false
        fmt.Println("Your car won't start and you are arrested by police")
      case 3:
        isHeistOn = false
        fmt.Println("You are betrayed by a gang member and killed before you can escape")
      case 4:
        fmt.Println("Start the getaway car!")
    }
  }

  if (isHeistOn) {
    amtStolen := 10000 + rand.Intn(1000000)
    fmt.Printf("Well done, you stoled %d\n", amtStolen)
  }

  fmt.Println(isHeistOn)
}

