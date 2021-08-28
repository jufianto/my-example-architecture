package httplib

import (
	"encoding/json"
	"fmt"
	"net/http"

	validator "github.com/go-playground/validator/v10"

	"github.com/labstack/echo"
)

type HttpLib struct {
	Request  *http.Request
	Response http.ResponseWriter
}

func (h *HttpLib) ReadBodyJSON(model interface{}) error {
	err := json.NewDecoder(h.Request.Body).Decode(&model)
	if err != nil {
		return err
	}
	return nil
}

func SendFailedBindEcho(c echo.Context, err error) error {
	errMsg := ErrorResponse{
		ErrorMessage: "Bad request body, please check your body request data",
		ErrorData:    err,
	}
	return c.JSON(400, errMsg)
}

func SendFailedCreate(c echo.Context, err error, data interface{}) error {
	errMsg := ErrorResponse{
		ErrorMessage: fmt.Sprintf("failed to create data %s", err.Error()),
		ErrorData:    data,
	}
	return c.JSON(400, errMsg)
}

type ResponseData struct {
	Metadata Metadata    `json:"metadata,omitempty"`
	Data     interface{} `json:"data"`
}

type ErrorResponse struct {
	ErrorMessage string      `json:"error_message"`
	ErrorData    interface{} `json:"error_data,omitempty"`
}

type Metadata struct {
	LengthData uint64 `json:"length_data,omitempty"`
	TotalData  uint64 `json:"total_data,omitempty"`
}

type CustomValidator struct {
	Validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	if err := cv.Validator.Struct(i); err != nil {
		errMsg := ErrorResponse{
			ErrorMessage: "Bad request form, please check your input",
			ErrorData:    err.Error(),
		}
		return echo.NewHTTPError(http.StatusBadRequest, errMsg)
	}
	return nil
}
