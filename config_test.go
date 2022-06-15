package main

func (suite *SkeletonTestSuite) TestNewConfig() {
	conf, err := NewConfig()
	suite.NoError(err)
	suite.IsType(Config{}, conf)
}

func (suite *SkeletonTestSuite) TestConfig() {
	suite.NotEmpty(suite.Conf.AppName)
	suite.Equal("test", suite.Conf.AppEnv)
}
