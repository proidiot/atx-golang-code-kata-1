package main

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type SkeletonTestSuite struct {
	suite.Suite
	Logger *Logger
	Conf   Config
}

func (suite *SkeletonTestSuite) SetupTest() {
	suite.Conf, _ = NewConfig()
	suite.Logger = NewLogger(suite.Conf)
}

func (suite *SkeletonTestSuite) TestVersion() {
	suite.NotEqual("", Version())
}

func TestTestSuite(t *testing.T) {
	suite.Run(t, new(SkeletonTestSuite))
}
