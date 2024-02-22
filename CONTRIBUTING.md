# Contributing

## Summary

This document covers how to contribute to the Calculator Project written in GoLand ```Version 1.22.0```

## Making Changes

Navigate to the project's root directory and you will find the core package named ```calculator```.  Inside you would find the source code for the calculator.
The test files are alongside it should follow the naming conventions as ```<test-file-name>_test.go```

## Test Framework - Testify

This project uses Testify as the standard for both Unit and Integration tests.
They are run as a Suite at the ```runner_test.go``` file.

Below are the naming conventions that need to be follower to run Unit and Integration Tests separately

1) Unit Test naming convention - ```UnitTestSuite```
2) Integration Test naming convention - ```TestCalculatorIntegration```

## Review and Pull Requests

After making necessary changes on you remote branch, raise a merge request and that should trigger the CI/CD test automatically.
