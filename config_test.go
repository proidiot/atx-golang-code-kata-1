package main

func (suite *KataTestSuite) TestNewConfig() {
	conf, err := NewConfig()
	suite.NoError(err)
	suite.IsType(Config{}, conf)
}

func (suite *KataTestSuite) TestConfig() {
	suite.NotEmpty(suite.Conf.AppName)
	suite.Equal("test", suite.Conf.AppEnv)
}
