package iot

import (
    "fmt"
    MQTT "github.com/eclipse/paho.mqtt.golang"
)

func InitMQTT() {
    opts := MQTT.NewClientOptions().AddBroker("tcp://broker:1883").SetClientID("earth-optimizer")
    client := MQTT.NewClient(opts)
    if token := client.Connect(); token.Wait() && token.Error() != nil {
        panic(token.Error())
    }

    fmt.Println("🔌 MQTT Connected - Subscribing to sensors...")
    client.Subscribe("city/+/resource/#", 0, handleMQTTMessage)
}

func handleMQTTMessage(client MQTT.Client, msg MQTT.Message) {
    fmt.Printf("[MQTT] %s: %s\n", msg.Topic(), msg.Payload())
}
