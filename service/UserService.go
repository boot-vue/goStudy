package service

type IUserService interface {
	GetUser(id int) string
}

type UserService struct {
}

func (u *UserService) GetUser(id int) string {
	if id == 1 {
		return "user1"
	}
	return "others"
}
