package controllers

import "SberDS/internal/store"

type BaseController struct {
	store *store.Repository
}

func NewBaseController(store *store.Repository) *BaseController {
	return &BaseController{store: store}
}
