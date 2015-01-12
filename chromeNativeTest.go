package main

import (
	"encoding/binary"
	"encoding/json"
	"io"
	"log"
	"os"
	"time"
	"fmt"
)

type inputMsg struct {
	Test string
}

type outputMsg struct {
	Reverse string
	Time JSONTime
}

type JSONTime time.Time

func (t JSONTime) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("\"%s\"", time.Time(t).Format("Mon Jan 2 15:04:05"))
	return []byte(stamp), nil
}

// A simple script to act as the native host for chrome plugin.
// It receives messages from chrome which include json object with one
// value "test" that has a string value. It then responses with a json
// message that includes one value "reverse" that is the received string
// but reversed.
func main() {
	var startTime = time.Now()
	for {
		var size uint32
		err := binary.Read(os.Stdin, binary.LittleEndian, &size)
		if err != nil {
			log.Fatalln(err)
		}

		var jsonBuf = make([]byte, size)
		_, err = io.ReadFull(os.Stdin, jsonBuf)
		if err != nil {
			log.Fatalln(err)
		}

		var msg inputMsg

		json.Unmarshal(jsonBuf, &msg)

		n := len(msg.Test)
		runes := make([]rune, n)
		for _, rune := range msg.Test {
			n--
			runes[n] = rune
		}
		outMsg := outputMsg{string(runes[n:]), JSONTime(startTime)}

		outBytes, err := json.Marshal(outMsg)

		if err != nil {
			log.Fatalln(err)
		}

		binary.Write(os.Stdout, binary.LittleEndian, uint32(len(outBytes)))
		os.Stdout.Write(outBytes)
	}
}
