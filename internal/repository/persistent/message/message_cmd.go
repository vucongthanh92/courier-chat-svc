package message

import (
	"context"

	"github.com/vucongthanh92/courier-chat-svc/database"
	"github.com/vucongthanh92/courier-chat-svc/internal/domain/entities"
	"github.com/vucongthanh92/courier-chat-svc/internal/domain/interfaces"
	errHandler "github.com/vucongthanh92/go-base-utils/helper/http_error"
	"github.com/vucongthanh92/go-base-utils/helper/transaction"
	"gorm.io/gorm"
)

type repoCmdMessage struct {
	writeDB *gorm.DB
}

func InitMessageCmdRepo(readDb *database.GormReadDb, writeDb *database.GormWriteDb) interfaces.MessageCmdRepoI {
	return &repoCmdMessage{
		writeDB: *writeDb,
	}
}

func (r *repoCmdMessage) CreateMessage(ctx context.Context, entity *entities.Message) *errHandler.ErrorBuilder {
	run := transaction.RunnerFromCtx(ctx, r.writeDB)
	if err := run.Model(&entities.Message{}).Create(entity).Error; err != nil {
		return errHandler.InitErrorBuilder(ctx).ValidateError(err)
	}
	return nil
}
