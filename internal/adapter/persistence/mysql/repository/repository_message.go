package repository

import (
	"toktok-backend/internal/adapter/persistence/mysql"
)

type MessageRepository struct {
	client *mysql.Client
}

func NewMessageRepository(client *mysql.Client) *MessageRepository {
	messageRepository := MessageRepository{
		client: client,
	}

	return &messageRepository
}

// func (r *MessageRepository) CreateMessage(ctx context.Context, message *domain.Message) (*domain.Message, error) {
// 	createdMessage, err := r.client.Message.Create().
// 		SetAvatarID(message.AvatarID).
// 		SetRelationID(message.RelationID).
// 		SetState(entmessage.State(message.State)).
// 		SetContent(message.Content).
// 		SetEnteredAt(message.EnteredAt).
// 		Save(ctx)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return utils.ToDomainMessage(createdMessage), nil
// }

// func (r *MessageRepository) GetMessage(ctx context.Context, id int) (*domain.Message, error) {
// 	queriedMessage, err := r.client.Message.Query().
// 		Where(entmessage.IDEQ(id)).
// 		WithAvatar().
// 		WithRelation().
// 		Only(ctx)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return utils.ToDomainMessage(queriedMessage), nil
// }

// func (r *MessageRepository) ListMessaeeByRealtionID(ctx context.Context, relationId, skip, limit int, order, critertion string) ([]*domain.Message, error) {
// 	builder := r.client.Message.Query().Where(
// 		entmessage.HasRelationWith(entrelation.ID(relationId))).
// 		Limit(limit).
// 		Offset((skip - 1) * limit)

// 	orderTermOption := sql.OrderAsc()
// 	if order == "desc" {
// 		orderTermOption = sql.OrderDesc()
// 	}

// 	switch critertion {
// 	case "entered_at":
// 		builder.Order(
// 			entmessage.ByEnteredAt(
// 				orderTermOption,
// 			),
// 		)

// 	default:
// 		builder.Order(
// 			entmessage.ByID(
// 				orderTermOption,
// 			),
// 		)
// 	}

// 	messages, err := builder.All(ctx)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return utils.ToDomainMesssages(messages), nil
// }

// // func (r *MessageRepository) UpdateMessage(ctx context.Context, message *domain.Message) (*domain.Message, error) {

// // }

// // func (r *MessageRepository) DeleteMessage(ctx context.Context, id int) error {}
