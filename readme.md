# AEP Powerley Energy Bridge Golang + MQTT Example

This is an example console app that connects to the AEP Ohio Powerly Energy bridge's MQTT protocol, receives messages and prints out realtime electricity usage to the console.

## Usage

`go run main.go -ip 192.168.1.100`

`go build`

`aep-golang-bridge.exe -ip 192.168.1.100` (IP to energy bridge)

## Geeky Details

* Username/password to the queue is: `admin` / `trinity` (not sure how this was discovered)
* The MQTT port was changed to TCP 2883
* The topic is: `_zigbee_metering/event/metering/instantaneous_demand`
* The messages pushed are JSON in this format `{"time":1555516806792,"demand":315}`
* This console app parses the JSON and also converts Unix epoch time to a human readable datetimestamp and prints the data to the console

## License

```txt
 Copyright (c) 2013 IBM Corp.

 All rights reserved. This program and the accompanying materials
 are made available under the terms of the Eclipse Public License v1.0
 which accompanies this distribution, and is available at
 http://www.eclipse.org/legal/epl-v10.html

 Contributors:
    Seth Hoenig
    Allan Stockdill-Mander
    Mike Robertson
```

Most of the code was taken from [https://github.com/eclipse/paho.mqtt.golang/blob/master/cmd/stdoutsub/main.go](https://github.com/eclipse/paho.mqtt.golang/blob/master/cmd/stdoutsub/main.go) and was modified to accept the -ip parameter and parse the JSON.