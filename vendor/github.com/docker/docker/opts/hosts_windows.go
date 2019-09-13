package opts // import "github.com/docker/docker/opts"

// DefaultHost constant defines the default host string used by docker on Windows
var DefaultHost = "npipe://" + DefaultNamedPipe
