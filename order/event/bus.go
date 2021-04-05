package event

import (
	"os"

	goes "github.com/jetbasrawi/go.geteventstore"
)

type Service struct {
	client     *goes.Client
	streamName string
}

func New() (*Service, error) {
	client, err := goes.NewClient(nil, os.Getenv("EVENTSTORE_URL"))
	if err != nil {
		return &Service{}, err
	}

	return &Service{
		client:     client,
		streamName: os.Getenv("EVENTSTORE_DATA_STREAM"),
	}, nil
}

func (s *Service) Publish(name string, data interface{}) error {
	meta := make(map[string]string)
	meta["Name"] = name

	ev := goes.NewEvent(goes.NewUUID(), name, &data, &meta)

	writer := s.client.NewStreamWriter(s.streamName)

	err := writer.Append(nil, ev)
	if err != nil {
		return err
	}

	return nil
}
