package main

import "pizza-tracker-go/internal/models"

type Handle struct {
	orders *models.OrderModel
}

func NewHandle(dbModel *models.DBModel) *Handle {

	return &Handle{
		orders: &dbModel.Order,
	}
}
