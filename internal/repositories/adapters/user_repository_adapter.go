package repositories_adapters

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/google/uuid"
	"github.com/vm-championships-manager/backend-service/config"
	"github.com/vm-championships-manager/backend-service/internal/dtos"
	"github.com/vm-championships-manager/backend-service/internal/entities"
	pkg_adapter "github.com/vm-championships-manager/backend-service/internal/pkg/adapters"
	pkg_protocols "github.com/vm-championships-manager/backend-service/internal/pkg/protocols"
	protocols_repositories "github.com/vm-championships-manager/backend-service/internal/repositories/protocols"
	utils_adapters "github.com/vm-championships-manager/backend-service/internal/utils/adapters"
	utils_protocols "github.com/vm-championships-manager/backend-service/internal/utils/protocols"
)

type input = protocols_repositories.UserRepositoryInput
type output = protocols_repositories.UserRepositoryOutput

type UserRepositoryAdapter struct {
	tableName         string
	db                dynamodb.Client
	log               utils_protocols.Log
	panicIfErrorOccur pkg_protocols.ErrorHandler
}

func NewUserRepository(db dynamodb.Client) protocols_repositories.UserRepository {
	log := utils_adapters.NewLoggerAdapter()

	return &UserRepositoryAdapter{
		tableName:         config.UserTableName,
		db:                db,
		log:               log,
		panicIfErrorOccur: pkg_adapter.NewPanicIfErrorOccur(log),
	}
}

func (r *UserRepositoryAdapter) Exec(parentCtx context.Context, in input) output {
	ctx, cancel := context.WithCancel(parentCtx)
	defer cancel()

	done := make(chan output, 1)
	go func() {
		var out output = output{}
		switch in.Action {
		case protocols_repositories.ActionUserCreate:
			out = r.create(ctx, in.Payload.User)
		case protocols_repositories.ActionUserFindByEmail:
			out = r.findByEmail(ctx, in.Payload.User.Email)
		default:
			r.log.Metadata(map[string]interface{}{"input": in}).Error("[UserRepository.Exec] unrecognized command")
			panic(fmt.Sprintf("[UserRepository.Exec] unrecognized command %s", in.Action))
		}

		done <- out
	}()

	select {
	case <-ctx.Done():
		r.log.Metadata(map[string]interface{}{"input": in}).Error("[UserRepository.Exec] context done")
		panic("[UserRepository.Exec] context done")
	case out := <-done:
		return out
	}
}

func (r *UserRepositoryAdapter) create(ctx context.Context, u entities.User) (out output) {
	logParams := func(msg string) (string, map[string]interface{}) {
		return fmt.Sprintf("[UserRepository.create] %s", msg), map[string]interface{}{"user": u}
	}

	uuid := r.panicIfErrorOccur.Double(uuid.NewUUID())(
		logParams("generate uuid get an error"),
	).(uuid.UUID)

	out.UserDto = &dtos.UserDto{
		Id:        uuid.String(),
		CreatedAt: time.Now().Unix(),
		UpdatedAt: time.Now().Unix(),
		Name:      u.Name,
		LastName:  u.LastName,
		Email:     u.Email,
		Birthdate: u.Birthdate,
		Phone:     u.Phone,
	}

	var jsonMap map[string]interface{}
	jsonBytes := r.panicIfErrorOccur.Double(json.Marshal(out.UserDto))(
		logParams("marshal got an error"),
	).([]byte)

	r.panicIfErrorOccur.Single(json.Unmarshal(jsonBytes, &jsonMap))(
		logParams("unmarshal json got an error"),
	)

	mm := r.panicIfErrorOccur.Double(
		attributevalue.MarshalMap(jsonMap),
	)(
		logParams("marshal map json to dynamodb format got an error"),
	).(map[string]types.AttributeValue)

	r.panicIfErrorOccur.Double(r.db.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(r.tableName),
		Item:      mm,
	}))(logParams("put item into dynamodb got an error"))

	return
}

func (r *UserRepositoryAdapter) findByEmail(ctx context.Context, email string) (out output) {
	logParams := func(msg string) (string, map[string]interface{}) {
		return fmt.Sprintf("[UserRepository.findByEmail] %s", msg),
			map[string]interface{}{"email": email}
	}

	result := r.panicIfErrorOccur.Double(r.db.Query(ctx, &dynamodb.QueryInput{
		IndexName:              aws.String("gsi-user-email"),
		TableName:              aws.String(r.tableName),
		KeyConditionExpression: aws.String("#email = :email"),
		ExpressionAttributeNames: aws.ToStringMap(map[string]*string{
			"#email": aws.String("email"),
		}),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":email": &types.AttributeValueMemberS{
				Value: email,
			},
		},
	}))(logParams("query got an error")).(*dynamodb.QueryOutput)

	if len(result.Items) == 0 {
		return
	}

	out.UserDto = &dtos.UserDto{}
	r.panicIfErrorOccur.Single(
		attributevalue.UnmarshalMap(result.Items[0], out.UserDto),
	)(logParams("unmarshal map got an error"))

	return
}
