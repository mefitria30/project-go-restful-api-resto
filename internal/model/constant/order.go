package constant

import "github.com/kelas.work/project-go-restful-api/internal/model"

const (
	OrderStatusProcessed model.OrderStatus =  "processed"
	OrderStatusFinished model.OrderStatus =  "finished"
	OrderStatusFailed model.OrderStatus =  "failed"
)

const (
	ProductOrderStatusPreparing model.ProductOrderStatus = "preparing"
	ProductOrderStatusFinished model.ProductOrderStatus = "finished"
)