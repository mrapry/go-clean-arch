package helper

import (
	"reflect"
	"testing"

	log "github.com/sirupsen/logrus"
)

func TestLogContext(t *testing.T) {
	type args struct {
		c string
		s string
	}
	tests := []struct {
		name string
		args args
		want *log.Entry
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LogContext(tt.args.c, tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LogContext() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLog(t *testing.T) {
	type args struct {
		level   log.Level
		message string
		context string
		scope   string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Log(tt.args.level, tt.args.message, tt.args.context, tt.args.scope)
		})
	}
}
