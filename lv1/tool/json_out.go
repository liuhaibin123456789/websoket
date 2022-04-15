package tool

import "github.com/gin-gonic/gin"

func JsonOut(c *gin.Context, err error, info interface{}) {
	if err != nil {
		c.JSON(200, gin.H{
			"err": err,
			"msg": info,
		})
		return
	}

	c.JSON(200, gin.H{
		"err": "",
		"msg": info,
	})
}
