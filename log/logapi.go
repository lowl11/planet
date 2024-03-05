package log

func Debug(args ...any) {
	get().Debug(args...)
}

func Debugf(template string, args ...any) {
	get().Debugf(template, args...)
}

func Info(args ...any) {
	get().Info(args...)
}

func Infof(template string, args ...any) {
	get().Infof(template, args...)
}

func Warn(args ...any) {
	get().Warn(args...)
}

func Warnf(template string, args ...any) {
	get().Warnf(template, args...)
}

func Error(args ...any) {
	get().Error(args...)
}

func Errorf(template string, args ...any) {
	get().Errorf(template, args...)
}

func Fatal(args ...any) {
	get().Fatal(args...)
}

func Fatalf(template string, args ...any) {
	get().Fatalf(template, args...)
}
