package logger

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/middleware"
	"io"
	"net"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"os"
	"time"

	"github.com/akleinloog/http-logger/config"

	"github.com/rs/zerolog"
)

var (
	// LogEntryCtxKey is the context.Context key to store the request log entry.
	//LogEntryCtxKey = &contextKey{"LogEntry"}

	// DefaultLogger is called by the Logger middleware handler to log each request.
	// Its made a package-level variable so that it can be reconfigured for custom
	// logging configurations.
	//DefaultLogger = New(config.DefaultConfig)

	DefaultLogger = New(config.DefaultConfig)
)

// Logger is used for logging.
type Logger struct {
	logger *zerolog.Logger
}

type RequestLogEntry struct {
	Method       string
	URL          string
	UserAgent    string
	Referer      string
	Protocol     string
	RemoteIP     string
	ServerIP     string
	Status       int
	Latency      time.Duration
	RequestBody  string
	ResponseBody string
	RequestId    string
}

// New initializes a new logger
func New(config *config.Config) *Logger {

	logLevel := zerolog.InfoLevel

	if config.Debug {
		logLevel = zerolog.DebugLevel
	}

	zerolog.SetGlobalLevel(logLevel)

	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()

	return &Logger{logger: &logger}
}

// initRequestEntry initializes a new log entry for
func initRequestEntry(request *http.Request) *RequestLogEntry {

	requestId := ""

	if request, ok := request.Context().Value(middleware.RequestIDKey).(string); ok {
		requestId = request
	}

	entry := &RequestLogEntry{
		RequestId: requestId,
		Method:    request.Method,
		URL:       request.URL.String(),
		UserAgent: request.UserAgent(),
		Referer:   request.Referer(),
		Protocol:  request.Proto,
		RemoteIP:  ipFromHostPort(request.RemoteAddr),
	}

	if localAddress, ok := request.Context().Value(http.LocalAddrContextKey).(net.Addr); ok {
		entry.ServerIP = ipFromHostPort(localAddress.String())
	}

	return entry
}

// RequestLogger is a middleware that logs the start and end of each request, along
// with some useful data about what was requested, what the response status was,
// and how long it took to return. When standard output is a TTY, Logger will
// print in color, otherwise it will print in black and white. Logger prints a
// request ID if one is provided.
//
// Alternatively, look at https://github.com/goware/httplog for a more in-depth
// http logger with structured logging support.
func RequestLogger(next http.Handler) http.Handler {

	fn := func(w http.ResponseWriter, r *http.Request) {

		entry := initRequestEntry(r)

		start := time.Now()

		request, err := httputil.DumpRequest(r, true)
		if err != nil {
			http.Error(w, fmt.Sprint(err), http.StatusInternalServerError)
			return
		}
		entry.RequestBody = fmt.Sprintf("%q", request)

		rec := httptest.NewRecorder()

		defer func() {

			entry.Latency = time.Since(start)

			entry.Status = rec.Code

			if entry.Status == 0 {
				entry.Status = http.StatusOK
			}

			entry.ResponseBody = fmt.Sprintf("%q", rec.Body)

			// this copies the recorded response to the response writer
			for k, v := range rec.HeaderMap {
				w.Header()[k] = v
			}
			w.WriteHeader(rec.Code)
			rec.Body.WriteTo(w)

			DefaultLogger.Info().
				Str("method", entry.Method).
				Str("url", entry.URL).
				Str("agent", entry.UserAgent).
				Str("referer", entry.Referer).
				Str("protocol", entry.Protocol).
				Str("remoteIp", entry.RemoteIP).
				Str("serverIp", entry.ServerIP).
				Int("status", entry.Status).
				Dur("latency", entry.Latency).
				Str("request", entry.RequestBody).
				Str("response", entry.ResponseBody).
				Str("requestId", entry.RequestId).
				Msg("")
		}()

		next.ServeHTTP(rec, WithLogEntry(r, entry))
	}

	return http.HandlerFunc(fn)
}

func WithLogEntry(r *http.Request, entry *RequestLogEntry) *http.Request {
	r = r.WithContext(context.WithValue(r.Context(), "requestLog", entry))
	return r
}

func ipFromHostPort(hp string) string {
	h, _, err := net.SplitHostPort(hp)
	if err != nil {
		return ""
	}
	if len(h) > 0 && h[0] == '[' {
		return h[1 : len(h)-1]
	}
	return h
}

// Output duplicates the global logger and sets w as its output.
func (l *Logger) Output(w io.Writer) zerolog.Logger {
	return l.logger.Output(w)
}

// With creates a child logger with the field added to its context.
func (l *Logger) With() zerolog.Context {
	return l.logger.With()
}

// Level creates a child logger with the minimum accepted level set to level.
func (l *Logger) Level(level zerolog.Level) zerolog.Logger {
	return l.logger.Level(level)
}

// Sample returns a logger with the s sampler.
func (l *Logger) Sample(s zerolog.Sampler) zerolog.Logger {
	return l.logger.Sample(s)
}

// Hook returns a logger with the h Hook.
func (l *Logger) Hook(h zerolog.Hook) zerolog.Logger {
	return l.logger.Hook(h)
}

// Debug starts a new message with debug level.
//
// You must call Msg on the returned event in order to send the event.
func (l *Logger) Debug() *zerolog.Event {
	return l.logger.Debug()
}

// Info starts a new message with info level.
//
// You must call Msg on the returned event in order to send the event.
func (l *Logger) Info() *zerolog.Event {
	return l.logger.Info()
}

// Warn starts a new message with warn level.
//
// You must call Msg on the returned event in order to send the event.
func (l *Logger) Warn() *zerolog.Event {
	return l.logger.Warn()
}

// Error starts a new message with error level.
//
// You must call Msg on the returned event in order to send the event.
func (l *Logger) Error() *zerolog.Event {
	return l.logger.Error()
}

// Fatal starts a new message with fatal level. The os.Exit(1) function
// is called by the Msg method.
//
// You must call Msg on the returned event in order to send the event.
func (l *Logger) Fatal() *zerolog.Event {
	return l.logger.Fatal()
}

// Panic starts a new message with panic level. The message is also sent
// to the panic function.
//
// You must call Msg on the returned event in order to send the event.
func (l *Logger) Panic() *zerolog.Event {
	return l.logger.Panic()
}

// WithLevel starts a new message with level.
//
// You must call Msg on the returned event in order to send the event.
func (l *Logger) WithLevel(level zerolog.Level) *zerolog.Event {
	return l.logger.WithLevel(level)
}

// Log starts a new message with no level. Setting zerolog.GlobalLevel to
// zerolog.Disabled will still disable events produced by this method.
//
// You must call Msg on the returned event in order to send the event.
func (l *Logger) Log() *zerolog.Event {
	return l.logger.Log()
}

// Print sends a log event using debug level and no extra field.
// Arguments are handled in the manner of fmt.Print.
func (l *Logger) Print(v ...interface{}) {
	l.logger.Print(v...)
}

// Printf sends a log event using debug level and no extra field.
// Arguments are handled in the manner of fmt.Printf.
func (l *Logger) Printf(format string, v ...interface{}) {
	l.logger.Printf(format, v...)
}

// Ctx returns the Logger associated with the ctx. If no logger
// is associated, a disabled logger is returned.
func (l *Logger) Ctx(ctx context.Context) *Logger {
	return &Logger{logger: zerolog.Ctx(ctx)}
}
