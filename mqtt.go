package main

import (
  "gobot.io/x/gobot"
  "gobot.io/x/gobot/platforms/mqtt"
  "fmt"
  "time"
)

func main() {
  mqttAdaptor := mqtt.NewAdaptor("tcp://0.0.0.0:1883", "pinger")

  work := func() {
    mqttAdaptor.On("hello", func(msg mqtt.Message) {
      fmt.Println(msg)
    })
    mqttAdaptor.On("hola", func(msg mqtt.Message) {
      fmt.Println(msg)
    })
    data := []byte("o")
    gobot.Every(1*time.Second, func() {
      mqttAdaptor.Publish("hello", data)
    })
    gobot.Every(5*time.Second, func() {
      mqttAdaptor.Publish("hola", data)
    })
  }

  robot := gobot.NewRobot("mqttBot",
    []gobot.Connection{mqttAdaptor},
    work,
  )

  robot.Start()
}