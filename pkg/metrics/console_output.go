package metrics

import "fmt"

type ConsoleOutputMetric struct {
	output MetricOutput
}

var (
	_ IMetric = (*ConsoleOutputMetric)(nil)
)

func newConsoleOutputMetric() (IMetric, error) {

	return &ConsoleOutputMetric{output: CONSOLE_OUTPUT}, nil
}

func (c *ConsoleOutputMetric) Store(snapshot *Snapshot) error {

	fmt.Printf("entry: id=%s, correlation_id=%s, component=%s, start=%v, end=%v, duration=%v, metadata=%v\n",
		snapshot.ID, snapshot.CorrelationID, snapshot.Component, snapshot.Start, snapshot.End,
		snapshot.Duration, snapshot.Metadata)

	return nil
}

func (c *ConsoleOutputMetric) GetOutputName() string {
	return string(c.output)
}
