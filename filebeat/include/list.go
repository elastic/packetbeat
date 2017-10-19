/*
Package include imports all prospector packages so that they register
their factories with the global registry. This package can be imported in the
main package to automatically register all of the standard supported prospectors
modules.
*/
package include

import (
	// This list is automatically generated by `make imports`
	_ "github.com/elastic/beats/filebeat/prospector/log"
	_ "github.com/elastic/beats/filebeat/prospector/redis"
	_ "github.com/elastic/beats/filebeat/prospector/stdin"
	_ "github.com/elastic/beats/filebeat/prospector/udp"
)
