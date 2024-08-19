package main

import (
	"UserManagement/commons"
	"context"
)

func main() {
	_, logger := commons.NewLoggerWithCorrelationid(context.Background(), "")
	logger.Info("This is Sample Project")
}
