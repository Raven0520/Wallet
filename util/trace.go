package util

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"math/rand"
	"net"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

// TraceBody Trace Struct
type TraceBody struct {
	TraceID     string
	SpanID      string
	Caller      string
	SrcMethod   string
	HintCode    int64
	HintContent string
}

// TContext Trace Context
type TContext struct {
	TraceBody
	CSpanID string
}

// NewTrace New Trace
func NewTrace() *TContext {
	trace := &TContext{}
	trace.TraceID = GetTraceID()
	trace.SpanID = NewSpanID()
	return trace
}

// NewSpanID SpanID
func NewSpanID() string {
	timestamp := uint32(time.Now().Unix())
	ipToLong := binary.BigEndian.Uint32(LocalIP.To4())
	b := bytes.Buffer{}
	b.WriteString(fmt.Sprintf("%08x", ipToLong^timestamp))
	b.WriteString(fmt.Sprintf("%08x", rand.Int31()))
	return b.String()
}

func calcTraceID(ip string) (traceID string) {
	now := time.Now()
	timestamp := uint32(now.Unix())
	timeNano := now.UnixNano()
	pid := os.Getpid()

	b := bytes.Buffer{}
	netIP := net.ParseIP(ip)
	if netIP == nil {
		b.WriteString("00000000")
	} else {
		b.WriteString(hex.EncodeToString(netIP.To4()))
	}
	b.WriteString(fmt.Sprintf("%08x", timestamp&0xffffffff))
	b.WriteString(fmt.Sprintf("%04x", timeNano&0xffff))
	b.WriteString(fmt.Sprintf("%04x", pid&0xffff))
	b.WriteString(fmt.Sprintf("%06x", rand.Int31n(1<<24)))
	b.WriteString("b0") // 末两位标记来源,b0为go

	return b.String()
}

// GetTraceId Calculate ID by ip
func GetTraceID() (traceID string) {
	return calcTraceID(LocalIP.String())
}

// GetGinTraceContext get trace from gin
func GetGinTraceContext(context *gin.Context) *TContext {
	if context == nil {
		return NewTrace()
	}
	traceContext, exists := context.Get("trace")
	if exists {
		if tc, ok := traceContext.(*TContext); ok {
			return tc
		}
	}
	return NewTrace()
}
