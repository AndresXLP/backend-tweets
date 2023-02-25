package handler_test

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/andresxlp/backend-twitter/internal/infra/api/handler"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

var (
	healthJson = `{"status":200,"message":"Active!"}`
	ctxTest    = context.Background()
)

type ControllerCase struct {
	Req     *http.Request
	Res     *httptest.ResponseRecorder
	context echo.Context
}

func SetupControllerCase(method, url string, body io.Reader) ControllerCase {
	e := echo.New()
	req := httptest.NewRequest(method, url, body)
	res := httptest.NewRecorder()
	ctx := e.NewContext(req, res)

	return ControllerCase{
		Req:     req,
		Res:     res,
		context: ctx,
	}
}

func TestHealthCheck(t *testing.T) {
	controller := SetupControllerCase(http.MethodGet, "/health", nil)

	if assert.NoError(t, handler.HealthCheck(controller.context)) {
		assert.Equal(t, http.StatusOK, controller.Res.Code)
		assert.Equal(t, healthJson, strings.TrimSpace(controller.Res.Body.String()))
	}
}
