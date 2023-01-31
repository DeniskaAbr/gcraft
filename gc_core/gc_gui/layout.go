package gc_gui

const layoutDebug = false

type Layout struct {
	entries []*layoutEntry
}

func (l *Layout) advance(elapsed float64) error {
	for _, entry := range l.entries {
		if err := entry.widget.advance(elapsed); err != nil {
			return err
		}
	}

	return nil
}
