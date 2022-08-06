package door100

import (
	"strconv"
	"strings"
	"testing"
)

func Testドア初期化時はすべてClosed(t *testing.T) {
	for _, d := range InitDoors() {
		if d != false {
			t.Error("初期化時に空いてるドアがある！")
		}
	}
}

func Testドア初期化時の数は100個(t *testing.T) {
	if len(InitDoors()) != 100 {
		t.Error("ドア数が初期化時に100個じゃない！")
	}
}

func Test一つの閉まったドアをvisitすると開く(t *testing.T) {
	doors := InitDoors()

	doors.visit(1)

	if doors[1] != true {
		t.Error("visitしたドアが開いてない")
	}
}

func Test一つの閉まったドアを2回visitすると閉じる(t *testing.T) {
	doors := InitDoors()

	doors.visit(1)
	doors.visit(1)

	if doors[1] == true {
		t.Error("2回visitしたドアが閉じていない")
	}
}

func Test1回めのループではすべてのドアを訪れる(t *testing.T) {
	doors := InitDoors()
	doors.VisitInPhase(1)

	for i, door := range doors {
		if door != true {
			t.Error(strconv.Itoa(i+1) + "番目のドアが閉じたまま！")
		}
	}
}

func Test2回目のループでは2_4_6___100のドアを訪れる(t *testing.T) {
	doors := InitDoors()
	doors.VisitInPhase(2)

	for i, door := range doors {
		if (i+1)%2 == 0 {
			if door != true {
				t.Error(strconv.Itoa(i+1) + "番目のドアが閉じたまま！")
			}
		} else {
			if door == true {
				t.Error(strconv.Itoa(i+1) + "番目のドアが開いている！")
			}
		}
	}
}

func Test50回目のループでは50_100のドアを訪れる(t *testing.T) {
	doors := InitDoors()
	doors.VisitInPhase(50)

	for i, door := range doors {
		if (i+1) == 50 || (i+1) == 100 {
			if door != true {
				t.Error(strconv.Itoa(i+1) + "番目のドアが閉じたまま！")
			}
		} else {
			if door == true {
				t.Error(strconv.Itoa(i+1) + "番目のドアが開いている！")
			}
		}
	}
}

func Test51回目のループでは51のドアのみ訪れる(t *testing.T) {
	doors := InitDoors()
	doors.VisitInPhase(51)

	for i, door := range doors {
		if (i + 1) == 51 {
			if door != true {
				t.Error(strconv.Itoa(i+1) + "番目のドアが閉じたまま！")
			}
		} else {
			if door == true {
				t.Error(strconv.Itoa(i+1) + "番目のドアが開いている！")
			}
		}
	}
}

func Test閉じているドアを文字列で出力できる(t *testing.T) {
	doors := InitDoors()

	closedDoors := doors.ShowClosed()

	if !strings.Contains(closedDoors, "Closed:") {
		t.Error("Closed:　が表示されていない")
	}
	for i := 1; i <= len(doors); i++ {
		istr := strconv.Itoa(i)
		if !strings.Contains(closedDoors, istr) {
			t.Error(istr + "が閉じられているドアに含まれていない")
		}
	}
}

func Test開いているドアを文字列で出力できる(t *testing.T) {
	doors := InitDoors()

	doors.VisitInPhase(1)
	openDoors := doors.ShowOpen()

	if !strings.Contains(openDoors, "Open:") {
		t.Error("Open:　が表示されていない")
	}
	for i := 1; i <= len(doors); i++ {
		istr := strconv.Itoa(i)
		if !strings.Contains(openDoors, istr) {
			t.Error(istr + "が開いているドアに含まれていない")
		}
	}

}
