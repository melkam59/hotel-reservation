package db

import (
	"context"

	"github.com/fulltimegodev/hotel-reservation/ent"
	"github.com/fulltimegodev/hotel-reservation/ent/user"
	"github.com/fulltimegodev/hotel-reservation/types"
)

type Map map[string]any

type Dropper interface {
	Drop(context.Context) error
}

type UserStore interface {
	Dropper

	GetUserByEmail(context.Context, string) (*types.User, error)
	GetUserByID(context.Context, int) (*types.User, error)
	GetUsers(context.Context) ([]*types.User, error)
	InsertUser(context.Context, *types.User) (*types.User, error)
	DeleteUser(context.Context, int) error
	UpdateUser(ctx context.Context, filter Map, params types.UpdateUserParams) error
}

type EntUserStore struct {
	client *ent.Client
}

func NewEntUserStore(client *ent.Client) *EntUserStore {
	return &EntUserStore{
		client: client,
	}
}

func (s *EntUserStore) Drop(ctx context.Context) error {
	_, err := s.client.User.Delete().Exec(ctx)
	return err
}

func (s *EntUserStore) UpdateUser(ctx context.Context, filter Map, params types.UpdateUserParams) error {
	id, ok := filter["_id"].(int)
	if !ok {
		return nil // or error
	}
	upd := s.client.User.UpdateOneID(id)
	if params.FirstName != "" {
		upd.SetFirstName(params.FirstName)
	}
	if params.LastName != "" {
		upd.SetLastName(params.LastName)
	}
	_, err := upd.Save(ctx)
	return err
}

func (s *EntUserStore) DeleteUser(ctx context.Context, id int) error {
	return s.client.User.DeleteOneID(id).Exec(ctx)
}

func (s *EntUserStore) InsertUser(ctx context.Context, userType *types.User) (*types.User, error) {
	u, err := s.client.User.Create().
		SetFirstName(userType.FirstName).
		SetLastName(userType.LastName).
		SetEmail(userType.Email).
		SetEncryptedPassword(userType.EncryptedPassword).
		SetIsAdmin(userType.IsAdmin).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	userType.ID = u.ID
	return userType, nil
}

func (s *EntUserStore) GetUsers(ctx context.Context) ([]*types.User, error) {
	users, err := s.client.User.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	var res []*types.User
	for _, u := range users {
		res = append(res, toTypeUser(u))
	}
	return res, nil
}

func (s *EntUserStore) GetUserByEmail(ctx context.Context, email string) (*types.User, error) {
	u, err := s.client.User.Query().Where(user.Email(email)).Only(ctx)
	if err != nil {
		return nil, err
	}
	return toTypeUser(u), nil
}

func (s *EntUserStore) GetUserByID(ctx context.Context, id int) (*types.User, error) {
	u, err := s.client.User.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return toTypeUser(u), nil
}

func toTypeUser(u *ent.User) *types.User {
	if u == nil {
		return nil
	}
	return &types.User{
		ID:                u.ID,
		FirstName:         u.FirstName,
		LastName:          u.LastName,
		Email:             u.Email,
		EncryptedPassword: u.EncryptedPassword,
		IsAdmin:           u.IsAdmin,
	}
}
