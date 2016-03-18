// Sample program demonstrating struct composition.
package main

import "fmt"

// Board represents a surface we can work on.
type Board struct {
	NailsNeeded int
	NailsDriven int
}

// =============================================================================

// Mallet is a tool that pounds in nails.
type Mallet struct{}

// DriveNail pounds a nail into the specified board.
func (Mallet) DriveNail(nailSupply *int, b *Board) {
	*nailSupply--
	b.NailsDriven++
	fmt.Println("Mallet: pounded nail into the board.")
}

// Crowbar is a tool that removes nails.
type Crowbar struct{}

// PullNail yanks a nail out of the specified board.
func (Crowbar) PullNail(nailSupply *int, b *Board) {
	b.NailsDriven--
	*nailSupply++
	fmt.Println("Crowbar: yanked nail out of the board.")
}

// =============================================================================

// Toolbox can contains a Mallet and a Crowbar.
type Toolbox struct {
	Mallet
	Crowbar

	nails int
}

// =============================================================================

// Contractor carries out the task of securing boards.
type Contractor struct{}

// Fasten will drive nails into a board.
func (Contractor) Fasten(m Mallet, nailSupply *int, b *Board) {
	for b.NailsDriven < b.NailsNeeded {
		m.DriveNail(nailSupply, b)
	}
}

// Unfasten will remove nails from a board.
func (Contractor) Unfasten(cb Crowbar, nailSupply *int, b *Board) {
	for b.NailsDriven > b.NailsNeeded {
		cb.PullNail(nailSupply, b)
	}
}

// ProcessBoards works against boards.
func (c Contractor) ProcessBoards(tb *Toolbox, nailSupply *int, boards []Board) {
	for i := range boards {
		b := &boards[i]

		fmt.Printf("Contractor: examining board #%d: %+v\n", i+1, b)

		switch {
		case b.NailsDriven < b.NailsNeeded:
			c.Fasten(tb.Mallet, nailSupply, b)

		case b.NailsDriven > b.NailsNeeded:
			c.Unfasten(tb.Crowbar, nailSupply, b)
		}
	}
}

// =============================================================================

// main is the entry point for the application.
func main() {
	boards := []Board{
		{NailsDriven: 3},
		{NailsNeeded: 2},
	}

	tb := Toolbox{
		Mallet:  Mallet{},
		Crowbar: Crowbar{},
		nails:   10,
	}

	var c Contractor
	c.ProcessBoards(&tb, &tb.nails, boards)
}
