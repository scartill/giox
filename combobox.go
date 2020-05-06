package giox

// Combo holds combobox state
type Combo struct {
    items    []string
    hint string
    selected int
    expanded bool
}

// MakeCombo Creates new combobox widget
func MakeCombo(items []string, hint string) Combo {
    c := Combo{
        items:    items,
        hint: hint,
        selected: -1,
        expanded: false,
        
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

// Items returns current list of items
func (c *Combo) Items() []string {
    return c.items
}

// Hint returns control's hint test
func (c *Combo) Hint() string {
    return c.hint
}

// SelectedText returns currently selected item
func (c *Combo) SelectedText() string {
    if c.selected == -1 {
        return c.hint
    }

    return c.items[c.selected]
}
