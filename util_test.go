package main

func (suite *SkeletonTestSuite) TestISO8601() {
	suite.Equal("2006-01-02T15:04:05Z07:00", ISO8601)
}
