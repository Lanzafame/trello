package main

import "github.com/jroimartin/trello"

func GetLabelsCmd(boardID string, labels []string) ([]trello.Label, error) {
	tc := trello.NewClient(cfg.Key, cfg.Token)
	labels, err := tc.Labels(boardID)
	if err != nil {
		return nil, err
	}
	logf("returned labels: %+v", labels)

	return labels, nil
}
