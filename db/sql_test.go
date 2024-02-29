package db

import (
	"reflect"
	"testing"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func TestInitializeDatabaseConnector(t *testing.T) {
	tests := []struct {
		name string
		want *sqlConn
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := InitializeDatabaseConnector(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("InitializeDatabaseConnector() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_initializeDatabaseConnector(t *testing.T) {
	tests := []struct {
		name    string
		want    *sqlConn
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := initializeDatabaseConnector()
			if (err != nil) != tt.wantErr {
				t.Errorf("initializeDatabaseConnector() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("initializeDatabaseConnector() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetDBConnection(t *testing.T) {
	tests := []struct {
		name string
		want *gorm.DB
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetDBConnection(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetDBConnection() = %v, want %v", got, tt.want)
			}
		})
	}
}
