/*
 * Copyright (c) 2013 IBM Corp.
 *
 * All rights reserved. This program and the accompanying materials
 * are made available under the terms of the Eclipse Public License v1.0
 * which accompanies this distribution, and is available at
 * http://www.eclipse.org/legal/epl-v10.html
 *
 * Contributors:
 *    Seth Hoenig
 *    Allan Stockdill-Mander
 *    Mike Robertson
 */

 // This example utility was forked from https://github.com/eclipse/paho.mqtt.golang/blob/master/cmd/stdoutsub/main.go

package main

import (
    "fmt"
    MQTT "github.com/eclipse/paho.mqtt.golang"
    "os"
    "os/signal"
    "syscall"
    "flag"
    "log"
    "encoding/json"
    "strconv"
    "time"

)

var loopster int = 0

func substring(s string, start int, end int) string {
    start_str_idx := 0
    i := 0
    for j := range s {
        if i == start {
            start_str_idx = j
        }
        if i == end {
            return s[start_str_idx:j]
        }
        i++
    }
    return s[start_str_idx:]
}

var f MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {

    jsonMessage := msg.Payload()

    type bridgePayload struct {
        Time   int `json:"time"`
        Demand int   `json:"demand"`
}

    var bridge bridgePayload	
    json.Unmarshal([]byte(jsonMessage), &bridge)

    wattsUsed := strconv.Itoa(bridge.Demand)
    unixTimestamp := strconv.Itoa(bridge.Time)

    i, err := strconv.ParseInt(unixTimestamp[0:10], 10, 64)
    if err != nil {
        panic(err)
    }
    tm := time.Unix(i, 0)


    fmt.Printf("Watts: %s | %s\n", wattsUsed, tm)

    loopster++
}

func main() {
    var ip string

    flag.StringVar(&ip, "ip", "", "IP address of energy bridge")
    flag.Parse()

    if ip == "" {
        log.Fatal("-ip must be provided")
    }

    c := make(chan os.Signal, 1)
    signal.Notify(c, os.Interrupt, syscall.SIGTERM)

    opts := MQTT.NewClientOptions().AddBroker("tcp://" + ip + ":2883")
    opts.SetUsername("admin")
    opts.SetPassword("trinity")
    opts.SetClientID("powerley-energybridge")
    opts.SetDefaultPublishHandler(f)
    topic := "remote/event/metering/instantaneous_demand"

    opts.OnConnect = func(c MQTT.Client) {
            if token := c.Subscribe(topic, 0, f); token.Wait() && token.Error() != nil {
                    panic(token.Error())
            }
    }
    client := MQTT.NewClient(opts)
    
    if token := client.Connect(); token.Wait() && token.Error() != nil {
            panic(token.Error())
    } else {
            fmt.Printf("Connected to Energy Bridge!\n")
    }
    <-c
}