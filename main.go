package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	mqttLib "github.com/Lilanga/sensor-data-processing-service/pkg/mqtt"
	"github.com/go-co-op/gocron"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	delay, err := strconv.Atoi(os.Getenv("DELAY"))

	if err != nil {
		log.Fatal(err)
	}

	client := mqttLib.GetMqttClient(os.Getenv("MQTT_CLIENT_ID"))
	topic := os.Getenv("MQTT_TOPIC")

	s := gocron.NewScheduler(time.UTC)
	_, _ = s.Every(delay).Seconds().Do(sendSensorReadings, client, topic)
	s.StartBlocking()
}

func sendSensorReadings(client *mqttLib.MqttClient, topic string) {
	sensors := strings.Split(os.Getenv("IDS"), ",")

	for _, sensor := range sensors {
		temp := randomSeed(20, 30)
		humidity := randomSeed(35, 80)

		message := fmt.Sprintf("h:%f,t:%f,n:%s", humidity, temp, sensor)
		client.Publish(topic, 1, message)
		time.Sleep(2 * time.Second)
	}
}

func randomSeed(min float64, max float64) float64 {
	return (rand.Float64() * (max - min)) + min
}
