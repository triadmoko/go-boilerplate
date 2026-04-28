package dto_auth

import (
	"app/internal/shared/constants"
	"app/internal/shared/domain/entity"
	"fmt"
	"time"
)

// RegisterRequest represents the request for user registration
type RegisterRequest struct {
	Email     string `json:"email"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	Name      string `json:"name"`
}

// Validate validates RegisterRequest fields
func (r *RegisterRequest) Validate(lang constants.Lang) map[string][]string {
	errors := make(map[string][]string)

	// Email validation
	if r.Email == "" {
		errors["email"] = append(errors["email"], fmt.Sprintf(constants.GetValidationMessage(constants.Required, lang), "email"))
	} else if !constants.IsValidEmail(r.Email) {
		errors["email"] = append(errors["email"], constants.GetValidationMessage(constants.InvalidEmail, lang))
	}

	// Username validation
	if r.Username == "" {
		errors["username"] = append(errors["username"], fmt.Sprintf(constants.GetValidationMessage(constants.Required, lang), "username"))
	} else {
		if !constants.MinLength(r.Username, 3) {
			errors["username"] = append(errors["username"], fmt.Sprintf(constants.GetValidationMessage(constants.UsernameTooShort, lang), 3))
		}
		if !constants.MaxLength(r.Username, 20) {
			errors["username"] = append(errors["username"], fmt.Sprintf(constants.GetValidationMessage(constants.UsernameTooLong, lang), 20))
		}
	}

	// Password validation
	if r.Password == "" {
		errors["password"] = append(errors["password"], fmt.Sprintf(constants.GetValidationMessage(constants.Required, lang), "password"))
	} else if !constants.MinLength(r.Password, 6) {
		errors["password"] = append(errors["password"], fmt.Sprintf(constants.GetValidationMessage(constants.PasswordTooShort, lang), 6))
	}

	// Name validation
	if r.Name == "" {
		errors["name"] = append(errors["name"], fmt.Sprintf(constants.GetValidationMessage(constants.Required, lang), "name"))
	}

	return errors
}

// LoginRequest represents the request for user login
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Validate validates LoginRequest fields
func (r *LoginRequest) Validate(lang constants.Lang) map[string][]string {
	errors := make(map[string][]string)

	// Email validation
	if r.Email == "" {
		errors["email"] = append(errors["email"], fmt.Sprintf(constants.GetValidationMessage(constants.Required, lang), "email"))
	} else if !constants.IsValidEmail(r.Email) {
		errors["email"] = append(errors["email"], constants.GetValidationMessage(constants.InvalidEmail, lang))
	}

	// Password validation
	if r.Password == "" {
		errors["password"] = append(errors["password"], fmt.Sprintf(constants.GetValidationMessage(constants.Required, lang), "password"))
	}

	return errors
}

// RegisterResponse represents the response for user registration
type RegisterResponse struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	Name      string    `json:"name"`
	Status    string    `json:"status"`
	Role      string    `json:"role"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
}

// ToRegisterResponse converts entity.User to RegisterResponse
func ToRegisterResponse(user *entity.User) *RegisterResponse {
	return &RegisterResponse{
		ID:        user.ID,
		Email:     user.Email,
		Username:  user.Username,
		Name:      user.Name,
		Status:    user.Status,
		Role:      user.Role,
		IsActive:  user.IsActive,
		CreatedAt: user.CreatedAt,
	}
}

// LoginResponse represents the response for user login
type LoginResponse struct {
	User  *RegisterResponse `json:"user"`
	Token string            `json:"token"`
}
