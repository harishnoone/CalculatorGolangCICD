# GoLang-Calculator-CI/CD

## Getting started

This project is a demonstration of a simple calculator written in GoLang and it's CI/CD integration.

## Install and build

Navigate to the Project's root directory and run the following command

```go
go install -v -n -a
``` 

## Test

Navigate to the Project's root directory to the calculator package and run the following commands to Test the code.

To run Unit Tests run the below command
```go
go test -run Unit
```
To run Integration Tests run the below command

```go
go test -run Integration
``` 

## CI/CD

This project's CI/CD pipeline is built on the ```.gitlab-ci.yml```.  This pipeline has 4 stages, they are:
1) Build
2) Test
   1) Unit Test
   2) Integration Test
3) Report
4) Deploy

The Pipeline Test phase is Automatically Triggered only during a Merge Request.  For normal commits/pushes it will be a manual button that can be triggered to run the tests. 

Reporting Phase requires the Unit Test and Integration Test jobs to be run as they hold artifacts to report test results.

## Future Enhancements

1) CI/CD Enhancements
    1) The result artifacts `````<test-name>.html````` files can be uploaded to a test reporting portal.
   2) The Deploy phase will be able to deploy to a target environment.
   3) HELM Package and HELM Deploy/build stages will be introduced once the application is containerized
2) Project Build
   1) At the Project root, Dockerfile can be built to containerize this application.
   2) At the Project root, HELM directory with Chart.yaml, values.yaml and templates directory to manage kubernetes manifests.
   3) At the Project root, CHANGELOG.md can be introduced for Versioning.
   
## Results

Please refer the `images` directory for result screenshots.

1) File ``branch-commit-manual-triggers.png`` highlights a single pipeline during a push event with manual trigger to initiate the tests.
2) File ``merge-request-commit-event.png`` highlights a single pipeline during a merge event with tests being triggered automatically.
3) File ``two-pipeliens.png highlights both`` pipelines to differentiate.
4) File ``invalid-test-pipeline-fail.png`` highlights the test failing and failing in pipeline.
