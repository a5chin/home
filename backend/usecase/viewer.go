package usecase

import "context"

type LPUseCase struct {
	LPRepository
}

func NewLPUseCase(r LPRepository) *LPUseCase {
	return &LPUseCase{r}
}

func (u LPUseCase) GetTotalViewers(ctx context.Context) (uint, error) {
	return u.LPRepository.GetTotalViewers(ctx)
}
