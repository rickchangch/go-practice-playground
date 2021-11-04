package suitepkg_test

import (
	"fmt"
	"go-practice-playground/localTest/db"
	"testing"

	"github.com/stretchr/testify/suite"
)

type TestSuite struct {
	suite.Suite
}

// entry of the suite test
func TestRunTestSuite(t *testing.T) {
	s := new(TestSuite)
	suite.Run(t, s)
}

// setup func: will run before the tests in the suite are run
func (s *TestSuite) SetupSuite() {
	// prepare test environment
	db.Init()
}

// setup func: will run before the tests in the suite are run
func (s *TestSuite) TearDownSuite() {
	// teardown test DB
	db.Teardown()
}

// test case A
func (s *TestSuite) TestTestCaseA() {
	fmt.Printf("-- case A\n")

	s.NoError(nil)
}
func (s *TestSuite) TestTestCaseB() {
	fmt.Printf("-- case B\n")

	s.NoError(nil)
}
