package response

type LoginResponse struct {
	// Add more user info or related here ...
	AccessToken string `json:"access_token"`
	ExpiresAt   int64  `json:"expires_at"`
}
