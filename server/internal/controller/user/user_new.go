package user

import "chat/internal/controller"

type ControllerV1 struct{}

func NewV1() controller.IUserV1 {
	return &ControllerV1{}
}
