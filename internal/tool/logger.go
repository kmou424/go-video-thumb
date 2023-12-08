package tool

import (
	"fmt"
	"github.com/gookit/goutil/strutil"
	"github.com/gookit/slog"
	"github.com/gookit/slog/handler"
	"github.com/kmou424/go-video-thumb/internal/global"
	"github.com/valyala/bytebufferpool"
	"runtime"
	"strings"
)

var loggerFormat = []string{
	"[", "{{level}}", "] ",
	"[", "{{datetime}}", "] ",
	"[", "{{caller}}", "] ",
	"{{message}}", "\n",
}

type mLoggerFormatter struct {
	slog.TextFormatter
}

var Logger *slog.Logger

func NewLoggerFormatter() *mLoggerFormatter {
	return &mLoggerFormatter{
		TextFormatter: slog.TextFormatter{
			TimeFormat: "2006/01/02 - 15:04:05",
			// EnableColor: true,
			ColorTheme: slog.ColorTheme,
			// FullDisplay: false,
			EncodeFunc: slog.EncodeToString,
		},
	}
}

// from gookit/slog/formatter_test.go
var textPool bytebufferpool.Pool

func (l *mLoggerFormatter) renderColorByLevel(text string, level slog.Level) string {
	if theme, ok := l.ColorTheme[level]; ok {
		return theme.Render(text)
	}
	return text
}

//goland:noinspection GoUnhandledErrorResult
func (l *mLoggerFormatter) Format(r *slog.Record) ([]byte, error) {
	buf := textPool.Get()
	defer textPool.Put(buf)

	for _, field := range loggerFormat {
		if !(strutil.IsStartOf(field, "{{") && strutil.IsEndOf(field, "}}")) {
			buf.WriteString(field)
			continue
		}

		switch {
		case field == "{{datetime}}":
			buf.B = r.Time.AppendFormat(buf.B, l.TimeFormat)
		case field == "{{level}}":
			buf.WriteString(l.renderColorByLevel(r.LevelName(), r.Level))
		case field == "{{message}}":
			buf.WriteString(l.renderColorByLevel(r.Message, r.Level))
		case field == "{{caller}}":
			if !global.Debug {
				s := buf.String()
				buf.Reset()
				s = s[:len(s)-3]
				buf.WriteString(s)
				break
			}
			pc, file, line, ok := runtime.Caller(6)
			if !ok {
				break
			}
			funcName := runtime.FuncForPC(pc).Name()
			funcName = funcName[strings.LastIndex(funcName, "/")+1:]
			filePath := strings.Split(file, "/")
			fileName := filePath[len(filePath)-1]

			buf.WriteString(fmt.Sprintf("%s:%d %s", fileName, line, funcName))
		}
	}

	return buf.B, nil
}

func init() {
	loggerHandler := handler.NewConsoleHandler(slog.AllLevels)
	loggerHandler.SetFormatter(NewLoggerFormatter())
	Logger = slog.NewWithHandlers(loggerHandler)
}
