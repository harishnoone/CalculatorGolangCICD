package calculator_test

import (
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestMyUnitSuite(t *testing.T) {
	suite.Run(t, new(UnitTestSuite))
}

func TestMyIntegrationSuite(t *testing.T) {
	suite.Run(t, new(TestCalculatorIntegration))
}
