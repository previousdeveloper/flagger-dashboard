package controller

import (
	"github.com/labstack/echo"
	"github.com/previousdeveloper/flagger-dashboard/pkg/client"
	"net/http"
)

type CanaryController struct {
	K8sClient client.K8sOperation
}

func (controller *CanaryController) Hello(c echo.Context) error {
	namespace := c.Param("namespace")
	result, err := controller.K8sClient.GetResourceByNamespace(namespace)

	if err != nil {
		return c.JSON(500, err)
	}
	return c.JSON(http.StatusOK, result)
}
