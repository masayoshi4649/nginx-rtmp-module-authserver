package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	var r *gin.Engine = gin.Default()
	r.GET(":key", checkKey)

	r.Run(":1919")
}

// keylistにあったらtrueを返す
func checkKey(c *gin.Context) {
	key := c.Param("key")

	keylist := getAllowedKey()
	fmt.Println(keylist)

	var checkresult bool = false
	for i := 0; i < len(keylist); i++ {
		if key == keylist[i] {
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
