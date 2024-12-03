package impl

type LoggerSettings struct {
	AppName     string
	AppID       string
	Environment string
	Levels      map[string]struct{}
}
