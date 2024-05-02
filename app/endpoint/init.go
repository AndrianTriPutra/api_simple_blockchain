package endpoint

import (
	"atp/payment/app/usecase/blockchain"
	"atp/payment/pkg/repository/transaction"
	"atp/payment/pkg/utils/domain"
	"atp/payment/pkg/utils/echos/middleware"

	"github.com/labstack/echo/v4"
)

type handler struct {
	ucase blockchain.Usecase
	repo  transaction.RepositoryI
	bc    *domain.Blockchain
}

func NewHandler(e *echo.Echo, endpoint string, ucase blockchain.Usecase, repo transaction.RepositoryI, bc *domain.Blockchain) {
	handler := handler{
		ucase: ucase,
		repo:  repo,
		bc:    bc,
	}

	e.GET(endpoint+"transaction", middleware.ErrorMiddleware(handler.GetbyKey))
	e.GET(endpoint+"history", middleware.ErrorMiddleware(handler.GetALL))
	e.POST(endpoint+"payment", middleware.ErrorMiddleware(handler.Transaction))
}
