package main

import (
	"github.com/tidwall/wal"
	"log"
	"strconv"
	"time"
)

func test(name string) {
	num := 20 * 1000

	log.Println("[" + name + "]" + "start")
	// open a new log file
	walLog, err := wal.Open("mylog"+name, nil)
	if err != nil {
		println(err)
	}
	log.Println("[" + name + "]" + "wal.1")
	lastIndex, err := walLog.LastIndex()
	log.Println("["+name+"]"+"wal.lastIndex:", lastIndex)
	preStart := time.Now().UnixNano() / 1e6
	// write some entries
	var d []byte
	d = append(d, 1)
	d = append(d, 2)
	d = append(d, 4)
	d = append(d, "{\"id\":\"3fwfaf09344\",\"title\":\"this is a test\"}"...)
	for i := 0; i < num; i++ {
		// 写入性能，初始每秒100左右，修改写入10000+/s
		err = walLog.Write(uint64(i+1), d)
	}
	preEnd := time.Now().UnixNano() / 1e6
	log.Println("["+name+"]"+"wal.write1:", int64(num)*1.0/(preEnd-preStart), "/s")

	start := time.Now().UnixNano() / 1e6
	// write some entries
	var d2 []byte
	d2 = append(d, 1)
	d2 = append(d, 3)
	d2 = append(d, 4)
	d = append(d, "{\"id\":\"3fwfaf09344\",\"title\":\"this is an apple\"}"...)
	for i := 0; i < num; i++ {
		// 写入性能，每秒100左右
		err = walLog.Write(uint64(i+1), d2)
	}
	end := time.Now().UnixNano() / 1e6
	log.Println("["+name+"]"+"wal.write2:", int64(num)*1.0/(end-start), "/s")
	// read an entry
	readStart := time.Now().UnixNano() / 1e6
	for i := 0; i < num; i++ {
		walLog.Read(uint64(i))
	}
	readEnd := time.Now().UnixNano() / 1e6
	log.Println("["+name+"]"+"wal.read", int64(num)*1.0/(readEnd-readStart), "/s")

	log.Println("[" + name + "]" + "wal.4")
	// close the log
	err = walLog.Close()

	log.Println("[" + name + "]" + "wal.5")
}

func main() {

	for i := 0; i < 1; i++ {
		go test(strconv.Itoa(i))
	}
	select {}
}
