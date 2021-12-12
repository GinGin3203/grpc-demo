package repository

import (
	"context"
	"github.com/GinGin3203/grpc-demo/server/model"
	"github.com/GinGin3203/grpc-demo/server/repository/query"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4"
	"log"
)

type Repository interface {
	AddUser(ctx context.Context, role, name string) (*model.UserRecord, error)
	ChangeUserRole(ctx context.Context, uid int, newRole string) (*model.UserRecord, error)
	DeleteUser(ctx context.Context, uid int) (*model.UserRecord, error)
	GetUser(ctx context.Context, uid int) (*model.UserRecord, error)
	GetAllUsers(ctx context.Context) ([]*model.UserRecord, error)
}

type repository struct {
	*pgx.Conn
}

func New(conn *pgx.Conn) Repository {
	return &repository{
		conn,
	}
}
func (r repository) AddUser(ctx context.Context, role, name string) (insertedUser *model.UserRecord, err error) {
	log.Println(role, name)
	insertedUser = &model.UserRecord{}
	err = pgxscan.Get(ctx, r.Conn, insertedUser, query.AddUser, name, role)
	if err != nil {
		return nil, err
	}
	return insertedUser, nil
}

func (r *repository) ChangeUserRole(ctx context.Context, uid int, newRole string) (updatedUser *model.UserRecord, err error) {
	log.Println(uid, newRole)
	updatedUser = &model.UserRecord{}
	err = pgxscan.Get(ctx, r.Conn, updatedUser, query.UpdateUserRole, uid, newRole)
	if err != nil {
		return nil, err
	}

	return updatedUser, err
}

func (r *repository) DeleteUser(ctx context.Context, uid int) (removedUser *model.UserRecord, err error) {
	log.Println(uid)
	removedUser = &model.UserRecord{}
	err = pgxscan.Get(ctx, r.Conn, removedUser, query.DeleteUser, uid)
	if err != nil {
		return nil, err
	}
	return removedUser, nil
}

func (r *repository) GetUser(ctx context.Context, uid int) (user *model.UserRecord, err error) {
	log.Println(uid)
	user = &model.UserRecord{}
	err = pgxscan.Get(ctx, r.Conn, user, query.GetUser, uid)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *repository) GetAllUsers(ctx context.Context) (users []*model.UserRecord, err error) {
	users = make([]*model.UserRecord, 0)
	err = pgxscan.Select(ctx, r.Conn, &users, query.GetAllUsers)
	if err != nil {
		return nil, err
	}
	return users, nil
}
