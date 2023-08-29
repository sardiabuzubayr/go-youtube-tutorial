package main

import (
	"fmt"
	"time"

	"github.com/go-co-op/gocron"
)

func main() {
	local, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		local = time.UTC
	}

	s := gocron.NewScheduler(local)
	s.Every(1).Day().At("05:45").Do(func() {
		fmt.Println("Hello saya adalah tugas yang dijalankan setiap 05:45 di eksekusi")
	})

	s.Every(1).Day().At("00:00").Do(func() {
		fmt.Println("Hello saya adalah tugas yang dijalankan setiap 00:00 di eksekusi")
		// kode apa yang mau dieksekusi
	})

	s.Every(1).Monday().At("07:00").Do(func() {
		fmt.Println("Hello saya adalah tugas yang dijalankan setiap 07:00 di eksekusi")
		// kode apa yang mau dieksekusi
	})

	s.StartBlocking()
}
