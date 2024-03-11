package usecase

import "context"

type ViewerUseCase struct {
	ViewerRepository
}

func NewViewerUseCase(r ViewerRepository) *ViewerUseCase {
	return &ViewerUseCase{r}
}

func (u ViewerUseCase) GetViewers(ctx context.Context) (uint, error) {
	return u.ViewerRepository.GetViewers(ctx)
}
