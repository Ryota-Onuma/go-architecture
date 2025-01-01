package articles

import (
	"context"
	i "onion/internal/usecase/interfaces"
)

var _ i.FetchArticlesUsecase = (*FetchArticlesUsecaseImpl)(nil)

type FetchArticlesUsecaseImpl struct {
	i.BaseUsecase
}

func NewFetchArticlesUsecase(base i.BaseUsecase) *FetchArticlesUsecaseImpl {
	return &FetchArticlesUsecaseImpl{
		base,
	}
}

func (a *FetchArticlesUsecaseImpl) Run(ctx context.Context) error {
	a.Info(ctx, "Fetching articles")
	return a.NewInternalServerError(ctx, "not implemented")
}
