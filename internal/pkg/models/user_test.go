package models

/*
go test github.com/imayavgi/gows/internal/pkg/model
*/

import (
	"reflect"
	"testing"
)

func TestGetUsers(t *testing.T) {
	tests := []struct {
		name string
		want []*User
	}{
		{name: "Empty list", want: nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetUsers(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUsers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddUser(t *testing.T) {
	type args struct {
		u User
	}
	nu := User{FirstName: "Imaya", LastName: "Kulothungan"}
	enu := User{ID: 1, FirstName: "Imaya", LastName: "Kulothungan"}
	arg1 := args{nu}
	tests := []struct {
		name    string
		args    args
		want    User
		wantErr bool
	}{
		{"New User", arg1, enu, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AddUser(tt.args.u)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AddUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
