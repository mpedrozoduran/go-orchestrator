#!/bin/bash

# create DynamoDB tables
awslocal dynamodb create-table \
    --table-name payments \
    --key-schema AttributeName=KeyId,KeyType=HASH \
    --attribute-definitions AttributeName=KeyId,AttributeType=S \
    --billing-mode PAY_PER_REQUEST \
    --region us-east-1

awslocal dynamodb create-table \
    --table-name refunds \
    --key-schema AttributeName=KeyId,KeyType=HASH \
    --attribute-definitions AttributeName=KeyId,AttributeType=S \
    --billing-mode PAY_PER_REQUEST \
    --region us-east-1

awslocal dynamodb create-table \
    --table-name audit_trail \
    --key-schema AttributeName=KeyId,KeyType=HASH \
    --attribute-definitions AttributeName=KeyId,AttributeType=S \
    --billing-mode PAY_PER_REQUEST \
    --region us-east-1

awslocal dynamodb scan --table-name payments