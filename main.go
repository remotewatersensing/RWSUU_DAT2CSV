package main

import (
	"encoding/binary"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Data struct {
	Id     uint8
	Val    float32
	Day    uint8
	Month  uint8
	Year   uint16
	Hour   uint8
	Minute uint8
	Second uint8
}

func main() {
	const outputFile = "output.csv"
	const packetSize = 12

	data := Data{}
	datfile, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Couldn't open file")
	}
	defer datfile.Close()

	csvfile, err := os.Create(outputFile)
	if err != nil {
		log.Fatal(err)
	}
	defer csvfile.Close()

	writer := csv.NewWriter(csvfile)
	defer writer.Flush()
	line := []string{"day", "month", "year", "hour", "minute", "second", "id", "val"}
	writer.Write(line)

	stats, err := os.Stat(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Amount of entries: %d\n", stats.Size()/packetSize)

	var i int64
	for i = 0; i < stats.Size()/packetSize; i++ {
		binary.Read(datfile, binary.LittleEndian, &data)
		line := []string{
			strconv.Itoa(int(data.Day)),
			strconv.Itoa(int(data.Month)),
			strconv.Itoa(int(data.Year)),
			strconv.Itoa(int(data.Hour)),
			strconv.Itoa(int(data.Minute)),
			strconv.Itoa(int(data.Second)),
			strconv.Itoa(int(data.Id)),
			fmt.Sprintf("%f", data.Val)}
		writer.Write(line)
	}
	fmt.Printf("Written to ")
	fmt.Print(outputFile + "\n")

}
