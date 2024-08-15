package main

import (
	commons "UserManagement/Commons"
	"context"
)

func main() {
	_, logger := commons.NewLoggerWithCorrelationid(context.Background(), "")
	logger.Info("This is Sample Project")
}
