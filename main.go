package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

const (
	port = ":1934"
)

func main() {

	var r *gin.Engine = gin.Default()
	r.GET("/", checkKey)

	r.Run(port)
}

type publish struct {
	app      string
	flashver string
	swfurl   string
	tcurl    string
	pageurl  string
	call     string
	name     string
	livetype string
}

// keylistにあったらtrueを返す
func checkKey(c *gin.Context) {
	ld := publish{
		c.Query("app"),
		c.Query("flashver"),
		c.Query("swfurl"),
		c.Query("tcurl"),
		c.Query("pageurl"),
		c.Query("call"),
		c.Query("name"),
		c.Query("type"),
	}

	keylist := getAllowedKey()
	fmt.Println(keylist)

	var checkresult bool = false
	for i := 0; i < len(keylist); i++ {
		if ld.name == keylist[i] {
			checkresult = true
			break
		}
	}

	switch checkresult {
	case true:
		c.Status(200)
	case false:
		c.Status(404)
	}
}

// 許可されるStreamkeyを配列で返す
func getAllowedKey() []string {
	file, err := os.Open("AllowedKey.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	r := csv.NewReader(file)
	rows, err := r.ReadAll() // csvを一度に全て読み込む
	if err != nil {
		log.Fatal(err)
	}

	var keylist []string
	for i := 0; i < len(rows); i++ {

		keylist = append(keylist, rows[i][0])
	}

	return keylist
}
