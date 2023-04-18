package test

import (
	"AuthService/app/models"
	"AuthService/app/test/moks"
	"errors"
	"reflect"
	"testing"
)

func TestRepository_Create(t *testing.T) {
	type args struct {
		userAuth models.LoginUser
	}
	tests := []struct {
		name    string
		args    args
		want    models.User
		wantErr bool
	}{
		{
			name: "success_create_user",
			args: args{models.LoginUser{
				Login:    "tester@yandex.ru",
				Password: "tester",
			}},
			want: models.User{
				NdsLogin: "tester@yandex.ru",
				Password: "tester",
			},
		},
		{
			name:    "failed_create_user",
			args:    args{models.LoginUser{}},
			want:    models.User{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rep := moks.NewMockIUserRepository(t)
			if !tt.wantErr {
				rep.
					On("Create", tt.args.userAuth).
					Return(models.User{
						NdsLogin: "tester@yandex.ru",
						Password: "tester",
					}, nil)
			} else {
				rep.On("Create", tt.args.userAuth).Return(tt.want, errors.New("failed check data"))
			}
			got, err := rep.Create(tt.args.userAuth)
			if (err != nil) != tt.wantErr {
				t.Errorf("Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Create() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepository_Get(t *testing.T) {
	type args struct {
		userAuth models.LoginUser
	}
	tests := []struct {
		name    string
		args    args
		want    models.User
		wantErr bool
	}{
		{
			name: "success_get_user",
			args: args{models.LoginUser{
				Login:    "tester@yandex.ru",
				Password: "tester",
			}},
			want: models.User{
				NdsLogin: "tester@yandex.ru",
				Password: "tester",
			},
			wantErr: false,
		},
		{
			name:    "failed_get_user",
			args:    args{models.LoginUser{}},
			want:    models.User{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rep := moks.NewMockIUserRepository(t)
			if !tt.wantErr {
				rep.
					On("Get", tt.args.userAuth).
					Return(models.User{
						NdsLogin: "tester@yandex.ru",
						Password: "tester",
					}, nil)
			} else {
				rep.On("Get", tt.args.userAuth).Return(tt.want, errors.New("failed check data"))
			}
			got, err := rep.Get(tt.args.userAuth)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRepository_GetRoles(t *testing.T) {
	type args struct {
		user models.User
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "success_get_user",
			args: args{models.User{
				NdsLogin: "tester@yandex.ru",
				Password: "tester",
			}},
			want: []string{"parent", "admin"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rep := moks.NewMockIUserRepository(t)
			if !tt.wantErr {
				rep.
					On("GetRoles", tt.args.user).
					Return([]string{"parent", "admin"})
			} else {
				rep.On("GetRoles", tt.args.user).Return(tt.want, errors.New("failed check data"))
			}
			if got := rep.GetRoles(tt.args.user); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetRoles() = %v, want %v", got, tt.want)
			}
		})
	}
}
