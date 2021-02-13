package main

import (
	"time"

	"github.com/avereha/pod/pkg/bluetooth"
	"github.com/avereha/pod/pkg/pod"

	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
	//	"github.com/muka/go-bluetooth/api/service"
	//	"github.com/muka/go-bluetooth/bluez/profile/agent"
	//	"github.com/muka/go-bluetooth/bluez/profile/gatt"
)

func main() {
	log.SetLevel(log.DebugLevel)
	log.SetFormatter(&logrus.TextFormatter{
		DisableQuote: true,
		ForceColors:  true,
	})

	ble, err := bluetooth.New("hci0")
	defer ble.Close()
	if err != nil {
		log.Fatalf("Could not start BLE: ", err)
	}
	pod := pod.New(ble)
	pod.StartActivation()

	time.Sleep(9999 * time.Second)
}
