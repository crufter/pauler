package shared

import (
	"flag"
	cmap "github.com/streamrail/concurrent-map"
)

var (
	Dir      = flag.String("dir", "/etc/puller", "The dir to load service configs from")
	D        = flag.Bool("d", false, "Run as a daemon")
	Join     = flag.String("join", "", "Join a cluster")
	Port     = flag.Int("port", 7946, "Port is a port used for internal communication. Port + 1 is the port number of the http server")
	Node     = flag.String("node", "", "Name of the node. If left empty, defaults to os value")
	Interval = flag.Int64("interval", 30, "Time to sleep between runs of processing")
)

var (
	Services         = cmap.New() // map[string]types.Service
	ChangedServices  = cmap.New() // map[string]bool - service definition has changed.
	OutdatedServices = cmap.New() // map[string]bool - service was launched with an image that's older than the current one locally
	BadServiceFiles  = cmap.New() // map[string]bool - bad service files
)
