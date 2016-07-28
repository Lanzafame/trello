// Copyright 2016 The trello client Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package trello

import "fmt"

// A Board represents a Trello board, composed by lists.
type Board struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

func (b *Board) String() string {
	return fmt.Sprintf("Board ID: %s; Name: %s\n", b.ID, b.Name)
}

// A List represents a board list, composed by cards.
type List struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

func (l *List) String() string {
	return fmt.Sprintf("List ID: %s; Name: %s\n", l.ID, l.Name)
}

// A Label represents a card label.
type Label struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

func (l *Label) String() string {
	return fmt.Sprintf("Label ID: %s; Name: %s\n", l.ID, l.Name)
}

// A Card represents a task that can be added to a list.
type Card struct {
	Name     string `json:"name"`
	Desc     string `json:"desc"`
	IDList   string `json:"idList"`
	IDLabels string `json:"idLabels"`
}

func (c *Card) String() string {
	return fmt.Sprintf("Card Name: %s; Labels: %s; List: %s\n", c.Name, c.IDLabels, c.IDList)
}