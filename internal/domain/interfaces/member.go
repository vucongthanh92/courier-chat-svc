package interfaces

import (
	"context"

	"github.com/vucongthanh92/courier-chat-svc/internal/domain/entities"
	"github.com/vucongthanh92/courier-chat-svc/internal/domain/models"
	errHandler "github.com/vucongthanh92/go-base-utils/helper/http_error"
)

type MemberServiceI interface {
	ListConversationMembers(ctx context.Context, req *models.ListConversationMembersRequest) (*models.ListConversationMembersResponse, *errHandler.ErrorBuilder)
}

type MemberCmdRepoI interface {
	CreateMembers(ctx context.Context, entities []entities.ConversationMember) *errHandler.ErrorBuilder
	UpdateReadState(ctx context.Context, conversationID, userID, lastReadMessageID uint64) (*entities.ConversationMember, *errHandler.ErrorBuilder)
}

type MemberQueryRepoI interface {
	ListConversationMembers(ctx context.Context, conversationID uint64) ([]entities.ConversationMember, *errHandler.ErrorBuilder)
	GetConversationMember(ctx context.Context, conversationID, userID uint64) (*entities.ConversationMember, *errHandler.ErrorBuilder)
}
