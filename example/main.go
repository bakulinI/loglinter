package main

import "log/slog"

func main() {
	password := "12345"
	apiKey := "abcdef"
	token := "token123"

	slog.Info("Starting server")  //uppercase
	slog.Info("запуск сервера")   //кириллица
	slog.Info("server started!🚀") //спецсимвол / emoji
	slog.Info("server started")   //ок

	slog.Info("user password: " + password) //sensitive
	slog.Debug("api_key=" + apiKey)         //sensitive
	slog.Info("token: " + token)            //sensitive
	slog.Info("user authenticated")         //ок
}
