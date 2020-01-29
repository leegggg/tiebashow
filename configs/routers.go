package configs

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leegggg/tiebashow/common"
	"github.com/leegggg/tiebashow/threads"
)

func ConfigRegister(router *gin.RouterGroup) {
	router.GET("/status", ConfigRetrieve)
}

func ConfigRetrieve(c *gin.Context) {
	var err error

	if err != nil {
		c.JSON(http.StatusBadRequest, common.NewError("status", errors.New("Invalid pn or nb")))
		return
	}

	config, err := GetConfig()

	if err != nil {
		c.JSON(http.StatusOK, gin.H{"config": config, "kws": nil, "status": 501, "msg": "config not loaded"})
		return
	}

	kws, err := threads.GetKws()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"config": config, "kws": nil, "status": 502, "msg": "database error"})
		return
	}

	if len(kws) <= 0 {
		c.JSON(http.StatusOK, gin.H{"config": config, "kws": kws, "status": 503, "msg": "No kws found in db"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"config": config, "kws": kws, "status": 200, "msg": "OK"})
}
