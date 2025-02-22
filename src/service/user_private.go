package service

import (
	"context"
	pgmodel "gin-base/internal/models"
	responsemodel "gin-base/src/model/response"
	"sync"
)

func (s userImpl) getListUser(ctx context.Context, users []pgmodel.User) []responsemodel.UserBrief {
	var (
		wg     sync.WaitGroup
		result = make([]responsemodel.UserBrief, len(users))
	)

	wg.Add(len(users))
	for index, user := range users {
		go func(i int, u pgmodel.User) {
			defer wg.Done()
			result[i] = s.brief(ctx, u)
		}(index, user)
	}

	wg.Wait()
	return result
}

func (s userImpl) brief(ctx context.Context, u pgmodel.User) responsemodel.UserBrief {
	return responsemodel.UserBrief{
		ID:        u.ID,
		Name:      u.Name,
		Status:    u.Status,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

func (s userImpl) detail(ctx context.Context, u pgmodel.User) responsemodel.UserDetail {
	return responsemodel.UserDetail{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		Status:    u.Status,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}

}
