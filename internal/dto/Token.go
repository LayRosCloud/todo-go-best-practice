package dto

type TokenResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type RefreshAccessTokenRequest struct {
	RefreshToken string `json:"refreshToken" validate:"required,min=10,max=1000"`
}

type LogoutRequest struct {
	RefreshToken string `json:"refreshToken" validate:"required,min=10,max=1000"`
}