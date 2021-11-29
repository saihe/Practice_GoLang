package domain

import (
	"log"
	"reflect"
	"syscall"
	"testing"
	"time"
)

func TestFindById(t *testing.T) {
	type args struct {
		id uint64
	}
	syscall.Getenv("")
	createdAt, _ := time.Parse("2006-01-02 15:04:05.999999", "2021-11-26 23:57:30.200418")
	updatedAt, _ := time.Parse("2006-01-02 15:04:05.999999", "2021-11-26 23:57:30.200418")
	tests := []struct {
		name string
		args args
		want Employee
	}{
		{"case 1", args{1}, Employee{1, "齋藤恭平", 1, createdAt, updatedAt}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindById(tt.args.id); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindById() = %v, want %v", got, tt.want)
			} else {
				log.Println(got)
			}
		})
	}
}
