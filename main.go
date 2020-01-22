package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"github.com/leegggg/tiebashow/common"
	"github.com/leegggg/tiebashow/posts"
	"github.com/leegggg/tiebashow/threads"
)

var db *gorm.DB
var err error

func main() {
	// Open our jsonFile
	jsonFile, err := os.Open("config.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Println(err)
	}
	log.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	jsonBlob, _ := ioutil.ReadAll(jsonFile)

	var config interface{}
	err = json.Unmarshal(jsonBlob, &config)

	log.Println(config)

	db := common.Init("./data/all.data.tieba.baidu.com.db")
	// Migrate(db)
	defer db.Close()

	r := gin.Default()
	r.StaticFS("/static", http.Dir("./static/"))
	r.StaticFile("/f", "./static/f.html")
	r.StaticFile("/p", "./static/p.html")

	r.StaticFS("/img", http.Dir("C:\\Users\\yizho\\tiebawap-get\\data\\柯哀IMG\\"))
	v1 := r.Group("/api")
	threads.ThreadRegister(v1.Group("/f"))
	posts.PostRegister(v1.Group("/p"))
	r.Run()

}
