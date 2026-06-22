package proxy

import (
	"net/http"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

func ForwardRequest(c *gin.Context, target string) {
	targetURL, err := url.Parse(target)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "invalid target service url",
		})
		return
	}

	reverseProxy := httputil.NewSingleHostReverseProxy(targetURL)

	c.Request.Host = targetURL.Host

	reverseProxy.ErrorHandler = func(
		writer http.ResponseWriter,
		request *http.Request,
		err error,
	) {
		c.JSON(http.StatusBadGateway, gin.H{
			"error": "service unavailable",
		})
	}

	reverseProxy.ServeHTTP(c.Writer, c.Request)
}
