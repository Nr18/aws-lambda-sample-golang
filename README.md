# AWS Lambda - Sample Golang

This is my sample repository for deploying Golang AWS Lambda functions.

## Deploy the template

> Note: AWS SAM will create a bucket in your account used for uploading the artifacts.

```shell
sam build
sam deploy \
    --stack-name aws-lambda-sample-golang \
    --capabilities CAPABILITY_IAM \
    --resolve-s3
```
