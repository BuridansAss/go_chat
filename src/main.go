package src

import (
	chatProviders "buridansAss/chat/src/providers"
	"go.uber.org/fx"
)

func buildProviders() fx.Option {
	providers := fx.Provide(
		chatProviders.NewConfig,
		chatProviders.NewDbConnection,
		chatProviders.NewLogger,
	)

	return providers
}

func Run() {

	app := fx.New(buildProviders())

	app.Run()
}
