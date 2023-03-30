package main

import (
	"fmt"
	"slyfox-payload-controller/crsf"
	"slyfox-payload-controller/devices"
	"slyfox-payload-controller/utils"
)

func main() {
	rx_port := utils.GetEnv("RX_UART", "/dev/tty1")
	rx_baud := utils.GetEnvInt("RX_BAUD", crsf.CRSF_BAUDRATE)
	tx_port := utils.GetEnv("TX_UART", "/dev/tty2")
	tx_baud := utils.GetEnvInt("TX_BAUD", crsf.CRSF_BAUDRATE)

	rx_from_apparature, err := devices.NewCrsfSerial(rx_port, rx_baud)
	if err != nil {
		fmt.Println("Rx setting up finished with error:", err.Error())
		return
	}

	tx_to_payload, err := devices.NewCrsfSerial(tx_port, tx_baud)
	if err != nil {
		fmt.Println("Tx setting up finished with error:", err.Error())
		return
	}

	rx_from_apparature.SetOnPacketChannels(nil)
	tx_to_payload.SetOnPacketChannels(nil)
}
