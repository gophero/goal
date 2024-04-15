package logx

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

var Default = NewLog(&Zap{Director: "logs"})

const LoggerKey = "zapLogger"

type Logger struct {
	*zap.SugaredLogger
}

// NewLog 获取 zap.Logger
func NewLog(c *Zap) *Logger {
	zapObj = zapDef{c: c}

	// 如果日志文件夹没有，则创建
	if ok, _ := PathExists(c.Director); !ok {
		fmt.Printf("create %v directory\n", c.Director)
		_ = os.Mkdir(c.Director, os.ModePerm)
	}

	cores := zapObj.GetZapCores()
	logger := zap.New(zapcore.NewTee(cores...))

	if c.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	return &Logger{SugaredLogger: logger.Sugar()}
}

var zapObj zapDef

type zapDef struct {
	c *Zap
}

// GetEncoder 获取 zapcore.Encoder
func (z *zapDef) GetEncoder() zapcore.Encoder {
	switch z.c.Format {
	case ZapFormatJson:
		return zapcore.NewJSONEncoder(z.GetEncoderConfig())
	default:
		return zapcore.NewConsoleEncoder(z.GetEncoderConfig())
	}
}

// GetEncoderConfig 获取zapcore.EncoderConfig
func (z *zapDef) GetEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		MessageKey:     "msg",
		LevelKey:       "level",
		TimeKey:        "ts",
		NameKey:        "log",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		StacktraceKey:  z.c.StacktraceKey,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    z.c.ZapEncodeLevel(),
		EncodeTime:     z.CustomTimeEncoder, // 日志时间
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}

// GetEncoderCore 获取Encoder的 zapcore.Core
func (z *zapDef) GetEncoderCore(l zapcore.Level, level zap.LevelEnablerFunc) zapcore.Core {
	writer, err := FileRotatelogs.GetWriteSyncer(z.c, l.String()) // 使用file-rotatelogs进行日志分割
	if err != nil {
		fmt.Printf("Get Write Syncer Failed err:%v", err.Error())
		return nil
	}
	return zapcore.NewCore(z.GetEncoder(), writer, level)
}

// CustomTimeEncoder 自定义日志输出时间格式
func (z *zapDef) CustomTimeEncoder(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
	var prefix string
	if z.c.Prefix != "" {
		prefix = "[" + z.c.Prefix + "] "
	}
	encoder.AppendString(prefix + t.Format("2006/01/02 15:04:05.000"))
}

// GetZapCores 根据配置文件的Level获取 []zapcore.Core
func (z *zapDef) GetZapCores() []zapcore.Core {
	cores := make([]zapcore.Core, 0, 7)
	for level := z.c.TransportLevel(); level <= zapcore.FatalLevel; level++ {
		cores = append(cores, z.GetEncoderCore(level, z.GetLevelPriority(level)))
	}
	return cores
}

// GetLevelPriority 根据 zapcore.Level 获取 zap.LevelEnablerFunc
func (z *zapDef) GetLevelPriority(level zapcore.Level) zap.LevelEnablerFunc {
	switch level {
	case zapcore.DebugLevel:
		return func(level zapcore.Level) bool { // 调试级别
			return level == zap.DebugLevel
		}
	case zapcore.InfoLevel:
		return func(level zapcore.Level) bool { // 日志级别
			return level == zap.InfoLevel
		}
	case zapcore.WarnLevel:
		return func(level zapcore.Level) bool { // 警告级别
			return level == zap.WarnLevel
		}
	case zapcore.ErrorLevel:
		return func(level zapcore.Level) bool { // 错误级别
			return level == zap.ErrorLevel
		}
	case zapcore.DPanicLevel:
		return func(level zapcore.Level) bool { // dpanic级别
			return level == zap.DPanicLevel
		}
	case zapcore.PanicLevel:
		return func(level zapcore.Level) bool { // panic级别
			return level == zap.PanicLevel
		}
	case zapcore.FatalLevel:
		return func(level zapcore.Level) bool { // 终止级别
			return level == zap.FatalLevel
		}
	default:
		return func(level zapcore.Level) bool { // 调试级别
			return level == zap.DebugLevel
		}
	}
}

// NewContext 给指定的context添加字段
func (l *Logger) NewContext(ctx *gin.Context, fields ...any) {
	ctx.Set(LoggerKey, l.WithContext(ctx).With(fields...))
}

// WithContext 从指定的context返回一个zap实例
func (l *Logger) WithContext(ctx *gin.Context) *Logger {
	if ctx == nil {
		return l
	}
	zl, _ := ctx.Get(LoggerKey)
	ctxLogger, ok := zl.(*zap.SugaredLogger)
	if ok {
		return &Logger{ctxLogger}
	}
	return l
}
