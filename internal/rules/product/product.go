package product

import (
	"encoding/json"
	"errors"
	"io"

	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type Rules struct {
}

func NewRules() *Rules {
	return &Rules{}
}

func (r *Rules) ConvertIoReaderToStruct(data io.Reader, model interface{}) (interface{}, error) {
	if data == nil {
		return nil, errors.New("body is invalid")
	}
	return model, json.NewDecoder(data).Decode(&model)
}

func (r *Rules) GetMock() interface{} {
	// TODO: return mock product model
	return nil
}

func (r *Rules) Migrate(connection *dynamodb.DynamoDB) error {
	return r.CreateTable(connection)
}

func (r *Rules) Validate(model interface{}) error {
	// TODO: validate product model
	return nil
}

func (r *Rules) CreateTable(connection *dynamodb.DynamoDB) error {
	return nil
}