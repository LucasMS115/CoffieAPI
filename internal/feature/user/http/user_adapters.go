package http

import "coffie/internal/feature/user/domain"

func toRegisterRequest(req *RegisterUser) domain.RegisterRequest {
	return domain.RegisterRequest{
		Name:  req.Name,
		Email: req.Email,
	}
}

func toUserResponse(u *domain.User) *UserResponse {
	return &UserResponse{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		CreatedAt: u.CreatedAt,
	}
}
