package cartridge

import (
	"embed"
	"image"

	"github.com/TheMightyGit/marv/marvlib"
)

//go:embed "resources/*"
var Resources embed.FS

const (
	GfxBankFont = iota
	GfxBankMrmoText
)
const (
	MapBankStuff = iota
)
const (
	Sprite0 = iota
)

func Start() {
	area := marvlib.API.MapBanksGet(MapBankStuff).GetArea(0)
	marvlib.API.SpritesGet(Sprite0).ChangePos(image.Rectangle{
		Min: image.Point{X: 0, Y: 0},
		Max: image.Point{X: 240, Y: 240},
	})
	marvlib.API.SpritesGet(Sprite0).Show(GfxBankMrmoText, area)
}

func Update() {
}
