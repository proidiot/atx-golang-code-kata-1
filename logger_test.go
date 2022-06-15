package main

import (
	"errors"
	"io"
)

func (suite *SkeletonTestSuite) TestlogWriter() {
	suite.Implements((*io.Writer)(nil), logWriter("test"))
}

func (suite *SkeletonTestSuite) TestNewLogger() {
	suite.IsType(&Logger{}, NewLogger(suite.Conf))
}

func (suite *SkeletonTestSuite) TestLoggerLogError() {
	err := errors.New("test error")
	suite.Error(suite.Logger.LogError(err, ""))
}
