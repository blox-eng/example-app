package config

import (
	"reflect"
	"testing"
)

func TestInitialize(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Initialize(tt.args.path)
		})
	}
}

func TestGet(t *testing.T) {
	type args struct {
		configName string
	}
	tests := []struct {
		name string
		args args
		want interface{}
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Get(tt.args.configName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_readConfig(t *testing.T) {
	type args struct {
		filename string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			readConfig(tt.args.filename)
		})
	}
}

func TestGetYamlValues(t *testing.T) {
	tests := []struct {
		name string
		want *yamlConfig
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetYamlValues(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetYamlValues() = %v, want %v", got, tt.want)
			}
		})
	}
}
