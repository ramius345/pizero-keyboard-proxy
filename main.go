package main

import (
        "encoding/binary"
        "fmt"
        "os"
        "golang.org/x/sys/unix"
	"github.com/ramius345/keyboard-proxy/evthid"
)



//This function primarily creates a 16 byte read buffer and closes over it and the file pointer
func makePacketHandler(fp *os.File, hidfp *os.File) func() bool {
        packetBuf := make([]byte, 16)
	hidPacket := make([]byte, 8)
	hidmap := evthid.BuildEvtToHidMap()

	handlePressNormalKey := func( evt evthid.EvtHid ) {
		for i,value := range hidPacket[2:8] {
			if value == evt.Hidcode {
				break
			} else if value == 0 {
				hidPacket[i+2] = evt.Hidcode
				break
			}
		}
	}

	handleReleaseNormalKey := func( evt evthid.EvtHid ) {
		for i,value := range hidPacket[2:8] {
			if value == evt.Hidcode {
				offset := i+2
				for j,_ := range hidPacket[offset:7] {
					hidPacket[j+offset] = hidPacket[j+offset+1]
				}
				hidPacket[7] = 0
			}
		}
	}

	handlePressModifierKey := func( evt evthid.EvtHid ) {
		fmt.Printf("Hidmodbits: %d\n", evt.Hidmodbits )
		hidPacket[0] |= evt.Hidmodbits
	}

	handleReleaseModifierKey := func( evt evthid.EvtHid ) {
		mask := evt.Hidmodbits ^ 0xFF
		hidPacket[0] &= mask
	}
	
	emitPressHidPacket := func( evt evthid.EvtHid ) {
		fmt.Printf("Press registered\n")
		if evt.IsNormalKey() {
			handlePressNormalKey(evt)
		} else {
			handlePressModifierKey(evt)
		}
		fmt.Printf("hidbuf: %v, evtcode: %v\n" ,  hidPacket, evt)	
	}

	emitReleaseHidPacket := func ( evt evthid.EvtHid ) {
		fmt.Printf("Release registered\n")
		if evt.IsNormalKey() {
			handleReleaseNormalKey(evt)
		} else {
			handleReleaseModifierKey(evt)
		}
		fmt.Printf("hidbuf: %v, evtcode: %v\n" ,  hidPacket, evt)
	}

	writeHidPacket := func() {
		fmt.Printf("Writing hid packet\n")
		_,err:=hidfp.Write(hidPacket)
		if err != nil {
			fmt.Println(err)
		}
		hidfp.Sync()
	}
	
	handlePacket := func() bool {
                _, err := fp.Read(packetBuf)

                if err != nil {
                        return false
                }

		//fetch the fields we care about
                typ := int32(binary.LittleEndian.Uint16(packetBuf[8:10]))
                code := int32(binary.LittleEndian.Uint16(packetBuf[10:12]))
		value := uint32(binary.LittleEndian.Uint32(packetBuf[12:16]))
		if typ == 1 {
			if  hidElem,ok := hidmap[code]; ok == true {
				fmt.Printf("type: %d, code: %s, value: %d\n", typ, hidElem.GetName(), value )
				if value == 1 || value == 2 {
					//press or repeat
					emitPressHidPacket(hidElem)
					writeHidPacket()
				} else if value == 0 {
					//release
					emitReleaseHidPacket(hidElem)
					writeHidPacket()
				}
			} else {
				fmt.Printf("type: %d, code: %d, value: %d\n", typ, code, value)
			}
		}
		return true
	}

	return handlePacket
}

func readFromKeyboardDeviceFile(deviceFileName string, hidDeviceName string) {
        //open the ev device
	fp, err := os.Open(deviceFileName)
	if err != nil {
		panic(err)
	}

	hidfp, err := os.OpenFile(hidDeviceName,os.O_WRONLY,0666)
	if err != nil {
		panic(err)
	}
	
	//set up a device lock with the EVIOCGRAB ioctl
	const EVIOCGRAB  = 1074021776
	fd := int(fp.Fd())
	err = unix.IoctlSetInt(fd,EVIOCGRAB,1);
        if err != nil {
                panic(err)
        }

	//Unlock device on function exit
	defer unix.IoctlSetInt(fd,EVIOCGRAB,0)
        defer fp.Close()
	defer hidfp.Close()
	
	//create a handler function, then repeatedly call it
	handlePacket := makePacketHandler(fp,hidfp)
	keepGoing := true
        for keepGoing {
		keepGoing = handlePacket()
        }

}

func main() {
	readFromKeyboardDeviceFile("/dev/input/event0","/dev/hidg0")
}

