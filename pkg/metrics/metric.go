package metrics

import (
	"fmt"
	"sync"
	"time"

	"github.com/rs/xid"
)

type MetricOutput string

const (
	CONSOLE_OUTPUT MetricOutput = "CONSOLE"
)

var metricHandlers = map[MetricOutput]func() (IMetric, error){
	CONSOLE_OUTPUT: newConsoleOutputMetric,
}

type IMetric interface {
	Store(snapshot *Snapshot) error
	GetOutputName() string
}

type Metric struct {
	handler IMetric
	buffer  chan *Snapshot
	wg      sync.WaitGroup
}

func NewMetricHandler(metricOutput string, bufferSize int) (*Metric, error) {

	callback, found := metricHandlers[MetricOutput(metricOutput)]

	if !found {
		return nil, fmt.Errorf("no output could be defined for the metrics")
	}

	handler, err := callback()

	if err != nil {
		return nil, err
	}

	h := &Metric{
		handler: handler,
		buffer:  make(chan *Snapshot, bufferSize),
	}

	h.wg.Add(1)

	go h.process()

	return h, nil
}

func (m *Metric) process() {

	defer m.wg.Done()

	for mt := range m.buffer {
		m.handler.Store(mt)
	}

}

func (m *Metric) Start(correlationID, component string, metadata Metadata, parent *Snapshot) *Snapshot {

	snapshot := &Snapshot{
		ID:            xid.New().String(),
		Start:         time.Now(),
		Metadata:      metadata,
		CorrelationID: correlationID,
		Component:     component,
		Parent:        parent,
	}

	return snapshot
}

func (m *Metric) End(snapshot *Snapshot) {
	end := time.Now()
	snapshot.End = end
	snapshot.Duration = end.Sub(snapshot.Start)
}

func (m *Metric) Store(snapshot *Snapshot) {
	m.buffer <- snapshot
}

func (m *Metric) ShuttingDown(snapshot *Snapshot) {
	close(m.buffer)
	m.wg.Wait()
}

func (m *Metric) GetOutputName() string {
	return m.handler.GetOutputName()
}
