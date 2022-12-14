package user

import (
	"context"

	"github.com/google/uuid"
	"github.com/mukul1234567/Library-Management-System/db"
	"go.uber.org/zap"
)

type Service interface {
	List(ctx context.Context) (response ListResponse, err error)
	Show(ctx context.Context) (responser ListResponser, err error)
	Create(ctx context.Context, req CreateRequest) (err error)
	FindByID(ctx context.Context, id string) (response FindByIDResponse, err error)
	DeleteByID(ctx context.Context, id string) (err error)
	Update(ctx context.Context, req UpdateRequest) (err error)
	UpdatePassword(ctx context.Context, req UpdatePasswordStruct) (err error)
}

type userService struct {
	store  db.Storer
	logger *zap.SugaredLogger
}

// var userinfo []db.User
func (cs *userService) List(ctx context.Context) (response ListResponse, err error) {
	users, err := cs.store.ListUsers(ctx)
	if err == db.ErrUserNotExist {
		cs.logger.Error("No user present", "err", err.Error())
		return response, errNoUsers
	}
	if err != nil {
		cs.logger.Error("Error listing users", "err", err.Error())
		return
	}

	response.User = users
	// userinfo = users
	return
}

func (cs *userService) Show(ctx context.Context) (responser ListResponser, err error) {
	users, err := cs.store.ShowUsers(ctx)
	if err == db.ErrUserNotExist {
		cs.logger.Error("No user present", "err", err.Error())
		return responser, errNoUsers
	}
	if err != nil {
		cs.logger.Error("Error listing users", "err", err.Error())
		return
	}

	responser.Users = users
	// userinfo = users
	return
}

func (cs *userService) Create(ctx context.Context, c CreateRequest) (err error) {
	// err = c.Validate()
	// if err != nil {
	// 	cs.logger.Errorw("Invalid request for user create", "msg", err.Error(), "user", c)
	// 	return
	// }
	uuidgen := uuid.New()
	c.ID = uuidgen.String()
	err = cs.store.CreateUser(ctx, &db.User{

		ID:        c.ID,
		FirstName: c.FirstName,
		LastName:  c.LastName,
		Gender:    c.Gender,
		Address:   c.Address,
		Age:       c.Age,
		Email:     c.Email,
		Password:  c.Password,
		MobileNum: c.MobileNum,
		Role:      c.Role,
	})
	if err != nil {
		cs.logger.Error("Error creating user", "err", err.Error())
		return
	}
	return
}

func (cs *userService) Update(ctx context.Context, c UpdateRequest) (err error) {
	err = c.Validate()
	if err != nil {
		cs.logger.Error("Invalid Request for user update", "err", err.Error(), "user", c)
		return
	}

	err = cs.store.UpdateUser(ctx, &db.User{
		ID:        c.ID,
		FirstName: c.FirstName,
		LastName:  c.LastName,
		Gender:    c.Gender,
		Address:   c.Address,
		Age:       c.Age,
		Email:     c.Email,
		// Password:  c.NewPassword,
		Password:  c.Password,
		MobileNum: c.MobileNum,
		Role:      c.Role,
	})
	if err != nil {
		cs.logger.Error("Error updating user", "err", err.Error(), "user", c)
		return
	}

	return
}

func (cs *userService) UpdatePassword(ctx context.Context, c UpdatePasswordStruct) (err error) {
	// err = c.Validate()
	if err != nil {
		cs.logger.Error("Invalid Request for user update", "err", err.Error(), "user", c)
		return
	}

	err = cs.store.UpdatePassword(ctx, &db.User{
		ID:       c.ID,
		Password: c.NewPassword,
	})
	if err != nil {
		cs.logger.Error("Error updating user", "err", err.Error(), "user", c)
		return
	}

	return
}

func (cs *userService) FindByID(ctx context.Context, id string) (response FindByIDResponse, err error) {
	user, err := cs.store.FindUserByID(ctx, id)
	if err == db.ErrUserNotExist {
		cs.logger.Error("No user present", "err", err.Error())
		return response, errNoUserId
	}
	if err != nil {
		cs.logger.Error("Error finding user", "err", err.Error(), "id", id)
		return
	}

	response.User = user
	return
}

func (cs *userService) DeleteByID(ctx context.Context, id string) (err error) {
	err = cs.store.DeleteUserByID(ctx, id)
	if err == db.ErrUserNotExist {
		cs.logger.Error("user Not present", "err", err.Error(), "id", id)
		return errNoUserId
	}
	if err != nil {
		cs.logger.Error("Error deleting user", "err", err.Error(), "id", id)
		return
	}

	return
}

func NewService(s db.Storer, l *zap.SugaredLogger) Service {
	return &userService{
		store:  s,
		logger: l,
	}
}
