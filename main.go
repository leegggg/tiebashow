package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"github.com/leegggg/tiebashow/common"
	"github.com/leegggg/tiebashow/configs"
	"github.com/leegggg/tiebashow/posts"
	"github.com/leegggg/tiebashow/threads"
)

var db *gorm.DB
var err error
var srv *http.Server

func main() {
	config, err := configs.GetConfig()
	if err != nil {
		log.Fatalln("Config can not be loaded. Exiting")
		os.Exit(-1)
	}
	db := common.Init(config["db_path"])
	// Migrate(db)
	defer db.Close()

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.StaticFS("/static", http.Dir("./static/"))
	r.StaticFile("/f", "./static/f.html")
	r.StaticFile("/p", "./static/p.html")
	r.StaticFile("/", "./static/index.html")
	r.StaticFile("/favicon.ico", "./static/img/favicon.ico")

	r.StaticFS("/img", http.Dir(config["img_dir"]))
	v1 := r.Group("/api")
	v1.GET("/p/:kz", posts.PostRetrieve)
	v1.GET("/f", threads.ThreadRetrieve)
	v1.GET("/status", configs.ConfigRetrieve)
	//threads.ThreadRegister(v1.Group("/f"))
	//posts.PostRegister(v1.Group("/p"))
	//configs.ConfigRegister(v1.Group("/status"))

	srv = &http.Server{
		Handler: r,
	}

	r.Run(":45678")

	return
}
