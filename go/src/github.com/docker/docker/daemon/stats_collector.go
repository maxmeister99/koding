package daemon

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/docker/docker/daemon/execdriver"
	"github.com/docker/docker/pkg/pubsub"
	"github.com/docker/libcontainer/system"
)

// newStatsCollector returns a new statsCollector that collections
// network and cgroup stats for a registered container at the specified
// interval.  The collector allows non-running containers to be added
// and will start processing stats when they are started.
func newStatsCollector(interval time.Duration) *statsCollector {
	s := &statsCollector{
		interval:   interval,
		publishers: make(map[*Container]*pubsub.Publisher),
		clockTicks: uint64(system.GetClockTicks()),
	}
	go s.run()
	return s
}

// statsCollector manages and provides container resource stats
type statsCollector struct {
	m          sync.Mutex
	interval   time.Duration
	clockTicks uint64
	publishers map[*Container]*pubsub.Publisher
}

// collect registers the container with the collector and adds it to
// the event loop for collection on the specified interval returning
// a channel for the subscriber to receive on.
func (s *statsCollector) collect(c *Container) chan interface{} {
	s.m.Lock()
	defer s.m.Unlock()
	publisher, exists := s.publishers[c]
	if !exists {
		publisher = pubsub.NewPublisher(100*time.Millisecond, 1024)
		s.publishers[c] = publisher
	}
	return publisher.Subscribe()
}

// stopCollection closes the channels for all subscribers and removes
// the container from metrics collection.
func (s *statsCollector) stopCollection(c *Container) {
	s.m.Lock()
	if publisher, exists := s.publishers[c]; exists {
		publisher.Close()
		delete(s.publishers, c)
	}
	s.m.Unlock()
}

// unsubscribe removes a specific subscriber from receiving updates for a container's stats.
func (s *statsCollector) unsubscribe(c *Container, ch chan interface{}) {
	s.m.Lock()
	publisher := s.publishers[c]
	if publisher != nil {
		publisher.Evict(ch)
		if publisher.Len() == 0 {
			delete(s.publishers, c)
		}
	}
	s.m.Unlock()
}

func (s *statsCollector) run() {
	for _ = range time.Tick(s.interval) {
		for container, publisher := range s.publishers {
			systemUsage, err := s.getSystemCpuUsage()
			if err != nil {
				logrus.Errorf("collecting system cpu usage for %s: %v", container.ID, err)
				continue
			}
			stats, err := container.Stats()
			if err != nil {
				if err != execdriver.ErrNotRunning {
					logrus.Errorf("collecting stats for %s: %v", container.ID, err)
				}
				continue
			}
			stats.SystemUsage = systemUsage
			publisher.Publish(stats)
		}
	}
}

const nanoSeconds = 1e9

// getSystemCpuUSage returns the host system's cpu usage in nanoseconds
// for the system to match the cgroup readings are returned in the same format.
func (s *statsCollector) getSystemCpuUsage() (uint64, error) {
	f, err := os.Open("/proc/stat")
	if err != nil {
		return 0, err
	}
	defer f.Close()
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		parts := strings.Fields(sc.Text())
		switch parts[0] {
		case "cpu":
			if len(parts) < 8 {
				return 0, fmt.Errorf("invalid number of cpu fields")
			}
			var sum uint64
			for _, i := range parts[1:8] {
				v, err := strconv.ParseUint(i, 10, 64)
				if err != nil {
					return 0, fmt.Errorf("Unable to convert value %s to int: %s", i, err)
				}
				sum += v
			}
			return (sum * nanoSeconds) / s.clockTicks, nil
		}
	}
	return 0, fmt.Errorf("invalid stat format")
}
