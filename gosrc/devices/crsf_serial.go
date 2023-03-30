package devices

import (
	"slyfox-payload-controller/crsf"

	"github.com/tarm/serial"
)

type CrsfSerial struct {
	serialPort *serial.Port
	crsf       *crsf.Crsf
}

func NewCrsfSerial(port string, baud int) (*CrsfSerial, error) {
	crsf_serial := &CrsfSerial{
		crsf: crsf.NewCrsf(),
	}
	c := &serial.Config{Name: port, Baud: baud}
	serialPort, err := serial.OpenPort(c)
	if err != nil {
		return nil, err
	}

	crsf_serial.serialPort = serialPort

	return crsf_serial, nil
}

func (crsf_serial *CrsfSerial) Tick() {
	crsf_serial.handleSerialIn()
}

func (crsf_serial *CrsfSerial) handleSerialIn() {
	data := make([]byte, crsf.CRSF_MAX_PACKET_LEN)

	for {
		count, err := crsf_serial.serialPort.Read(data)

		if err != nil {
			break
		}

		crsf_serial.crsf.HandleBytes(data[:count])
	}
}

func (crsf_serial *CrsfSerial) SetOnPacketChannels(callback func()) {
	crsf_serial.crsf.OnPacketChannels = callback
}

func (crsf_serial *CrsfSerial) QueuePacket(addr, data_type uint8, payload []byte) *[crsf.CRSF_MAX_PACKET_LEN + 4]byte {
	return crsf_serial.crsf.QueuePacket(addr, data_type, payload)
}
