package src

import (
	"buridansAss/chat/src/http"
	chatProviders "buridansAss/chat/src/providers"
	"buridansAss/chat/src/repositories"
	"go.uber.org/fx"
)

func buildProviders() fx.Option {
	providers := fx.Provide(
		chatProviders.NewConfig,
		chatProviders.NewDbConnection,
		repositories.NewUser,
		chatProviders.NewLogger,
	)

	return providers
}

func Run() {
	app := fx.New(
		buildProviders(),
		fx.Invoke(http.New),
	)

	app.Run()
}
