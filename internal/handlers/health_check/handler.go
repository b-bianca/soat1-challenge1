package health

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Handler is used for providing healthcheck through HTTP
type Handler struct{}

// NewHandler is the HealthHandlerBuilder
func NewHandler() *Handler {
	return &Handler{}
}

// healthCheck is used to check if the service is up service.
func (hdl *Handler) healthCheck(c *gin.Context) {
	c.String(http.StatusOK, "UP")
}

// RegisterRoutes register routes and set its permissions
func (hdl *Handler) RegisterRoutes(routes *gin.RouterGroup) {
	routes.GET("/health", hdl.healthCheck)
}
