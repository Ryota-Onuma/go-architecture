package usecase

import "context"

type FetchArticlesUsecase interface {
	BaseUsecase
	Run(ctx context.Context) error
}
