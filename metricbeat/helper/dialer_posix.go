//+build !windows

package helper

import (
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/pkg/errors"

	"github.com/elastic/beats/libbeat/api/npipe"
	"github.com/elastic/beats/libbeat/outputs/transport"
)

func makeDialer(t time.Duration, uri string) (transport.Dialer, string, error) {
	if npipe.IsNPipe(uri) {
		return nil, "", fmt.Errorf(
			"cannot use %s as the URI, named pipes are only supported on Windows",
			uri,
		)
	}

	if strings.HasPrefix(uri, "http+unix://") || strings.HasPrefix(uri, "unix://") {
		s := strings.TrimPrefix(uri, "http+unix://")
		s = strings.TrimPrefix(s, "unix://")

		parts := strings.SplitN(s, "/", 2)

		sockFile, err := url.PathUnescape(parts[0])
		if err != nil {
			return nil, "", errors.Wrap(err, "could no decode path to the socket")
		}

		if len(parts) == 1 {
			return transport.UnixDialer(t, sockFile), "http://unix/", nil
		}

		return transport.UnixDialer(t, sockFile), "http://unix/" + parts[1], nil
	}

	return transport.NetDialer(t), uri, nil
}
