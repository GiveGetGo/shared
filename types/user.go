package types

// RegisterRequest - request body for the register endpoint
type RegisterRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Class    string `json:"class" binding:"required"`
	Major    string `json:"major" binding:"required"`
}

// LoginRequest - request body for the login endpoint
type LoginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// SetUserEmailVerifiedRequest - request body for setting user email to verified
type SetUserEmailVerifiedRequest struct {
	Email string `json:"email" binding:"required"`
}

// UserForgetPassRequest is the request body for the user forget password endpoint
type UserForgetPassRequest struct {
	Email string `json:"email" binding:"required"`
}

type UserResetPassRequest struct {
	Email       string `json:"email" binding:"required"`
	Newpassword string `json:"newpassword" binding:"required"`
}

type UserUpdateRequest struct {
	Username     *string `json:"username,omitempty"`
	Class        *string `json:"class,omitempty"`
	Major        *string `json:"major,omitempty"`
	ProfileImage *string `json:"profile_image,omitempty"`
	ProfileInfo  *string `json:"profile_info,omitempty"`
}

type UserInfoResponse struct {
	UserID        uint   `json:"userID"`
	Username      string `json:"username"`
	Email         string `json:"email"`
	Class         string `json:"class"`
	Major         string `json:"major"`
	ProfileImage  string `json:"profile_image"`
	ProfileInfo   string `json:"profile_info"`
	EmailVerified bool   `json:"email_verified"`
	MfaVerified   bool   `json:"mfa_verified"`
}
