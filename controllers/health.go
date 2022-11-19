package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/uptrace/opentelemetry-go-extra/otelplay"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/otel/trace"
	"net/http"
)

type HealthController struct{}

func (h HealthController) Status(c *gin.Context) {
	c.String(http.StatusOK, "Working!")
}

const (
	indexTmpl   = "index"
	profileTmpl = "profile"
)

func ExampleMonitor(c *gin.Context) {
	ctx := c.Request.Context()
	otelgin.HTML(c, http.StatusOK, profileTmpl, gin.H{
		"username": c.Param("username"),
		"traceURL": otelplay.TraceURL(trace.SpanFromContext(ctx)),
	})
}
