package main

import (
	"fmt"
	"log"

	"github.com/google/gousb"
)

// 020.002 13fe:3600 flash drive (4GB, EMTEC) (Kingston Technology Company Inc.)
//   Protocol: (Defined at Interface level)
//   Configuration 1:
//     --------------
//     Interface 0 alternate setting 0 (available endpoints: [0x02(2,OUT) 0x81(1,IN)])
//       Mass Storage (SCSI) Bulk-Only
//       ep #1 IN (address 0x81) bulk [512 bytes]
//       ep #2 OUT (address 0x02) bulk [512 bytes]
//     --------------

func main() {
	// Open a new context.
	ctx := gousb.NewContext()
	defer ctx.Close()

	// NB: Need to check for nil value, or it will panic on using the dev
	// on the next step.
	dev, err := ctx.OpenDeviceWithVIDPID(0x13fe, 0x3600)
	if err != nil || dev == nil {
		log.Printf("error: opening endpoint failed: %v\n", err)
		return
	}
	defer dev.Close()

	cfg, err := dev.Config(1)
	if err != nil {
		log.Printf("error: opening config failed: %v\n", err)
		return
	}
	defer cfg.Close()

	intf, err := cfg.Interface(0, 0)
	if err != nil {
		log.Printf("error: opening interface failed: %v\n", err)
		return
	}
	defer intf.Close()

	inEP, err := intf.InEndpoint(1)

	if err != nil {
		log.Printf("error: opening in endpoint failed: %v\n", err)
		return
	}

	data := make([]byte, 512)
	for {
		n, err := inEP.Read(data)
		if err != nil {
			fmt.Printf("error: failed to read inEP: %v\n", err)
			fmt.Printf("n = %v\n", n)
			break
		}

		fmt.Printf("data = %v\n", data)
	}

}
