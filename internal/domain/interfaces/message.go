package interfaces

import (
	"context"

	"github.com/vucongthanh92/courier-chat-svc/internal/domain/entities"
	"github.com/vucongthanh92/courier-chat-svc/internal/domain/models"
	errHandler "github.com/vucongthanh92/go-base-utils/helper/http_error"
)

type MessageServiceI interface {
	CreateMessage(ctx context.Context, req *models.SendMessageRequest) (*models.MessageResponse, bool, *errHandler.ErrorBuilder)
	CreateSystemMessage(ctx context.Context, req *models.CreateSystemMessageRequest) (*models.MessageResponse, bool, *errHandler.ErrorBuilder)
	ListMessages(ctx context.Context, req *models.ListMessagesRequest) (*models.ListMessagesResponse, *errHandler.ErrorBuilder)
}

type AssistantEventPublisherI interface {
	PublishAssistantRequested(ctx context.Context, payload models.AssistantRequestedPayload) error
}

type ChatEventPublisherI interface {
	PublishConversationCreated(ctx context.Context, payload models.ConversationCreatedPayload) error
}

type MessageQueryRepoI interface {
	GetMessageByClientMessageID(ctx context.Context, conversationID uint64, clientMessageID string) (*entities.Message, *errHandler.ErrorBuilder)
	ListMessages(ctx context.Context, conversationID uint64, req models.ListMessagesRequest) ([]entities.Message, *errHandler.ErrorBuilder)
	GetMessageByID(ctx context.Context, id uint64) (*entities.Message, *errHandler.ErrorBuilder)
}

type MessageCmdRepoI interface {
	CreateMessage(ctx context.Context, entity *entities.Message) *errHandler.ErrorBuilder
}
