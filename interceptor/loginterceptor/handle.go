package loginterceptor

import (
	"bufio"
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/sillyhatxu/gin-utils/codes"
	"github.com/sillyhatxu/gin-utils/response"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

type bufferedWriter struct {
	gin.ResponseWriter
	out    *bufio.Writer
	Buffer bytes.Buffer
}

func (g *bufferedWriter) Write(data []byte) (int, error) {
	g.Buffer.Write(data)
	return g.out.Write(data)
}

func Handle() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		w := bufio.NewWriter(ctx.Writer)
		buff := bytes.Buffer{}
		newWriter := &bufferedWriter{ctx.Writer, w, buff}
		ctx.Writer = newWriter
		defer func() {
			logrus.Infof("response status : %d; body : %s", ctx.Writer.Status(), newWriter.Buffer.Bytes())
			w.Flush()
		}()
		body, err := ctx.GetRawData()
		if err != nil {
			ctx.JSON(http.StatusOK, response.Errorf(codes.InvalidParameter, err))
			return
		}
		logrus.Infof("request [%s%s] body : %v", ctx.Request.Host, ctx.Request.URL, string(body))
		ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))
		ctx.Next()
	}
}
