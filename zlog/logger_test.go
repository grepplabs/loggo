package zlog

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

//nolint:funlen
func TestLogger(t *testing.T) {
	type test struct {
		name string
		cfg  *LogConfig
	}

	tests := []test{
		{name: "init", cfg: nil},
		{name: "debug default", cfg: &LogConfig{Level: "debug"}},
		{name: "debug text", cfg: &LogConfig{Level: LogLevelDebug, Format: LogFormatText}},
		{name: "debug json", cfg: &LogConfig{Level: LogLevelDebug, Format: LogFormatJson}},
		{name: "info text", cfg: &LogConfig{Level: LogLevelInfo, Format: LogFormatText}},
		{name: "info json", cfg: &LogConfig{Level: LogLevelInfo, Format: LogFormatJson}},
		{name: "warn text", cfg: &LogConfig{Level: LogLevelWarn, Format: LogFormatText}},
		{name: "warn json", cfg: &LogConfig{Level: LogLevelWarn, Format: LogFormatJson}},
		{name: "error text", cfg: &LogConfig{Level: LogLevelError, Format: LogFormatText}},
		{name: "error json", cfg: &LogConfig{Level: LogLevelError, Format: LogFormatJson}},
		{name: "fatal text", cfg: &LogConfig{Level: LogLevelFatal, Format: LogFormatText}},
		{name: "fatal json", cfg: &LogConfig{Level: LogLevelFatal, Format: LogFormatJson}},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if tc.cfg != nil {
				Init(*tc.cfg)
			}
			Vf(0, "hello %s", "f base")
			Vf(1, "hello %s", "f verbose")
			Printf("hello %s", "world")
			Infof("pi=%0.2f", 3.14159)
			Debugf("id=%d", 42)
			Warnf("oops: %s", "careful")
			Errorf("failed: %s", "boom")
			require.Panics(t, func() { Panicf("p %s", "f") })

			Printw("Printw", "k", "v")
			Infow("Infow", "k", "v")
			Debugw("Debugw", "k", "v")
			Warnw("Warnw", "k", "v")
			Errorw("Errorw", "k", "v")
			require.Panics(t, func() { Panicw("Panicw", "k", "v") })

			Printw("odd number Printw", "k")
			Infow("odd number Infow", "k")
			Debugw("odd number Debugw", "k")
			Warnw("odd number Warnw", "k")
			Errorw("odd number Errorw", "k")
			require.Panics(t, func() { Panicw("odd number Panicw", "k") })

			ctx := WithContext(context.Background(), Logger)

			VCf(ctx, 0, "hello %s", "f base")
			VCf(ctx, 1, "hello %s", "f verbose")
			PrintCf(ctx, "hello %s", "world")
			InfoCf(ctx, "pi=%0.2f", 3.14159)
			DebugCf(ctx, "id=%d", 42)
			WarnCf(ctx, "oops: %s", "careful")
			ErrorCf(ctx, "failed: %s", "boom")
			require.Panics(t, func() { PanicCf(ctx, "p %s", "f") })

			logger := Logger.WithValues("namespace", "default")
			ctx = WithContext(ctx, logger)

			PrintCw(ctx, "ctx-info", "k", "v")
			InfoCw(ctx, "ctx-info", "k", "v")
			DebugCw(ctx, "ctx-dbg", "k", "v")
			WarnCw(ctx, "ctx-warn", "k", "v")
			ErrorCw(ctx, "ctx-err", "k", "v")
			require.Panics(t, func() { PanicCw(ctx, "p", "k", "v") })

			PrintCw(ctx, "odd number PrintCw", "k")
			InfoCw(ctx, "odd number Infow", "k")
			DebugCw(ctx, "odd number Debugw", "k")
			WarnCw(ctx, "odd number Warnw", "k")
			ErrorCw(ctx, "odd number Errorw", "k")
			require.Panics(t, func() { PanicCw(ctx, "odd number Panicw", "k") })
		})
	}
}
