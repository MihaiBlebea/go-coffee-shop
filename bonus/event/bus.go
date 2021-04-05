package event

import (
	"fmt"
	"os"

	"github.com/MihaiBlebea/coffee-shop/bonus/stamp"
	goes "github.com/jetbasrawi/go.geteventstore"
)

type StampService interface {
	NewStamp(userID string) (*stamp.Stamp, error)
	AssignWelcomeStamps(userID string) error
}

type Service struct {
	client       *goes.Client
	streamName   string
	stampService StampService
}

func New(stampService StampService) (*Service, error) {
	client, err := goes.NewClient(nil, os.Getenv("EVENTSTORE_URL"))
	if err != nil {
		return &Service{}, err
	}

	return &Service{
		client:       client,
		streamName:   os.Getenv("EVENTSTORE_DATA_STREAM"),
		stampService: stampService,
	}, nil
}

// func (s *Service) Publish(name string, data interface{}) error {
// 	meta := make(map[string]string)
// 	meta["Name"] = name

// 	ev := goes.NewEvent(goes.NewUUID(), name, &data, &meta)

// 	writer := s.client.NewStreamWriter(s.streamName)

// 	err := writer.Append(nil, ev)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

func (s *Service) Listen() {
	reader := s.client.NewStreamReader(s.streamName)

	for reader.Next() {
		if reader.Err() != nil {
			if _, ok := reader.Err().(*goes.ErrNoMoreEvents); ok {
				reader.LongPoll(15)
			}
		} else {
			// get the data from meta first
			meta := make(map[string]string)
			if err := reader.Scan(nil, &meta); err != nil {
				fmt.Println(err)
				continue
			}

			if _, ok := meta["Name"]; ok == false {
				continue
			}

			switch meta["Name"] {
			case "account.created":
				ev := AccountCreated{}
				err := reader.Scan(&ev, nil)
				if err != nil {
					fmt.Println(err)
					break
				}

				if err := s.stampService.AssignWelcomeStamps(ev.UserID); err != nil {
					fmt.Println(err)
					break
				}
			}
		}
	}
}
