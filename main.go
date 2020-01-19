package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"

	"github.com/leegggg/tiebashow/common"
	"github.com/leegggg/tiebashow/posts"
	"github.com/leegggg/tiebashow/threads"
)

var db *gorm.DB
var err error

func main() {
	db := common.Init()
	// Migrate(db)
	defer db.Close()

	r := gin.Default()

	v1 := r.Group("/api")
	threads.ThreadRegister(v1.Group("/f"))
	posts.PostRegister(v1.Group("/p"))
	r.Run()

}
