package service

import (
	"context"

	"github.com/plagioriginal/user-microservice/domain"
)

// Gets a specific user that has a RefreshToken
func (s DefaultRefreshTokenService) GetUserByToken(ctx context.Context, token domain.RefreshToken) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, s.ContextTimeout)
	defer cancel()

	user, err := s.UserRepo.GetByRefreshToken(ctx, token.Id)
	if err != nil {
		s.Logger.Printf("error getting user from token {%s}: %v\n", token.Id.String(), err)
		return nil, err
	}

	return user, nil
}
