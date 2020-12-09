// Copyright 2020 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package charmhub

import (
	"context"
	"net/url"

	apicharmhub "github.com/juju/juju/api/charmhub"
	"github.com/juju/juju/charmhub"
	"github.com/juju/juju/charmhub/transport"
)

// Printer defines an interface for printing out values.
type Printer interface {
	Print() error
}

// Log describes a log format function to output to.
type Log = func(format string, params ...interface{})

// InfoCommandAPI describes API methods required to execute the info command.
type InfoCommandAPI interface {
	Info(string, string) (apicharmhub.InfoResponse, error)
	Close() error
}

// FindCommandAPI describes API methods required to execute the find command.
type FindCommandAPI interface {
	Find(string) ([]apicharmhub.FindResponse, error)
	Close() error
}

// DownloadCommandAPI describes API methods required to execute the download
// command.
type DownloadCommandAPI interface {
	Info(context.Context, string, ...charmhub.InfoOption) (transport.InfoResponse, error)
	Download(context.Context, *url.URL, string, ...charmhub.DownloadOption) error
}

type ModelConfigGetter interface {
	ModelGet() (map[string]interface{}, error)
}
