package consumers

import (
	"context"
	"strconv"

	"truenorth/packages/logger"
	opapi "truenorth/services/operations_service/api/operations"
)

func UserCreatedController(ctx context.Context, message map[string]any) error {
	opApi := opapi.NewOperationsApi()
	userId, _ := strconv.ParseInt(message["user_id"].(string), 10, 64)
	_, err := opApi.CreateUserBalance(ctx, userId)
	if err != nil {
		logger.GetLogger().Errorf(ctx, "UserCreatedController error")
		return err
	}
	delete(message, "content")
	logger.GetLogger().Infow(ctx, "UserCreatedController message received", "body", message)
	return nil
}
