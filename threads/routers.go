package threads

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/leegggg/tiebashow/common"
)

func ThreadRegister(router *gin.RouterGroup) {
	router.GET("/", ThreadRetrieve)
}

func ThreadRetrieve(c *gin.Context) {
	var err error
	kw := c.DefaultQuery("kw", "_")
	pnParam := c.DefaultQuery("pn", "0")
	nbParam := c.DefaultQuery("nb", "50")
	tabParam := c.DefaultQuery("tab", "")

	pn, err := strconv.ParseInt(pnParam, 10, 64)
	nb, err := strconv.ParseInt(nbParam, 10, 64)
	good := "good" == tabParam

	if err != nil {
		c.JSON(http.StatusBadRequest, common.NewError("thread", errors.New("Invalid pn or nb")))
		return
	}

	var threads []Thread
	threads, err = GetThreads(kw, pn, nb, good)

	if err != nil {
		c.JSON(http.StatusInternalServerError, common.NewError("thread", errors.New("DB Error")))
		return
	}

	c.JSON(http.StatusOK, gin.H{"threads": threads})
}
