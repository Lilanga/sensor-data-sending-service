### Sensor data sending service

This service is use to emulate sensor data aquestition part. Mock values are publishing to mqtt topic.
Service is designed to run on raspberry pi devices (or any supported image by balena) using balena cloud.

Please update following environment variables

```go
// comma seperated list of sensor IDs which need to be used as emulated device ids
IDS=

// integer number to specify delay in seconds between two sensor readings
DELAY=

// MQTT Host URL in 'tls://xxxxxxx....' format for private HiveMQ endpoints
MQTT_HOST=

// MQTT port number
MQTT_PORT=

// MQTT user name
MQTT_USER=

// MQTT password
MQTT_PASS=

// string to be used as the client for the connection
MQTT_CLIENT_ID=

// MQTT topic which message needs to be published
MQTT_TOPIC=
```
