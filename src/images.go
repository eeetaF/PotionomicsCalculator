package src

import (
	"fmt"
	"image"
	"image/color"
	"os"
	"path/filepath"

	"github.com/disintegration/imaging"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"io/ioutil"
)

var (
	fontFaceMain  font.Face
	fontFaceLarge font.Face
	fontFaceSmall font.Face
)

func init() {
	fontBytes, err := ioutil.ReadFile("assets/Roboto-Bold.ttf")
	if err != nil {
		panic("failed to load font: " + err.Error())
	}
	fnt, err := opentype.Parse(fontBytes)
	if err != nil {
		panic(err)
	}
	fontFaceMain, _ = opentype.NewFace(fnt, &opentype.FaceOptions{Size: 24, DPI: 72, Hinting: font.HintingFull})
	fontFaceLarge, _ = opentype.NewFace(fnt, &opentype.FaceOptions{Size: 36, DPI: 72, Hinting: font.HintingFull})
	fontFaceSmall, _ = opentype.NewFace(fnt, &opentype.FaceOptions{Size: 16, DPI: 72, Hinting: font.HintingFull})
}

func drawTextWithOutline(dst *ebiten.Image, str string, face font.Face, x, y int, fg color.Color, outline color.Color, outlineWidth int) {
	for dx := -outlineWidth; dx <= outlineWidth; dx++ {
		for dy := -outlineWidth; dy <= outlineWidth; dy++ {
			if dx == 0 && dy == 0 {
				continue
			}
			text.Draw(dst, str, face, x+dx, y+dy, outline)
		}
	}
	text.Draw(dst, str, face, x, y, fg)
}

func drawRect(dst *ebiten.Image, x, y, w, h int, col color.Color) {
	rect := ebiten.NewImage(w, h)
	rect.Fill(col)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(x), float64(y))
	dst.DrawImage(rect, op)
}

