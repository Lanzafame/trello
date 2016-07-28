package main

import (
	"fmt"

	"github.com/jroimartin/trello"
)

func GetListsCmd(args []string) error {
	boardID := args[0]
	lists, err := GetListsOp(boardID)
	if err != nil {
		return fmt.Errorf("getlistscmd: %v", err)
	}

	fmt.Println("Lists:")
	fmt.Print(lists)

	return nil
}

func GetListsOp(boardID string) ([]trello.List, error) {
	tc := trello.NewClient(cfg.Key, cfg.Token)
	lists, err := tc.Lists(boardID)
	if err != nil {
		return nil, err
	}

	return lists, nil
}

func getListID(boardID, list string) (string, error) {
	cli := trello.NewClient(cfg.Key, cfg.Token)
	lists, err := cli.Lists(boardID)
	if err != nil {
		return "", err
	}
	logf("returned lists: %+v", lists)

	for _, l := range lists {
		if l.Name == list {
			return l.ID, nil
		}
	}

	return "", fmt.Errorf("getlistid: could not find list: %s", list)
}
