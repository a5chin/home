package usecase

import "context"

type ViewerUseCase struct {
	ViewerRepository
}

func NewViewerUseCase(r ViewerRepository) *ViewerUseCase {
	return &ViewerUseCase{r}
}

func (u ViewerUseCase) GetTotalViewers(ctx context.Context) (uint, error) {
	return u.ViewerRepository.GetTotalViewers(ctx)
}
