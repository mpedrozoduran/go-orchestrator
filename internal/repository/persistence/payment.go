package persistence

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	ddbtypes "github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/mpedrozoduran/go-orchestrator/internal/repository/persistence/entities"
	"log"
)

type PaymentRepository[T entities.Payment | entities.Refund | entities.AuditTrail] struct {
	DbClient
	TableName string
}

func NewPaymentRepository[T entities.Payment | entities.Refund | entities.AuditTrail](dbClient DbClient, tableName string) PaymentRepository[T] {
	return PaymentRepository[T]{
		dbClient,
		tableName,
	}
}

func (p PaymentRepository[T]) Store(record T) error {

	av, err := attributevalue.MarshalMap(&record)
	if err != nil {
		return fmt.Errorf("got error marshalling item: %s", err)
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: &p.TableName,
	}

	_, err = p.svc.PutItem(context.TODO(), input)
	if err != nil {
		return fmt.Errorf("error storing payment: %v", err)
	}

	return nil
}

func (p PaymentRepository[T]) Get(id string) (*T, error) {
	result, err := p.svc.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: &p.TableName,
		Key: map[string]ddbtypes.AttributeValue{
			"KeyId": &ddbtypes.AttributeValueMemberS{Value: id},
		},
	})

	if err != nil {
		return nil, err
	}

	if result.Item == nil {
		return nil, nil
	}

	if err != nil {
		return nil, err
	}

	var item T

	err = attributevalue.UnmarshalMap(result.Item, &item)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal record, %v", err)
	}
	return &item, nil
}

func (p PaymentRepository[T]) GetAll() ([]T, error) {
	params := &dynamodb.ScanInput{
		TableName: &p.TableName,
	}
	result, err := p.svc.Scan(context.TODO(), params)
	if err != nil {
		return nil, err
	}
	var items []T
	for _, record := range result.Items {
		var item T

		err = attributevalue.UnmarshalMap(record, &item)

		if err != nil {
			log.Printf("error when unmarshalling: %s", err)
			return nil, err
		}
		items = append(items, item)
	}

	return items, nil
}
