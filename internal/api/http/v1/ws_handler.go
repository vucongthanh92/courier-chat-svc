package v1

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/vucongthanh92/courier-chat-svc/internal/api/ws"
	"github.com/vucongthanh92/courier-chat-svc/internal/domain/interfaces"
	"github.com/vucongthanh92/courier-chat-svc/internal/domain/models"
	"github.com/vucongthanh92/go-base-utils/helper/utils"

	errHandler "github.com/vucongthanh92/go-base-utils/helper/http_error"
	utilsMidd "github.com/vucongthanh92/go-base-utils/http/middlewares"
)

type WsHandler struct {
	hub         *ws.Hub
	upgrader    websocket.Upgrader
	tokenDeny   interfaces.TokenDenylistI
	keyResolver func(context.Context, string) (interface{}, *errHandler.ErrorBuilder)
}

func InitWsHandler(
	hub *ws.Hub,
	allowOrigins []string,
	tokenDeny interfaces.TokenDenylistI,
	keyResolver func(context.Context, string) (interface{}, *errHandler.ErrorBuilder),
) *WsHandler {
	return &WsHandler{
		hub:         hub,
		upgrader:    ws.Upgrader(allowOrigins),
		tokenDeny:   tokenDeny,
		keyResolver: keyResolver,
	}
}

// VerifyAndConnect verifies the JWT token from the query parameter, checks its validity, expiration, and denylist status.
// If valid, it upgrades the HTTP connection to a WebSocket connection and registers the user with the hub.
func (h *WsHandler) VerifyAndConnect(c *gin.Context) {
	token := c.Query("access_token")
	if token == "" {
		token = c.Query("token")
	}
	if token == "" {
		errHandler.InitErrorBuilder(c).
			SetStatus(http.StatusUnauthorized).
			SetError(models.ErrorDTO{Code: "token_missing", Message: "access_token query parameter missing"}).
			ExposeHttpError(c)
		return
	}

	claims, authErr := utilsMidd.VerifyToken(c.Request.Context(), token, h.tokenDeny, h.keyResolver)
	if authErr != nil {
		authErr.ExposeHttpError(c)
		return
	}

	userID := utils.ParseUserID(claims["sub"])
	if userID == 0 {
		errHandler.InitErrorBuilder(c).
			SetStatus(http.StatusUnauthorized).
			SetError(models.ErrorDTO{Code: "unauthorized", Message: "missing authenticated user"}).
			ExposeHttpError(c)
		return
	}

	conn, err := h.upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	h.hub.Register(userID, conn)
}
