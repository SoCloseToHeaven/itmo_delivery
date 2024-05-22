package service

import (
	"itmo_delivery/model"
	"itmo_delivery/repository"

	tgbotapi "github.com/Syfaro/telegram-bot-api"
	"gorm.io/gorm"
)

type UserService interface {
	GetOrCreateUser(u *tgbotapi.Update) (*model.User, error)
}

type userService struct {
	UserRepository repository.UserRepository
}

func NewUserService(db *gorm.DB) UserService {
	return &userService{
		UserRepository: repository.NewUserRepository(db),
	}
}

func (r *userService) GetOrCreateUser(u *tgbotapi.Update) (*model.User, error) {
	chatID := u.Message.Chat.ID
	tgID := u.Message.From.ID

	tx := r.UserRepository.DB().Begin()

	defer func() {
		if rec := recover(); rec != nil {
			tx.Rollback()
		}
	}()

	user, err := r.UserRepository.GetByChatID(chatID)

	if err == nil {
		tx.Rollback()
		return user, nil
	}

	user = &model.User{
		ChatID:     chatID,
		TelegramID: tgID,
		State:      model.Main,
	}

	if err := tx.Save(user).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Commit()
	return user, nil
}
