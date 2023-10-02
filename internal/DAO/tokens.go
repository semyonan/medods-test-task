package DAO

import "time"

type Tokens struct {
	AccessToken	 string  `json:"accessToken"`
	AccessTokenTTL time.Duration `json:"accessTokenTTL"`
	RefreshToken string  `json:"refreshToken"`
	RefreshTokenTTL time.Duration `json:"refreshTokenTTL"`
}