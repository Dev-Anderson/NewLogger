package newlogger

func LogInfo(cod, msg string) string {
	return `{"code": "` + cod + `" "message": "` + msg + `"}`
}
