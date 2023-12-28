package connection

import (
	"go-election/config"
	"testing"
)

var cfg config.DBConfig

func init() {
	all, _ := config.LoadConfig()
	cfg = all.DB
}

func TestConnectDB(t *testing.T) {
	_, _, err := ConnectDB()
	if err != nil {
		t.Errorf("Failed to connect to database: %v\n", err)
		return
	}
	t.Log("Connected to database")
}
