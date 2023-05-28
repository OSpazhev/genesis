package db

import (
	"bufio"
	"io"
	"log"
	"os"
)

type Database struct {
	file        *os.File
	subscribers map[string]bool
}

var dbPath string

const dbName = "subscribers.txt"

var Db *Database

func (d *Database) updateSubscribers() {
	d.file.Seek(io.SeekStart, 0)

	fileScanner := bufio.NewScanner(d.file)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		d.subscribers[fileScanner.Text()] = true
	}

	d.file.Seek(io.SeekStart, 0)
}

func (d *Database) AddSubscriber(id string) {
	d.subscribers[id] = true

	d.file.Truncate(0)
	writer := bufio.NewWriter(d.file)

	for key, active := range d.subscribers {
		if active == true {
			writer.WriteString(key + "\n")
		}
	}
	writer.Flush()
}

func (d *Database) IsSubscriber(id string) bool {
	if val, ok := d.subscribers[id]; ok == true {
		return val
	} else {
		return false
	}

}

func (d *Database) Subscribers() []string {
	var subscribers []string
	for key, value := range d.subscribers {
		if value == true {
			subscribers = append(subscribers, key)
		}
	}
	return subscribers
}

func init() {
	dbPath = os.Getenv("DATAPATH")
	if len(dbPath) == 0 {
		dbPath = "data/"
	}

	err := os.MkdirAll(dbPath, 0750)
	if err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}

	file, err := os.OpenFile(dbPath+dbName, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}

	Db = &Database{file: file, subscribers: map[string]bool{}}
	Db.updateSubscribers()
}
