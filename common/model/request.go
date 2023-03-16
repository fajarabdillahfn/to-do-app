package model

type AddActivityRequest struct {
	Title string `json:"title" validate:"required"`
	Email string `json:"email" validate:"email"`
}

type UpdateActivityRequest struct {
	Title string `json:"title"`
}

type AddTodoRequest struct {
	Title           string `json:"title" validate:"required"`
	ActivityGroupId int    `json:"activity_group_id"`
	IsActive        bool   `json:"is_active" default:"true"`
}

type UpdateTodoRequest struct {
	Title           string `json:"title"`
	Priority        string `json:"priority"`
	ActivityGroupId int    `json:"activity_group_id"`
	IsActive        bool   `json:"is_active" default:"true"`
	Status          string `json:"status"`
}
