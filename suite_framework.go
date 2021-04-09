package testutils

import (
	"github.com/stretchr/testify/suite"
)

type BDTestSuite struct {
	suite.Suite
}

func (s *BDTestSuite) AfterTest(suiteName, testName string) {
}

func (s *BDTestSuite) SetupTest() {
}

func (s *BDTestSuite) Scenario(name string, steps ...func(s **BDTestSuite)) {
	s.Run(name, func() {
		for _, step := range steps {
			s.NotPanics(func() {
				step(&s)
			})
		}
	})
}

func (s *BDTestSuite) Given(fn func(s **BDTestSuite)) func(s **BDTestSuite) {
	return fn
}

func (s *BDTestSuite) When(fn func(s **BDTestSuite)) func(s **BDTestSuite) {
	return fn
}

func (s *BDTestSuite) Then(fn func(s **BDTestSuite)) func(s **BDTestSuite) {
	return fn
}

func (s *BDTestSuite) And(fn func(s **BDTestSuite)) func(s **BDTestSuite) {
	return fn
}
