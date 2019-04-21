package main

/* termbox-go: cui library
 * installation: go get -u github.com/nsf/termbox-go
 * api: https://godoc.org/github.com/nsf/termbox-go
 */

import (
	"time"
	log "github.com/Sirupsen/logrus"
	"github.com/nsf/termbox-go"
)

/* 0. Configure the display setting */

const disp_color = termbox.ColorDefault
const (
	// Color
	Black   = termbox.ColorBlack
	Red     = termbox.ColorRed
	Green   = termbox.ColorGreen
	Yellow  = termbox.ColorYellow
	Blue    = termbox.ColorBlue
	Magenta = termbox.ColorMagenta
	Cyan    = termbox.ColorCyan
	White   = termbox.ColorWhite
	// Character (Add with OR to Color)
	Bold    = termbox.AttrBold
	Under   = termbox.AttrUnderline
	Reverse = termbox.AttrReverse
)

var disp_width, disp_height int

func validate_width(x int) int {
	if x < 0 && x > disp_width {
		print("width ", x, ":Out of Range")
		return 1
	}
	return 0
}

func validate_height(y int) int {
	if y < 0 && y > disp_height {
		print("height ", y, ":Out of Range")
		return 1
	}
	return 0
}

/* 1. Create Draw function */
/* In termbox, the drawrable objects are callod Cell.
 * Cell is the struct of a character, Fg(character color)
 * and Bg(background color).
 */

func DrawHbar(xs, xe, y int) bool {
	ret := 0
	ret += validate_width(xs)
	ret += validate_width(xe)
	ret += validate_height(y)
	if ret != 0 {
		return false
	}
	if xs > xe {
		xs, xe = xe, xs
	}
	// draw horizontal bar
	for x := xs; x != xe; x++ {
		// SetCell(x, y iny, ch rune,
		//         foreground, background termbox.Attribute)
		termbox.SetCell(x, y, ' ', Blue | Reverse, Blue)
	}
	return true
}

func DrawVbar(x, ys, ye int) bool {
	ret := 0
	ret += validate_width(x)
	ret += validate_height(ys)
	ret += validate_height(ye)
	if ret != 0 {
		return false
	}
	if ys > ye {
		ys, ye = ye, ys
	}
	//draw vertical bar
	for y := ys; y != ye; y++ {
		termbox.SetCell(x, y, ' ', Blue | Reverse, Blue)
	}
	return true
}

func DrawDot(x, y int) bool {
	ret := 0
	ret += validate_width(x)
	ret += validate_height(y)
	if ret != 0 {
		return false
	}
	termbox.SetCell(x, y, ' ', Blue | Reverse, Blue)
	return true
}

func drawStart() {
	hpad := int(disp_width / 30)
	vcenter := int(disp_height / 2)
	x := 0
	y := vcenter
	/* draw "start" */
	/* a charcater ocupies 5 * 7 cells. pad 1. */
	// S
	DrawHbar(x, x + 4, y - 3)
	DrawVbar(x, y, y - 3)
	DrawHbar(x, x + 4, y)
	DrawVbar(x + 4, y, y + 3)
	DrawHbar(x, x + 4, y + 3)
	// T
	x += hpad
	DrawHbar(x, x + 4, y + 3)
	DrawVbar(x + 2, y - 3, y + 3)
	// A
	x += hpad
	x = x + hpad
	DrawVbar(hpad, vcenter - 2, vcenter + 2)
	DrawHbar(hpad, hpad + 2, vcenter + 2)
	DrawHbar(hpad, hpad + 2, vcenter)
	// R
	x += hpad
	DrawHbar(x, x + 4, y - 3)
	DrawVbar(x, y + 3, y - 3)
	DrawVbar(x + 4, y, y - 3)
	DrawHbar(x, x + 4, y)
	// TODO: create DrawDigbar
	DrawDot(x + 2, y + 1)
	DrawDot(x + 3, y + 2)
	DrawDot(x + 4, y + 3)
	// T
	x += hpad
	DrawHbar(x, x + 4, y + 3)
	DrawVbar(x + 2, y - 3, y + 3)
}

/* 2. Define Event and Event Handler */

func handleKeyEvent(pressed_key termbox.Key) bool {
	ret := false
	switch pressed_key {
	case termbox.KeyEsc:
		ret = true
	}
	return ret
}

/* 3. Define Main Loop */

func main() {
	/* prepare the display buffer */
	if err := termbox.Init(); err != nil {
		log.Fatal(err)
	}
	defer termbox.Close()

	/* set the display size */
	disp_width, disp_height = termbox.Size()

	/* add dot per one loop */
	var startStr = []rune{
		's',
		't',
		'a',
		'r',
		't',
	}
	termbox.Clear(disp_color, disp_color)
	drawStart()
	for i := 1; i < 6; i++ {
		termbox.SetCell(i, 1, startStr[i -1],
				 White, Blue)
	}
	termbox.Flush()
	x, y := 0, 0

MAIN_LOOP:
	for {
		/* Clear previous display */
		termbox.Clear(disp_color, disp_color)
		/* Event Handling */
		switch event := termbox.PollEvent(); event.Type {
		case termbox.EventKey:
			if handleKeyEvent(event.Key) {
				break MAIN_LOOP
			}
		}
		/* Draw buffer */
		drawStart()
		DrawDot(x, y)
		/* Draw display */
		termbox.Flush()

		if y == disp_height {
			y = 0
		}
		if x == disp_width {
			x = 0
			y += 1
		}
		time.Sleep(1000)
	}
}

