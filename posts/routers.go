package posts

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/leegggg/tiebashow/common"
)

// PostRegister ...
func PostRegister(router *gin.RouterGroup) {
	router.GET("/:kz", PostRetrieve)
}

// PostRetrieve ...
func PostRetrieve(c *gin.Context) {
	var err error
	kzParam := c.Param("kz")
	pnParam := c.DefaultQuery("pn", "0")
	nbParam := c.DefaultQuery("nb", "50")
	flrParam := c.DefaultQuery("flr", "0")

	kz, err := strconv.ParseInt(kzParam, 10, 64)
	pn, err := strconv.ParseInt(pnParam, 10, 64)
	nb, err := strconv.ParseInt(nbParam, 10, 64)
	flr := "0" != flrParam

	if kz < 0 {
		c.JSON(http.StatusBadRequest, common.NewError("post", errors.New("kz is needed")))
		return
	}

	if err != nil {
		c.JSON(http.StatusBadRequest, common.NewError("post", errors.New("Invalid kz or pn or nb")))
		return
	}

	var posts []Post
	posts, err = GetPostsByParentAndFlrFlag(kz, flr, pn, nb)

	if err != nil {
		c.JSON(http.StatusInternalServerError, common.NewError("post", errors.New("DB Error")))
		return
	}

	c.JSON(http.StatusOK, gin.H{"posts": posts})
}
