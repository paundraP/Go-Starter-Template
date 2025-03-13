package domain

import (
	"errors"
	"mime/multipart"
)

var (
	MessageSuccessRegister             = "register success"
	MessageSuccessLogin                = "login success"
	MessageSuccessVerify               = "verify email success"
	MessageSuccessGetDetail            = "success get detail"
	MessageSuccessSendVerificationMail = "send verify email success"
	MessageSuccessUpdateUser           = "update user success"
	MessageSuccessGetProfile           = "get profile success"
	MessageSuccessAddEducation         = "add education success"
	MessageSuccessUpdateEducation      = "update education success"
	MessageSuccessDeleteEducation      = "delete education success"
	MessageSuccessDeleteExperience     = "delete experience success"
	MessageSuccessAddSkill             = "add skill success"

	MessageFailedBodyRequest      = "body request failed"
	MessageFailedRegister         = "register failed"
	MessageFailedLogin            = "login failed"
	MessageFailedGetDetail        = "failed get detail"
	MessageFailedUpdateUser       = "failed update user"
	MessageFailedGetProfile       = "failed get profile"
	MessageFailedAddEducation     = "failed add education success"
	MessageFailedDeleteExperience = "failed delete experience"
	MessageFailedUpdateEducation  = "failed delete education"
	MessageFailedDeleteEducation  = "failed delete education"
	MessageFailedAddSkill         = "failed add skill"

	ErrAccountAlreadyVerified = errors.New("account already verified")
	ErrEmailAlreadyExists     = errors.New("email already exists")
	ErrUserNotFound           = errors.New("user not found")
	CredentialInvalid         = errors.New("credential invalid")
	ErrUserNotVerified        = errors.New("user not verified")
	ErrRegisterUserFailed     = errors.New("register user failed")
	ErrTokenInvalid           = errors.New("token invalid")
	ErrTokenExpired           = errors.New("token expired")
	ErrUploadFile             = errors.New("upload file failed")
	ErrUpdateEducation        = errors.New("update education failed")
	ErrDeleteEducation        = errors.New("delete education failed")
	ErrPostExperience         = errors.New("post experience failed")
	ErrUpdateExperience       = errors.New("update experience failed")
	ErrDeleteExperience       = errors.New("delete experience failed")
	ErrPostSkill              = errors.New("add skill failed")
	ErrDeleteSkill            = errors.New("delete skill failed")
	ErrGetProfile             = errors.New("get profile failed")
)

type (
	UserProfile struct {
		PersonalInfo UserPersonalInfo `json:"personal_info"`
		Educations   []UserEducation  `json:"educations"`
		Experiences  []UserExperience `json:"experiences"`
		Skills       []UserSkill      `json:"skills"`
	}

	UserPersonalInfo struct {
		Name           string `json:"name"`
		About          string `json:"about"`
		Address        string `json:"address"`
		CurrentTitle   string `json:"current_title"`
		ProfilePicture string `json:"profile_picture"`
		Headline       string `json:"headline"`
	}

	UserEducation struct {
		ID           string `json:"id"`
		SchoolName   string `json:"school_name"`
		Degree       string `json:"degree"`
		FieldOfStudy string `json:"field_of_study"`
		Description  string `json:"description"`
	}

	UserExperience struct {
		ID          string `json:"id"`
		Title       string `json:"title"`
		CompanyID   string `json:"company_id"`
		CompanyName string `json:"company_name"`
		Location    string `json:"location"`
		StartDate   string `json:"started_at"`
		EndDate     string `json:"ended_at"`
		Description string `json:"description"`
	}

	UserSkill struct {
		ID      string `json:"id"`
		SkillID string `json:"skill_id"`
		Name    string `json:"name"`
	}

	UserRegisterRequest struct {
		Name           string                `json:"name" form:"name" validate:"required"`
		Password       string                `json:"password" form:"password" validate:"required"`
		Email          string                `json:"email" form:"email" validate:"required,email"`
		About          string                `json:"about" form:"about" validate:"required"`
		Address        string                `json:"address" form:"address" validate:"required"`
		CurrentTitle   string                `json:"current_title" form:"current_title"`
		ProfilePicture *multipart.FileHeader `json:"profile_picture" form:"profile_picture"`
		Headline       *multipart.FileHeader `json:"headline" form:"headline"`
	}

	UserRegisterResponse struct {
		Name           string `json:"name"`
		Email          string `json:"email"`
		About          string `json:"about"`
		Address        string `json:"address"`
		CurrentTitle   string `json:"current_title"`
		ProfilePicture string `json:"profile_picture"`
		Headline       string `json:"headline"`
		IsPremium      bool   `json:"is_premium"`
	}

	UserLoginRequest struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	}

	UserLoginResponse struct {
		Email string `json:"email"`
		Token string `json:"token"`
		Role  string `json:"role"`
	}

	UpdateUserRequest struct {
		Name           string                `json:"name" form:"name"`
		Email          string                `json:"email" form:"email" validate:"required,email"`
		NewEmail       string                `json:"new_email" form:"new_email" validate:"email"`
		About          string                `json:"about" form:"about"`
		Address        string                `json:"address" form:"address"`
		CurrentTitle   string                `json:"current_title" form:"current_title"`
		ProfilePicture *multipart.FileHeader `json:"profile_picture" form:"profile_picture"`
		Headline       *multipart.FileHeader `json:"headline" form:"headline"`
	}

	PostUserEducationRequest struct {
		SchoolName   string `json:"school_name" form:"school_name" validate:"required"`
		Degree       string `json:"degree" form:"degree"`
		FieldOfStudy string `json:"field_of_study" form:"field_of_study"`
		Description  string `json:"description" form:"description"`
	}

	UpdateUserEducationRequest struct {
		EducationID  string `json:"education_id" form:"education_id" validate:"required"`
		SchoolName   string `json:"school_name" form:"school_name" validate:"required"`
		Degree       string `json:"degree" form:"degree"`
		FieldOfStudy string `json:"field_of_study" form:"field_of_study"`
		Description  string `json:"description" form:"description"`
	}

	PostUserExperienceRequest struct {
		Title       string `json:"title" form:"title" validate:"required"`
		CompanyID   string `json:"company_id" form:"company_id" validate:"required"`
		Location    string `json:"location" form:"location" validate:"required"`
		StartDate   string `json:"start_date" form:"start_date" validate:"required"`
		EndDate     string `json:"end_date" form:"end_date"`
		Description string `json:"description" form:"description"`
	}

	UpdateUserExperienceRequest struct {
		ExperienceID string `json:"experience_id" form:"experience_id" validate:"required"`
		Title        string `json:"title" form:"title" validate:"required"`
		CompanyID    string `json:"company_id" form:"company_id" validate:"required"`
		Location     string `json:"location" form:"location" validate:"required"`
		StartDate    string `json:"start_date" form:"start_date" validate:"required"`
		EndDate      string `json:"end_date" form:"end_date"`
		Description  string `json:"description" form:"description"`
	}

	PostUserSkillRequest struct {
		SkillID string `json:"skill_id" form:"skill_id" validate:"required"`
	}
)
