package terminal

import (
    "fmt"
    "log"

    "go-terminal-ai/internal/ai"

    "github.com/gdamore/tcell/v2"
    "github.com/rivo/tview"
)

// StartTerminal starts the terminal emulation
func StartTerminal(app *tview.Application) {
    textView := tview.NewTextView().
        SetDynamicColors(true).
        SetRegions(true).
        SetWordWrap(true)

    textView.SetChangedFunc(func() {
        app.Draw()
    })

    go func() {
        for {
            fmt.Fprintf(textView, "[blue]$ [white]")
            var input string
            fmt.Scanln(&input)
            fmt.Fprintf(textView, "[yellow]%s\n", input)
            suggestion := ai.GetSuggestions(input)
            fmt.Fprintf(textView, "[green]AI Suggestion: %s\n", suggestion)
        }
    }()

    if err := app.SetRoot(textView, true).Run(); err != nil {
        log.Fatalf("Error running application: %v", err)
    }
}
