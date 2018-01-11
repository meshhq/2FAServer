package db_test

import (
	"testing"

	"github.com/stretchr/testify/mock"
)

type MockPGInterface struct {
	mock.Mock
}

var (
	testKeyID  = "123"
	testKey    = "c6c23f9b3fe8d06a22564aeafc78e2fc"
	testUserID = "dummyGitUser"
)

func TestGetModel(t *testing.T) {

}
