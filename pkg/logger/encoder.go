package logger

import (
	"fmt"
	"time"

	"go.uber.org/zap/buffer"
	"go.uber.org/zap/zapcore"
)

type appEncoder struct {
	zapcore.Encoder
}

func (c appEncoder) EncodeEntry(entry zapcore.Entry, fields []zapcore.Field) (*buffer.Buffer, error) {
	b, err := c.Encoder.EncodeEntry(entry, fields)
	if err != nil {
		return nil, err
	}

	b.Reset()
	b.AppendString(fmt.Sprintf("[%s] - %s - (%s)",
		entry.Level.CapitalString(), entry.Message, entry.Time.Format(time.RFC3339)))
	for _, field := range fields {
		field.AddTo(c.Encoder)
	}
	b.AppendString("\n")
	return b, nil
}
