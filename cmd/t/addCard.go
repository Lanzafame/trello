package main

import (
	"errors"
	"fmt"
	"github.com/jroimartin/trello"
	"regexp"
	"strings"
)

func AddCardCmd(args []string) error {
	var title string
	var description string
	switch len(args) {
	case 1:
		title = args[0]
		return addCardOp(title, "")
	case 2:
		title, args = args[0], args[1:]
		description, args = args[0], args[1:]
		if len(args) > 0 {
			return errors.New("too many args provided")
		}
		return addCardOp(title, description)
	default:
		usage()
	}

	return fmt.Errorf("add card cmd failed to run with title: %s, and description: %s", title, description)
}

func addCardOp(title, desc string) error {
	title, attr := extractAttr(title)
	logf("adding task %v - %v %+v", title, desc, attr)

	board := attr.board
	if board == "" {
		board = cfg.DefaultBoard
	}

	boardID, err := getBoard(board)
	if err != nil {
		return err
	}
	logf("found board %v: %v", board, boardID)

	list := attr.list
	if list == "" {
		list = cfg.DefaultList
	}

	listID, err := getList(boardID, list)
	if err != nil {
		return err
	}
	logf("found list %v: %v", list, listID)

	labelIDs := ""
	if len(attr.labels) > 0 {
		labelIDs, err = getLabels(boardID, attr.labels)
		if err != nil {
			return err
		}
		logf("found labels %v: %v", attr.labels, labelIDs)
	}

	return pushCard(listID, title, desc, labelIDs)
}

func pushCard(listID, title, desc, labelIDs string) error {
	card := trello.Card{
		Name:     title,
		Desc:     desc,
		IDList:   listID,
		IDLabels: labelIDs,
	}

	client := trello.NewClient(cfg.Key, cfg.Token)
	return client.PushCard(card)
}

func extractAttr(str string) (string, taskAttr) {
	attr := taskAttr{}

	re := regexp.MustCompile(`(?: )(@|#|\^)(\w+)`)
	matches := re.FindAllStringSubmatch(str, -1)
	str = re.ReplaceAllString(str, "")

	for _, m := range matches {
		switch m[1] {
		case "@":
			attr.labels = append(attr.labels, m[2])
		case "#":
			attr.board = m[2]
		case "^":
			attr.list = m[2]
		}
	}

	return strings.TrimSpace(str), attr
}