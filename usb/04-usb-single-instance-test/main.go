package main

import (
	"fmt"
	"io"
	"log"

	"github.com/creack/pty"
	"github.com/google/gousb"
)

// 001.002 067b:2303 PL2303 Serial Port (Prolific Technology, Inc.)
//   Protocol: (Defined at Interface level)
//   * Vendor: 067b, Product: 2303
//   Configuration 1:
//     --------------
//     Interface 0 alternate setting 0 (available endpoints: [0x02(2,OUT) 0x81(1,IN) 0x83(3,IN)])
//       Vendor Specific Class
//       ep #1 IN (address 0x81) interrupt - undefined usage [10 bytes], maxPacketSize=10
//       ep #2 OUT (address 0x02) bulk [64 bytes], maxPacketSize=64
//       ep #3 IN (address 0x83) bulk [64 bytes], maxPacketSize=64
//     --------------

type usb struct {
	ctx  *gousb.Context
	dev  *gousb.Device
	cfg  *gousb.Config
	intf *gousb.Interface

	inEP  *gousb.InEndpoint
	outEP *gousb.OutEndpoint
}

// newUsbEP will prepare and return an in and out endpoint of a usb device.
func usbEPStart(uCfg usbCfg) (*usb, error) {
	var err error
	u := &usb{}

	// Open a new context.
	u.ctx = gousb.NewContext()

	u.dev, err = u.ctx.OpenDeviceWithVIDPID(uCfg.vid, uCfg.pid)
	if err != nil {
		return nil, fmt.Errorf("error: opening endpoint failed: %v", err)
	}

	u.cfg, err = u.dev.Config(uCfg.config)
	if err != nil {
		return nil, fmt.Errorf("error: opening config failed: %v", err)
	}

	u.intf, err = u.cfg.Interface(uCfg.intfNr, uCfg.subIntfNr)
	if err != nil {
		return nil, fmt.Errorf("error: opening interface failed: %v", err)
	}

	u.inEP, err = u.intf.InEndpoint(uCfg.inEpNr)
	if err != nil {
		return nil, fmt.Errorf("error: opening in endpoint failed: %v", err)
	}
	fmt.Printf("created inEP: %v\n", u.inEP)

	u.outEP, err = u.intf.OutEndpoint(uCfg.outEpNr)
	if err != nil {
		return nil, fmt.Errorf("error: opening out endpoint failed: %v", err)
	}
	fmt.Printf("created outEP: %v\n", u.outEP)

	return u, nil
}

func usbEPStop(u *usb) error {
	var err error

	err = u.ctx.Close()
	if err != nil {
		return err
	}
	err = u.dev.Close()
	if err != nil {
		return err
	}
	err = u.cfg.Close()
	if err != nil {
		return err
	}
	u.intf.Close()

	return nil
}

type usbCfg struct {
	vid       gousb.ID
	pid       gousb.ID
	config    int
	intfNr    int
	subIntfNr int
	inEpNr    int
	outEpNr   int
	byteSize  int
}

func newUsbCfg(vid gousb.ID, pid gousb.ID, config int, intfNr int, subIntNr int, inEpNr int, outEpNr int, byteSize int) usbCfg {
	u := usbCfg{
		vid:       vid,
		pid:       pid,
		config:    config,
		intfNr:    intfNr,
		subIntfNr: subIntNr,
		inEpNr:    inEpNr,
		outEpNr:   outEpNr,
		byteSize:  byteSize,
	}

	return u
}

// Baud rate (TX/RX) is 9600/9600, no parity, 2 stopbits, 8 databits

func main() {
	// usb create

	uCfg := newUsbCfg(0x0557, 0x2008, 1, 0, 0, 3, 2, 64)

	u, err := usbEPStart(uCfg)
	if err != nil {
		log.Printf("%v\n", err)
		return // TODO: Decide if we should return here, or other..
	}
	defer usbEPStop(u)

	// pty create

	pt, tt, err := pty.Open()
	if err != nil {
		log.Printf("error: failed to pty.Open: %v\n", err)
	}
	defer pt.Close()
	defer tt.Close()

	fmt.Printf("pty: %v\n", pt.Name())
	fmt.Printf("tty: %v\n", tt.Name())

	// --- read/write ---

	// usb -> pty
	go func() {
		for {
			b := make([]byte, uCfg.byteSize)

			// USB Read
			n, err := u.inEP.Read(b)
			if err != nil {
				log.Printf("error: failed to read inEP: %v\n", err)
				log.Printf("n = %v\n", n)
				continue
			}

			// write to pty

			{
				n, err := pt.Write(b)
				if err != nil {
					log.Printf("error: writing to pty: %v\n", err)
				}

				fmt.Printf(" * pty: wrote %v bytes\n", n)
			}

		}
	}()

	// pty -> usb
	for {
		b := make([]byte, 64)
		_, err := pt.Read(b)
		if err != nil && err != io.EOF {
			log.Printf("error: failed to read conn : %v\n", err)
			continue
		}
		if err == io.EOF {
			log.Printf("error: got io.EOF: %v\n", err)
			return
		}

		{
			fmt.Printf(" * debug : read from pty: %v\n", b)
			n, err := u.outEP.Write(b)
			if err != nil {
				log.Printf("error: writing to ep: %v\n", err)
			}

			fmt.Printf(" * ep: wrote %v bytes\n", n)
		}
	}

}
