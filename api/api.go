package api

import (
	"context"
	"fmt"
)

const (
	NGESDataServer = "DataServer.DataService/GraphqlQuery"
)

type App interface {
	AppInit(interface{})
	GraphQuery(context.Context, string) (string, error)

	GetToken(ctx context.Context) string
	GetTenantID(context.Context) uint64
	GetUin(context.Context) uint64
	NgesInvoke(ctx context.Context, cmd, payload string) (string, error)
}

var apps []App

var DefaultApp App

func Register(a App) {
	apps = append(apps, a)
}

func Start(h interface{}) error {
	if len(apps) <= 0 {
		return fmt.Errorf("no app was found ! ")
	}
	DefaultApp = apps[len(apps)-1]
	DefaultApp.AppInit(h)
	return nil
}

/*
type Handler interface {
	Handle(context.Context, string) (string, error)
}
type FaasHandler func(context.Context, string) (string, error)

func (handler FaasHandler) Handle(ctx context.Context, payload string) (string, error) {
	response, err := handler(ctx, payload)
	if err != nil {
		return "", err
	}
	return response, nil
}
*/
