package main

type (
	// AccessRefreshTokenResponse is the response for the access refresh token
	AccessRefreshTokenResponse struct {
		JwtID        string               `json:"jwt_id"`
		ClientID     uint64               `json:"client_id"`
		TokenType    string               `json:"token_type"`
		Name         string               `json:"name"`
		Email        string               `json:"email"`
		Handphone    string               `json:"handphone"`
		AccessToken  AccessTokenResponse  `json:"access_token"`
		RefreshToken RefreshTokenResponse `json:"refresh_token"`
		Audience     string               `json:"audience"`
		Issuer       string               `json:"issuer"`
		Subject      string               `json:"subject"`
	}

	// AccessTokenResponse is the response for the access token
	AccessTokenResponse struct {
		Value     string `json:"value"`
		ExpiredAt int64  `json:"expired_at"`
	}

	// RefreshTokenResponse is the response for the refresh token
	RefreshTokenResponse struct {
		Value     string `json:"value"`
		ExpiredAt int64  `json:"expired_at"`
	}
)

func main() {

}
