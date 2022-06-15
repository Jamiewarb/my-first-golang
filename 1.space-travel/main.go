package main

import "fmt"

func fuelGauge(fuel int) {
  fmt.Printf("You've got %d fuel units left\n", fuel)
}

func calculateFuel(planet string) int {
  switch planet {
    case "venus":
      return 300000
    case "mercury":
      return 500000
    case "mars":
      return 700000
    default:
      return 0
  }
}

func greetPlanet(planet string) {
  fmt.Printf("Good evening passengers. We will shortly be departing for %s.\n", planet)
}

func cantFly() {
  fmt.Println("We do not have the available fuel to fly there.")
}

func flyToPlanet(planet string, fuel int) int {
  fuelRemaining, fuelCost := fuel, calculateFuel(planet)

  if (fuelRemaining >= fuelCost) {
    greetPlanet(planet)
    fuelRemaining -= fuelCost
  } else {
    cantFly()
  }

  return fuelRemaining
}

func main() {
  fuel := 1000000
  planetChoice := "venus"

  fuel = flyToPlanet(planetChoice, fuel)
  fuelGauge(fuel)
}
