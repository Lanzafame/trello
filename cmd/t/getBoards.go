package main

import "github.com/jroimartin/trello"

func GetBoards() ([]trello.Board, error) {
	tc := trello.NewClient(cfg.Key, cfg.Token)
	boards, err := tc.Boards("me")
	if err != nil {
		return nil, err
	}
	logf("returned boards: %+v", boards)

	return boards, nil
}

func GetBoard(board string) ([]trello.Board, error) {
	tc := trello.NewClient(cfg.Key, cfg.Token)
	boards, err := tc.Boards("me")
	if err != nil {
		return nil, err
	}
	logf("returned boards: %+v", boards)

	return boards, nil
}
