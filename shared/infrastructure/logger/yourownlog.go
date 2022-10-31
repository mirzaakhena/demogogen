package logger

import (
	"context"
	"demogogen1/shared/gogen"
	"fmt"
)

func NewYourLog(appData gogen.ApplicationData) Logger {
	return &yourLog{appData: appData}
}

type yourLog struct {
	appData gogen.ApplicationData
}

func (y yourLog) Info(ctx context.Context, message string, args ...any) {
	fmt.Printf("INFO >>>> %v\n", fmt.Sprintf(message, args))
}

func (y yourLog) Error(ctx context.Context, message string, args ...any) {
	fmt.Printf("ERROR >>>> %v\n", fmt.Sprintf(message, args))
}
