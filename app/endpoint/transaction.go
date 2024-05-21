package endpoint

import (
	"atp/payment/pkg/utils/domain"
	"atp/payment/pkg/utils/echos/util"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h handler) Transaction(c echo.Context) error {
	ctx := c.Request().Context()
	var data domain.Data
	err := c.Bind(&data)
	if err != nil {
		return util.CustomError{
			ErrorType: util.ErrBadRequest,
			Message:   "The given data was invalid",
			Cause:     "failed decode input",
		}
	}

	h.bc.GiveData(data)

	block := h.ucase.CreateBlock(ctx, h.bc, "")

	response := util.WrapSuccessResponse("success", block)
	return c.JSON(http.StatusOK, response)
}
