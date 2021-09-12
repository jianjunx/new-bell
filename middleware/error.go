package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Newcontent struct {
	gin.Context
	ErrorHandler func(error)
}

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		nc := new(Newcontent)
		nc.ErrorHandler = func(e error) {
			fmt.Println("e", e)
		}
		c.Set("Content", nc)
		// c.ErrHandler = func() {

		// }
	}
}