type Game struct {
	sprites     map[string]*ebiten.Image
	keys        []string // To keep order
	results     *[]SearchResult
	staticImage *ebiten.Image
	initialized bool
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if !g.initialized {
		w, h := 1920, 1080
		g.staticImage = ebiten.NewImage(w, h)
		bg := color.RGBA{196, 167, 167, 255}
		g.staticImage.Fill(bg)

		const spacing = 110
		const labelOffset = 105
		const rowSpacing = 140
		const ingredientSpacing = 110
		const infoOffset = 120

		if g.results != nil {
			for rowIdx, sr := range *g.results {
				mainKey := sr.ResultingPotion.Name
				img, ok := g.sprites[mainKey]
				if !ok {
					continue
				}
				x := 10.0
				y := float64(rowIdx*rowSpacing + 10)
				op := &ebiten.DrawImageOptions{}
				op.GeoM.Translate(x, y)
				g.staticImage.DrawImage(img, op)

				mainText := mainKey
				textWidth := text.BoundString(fontFaceMain, mainText).Dx()
				textX := int(x) + 50 - textWidth/2
				textY := int(y) + labelOffset
				drawTextWithOutline(g.staticImage, mainText, fontFaceMain, textX, textY, color.White, color.Black, 2)

				// Move info higher
				infoX := int(x) + 100 + 10
				infoY := int(y) + 10

				// Calculate block height for background
				blockW := 320
				blockH := 36 + 18 + 18 + len(sr.Traits)*20 + 10
				bgCol := color.RGBA{0, 0, 0, 128}
				drawRect(g.staticImage, infoX-8, infoY-8, blockW, blockH, bgCol)

				// Magimints
				drawTextWithOutline(g.staticImage, fmt.Sprintf("M: %d", sr.TotalMagimints), fontFaceMain, infoX, infoY, color.White, color.Black, 2)

				// Ingredients
				infoY2 := infoY + 30
				drawTextWithOutline(g.staticImage, fmt.Sprintf("I: %d", sr.NumberIngreds), fontFaceMain, infoX, infoY2, color.White, color.Black, 2)

				// Traits heading
				traitsY := infoY2 + 28
				drawTextWithOutline(g.staticImage, "Traits:", fontFaceMain, infoX, traitsY, color.White, color.Black, 2)

				// Each trait on a new line, colored
				traitLineY := traitsY + 24
				for _, t := range sr.Traits {
					traitLabel := ""
					traitColor := color.RGBA{200, 255, 200, 255}
					if t.IsGood {
						traitLabel += "+"
						traitColor = color.RGBA{80, 255, 80, 255}
					} else {
						traitLabel += "-"
						traitColor = color.RGBA{255, 80, 80, 255}
					}
					traitLabel += t.Trait.String()
					drawTextWithOutline(g.staticImage, traitLabel, fontFaceMain, infoX+16, traitLineY, traitColor, color.Black, 2)
					traitLineY += 24
				}

				// Draw ingredients as before, but quantity larger and more visible
				ingrX := infoX + blockW + 16
				for ingrKey, qty := range sr.Ingredients {
					ingrImg, ok := g.sprites[ingrKey]
					if !ok {
						continue
					}
					op := &ebiten.DrawImageOptions{}
					op.GeoM.Translate(float64(ingrX), y)
					g.staticImage.DrawImage(ingrImg, op)
					label := ingrKey
					labelWidth := text.BoundString(fontFaceSmall, label).Dx()
					labelX := ingrX + 50 - labelWidth/2
					labelY := int(y) + 110
					drawTextWithOutline(g.staticImage, label, fontFaceSmall, labelX, labelY, color.White, color.Black, 2)
					// Draw quantity at top-right, bigger font and colored
					qtyStr := fmt.Sprintf("%d", qty)
					qtyWidth := text.BoundString(fontFaceLarge, qtyStr).Dx()
					qtyX := ingrX + 100 - qtyWidth - 4
					qtyY := int(y) + 40
					qtyColor := color.RGBA{255, 255, 128, 255}
					drawTextWithOutline(g.staticImage, qtyStr, fontFaceLarge, qtyX, qtyY, qtyColor, color.Black, 3)
					ingrX += 120
				}
			}
		}
		g.initialized = true
	}
	screen.DrawImage(g.staticImage, nil)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 1920, 1080
}

func loadSprites(dir string) (map[string]*ebiten.Image, []string, error) {
	sprites := make(map[string]*ebiten.Image)
	var keys []string

	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, nil, err
	}

	for _, file := range files {
		if file.IsDir() || filepath.Ext(file.Name()) != ".png" {
			continue
		}

		name := file.Name()
		key := name[:len(name)-len(filepath.Ext(name))]
		f, err := os.Open(filepath.Join(dir, name))
		if err != nil {
			continue
		}
		img, _, err := image.Decode(f)
		f.Close()
		if err != nil {
			continue
		}

		resized := imaging.Resize(img, 100, 100, imaging.Lanczos)
		circled := cropCircle(resized)
		eimg := ebiten.NewImageFromImage(circled)
		sprites[key] = eimg
		keys = append(keys, key)
	}
	return sprites, keys, nil
}

func cropCircle(src image.Image) image.Image {
	dst := image.NewRGBA(image.Rect(0, 0, 100, 100))
	center := 50
	radius := 50
	for y := 0; y < 100; y++ {
		for x := 0; x < 100; x++ {
			dx := x - center
			dy := y - center
			if dx*dx+dy*dy <= radius*radius {
				dst.Set(x, y, src.At(x, y))
			} else {
				dst.Set(x, y, color.Transparent)
			}
		}
	}
	return dst
}

func run(sr *[]SearchResult) {
	sprites, keys, err := loadSprites("data")
	if err != nil {
		panic(err)
	}
	g := &Game{sprites: sprites, keys: keys, results: sr}
	ebiten.SetWindowSize(1920, 1080)
	ebiten.SetWindowTitle("PNG Viewer")
	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}
