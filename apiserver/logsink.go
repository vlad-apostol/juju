// Copyright 2015 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package apiserver

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/juju/errors"
	"github.com/juju/loggo/v2"

	"github.com/juju/juju/apiserver/logsink"
	corelogger "github.com/juju/juju/core/logger"
	"github.com/juju/juju/rpc/params"
)

type agentLoggingStrategy struct {
	modelLogger corelogger.ModelLogger
	fileLogger  io.Writer

	recordLogger corelogger.Logger
	releaser     func() error
	entity       string
	modelUUID    string
}

// newAgentLogWriteCloserFunc returns a function that will create a
// logsink.LoggingStrategy given an *http.Request, that writes log
// messages to the given writer and also to the state database.
func newAgentLogWriteCloserFunc(
	ctxt httpContext,
	fileLogger io.Writer,
	modelLogger corelogger.ModelLogger,
) logsink.NewLogWriteCloserFunc {
	return func(req *http.Request) (logsink.LogWriteCloser, error) {
		strategy := &agentLoggingStrategy{
			modelLogger: modelLogger,
			fileLogger:  fileLogger,
		}
		if err := strategy.init(ctxt, req); err != nil {
			return nil, errors.Annotate(err, "initialising agent logsink session")
		}
		return strategy, nil
	}
}

func (s *agentLoggingStrategy) init(ctxt httpContext, req *http.Request) error {
	st, entity, err := ctxt.stateForRequestAuthenticatedAgent(req)
	if err != nil {
		return errors.Trace(err)
	}

	s.entity = entity.Tag().String()
	s.modelUUID = st.ModelUUID()
	m, err := st.Model()
	if err != nil {
		return errors.Trace(err)
	}
	if s.recordLogger, err = s.modelLogger.GetLogger(s.modelUUID, m.Name()); err != nil {
		return errors.Trace(err)
	}
	s.releaser = func() error {
		if removed := st.Release(); removed {
			return s.modelLogger.RemoveLogger(s.modelUUID)
		}
		return nil
	}
	return nil
}

func (s *agentLoggingStrategy) filePrefix() string {
	return s.modelUUID + ":"
}

// Close is part of the logsink.LogWriteCloser interface.
//
// Close releases the StatePool entry, closing the DB logger
// if the State is closed/removed. The file logger is owned
// by the apiserver, so it is not closed.
func (s *agentLoggingStrategy) Close() error {
	return s.releaser()
}

// WriteLog is part of the logsink.LogWriteCloser interface.
func (s *agentLoggingStrategy) WriteLog(m params.LogRecord) error {
	level, _ := loggo.ParseLevel(m.Level)
	dbErr := errors.Annotate(s.recordLogger.Log([]corelogger.LogRecord{{
		Time:      m.Time,
		Entity:    s.entity,
		Module:    m.Module,
		Location:  m.Location,
		Level:     level,
		Message:   m.Message,
		Labels:    m.Labels,
		ModelUUID: s.modelUUID,
	}}), "logging to DB failed")

	// If the log entries cannot be inserted to the DB log out an error
	// to let users know. See LP1930899.
	if dbErr != nil {
		// If this fails then the next logToFile will fail as well; no
		// need to check for errors here.
		_ = logToFile(s.fileLogger, s.filePrefix(), params.LogRecord{
			Time:    time.Now(),
			Module:  "juju.apiserver",
			Level:   loggo.ERROR.String(),
			Message: errors.Annotate(dbErr, "unable to persist log entry to the database").Error(),
		})
	}

	m.Entity = s.entity
	fileErr := errors.Annotate(
		logToFile(s.fileLogger, s.filePrefix(), m),
		"logging to logsink.log failed",
	)
	err := dbErr
	if err == nil {
		err = fileErr
	} else if fileErr != nil {
		err = errors.Errorf("%s; %s", dbErr, fileErr)
	}
	return err
}

// logToFile writes a single log message to the logsink log file.
func logToFile(writer io.Writer, prefix string, m params.LogRecord) error {
	//TODO(debug-log) - we'll move to newline delimited json
	var labelsOut []string
	for k, v := range m.Labels {
		labelsOut = append(labelsOut, fmt.Sprintf("%s:%s", k, v))
	}
	_, err := writer.Write([]byte(strings.Join([]string{
		prefix,
		m.Entity,
		m.Time.In(time.UTC).Format("2006-01-02 15:04:05"),
		m.Level,
		m.Module,
		m.Location,
		m.Message,
		strings.Join(labelsOut, ","),
	}, " ") + "\n"))
	return err
}
