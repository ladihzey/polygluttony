package logger

import (
	"channels/utils"
	"fmt"
	"os"
	"sort"
	"strings"
	"sync"
	"time"
)

type logEntry struct {
	elapsed time.Duration
	source  string
	message string
}

type FileLogger struct {
	entries     []logEntry
	sourceOrder []string
	mu          sync.Mutex
}

func newFileLogger() *FileLogger {
	return &FileLogger{}
}

func (f *FileLogger) registerSource(name string) {
	f.mu.Lock()
	f.sourceOrder = append(f.sourceOrder, name)
	f.mu.Unlock()
}

func (f *FileLogger) log(source, message string, elapsed time.Duration) {
	f.mu.Lock()
	f.entries = append(f.entries, logEntry{elapsed, source, message})
	f.mu.Unlock()
}

func (f *FileLogger) flush(filename string) error {
	f.mu.Lock()
	defer f.mu.Unlock()

	if len(f.entries) == 0 {
		return nil
	}

	timeColWidth := 10
	colWidth := 50
	numCols := len(f.sourceOrder)

	// Build source index map
	sourceCol := make(map[string]int)
	for i, s := range f.sourceOrder {
		sourceCol[s] = i
	}

	// Sort entries by time, then by source order
	sort.Slice(f.entries, func(i, j int) bool {
		if f.entries[i].elapsed != f.entries[j].elapsed {
			return f.entries[i].elapsed < f.entries[j].elapsed
		}
		return sourceCol[f.entries[i].source] < sourceCol[f.entries[j].source]
	})

	var sb strings.Builder

	// Header row
	sb.WriteString(utils.Center("TIME", timeColWidth))
	sb.WriteString("│")
	for i, src := range f.sourceOrder {
		header := fmt.Sprintf("[%s]", strings.ToUpper(src))
		sb.WriteString(utils.Center(header, colWidth))
		if i < numCols-1 {
			sb.WriteString("│")
		}
	}
	sb.WriteString("\n")
	sb.WriteString(strings.Repeat("─", timeColWidth+1+colWidth*numCols+numCols-1) + "\n")

	for _, e := range f.entries {
		// Time column (round for display, but sorting used full precision)
		sb.WriteString(fmt.Sprintf("%9s ", e.elapsed.Round(time.Millisecond)))
		sb.WriteString("│")

		// Build the row with message in correct column
		col := sourceCol[e.source]
		text := e.message
		if len(text) > colWidth-2 {
			text = text[:colWidth-5] + "..."
		}

		for i := 0; i < numCols; i++ {
			if i == col {
				sb.WriteString(fmt.Sprintf(" %-*s", colWidth-1, text))
			} else {
				sb.WriteString(strings.Repeat(" ", colWidth))
			}
			if i < numCols-1 {
				sb.WriteString("│")
			}
		}
		sb.WriteString("\n")
	}

	return os.WriteFile(filename, []byte(sb.String()), 0644)
}
