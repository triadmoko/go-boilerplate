package dto_user

import (
	"app/internal/shared/constants"
	"app/internal/shared/domain/entity"
	"app/pkg"
	"time"
)

// UpdateProfileRequest represents the request for updating user profile
type UpdateProfileRequest struct {
	Name string `json:"name"`
}

// Validate validates UpdateProfileRequest fields
func (r *UpdateProfileRequest) Validate(lang constants.Lang) map[string][]string {
	errors := make(map[string][]string)

	if r.Name == "" {
		errors["name"] = append(errors["name"], constants.GetValidationMessage(constants.Required, lang))
	}

	return errors
}

// UserResponse represents a user data in response
type UserResponse struct {
	ID        string    `json:"id"`
	Email     string    `json:"email"`
	Username  string    `json:"username"`
	Name      string    `json:"name"`
	Phone     *string   `json:"phone,omitempty"`
	Status    string    `json:"status"`
	BirthDate *string   `json:"birth_date,omitempty"`
	Gender    string    `json:"gender,omitempty"`
	Role      string    `json:"role"`
	Provider  string    `json:"provider,omitempty"`
	IsActive  bool      `json:"is_active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// ToUserResponse converts entity.User to UserResponse
func ToUserResponse(user *entity.User) *UserResponse {
	if user == nil {
		return nil
	}

	response := &UserResponse{
		ID:        user.ID,
		Email:     user.Email,
		Username:  user.Username,
		Name:      user.Name,
		Phone:     user.Phone,
		Status:    user.Status,
		Gender:    user.Gender,
		Role:      user.Role,
		Provider:  user.Provider,
		IsActive:  user.IsActive,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}

	// Format birth date if exists
	if user.BirthDate != nil {
		birthDate := user.BirthDate.Format("2006-01-02")
		response.BirthDate = &birthDate
	}

	return response
}

// UserListResponse represents the response for listing users with pagination
type UserListResponse struct {
	Users      []*UserResponse        `json:"users"`
	Pagination pkg.PaginationResponse `json:"pagination"`
}
