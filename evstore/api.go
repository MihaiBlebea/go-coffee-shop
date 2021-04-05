package evstore

import (
	"os"

	goes "github.com/jetbasrawi/go.geteventstore"
)

type EventName string

const (
	UserCreatedEvent EventName = "user.created"
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

func (s *Service) Listen() {
	reader := s.client.NewStreamReader(s.streamName)

	for reader.Next() {
		if reader.Err() != nil {
			if _, ok := reader.Err().(*goes.ErrNoMoreEvents); ok {
				reader.LongPoll(15)
			}
		} else {
			meta := make(map[string]string)
			_ = reader.Scan(nil, &meta)

			if _, ok := meta["Name"]; ok == false {
				continue
			}

			switch EventName(meta["Name"]) {
			case UserCreatedEvent:
				// ev := UserCreated{}
				// err := reader.Scan(&ev, nil)
				// if err != nil {
				// 	break
				// }

				// if err := s.orderViewStore.AddUser(ev.FirstName, ev.LastName, ev.ID, int(ev.Age)); err != nil {
				// 	break
				// }
			}
		}
	}
}
