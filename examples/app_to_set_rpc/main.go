package main

import (
	"fmt"
	"log"
	"time"

	"github.com/andlabs/ui"
	rpc "github.com/superbizonsdevelopment/bizonrpc"
)

func setupUI() {
	mainwin := ui.NewWindow("BizonRPC", 600, 600, true)
	mainwin.SetMargined(true)

	mainwin.OnClosing(func(*ui.Window) bool {
		mainwin.Destroy()
		ui.Quit()
		return false
	})

	ui.OnShouldQuit(func() bool {
		mainwin.Destroy()
		return true
	})

	hbox := ui.NewHorizontalBox()
	hbox.SetPadded(true)
	mainwin.SetChild(hbox)

	vbox := ui.NewVerticalBox()
	vbox.SetPadded(true)
	hbox.Append(vbox, false)

	grid := ui.NewGrid()
	grid.SetPadded(true)

	client_id_label := ui.NewLabel("ID")
	state_label := ui.NewLabel("State")
	details_label := ui.NewLabel("Details")

	client_id_entry := ui.NewEntry()
	state_entry := ui.NewEntry()
	details_entry := ui.NewEntry()

	submitButton := ui.NewButton("Submit")
	submitButton.OnClicked(func(b *ui.Button) {

		if sthDatasSend {
			quit <- true
		}
		sthDatasSend = true
		go updateRPC(client_id_entry.Text(), state_entry.Text(), details_entry.Text())
	})

	vbox.Append(grid, false)

	grid.Append(client_id_label, 0, 0, 1, 1, false, ui.AlignFill, false, ui.AlignFill)
	grid.Append(client_id_entry, 1, 0, 1, 1, false, ui.AlignFill, false, ui.AlignFill)

	grid.Append(state_label, 0, 1, 1, 1, false, ui.AlignFill, false, ui.AlignFill)
	grid.Append(state_entry, 1, 1, 1, 1, false, ui.AlignFill, false, ui.AlignFill)

	grid.Append(details_label, 0, 2, 1, 1, false, ui.AlignFill, false, ui.AlignFill)
	grid.Append(details_entry, 1, 2, 1, 1, false, ui.AlignFill, false, ui.AlignFill)

	grid.Append(submitButton, 1, 3, 1, 1, false, ui.AlignFill, false, ui.AlignFill)

	mainwin.Show()
}

func updateRPC(client_id string, state string, details string) {
	if client_id == "" {
		log.Println("Error: You don't set client id!")
		return
	}

	win, err := rpc.New(client_id)

	if err != nil {
		log.Println("Error: ", err.Error())
		return
	}

	win.Open()

	for {
		select {
		case <-quit:
			return
		default:
			activity := &rpc.Activity{
				Details: details,
				State:   state,
				TimeStamps: &rpc.TimeStamps{
					StartTimestamp: time.Now().Unix(),
					EndTimestamp:   time.Now().Add(20 * time.Second).Unix(),
				},
			}

			err := win.SetRichPresence(activity)
			if err != nil {
				fmt.Println("Error: ", err.Error())
				continue
			}

			str, err := win.Connection.Read()

			if err != nil {
				fmt.Println("Error: ", err.Error())
				return
			}

			fmt.Println("STR1", str)

			fmt.Println("---\nDone?")
			time.Sleep(time.Second * 20)
		}
	}

}

var (
	start        time.Time
	sthDatasSend bool
	quit         = make(chan bool)
)

func main() {
	start = time.Now()
	ui.Main(setupUI)
}
