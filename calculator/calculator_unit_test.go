package calculator_test

import (
	"TestCalculator/calculator"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type UnitTestSuite struct {
	suite.Suite
}

func (t *UnitTestSuite) TestCalculateFunction() {
	tests := []struct {
		operand1 float64
		operator string
		operand2 float64
		expected float64
		err      error
	}{
		{1, "+", 1, 2, nil},
		{5, "-", 3, 2, nil},
		{2, "*", 3, 6, nil},
		{6, "/", 2, 3, nil},
		{6, "/", 0, 0, errors.New("division by zero is not allowed")},
		{1, "%", 1, 0, errors.New("invalid operator")},
	}
	var a int = 0
	var totalTests = len(tests)
	for _, test := range tests {
		result, err := calculator.Calculate(test.operand1, test.operand2, test.operator)

		if test.err == nil {
			assert.Equal(t.T(), test.expected, result, "Basic unit _test operation failed for"+test.operator)
			a = a + 1
		}

		if test.err != nil {
			assert.Error(t.T(), test.err, err.Error(), "Basic error handling failed for "+test.err.Error())
			a = a + 1
		}
	}
	if assert.Equal(t.T(), totalTests, a, "[FAIL] ---->>>> Unit Test No:", a, "did not pass") {
		println("[SUCCESS] ---->>>> All Unit tests passed in Calculate Function")
	}
}

func (t *UnitTestSuite) TestParseOperandFunction() {
	tests := []struct {
		input    string
		expected float64
		err      error
	}{
		{"1", 1, nil},
		{"3.14", 3.14, nil},
		{"abc", 0, errors.New("invalid operand")},
	}
	var b int = 0
	var total = len(tests)
	for _, test := range tests {
		result, err := calculator.ParseOperand(test.input)

		if test.err == nil {
			assert.Equal(t.T(), test.expected, result)
			b = b + 1
		}

		if test.err != nil {
			assert.Error(t.T(), test.err, err)
			b = b + 1
		}
	}
	if assert.Equal(t.T(), total, b, "[FAIL] ---->>>>  Test No:", b, "did not pass") {
		println("[SUCCESS] ---->>>> All Unit tests passed in ParseOperand Function")
	}
}
