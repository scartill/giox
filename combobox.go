package giox

import (
	"errors"

	"gioui.org/widget"
)

// Combo holds combobox state
type Combo struct {
	items    []string
	hint string
	selected int
	expanded bool
	selectButton widget.Button
	buttons []widget.Button
}

// MakeCombo Creates new combobox widget
func MakeCombo(items []string, hint string) Combo {
	c := Combo{
		items:    items,
		hint: hint,
		selected: -1,
		expanded: false,
		buttons: make([]widget.Button, len(items)),
	}

	return c
}

// HasSelected returns true if an item is selected
func (c *Combo) HasSelected() bool {
	return c.selected != -1
}

// IsExpanded checks wheather box is expanded
func (c *Combo) IsExpanded() bool {
	return c.expanded
}

// Toggle expands and collapses a combobox
func (c *Combo) Toggle() bool {
	c.expanded = !c.expanded
	return c.expanded
}

// Len returns number of items
func (c *Combo) Len() int {
	return len(c.items)
}

// Items returns current list of items
func (c *Combo) Items() []string {
	return c.items
}

// Hint returns control's hint test
func (c *Combo) Hint() string {
	return c.hint
}

// Item returns a text for corresponding item
func (c *Combo) Item(index int) string {
	return c.items[index]
}

// SelectButton returns a points to main (open) combobox button
func (c *Combo) SelectButton() *widget.Button {
	return &c.selectButton
}

// Button returns a pointer to correspoding button widget
func (c *Combo) Button(index int) *widget.Button {
	return &(c.buttons[index])
}

// SelectedText returns currently selected item
func (c *Combo) SelectedText() string {
	if c.selected == -1 {
		return c.hint
	}

	return c.items[c.selected]
}

// SelectIndex sets currenly selected item by index
func (c *Combo) SelectIndex(index int) error {
	N := len(c.items)
	if index != -1 && (index < 0 || index >= N) {
		return errors.New("Combobox: bad index")
	}

	c.selected = index
	return nil
}

// SelectItem sets currenly selected item by value
func (c *Combo) SelectItem(item string) error {
	for i, val := range c.items {
		if val == item {
			c.selected = i
			return nil
		}
	}

	return errors.New("Combobox: bad index")
}

// Unselect removes current selection
func (c *Combo) Unselect() {
	c.selected = -1
}
