package testdata

import "log/slog"

func main() {
	slog.Info("server started")         //корректно
	slog.Debug("api request completed") //корректно
	slog.Info("user authenticated")     //корректно
}
