package testdata

import "log/slog"

func main() {
	password := "12345"
	apiKey := "abcdef"
	token := "token123"

	//нарушает правила
	slog.Info("Starting server")
	slog.Info("запуск сервера")
	slog.Info("server started!🚀")

	slog.Info("user password: " + password)
	slog.Debug("api_key=" + apiKey)
	slog.Info("token: " + token)
}
