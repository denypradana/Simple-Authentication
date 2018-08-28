package main

import (
	"net/http"
	"reflect"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func Test_connectDb(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for range tests {
		t.Run(tt.name, func(t *testing.T) {
			connectDb()
		})
	}
}

func Test_routes(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for range tests {
		t.Run(tt.name, func(t *testing.T) {
			routes()
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}

func Test_checkErr(t *testing.T) {
	type args struct {
		w   http.ResponseWriter
		r   *http.Request
		err error
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkErr(tt.args.w, tt.args.r, tt.args.err); got != tt.want {
				t.Errorf("checkErr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestQueryUser(t *testing.T) {
	type args struct {
		username string
	}
	tests := []struct {
		name string
		args args
		want User
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := QueryUser(tt.args.username); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("QueryUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_register(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			register(tt.args.w, tt.args.r)
		})
	}
}

func Test_login(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			login(tt.args.w, tt.args.r)
		})
	}
}

func Test_home(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			home(tt.args.w, tt.args.r)
		})
	}
}

func Test_logout(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			logout(tt.args.w, tt.args.r)
		})
	}
}
