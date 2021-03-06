// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/plagioriginal/user-microservice/domain"
	mock "github.com/stretchr/testify/mock"

	uuid "github.com/google/uuid"
)

// RefreshTokenService is an autogenerated mock type for the RefreshTokenService type
type RefreshTokenService struct {
	mock.Mock
}

// DeleteToken provides a mock function with given fields: ctx, token
func (_m *RefreshTokenService) DeleteToken(ctx context.Context, token domain.RefreshToken) error {
	ret := _m.Called(ctx, token)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, domain.RefreshToken) error); ok {
		r0 = rf(ctx, token)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GenerateRefreshToken provides a mock function with given fields: ctx, user
func (_m *RefreshTokenService) GenerateRefreshToken(ctx context.Context, user *domain.User) (domain.RefreshToken, error) {
	ret := _m.Called(ctx, user)

	var r0 domain.RefreshToken
	if rf, ok := ret.Get(0).(func(context.Context, *domain.User) domain.RefreshToken); ok {
		r0 = rf(ctx, user)
	} else {
		r0 = ret.Get(0).(domain.RefreshToken)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *domain.User) error); ok {
		r1 = rf(ctx, user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTokenFromRepo provides a mock function with given fields: ctx, token
func (_m *RefreshTokenService) GetTokenFromRepo(ctx context.Context, token uuid.UUID) (domain.RefreshToken, error) {
	ret := _m.Called(ctx, token)

	var r0 domain.RefreshToken
	if rf, ok := ret.Get(0).(func(context.Context, uuid.UUID) domain.RefreshToken); ok {
		r0 = rf(ctx, token)
	} else {
		r0 = ret.Get(0).(domain.RefreshToken)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uuid.UUID) error); ok {
		r1 = rf(ctx, token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserByToken provides a mock function with given fields: ctx, token
func (_m *RefreshTokenService) GetUserByToken(ctx context.Context, token domain.RefreshToken) (*domain.User, error) {
	ret := _m.Called(ctx, token)

	var r0 *domain.User
	if rf, ok := ret.Get(0).(func(context.Context, domain.RefreshToken) *domain.User); ok {
		r0 = rf(ctx, token)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.User)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, domain.RefreshToken) error); ok {
		r1 = rf(ctx, token)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsTokenValid provides a mock function with given fields: token
func (_m *RefreshTokenService) IsTokenValid(token domain.RefreshToken) bool {
	ret := _m.Called(token)

	var r0 bool
	if rf, ok := ret.Get(0).(func(domain.RefreshToken) bool); ok {
		r0 = rf(token)
	} else {
		r0 = ret.Get(0).(bool)
	}

	return r0
}
