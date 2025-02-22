package user

import "ecommerce/helper"

type UserAuthFormatter struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Token    string `json:"token"`
	ImageURL string `json:"image_url"`
	Role     string `json:"role"`
}

func FormatAuthUser(user User, token string) UserAuthFormatter {
	imageURL := ""
	if user.AvatarFileName != "" {
		imageURL = helper.GetAsset(user.AvatarFileName)
	}

	formatter := UserAuthFormatter{
		ID:       user.ID,
		Name:     user.Name,
		Email:    user.Email,
		Token:    token,
		ImageURL: imageURL,
		Role:     user.Role,
	}
	return formatter
}

type UserFormatter struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Role        string `json:"role"`
	PhoneNumber string `json:"phone_number"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

func FormatUser(user User) UserFormatter {
	formatter := UserFormatter{
		ID:          user.ID,
		Name:        user.Name,
		Email:       user.Email,
		Role:        user.Role,
		PhoneNumber: user.PhoneNumber,
		CreatedAt:   user.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:   user.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
	return formatter
}

func FormatUsers(users []User) []UserFormatter {
	var usersFormatter []UserFormatter

	for _, user := range users {
		userFormatter := FormatUser(user)
		usersFormatter = append(usersFormatter, userFormatter)
	}

	return usersFormatter
}
