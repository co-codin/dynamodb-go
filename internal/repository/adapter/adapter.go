package adapter

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
)

type Database struct {
	connection *dynamodb.DynamoDB
	logMode    bool
}

// CreateOrUpdate implements [Interface].
func (db *Database) CreateOrUpdate(entity interface{}, tableName string) (response *dynamodb.PutItemOutput, err error) {
	entityParsed, err := dynamodbattribute.MarshalMap(entity)
	if err != nil {
		return nil, err
	}
	return db.connection.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item:      entityParsed,
	})
}

// Delete implements [Interface].
func (db *Database) Delete(condition map[string]interface{}, tableName string) (response *dynamodb.DeleteItemOutput, err error) {
	conditionParsed, err := dynamodbattribute.MarshalMap(condition)
	if err != nil {
		return nil, err
	}
	return db.connection.DeleteItem(&dynamodb.DeleteItemInput{
		TableName: aws.String(tableName),
		Key:       conditionParsed,
	})
}

// FindAll implements [Interface].
func (db *Database) FindAll(condition expression.Expression, tableName string) (response *dynamodb.ScanOutput, err error) {
	input := &dynamodb.ScanInput{
		ExpressionAttributeNames:  condition.Names(),
		ExpressionAttributeValues: condition.Values(),
		FilterExpression:          condition.Filter(),
		ProjectionExpression:      condition.Projection(),
		TableName:                 aws.String(tableName),
	}
	return db.connection.Scan(input)
}

// FindOne implements [Interface].
func (db *Database) FindOne(condition map[string]interface{}, tableName string) (response *dynamodb.GetItemOutput, err error) {
	conditionParsed, err := dynamodbattribute.MarshalMap(condition)
	if err != nil {
		return nil, err
	}
	return db.connection.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key:       conditionParsed,
	})
}

// Health implements [Interface].
func (db *Database) Health() bool {
	_, err := db.connection.ListTables(&dynamodb.ListTablesInput{})
	return err == nil
}

type Interface interface {
	Health() bool
	FindAll(condition expression.Expression, tableName string) (response *dynamodb.ScanOutput, err error)
	FindOne(condition map[string]interface{}, tableName string) (response *dynamodb.GetItemOutput, err error)
	CreateOrUpdate(entity interface{}, tableName string) (response *dynamodb.PutItemOutput, err error)
	Delete(condition map[string]interface{}, tableName string) (response *dynamodb.DeleteItemOutput, err error)
}

func NewAdapter(con *dynamodb.DynamoDB) Interface {
	return &Database{
		connection: con,
		logMode:    false,
	}
}

