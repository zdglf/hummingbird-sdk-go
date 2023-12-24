/*******************************************************************************
 * Copyright 2017.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the License
 * is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
 * or implied. See the License for the specific language governing permissions and limitations under
 * the License.
 *******************************************************************************/

package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Logger interface {
	SetLogLevel(level string)
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	Infof(template string, args ...interface{})
	Warnf(template string, args ...interface{})
	Errorf(template string, args ...interface{})
}

type LoggerConfig struct {
	FileName   string
	Prefix     string // service name
	LogLevel   string
	MaxSize    int
	MaxBackups int
	MaxAge     int
	Compress   bool
}

var LogLevelMap = map[string]zapcore.Level{
	"DEBUG": zapcore.DebugLevel,
	"INFO":  zapcore.InfoLevel,
	"WARN":  zapcore.WarnLevel,
	"ERROR": zapcore.ErrorLevel,
}

type defaultLogger struct {
	level *zap.AtomicLevel
	*zap.SugaredLogger
}

func NewLogger(fileName, logLevel, serviceName string) Logger {
	newCfg := &LoggerConfig{
		FileName:   fileName,
		LogLevel:   logLevel,
		MaxSize:    30, // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: 10, // 日志文件最多保存多少个备份
		MaxAge:     3,  // 文件最多保存多少天
		Compress:   false,
	}
	return initLogger(newCfg)
}

func initLogger(lc *LoggerConfig) Logger {
	var ll zapcore.Level
	if err := ll.UnmarshalText([]byte(lc.LogLevel)); err != nil {
		ll = zapcore.InfoLevel
	}
	var level = zap.NewAtomicLevelAt(ll)
	if lc.FileName == "" {
		cfg := zap.NewDevelopmentConfig()
		cfg.Level = level
		cfg.EncoderConfig.ConsoleSeparator = " "
		cfg.EncoderConfig.LineEnding = zapcore.DefaultLineEnding
		cfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
		cfg.EncoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
		cfg.EncoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
		cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		logger, _ := cfg.Build(zap.AddCallerSkip(1), zap.AddStacktrace(zapcore.PanicLevel))
		return &defaultLogger{
			level:         &level,
			SugaredLogger: logger.Sugar().Named(lc.Prefix),
		}
	}

	writeSyncer := getLogWriter(lc)
	encoder := getEncoder()
	core := zapcore.NewCore(encoder, writeSyncer, level.Level())
	logger := zap.New(core, zap.AddCallerSkip(1), zap.AddCaller(), zap.AddStacktrace(zapcore.PanicLevel))
	return &defaultLogger{
		level:         &level,
		SugaredLogger: logger.Sugar().Named(lc.Prefix),
	}
}

func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.LineEnding = zapcore.DefaultLineEnding
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder // zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000")
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	encoderConfig.ConsoleSeparator = " "
	return zapcore.NewConsoleEncoder(encoderConfig)
}

func getLogWriter(cfg *LoggerConfig) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   cfg.FileName,
		MaxSize:    cfg.MaxSize,
		MaxBackups: cfg.MaxBackups,
		MaxAge:     cfg.MaxAge,
		Compress:   cfg.Compress,
		LocalTime:  true,
	}
	return zapcore.AddSync(lumberJackLogger)
}

func (dl *defaultLogger) SetLogLevel(level string) {
	v, ok := LogLevelMap[level]
	if !ok {
		return
	}
	dl.level.SetLevel(v)
}

func (dl *defaultLogger) Debug(args ...interface{}) {
	dl.SugaredLogger.Debug(args...)
}

func (dl *defaultLogger) Info(args ...interface{}) {
	dl.SugaredLogger.Info(args...)
}

func (dl *defaultLogger) Warn(args ...interface{}) {
	dl.SugaredLogger.Warn(args...)
}

func (dl *defaultLogger) Error(args ...interface{}) {
	dl.SugaredLogger.Error(args...)
}

func (dl *defaultLogger) Infof(template string, args ...interface{}) {
	dl.SugaredLogger.Infof(template, args...)
}

func (dl *defaultLogger) Warnf(template string, args ...interface{}) {
	dl.SugaredLogger.Warnf(template, args...)
}

func (dl *defaultLogger) Errorf(template string, args ...interface{}) {
	dl.SugaredLogger.Errorf(template, args...)
}
