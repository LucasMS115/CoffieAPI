package http

import "coffie/internal/feature/user/domain"

func toRegisterRequest(registerUserRequest *RegisterUser) domain.RegisterRequest {
	return domain.RegisterRequest{
		Name:  registerUserRequest.Name,
		Email: registerUserRequest.Email,
	}
}

func toUserResponse(user *domain.User) *UserResponse {
	return &UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
	}
}
