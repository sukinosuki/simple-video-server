package log

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"path/filepath"
	"simple-video-server/config"
	"simple-video-server/pkg/app_ctx"
	"time"
)

var Logger *zap.Logger

var traceIdKey = "trace-id"
var uidKey = "uid"

func Info(c *gin.Context, msg string, field ...zap.Field) {
	traceId, _ := app_ctx.GetTraceId(c)
	uid, _ := app_ctx.GetUid(c)

	logger := Logger.With(
		zap.String(traceIdKey, traceId),
		zap.Uint(uidKey, uid),
	)

	logger.Info(msg, field...)
}

//func Debug(args ...interface{}) {
//	Logger.Debug(args)
//}
//
//func Warn(args ...interface{}) {
//	Logger.Warn(args)
//}

//func Error(args ...interface{}) {
//	Logger.Error(args)
//}

// func InitLogger() *zap.SugaredLogger {
func init() {
	logMode := zapcore.DebugLevel

	if !config.Env.Debug {
		logMode = zapcore.InfoLevel
	}

	//core := zapcore.NewCore(getEncoder(), getWriteSyncer(), logMode)
	//core := zapcore.NewCore(getEncoder(), zapcore.NewMultiWriteSyncer(getWriteSyncer(), zapcore.AddSync(os.Stdout)), logMode)
	core := zapcore.NewCore(getEncoder(), zapcore.NewMultiWriteSyncer(getWriteSyncer(), zapcore.AddSync(os.Stdout)), logMode)

	//Logger = zap.New(core).Sugar()
	Logger = zap.New(core)
}

func getEncoder() zapcore.Encoder {

	//return zapcore.NewConsoleEncoder(zapcore.EncoderConfig{
	//	TimeKey:       "ts",
	//	LevelKey:      "level",
	//	NameKey:       "logger",
	//	CallerKey:     "caller_line",
	//	FunctionKey:   zapcore.OmitKey,
	//	MessageKey:    "msg",
	//	StacktraceKey: "stacktrace",
	//	LineEnding:    "  ",
	//})

	config := zap.NewProductionEncoderConfig()

	config.TimeKey = "time"

	config.EncodeLevel = zapcore.CapitalLevelEncoder

	config.EncodeTime = func(time time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		//encoder.AppendString(time.Local().Format(time.DateTime))
		encoder.AppendString(time.Local().Format("2006-01-02 15:04:05.000"))
	}

	return zapcore.NewJSONEncoder(config)
}

func getWriteSyncer() zapcore.WriteSyncer {
	separator := string(filepath.Separator)

	rootDir, err := os.Getwd()

	if err != nil {
		panic(err)
	}

	//logFilePath := rootDir + separator + "log" + separator + time.Now().Format(time.DateOnly) + ".txt"
	logFilePath := rootDir + separator + "log" + separator + time.Now().Format("2006-01-02") + ".txt"
	fmt.Println("logFilePath ", logFilePath)

	syncer := &lumberjack.Logger{
		Filename:   logFilePath,
		MaxAge:     config.Log.MaxAge, // days
		Compress:   true,              // disable by default
		MaxBackups: config.Log.MaxBackups,
		MaxSize:    config.Log.MaxSize, // megabytes
	}

	return zapcore.AddSync(syncer)
}
