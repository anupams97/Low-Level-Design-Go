package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof" // Enables /debug/pprof endpoints
	"runtime"
	"time"
)

var leak [][]byte // Simulated memory leak

func memoryLeak() {
	for i := 0; i < 1000000; i++ {
		chunk := make([]byte, 1024*100) // 100 KB
		leak = append(leak, chunk)      // Never freed — leak grows
		time.Sleep(10 * time.Millisecond)
		if i%100 == 0 {
			printMem()
		}
	}
}

func printMem() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("HeapAlloc = %v MiB\n", m.HeapAlloc/1024/1024)
}

func main() {
	go func() {
		log.Println("pprof listening at :6060")
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	memoryLeak()
}
