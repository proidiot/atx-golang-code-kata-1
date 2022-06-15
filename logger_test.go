package main

import (
	"errors"
	"io"
)

func (suite *KataTestSuite) TestlogWriter() {
	suite.Implements((*io.Writer)(nil), logWriter("test"))
}

func (suite *KataTestSuite) TestNewLogger() {
	suite.IsType(&Logger{}, NewLogger(suite.Conf))
}

func (suite *KataTestSuite) TestLoggerLogError() {
	err := errors.New("test error")
	suite.Error(suite.Logger.LogError(err, ""))
}
