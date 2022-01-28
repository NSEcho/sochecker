package checker

import (
	"fmt"
	"net/http"
	"os"
	"sort"
	"strings"
	"sync"
	"time"

	"github.com/jedib0t/go-pretty/table"
	"github.com/jedib0t/go-pretty/text"
	"github.com/muesli/termenv"
)

var (
	mutex  sync.RWMutex
	checks = make(map[string]Checker)
	wg     sync.WaitGroup
	resCh  = make(chan result, 10)
	done   = make(chan struct{})
	p      = termenv.ColorProfile()
)

type result struct {
	driverName string
	found      bool
	link       string
}

type Checker interface {
	Check(client *http.Client, name string) bool
	Info() string
	Link() string
}

func Register(name string, driver Checker) {
	mutex.Lock()
	defer mutex.Unlock()
	checks[name] = driver
}

func Checks() []string {
	mutex.RLock()
	defer mutex.RUnlock()
	list := make([]string, 0, len(checks))
	for name := range checks {
		list = append(list, name)
	}
	sort.Strings(list)
	return list
}

func Info(moduleName string) string {
	if moduleName != "all" && moduleName != "" {
		return checks[moduleName].Info()
	}
	var infos []string
	for name, check := range checks {
		info := fmt.Sprintf("\"%s\": %s", name, check.Info())
		infos = append(infos, info)
	}
	return strings.Join(infos, "\n")
}

func RunAll(name string) {
	mutex.RLock()
	defer mutex.RUnlock()

	client := &http.Client{
		Timeout: 3 * time.Second,
	}

	var results []result
	go func() {
		for res := range resCh {
			results = append(results, res)
		}
		done <- struct{}{}
	}()

	wg.Add(len(checks))
	for k, v := range checks {
		var driverName = k
		var driver = v
		go func() {
			fmt.Printf("Checking %s for \"%s\"\n", driverName, name)
			defer wg.Done()
			found := driver.Check(client, name)
			resCh <- result{
				driverName: driverName,
				found:      found,
				link:       driver.Link(),
			}
		}()
	}
	wg.Wait()
	close(resCh)
	<-done

	termenv.ClearScreen()

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.SetStyle(table.StyleColoredYellowWhiteOnBlack)
	t.SetIndexColumn(1)

	t.AppendHeader(table.Row{"#", "Driver", "Found", "Link"})

	var foundCounter = 0

	for i, res := range results {
		var found termenv.Style
		var link string
		if res.found {
			foundCounter++
			link = res.link
			found = termenv.String(fmt.Sprintf("%+v", res.found)).Foreground(p.Color("#00ff00"))
		} else {
			link = ""
			found = termenv.String(fmt.Sprintf("%+v", res.found)).Foreground(p.Color("#ff0000"))
		}
		t.AppendRow(table.Row{i + 1, res.driverName, found, link})
	}

	title := fmt.Sprintf("Results of checking \"%s\" \n(Drivers count: %d, found in: %d)", name, len(results), foundCounter)
	t.SetTitle(title)
	t.Style().Title.Align = text.AlignCenter

	t.Render()
}
