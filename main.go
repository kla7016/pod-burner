package main

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"time"
)

// ฟังก์ชันใช้ CPU โดยการวนลูปไม่หยุด
func consumeCPU() {
	for {
	}
}

// ฟังก์ชันใช้ RAM โดยการสร้าง Array ขนาดใหญ่
func consumeRAM(sizeMB int) {
	size := sizeMB * 1024 * 1024 / 8 // แปลงเป็นจำนวน element ของ int64
	mem := make([]int64, size)
	for i := range mem {
		mem[i] = int64(i)
	}
	fmt.Printf("Allocated %d MB of RAM\n", sizeMB)
}

func main() {
	// อ่านค่า ENV
	cpuUsage, err := strconv.Atoi(os.Getenv("CPU_USAGE"))
	if err != nil {
		cpuUsage = 1 // Default ใช้ 1 Core
	}

	ramUsageMB, err := strconv.Atoi(os.Getenv("RAM_USAGE_MB"))
	if err != nil {
		ramUsageMB = 100 // Default ใช้ 100 MB
	}

	// กำหนดจำนวน CPU ที่จะใช้
	runtime.GOMAXPROCS(cpuUsage)
	fmt.Printf("Using %d CPU cores\n", cpuUsage)

	// เริ่มใช้ RAM
	consumeRAM(ramUsageMB)

	// เริ่มใช้ CPU โดยการเปิด goroutine ตามจำนวนที่กำหนด
	for i := 0; i < cpuUsage; i++ {
		go consumeCPU()
	}

	// ทำให้ Service ทำงานต่อเนื่อง
	for {
		fmt.Println("Service is running...")
		time.Sleep(10 * time.Second)
	}
}
