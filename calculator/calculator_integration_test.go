package calculator_test

import (
	"TestCalculator/calculator"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type TestCalculatorIntegration struct {
	suite.Suite
}

func (t *TestCalculatorIntegration) TestCalculatorIntegration() {
	tests := []struct {
		operand1       float64
		operator       string
		operand2       float64
		expected       float64
		err            error
		failMessage    string
		invalidOperand string
	}{
		{operand1: 1, operator: "+", operand2: 1, expected: 2, failMessage: "Addition Operation Failed", invalidOperand: ""},
		{5, "-", 3, 2, nil, "Subtraction Operation Failed", ""},
		{5, "-", 10, -5, nil, "Subtraction Operation Failed", ""},
		{2, "*", 3, 6, nil, "Multiplication Operation Failed", ""},
		{6, "/", 2, 3, nil, "Division Operation Failed for Decimals", ""},
		{6.6, "/", 2, 3.3, nil, "Division Operation Failed", ""},
		{6, "/", 0, 0, errors.New("division by zero is not allowed"), "Error handling for divide by 0 failed", "abc"},
		{1, "%", 1, 0, errors.New("invalid operator"), "Error handling for invalid Operator failed", "bed"},
		{1, "**", 1, 0, errors.New("invalid operator"), "Error handling for invalid Operator failed", "cwer"},
		{1, "*", 11111111111111111111111, 11111111111111111111112, nil, "Error handling for invalid Operator failed", "dqw"},
		{0, "-", 3, -3, nil, "Subtraction Operation Failed", ""},
	}
	var a int = 0
	var totalTests = len(tests)
	for _, test := range tests {
		result, testError := calculator.Calculate(test.operand1, test.operand2, test.operator)
		if test.err == nil {
			assert.Equalf(t.T(), test.expected, result, test.failMessage)
			println("Test No:", a+1, "Passed for ", test.operator, "operator")
		}
		if test.err != nil {
			assert.EqualError(t.T(), testError, test.err.Error())
			println("Test No:", a+1, "Passed for ", test.err.Error(), "exception")

			if test.invalidOperand != "" {
				result1, testError1 := calculator.ParseOperand(test.invalidOperand)
				assert.EqualError(t.T(), testError1, "invalid operand", "error handling failed")
				assert.Equal(t.T(), float64(0), result1, "invalid return 0 failed")
			}

		}
		a = a + 1
	}
	if assert.Equalf(t.T(), totalTests, a, "[FAIL] ---->>>> Did not run all the tests/ Failures in tests") {
		println("Ran a total of ", a, "Tests")
		println("[SUCCESS] ---->>>> All Integration Tests passed")
	}
}
