package user

import (
	"Go-Starter-Template/domain"
	"Go-Starter-Template/entities"
	"context"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type (
	UserRepository interface {
		RegisterUser(ctx context.Context, req entities.User) (entities.User, error)
		CheckUserByEmail(ctx context.Context, email string) bool
		GetUserByEmail(ctx context.Context, email string) (entities.User, error)
		CheckUserByID(ctx context.Context, id string) bool
		UpdateSubscriptionStatus(ctx context.Context, userID string) error
		GetProfile(ctx context.Context, slug string) (domain.UserProfileResponse, error)
		UpdateProfile(ctx context.Context, user entities.User) error
		PostEducation(ctx context.Context, req entities.UserEducation) error
		UpdateEducation(ctx context.Context, req entities.UserEducation) error
		DeleteEducation(ctx context.Context, id uuid.UUID) error
		PostExperience(ctx context.Context, req entities.UserExperience) error
		UpdateExperience(ctx context.Context, req entities.UserExperience) error
		DeleteExperience(ctx context.Context, id uuid.UUID) error
		PostSkill(ctx context.Context, req entities.UserSkill) error
		DeleteSkill(ctx context.Context, id uuid.UUID) error
	}
	userRepository struct {
		db *gorm.DB
	}
)

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) RegisterUser(ctx context.Context, req entities.User) (entities.User, error) {
	if err := r.db.WithContext(ctx).Create(&req).Error; err != nil {
		return entities.User{}, err
	}
	return req, nil
}

func (r *userRepository) CheckUserByEmail(ctx context.Context, email string) bool {
	var user entities.User
	if err := r.db.WithContext(ctx).First(user, "email = ?", email).Error; err != nil {
		return false
	}
	if user.Email != email {
		return false
	}
	return true
}

func (r *userRepository) CheckUserByID(ctx context.Context, id string) bool {
	var user entities.User
	if err := r.db.WithContext(ctx).First(&user, "id = ?", id).Error; err != nil {
		return false
	}

	if user.ID.String() != id {
		return false
	}
	return true
}

func (r *userRepository) GetUserByEmail(ctx context.Context, email string) (entities.User, error) {
	var user entities.User
	if err := r.db.WithContext(ctx).First(&user, "email = ?", email).Error; err != nil {
		return entities.User{}, err
	}
	return user, nil
}

func (r *userRepository) UpdateSubscriptionStatus(ctx context.Context, userID string) error {
	if err := r.db.WithContext(ctx).
		Model(&entities.User{}).
		Where("id = ?", userID).
		Updates(map[string]interface{}{"subscribe": true}).Error; err != nil {
		return err
	}
	return nil
}

func (r *userRepository) GetProfile(ctx context.Context, slug string) (domain.UserProfileResponse, error) {
	var user entities.User
	var education []entities.UserEducation
	var experience []entities.UserExperience
	var skill []entities.UserSkill

	if err := r.db.WithContext(ctx).
		Where("slug = ?", slug).
		First(&user).Error; err != nil {
		return domain.UserProfileResponse{}, err
	}

	if err := r.db.WithContext(ctx).Find(&education, "user_id = ?", user.ID).Error; err != nil {
		return domain.UserProfileResponse{}, err
	}

	if err := r.db.WithContext(ctx).Preload("Company").Find(&experience, "user_id = ?", user.ID).Error; err != nil {
		return domain.UserProfileResponse{}, err
	}

	if err := r.db.WithContext(ctx).Preload("Skill").Find(&skill, "user_id = ?", user.ID).Error; err != nil {
		return domain.UserProfileResponse{}, err
	}

	formattedEducations := make([]domain.UserEducationsResponse, len(education))
	for i, edu := range education {
		formattedEducations[i] = domain.UserEducationsResponse{
			ID:           edu.ID.String(),
			SchoolName:   edu.SchoolName,
			Degree:       edu.Degree,
			FieldOfStudy: edu.FieldOfStudy,
			Description:  edu.Description,
			StartDate:    edu.StartedAt.Format("01-02-2006"),
			EndDate:      edu.EndedAt.Format("01-02-2006"),
		}
	}

	formattedExperiences := make([]domain.UserExperiencesResponse, len(experience))
	for i, exp := range experience {
		formattedExperiences[i] = domain.UserExperiencesResponse{
			ID:          exp.ID.String(),
			Title:       exp.Title,
			CompanyID:   exp.CompanyID.String(),
			CompanyName: exp.Company.Name,
			Location:    exp.Location,
			StartDate:   exp.StartedAt.Format("01-02-2006"),
			EndDate:     exp.EndedAt.Format("01-02-2006"),
			Description: exp.Description,
		}
	}

	formattedSkills := make([]domain.UserSkillsResponse, len(skill))
	for i, sk := range skill {
		formattedSkills[i] = domain.UserSkillsResponse{
			ID:      sk.ID.String(),
			SkillID: sk.SkillID.String(),
			Name:    sk.Skill.Name,
		}
	}

	return domain.UserProfileResponse{
		PersonalInfo: domain.UserPersonalInfoResponse{
			Name:           user.Name,
			About:          user.About,
			Address:        user.Address,
			CurrentTitle:   user.CurrentTitle,
			ProfilePicture: user.ProfilePicture,
			Headline:       user.Headline,
		},
		Educations:  formattedEducations,
		Experiences: formattedExperiences,
		Skills:      formattedSkills,
	}, nil
}

func (r *userRepository) UpdateProfile(ctx context.Context, user entities.User) error {
	if err := r.db.WithContext(ctx).Updates(&user).Error; err != nil {
		return err
	}
	return nil
}

func (r *userRepository) PostEducation(ctx context.Context, req entities.UserEducation) error {
	if err := r.db.WithContext(ctx).Create(&req).Error; err != nil {
		return err
	}
	return nil
}

func (r *userRepository) UpdateEducation(ctx context.Context, req entities.UserEducation) error {
	if err := r.db.WithContext(ctx).Updates(&req).Error; err != nil {
		return err
	}
	return nil
}

func (r *userRepository) DeleteEducation(ctx context.Context, id uuid.UUID) error {
	if err := r.db.WithContext(ctx).Delete(&entities.UserEducation{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *userRepository) PostExperience(ctx context.Context, req entities.UserExperience) error {
	if err := r.db.WithContext(ctx).Create(&req).Error; err != nil {
		return err
	}
	return nil
}

func (r *userRepository) UpdateExperience(ctx context.Context, req entities.UserExperience) error {
	if err := r.db.WithContext(ctx).Updates(&req).Error; err != nil {
		return err
	}
	return nil
}

func (r *userRepository) DeleteExperience(ctx context.Context, id uuid.UUID) error {
	if err := r.db.WithContext(ctx).Delete(&entities.UserExperience{}, id).Error; err != nil {
		return err
	}
	return nil
}

func (r *userRepository) PostSkill(ctx context.Context, req entities.UserSkill) error {
	if err := r.db.WithContext(ctx).Create(&req).Error; err != nil {
		return err
	}
	return nil
}

func (r *userRepository) DeleteSkill(ctx context.Context, id uuid.UUID) error {
	if err := r.db.WithContext(ctx).Delete(&entities.UserSkill{}, id).Error; err != nil {
		return err
	}
	return nil
}
