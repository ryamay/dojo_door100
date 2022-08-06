package door100

import "fmt"

/**
100 doors in a row are all initially closed. You make 100 passes by the doors. The first time through, you visit every door and toggle the door (if the door is closed, you open it; if it is open, you close it).
The second time you only visit every 2nd door (door #2, #4, #6, ...).
The third time, every 3rd door (door #3, #6, #9, ...), etc, until you only visit the 100th door.

Question: What state are the doors in after the last pass? Which are open, which are closed?
**/

type Door bool

const (
	Open   Door = true
	Closed Door = false
)

func (d Door) visit() Door {
	return !d
}

func (d Door) isClosed() bool {
	return d == Closed
}

func (d Door) isOpen() bool {
	return d == Open
}

type Doors []Door

func InitDoors() Doors {
	return make([]Door, 100)
}

func (ds Doors) visit(i int) {
	ds[i] = ds[i].visit()
}

func (ds Doors) VisitInPhase(p int) {
	for i := p - 1; i < len(ds); i += p {
		ds.visit(i)
	}
}

func (ds Doors) ShowClosed() string {
	var closed []int

	for i, d := range ds {
		if d.isClosed() {
			closed = append(closed, i+1)
		}
	}

	return fmt.Sprintf("Closed: %v\n", closed)
}

func (ds Doors) ShowOpen() string {
	var opens []int

	for i, d := range ds {
		if d.isOpen() {
			opens = append(opens, i+1)
		}
	}

	return fmt.Sprintf("Open: %v\n", opens)
}
