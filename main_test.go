package main

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type KataTestSuite struct {
	suite.Suite
	Logger *Logger
	Conf   Config
}

func (suite *KataTestSuite) SetupTest() {
	suite.Conf, _ = NewConfig()
	suite.Logger = NewLogger(suite.Conf)
}

func (suite *KataTestSuite) TestVersion() {
	suite.NotEqual("", Version())
}

func TestTestSuite(t *testing.T) {
	suite.Run(t, new(KataTestSuite))
}
