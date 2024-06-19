package users

import (
	"context"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/odanraujo/financial-organizer-api/infrastructure/excp"
	entity "github.com/odanraujo/financial-organizer-api/internal/entity/users"
	"github.com/odanraujo/financial-organizer-api/internal/repository/users/mocks"
	"github.com/stretchr/testify/assert"
)

func TestNewUserUsecase(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	rp := mocks.NewMockUserRepository(ctrl)
	instance := NewUserUsecase(rp)

	assert.NotNil(t, instance)
}

func TestCreateUser(t *testing.T) {
	ctx := context.Background()

	address := entity.Address{
		Street:  "rua sem nome",
		City:    "Recife",
		State:   "PE",
		ZipCode: "50050202",
		Country: "Brazil",
	}

	user := entity.CreateUser{
		Name:      "test",
		CPF:       "99999999999",
		BirthDate: time.Now(),
		Address:   address,
	}

	type fields struct {
		userRepo func(mocks *mocks.MockUserRepository) *mocks.MockUserRepository
	}

	type args struct {
		ctx  context.Context
		user entity.CreateUser
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr *excp.Exception
	}{
		{
			name: "Should return create user success",
			fields: fields{
				userRepo: func(mocks *mocks.MockUserRepository) *mocks.MockUserRepository {
					mocks.EXPECT().CreateUser(gomock.AssignableToTypeOf(ctx), gomock.Any()).
						DoAndReturn(func(ctx context.Context, u entity.CreateUser) (entity.CreateUser, error) {
							assert.Equal(t, "test", u.Name)
							assert.Equal(t, "99999999999", u.CPF)
							assert.Equal(t, address, u.Address)
							return entity.CreateUser{Name: u.Name}, nil
						}).Times(1)
					return mocks
				},
			},
			args: args{
				ctx:  ctx,
				user: user,
			},
			wantErr: nil,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			userRepo := mocks.NewMockUserRepository(ctrl)

			if test.fields.userRepo != nil {
				userRepo = test.fields.userRepo(userRepo)
			}

			instance := NewUserUsecase(userRepo)

			_, err := instance.CreateUser(test.args.ctx, test.args.user)

			assert.Equal(t, test.wantErr, err)
		})
	}
}
