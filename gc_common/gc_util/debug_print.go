package gc_util

import (
	"image"

	"gcraft/gc_common/gc_util/assets"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	cw = assets.CharWidth
	ch = assets.CharHeight
)

func NewDebugPrinter() *GlyphPrinter {
	img := ebiten.NewImageFromImage(assets.CreateTextImage())

	printer := &GlyphPrinter{
		glyphImageTable: img,
		glyphs:          make(map[rune]*ebiten.Image),
	}

	return printer
}

type GlyphPrinter struct {
	glyphImageTable *ebiten.Image
	glyphs          map[rune]*ebiten.Image
}

func (p *GlyphPrinter) PrintAt(target interface{}, str string, x, y int) {
	p.drawDebugText(target.(*ebiten.Image), str, x, y, false)
}

func (p *GlyphPrinter) drawDebugText(target *ebiten.Image, str string, ox, oy int, shadow bool) {
	op := &ebiten.DrawImageOptions{}

	if shadow {
		op.ColorM.Scale(0, 0, 0, 0.5)
	}

	x := 0
	y := 0

	w, _ := p.glyphImageTable.Size()

	for _, c := range str {
		if c == '\n' {
			x = 0
			y += ch

			continue
		}

		s, ok := p.glyphs[c]
		if !ok {
			n := w / cw
			sx := (int(c) % n) * cw
			sy := (int(c) / n) * ch
			rect := image.Rect(sx, sy, sx+cw, sy+ch)
			s = p.glyphImageTable.SubImage(rect).(*ebiten.Image)
			p.glyphs[c] = s
		}

		op.GeoM.Reset()
		op.GeoM.Translate(float64(x), float64(y))
		op.GeoM.Translate(float64(ox+1), float64(oy))

		op.CompositeMode = ebiten.CompositeModeLighter
		target.DrawImage(s, op)

		x += cw
	}
}
