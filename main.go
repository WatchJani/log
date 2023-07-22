package main

import (
	"log"
	"math/rand"
	"os"
	"time"
)

const DEFAULT_ERROR string = ""

type User struct {
	Name     string
	Password string
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	// var sleep time.Duration
	// sender := make(chan time.Duration)

	// go Receiver(sender, "/home/janko/Desktop/loger.log")

	// for {
	// 	time.Sleep(sleep)
	// 	sleep = time.Duration(rand.Intn(5)) * time.Second
	// 	sender <- sleep
	// }

	// Postavljanje različitih mjesta za logiranje

	file, err := os.Create("logfile.log")
	if err != nil {
		log.Fatal("Nije moguće otvoriti datoteku za logiranje:", err)
	}
	defer file.Close()

	// Kreiranje novog log.Logger objekta za logiranje u file
	fileLogger := log.New(file, "FILE LOG: ", log.Ldate|log.Ltime)

	// Kreiranje novog log.Logger objekta za logiranje na standardni izlaz (terminal)
	terminalLogger := log.New(os.Stdout, "TERMINAL LOG: ", log.Ldate|log.Ltime)

	// Logiranje na različita mjesta koristeći odgovarajuće log.Logger objekte
	fileLogger.Println("Ovo je zapis u log datoteku.")
	terminalLogger.Println("Ovo je zapis na standardni izlaz (terminal).")
	log.Println("Ovo je zapis kroz uobičajeni log.Logger, koji se također ispisuje na standardni izlaz (terminal).")
}

// func NewBuffer(path string) *bufio.Writer {
// 	file, err := os.OpenFile(path, os.O_WRONLY, 0766)
// 	e.ErrorHandler(err, DEFAULT_ERROR)

// 	return bufio.NewWriter(file)
// }

// func Receiver(Receiver chan time.Duration, path string) {
// 	saveToFile := time.Tick(5 * time.Second)
// 	buf := NewBuffer(path)

// 	for {
// 		select {
// 		case saveToBuffer := <-Receiver:
// 			buf.WriteString(fmt.Sprintf("save super risk \n")) //make format
// 			log.Println(saveToBuffer)
// 		case <-saveToFile:
// 			buf.Flush() //write data on disk
// 		}
// 	}
// }
