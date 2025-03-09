package user

type (
	UserService interface {
	}

	userService struct {
		userRepository UserRepository
	}
)

func NewUserService(userRepository UserRepository) UserService {
	return &userService{userRepository: userRepository}
}

var VerifyEmailRoute = "api/verify_email/user"
