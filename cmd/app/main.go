package main

import (
	"dynamo-golang/config"
	"dynamo-golang/internal/repository/adapter"
	"dynamo-golang/internal/repository/instance"

	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func main() {
	configs := config.GetConfig()
	connection := instance.GetConnection()
	repository := adapter.NewAdapter(connection)

}

func Migrate() []error {
	var errors []error
}

func callMigrateAndAppendError(errors *[]error, connection *dynamodb.DynamoDB, rule rules.Interface)

func checkTables() {

}
