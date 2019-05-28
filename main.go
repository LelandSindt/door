package main

import (
  "os"
  "github.com/stianeikeland/go-rpio"
  "fmt"
  "time"
  "net/http"
)

var (
  frontDoor = rpio.Pin(4)
  garageDoor = rpio.Pin(17)
)

func setLeds(state rpio.State) {
  if state == 0 {
    _, _ = http.Get("http://192.168.1.93/setled?pixel=all&red=0&green=0&blue=0")
  } else {
    _, _ = http.Get("http://192.168.1.93/setled?pixel=all&red=255&green=255&blue=255")
  }
}

func main() {
  // Open and map memory to access gpio, check for errors
  if err := rpio.Open(); err != nil {
    fmt.Println(err)
    os.Exit(1)
  }
  // Unmap gpio memory when done
  defer rpio.Close()

  frontDoor.Input()
  frontDoor.PullUp()
  garageDoor.Input()
  garageDoor.PullUp()
  // get initial state
  frontDoorState :=  frontDoor.Read()
  frontDoorPreviousState := frontDoorState
  garageDoorState := garageDoor.Read()
  garageDoorPreviousState := garageDoorState
  fmt.Printf(time.Now().String() + " Front Door Open: %d \n", frontDoorState)
  fmt.Printf(time.Now().String() + " Garage Door Open: %d \n", garageDoorState)
  setLeds(garageDoorState)
  for {
    frontDoorState = frontDoor.Read()
    if frontDoorState != frontDoorPreviousState {
      fmt.Printf(time.Now().String() + " Front Door Open: %d \n", frontDoorState)
      frontDoorPreviousState = frontDoorState
    }
    garageDoorState = garageDoor.Read()
    if garageDoorState != garageDoorPreviousState {
      fmt.Printf(time.Now().String() + " Garage Door Open: %d \n", garageDoorState)
      garageDoorPreviousState = garageDoorState
      setLeds(garageDoorState)
    }
    time.Sleep(100 * time.Millisecond)
  }
}
