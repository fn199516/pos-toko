package handler

import "pos-toko/internal/domain"

type Base struct {
	Helper *domain.Helper
}

func NewHandler(hlp *domain.Helper) Base {
	return Base{Helper: hlp}
}
