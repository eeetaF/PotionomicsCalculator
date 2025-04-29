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

func drawRoundedRect(dst *ebiten.Image, x, y, w, h, r int, col color.Color) {
	// Center rectangle
	rect := ebiten.NewImage(w-2*r, h)
	rect.Fill(col)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(float64(x+r), float64(y))
	dst.DrawImage(rect, op)
	// Side rectangles
	rectV := ebiten.NewImage(r, h-2*r)
	rectV.Fill(col)
	op1 := &ebiten.DrawImageOptions{}
	op1.GeoM.Translate(float64(x), float64(y+r))
	dst.DrawImage(rectV, op1)
	op2 := &ebiten.DrawImageOptions{}
	op2.GeoM.Translate(float64(x+w-r), float64(y+r))
	dst.DrawImage(rectV, op2)
	// Four corners (circles)
	circ := ebiten.NewImage(r*2, r*2)
	for cy := 0; cy < r*2; cy++ {
		for cx := 0; cx < r*2; cx++ {
			dx := cx - r
			dy := cy - r
			if dx*dx+dy*dy <= r*r {
				circ.Set(cx, cy, col)
			}
		}
	}
	// Top-left
	opTL := &ebiten.DrawImageOptions{}
	geoTL := ebiten.GeoM{}
	geoTL.Translate(float64(x), float64(y))
	opTL.GeoM = geoTL
	dst.DrawImage(circ, opTL)
	// Top-right
	opTR := &ebiten.DrawImageOptions{}
	geoTR := ebiten.GeoM{}
	geoTR.Translate(float64(x+w-2*r), float64(y))
	opTR.GeoM = geoTR
	dst.DrawImage(circ, opTR)
	// Bottom-left
	opBL := &ebiten.DrawImageOptions{}
	geoBL := ebiten.GeoM{}
	geoBL.Translate(float64(x), float64(y+h-2*r))
	opBL.GeoM = geoBL
	dst.DrawImage(circ, opBL)
	// Bottom-right
	opBR := &ebiten.DrawImageOptions{}
	geoBR := ebiten.GeoM{}
	geoBR.Translate(float64(x+w-2*r), float64(y+h-2*r))
	opBR.GeoM = geoBR
	dst.DrawImage(circ, opBR)
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

				// Magimints (smaller font)
				drawTextWithOutline(g.staticImage, fmt.Sprintf("M: %d", sr.TotalMagimints), fontFaceSmall, infoX, infoY, color.White, color.Black, 2)

				// Ingredients (smaller font)
				infoY2 := infoY + 18
				drawTextWithOutline(g.staticImage, fmt.Sprintf("I: %d", sr.NumberIngreds), fontFaceSmall, infoX, infoY2, color.White, color.Black, 2)

				// Each trait on a new line, colored, smaller font
				traitLineY := infoY2 + 18
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
					drawTextWithOutline(g.staticImage, traitLabel, fontFaceSmall, infoX+12, traitLineY, traitColor, color.Black, 2)
					traitLineY += 18
				}

				// Draw ingredients as before, but quantity larger and more visible
				ingrX := infoX + 120
				for ingrKey, qty := range sr.Ingredients {
					ingrImg, ok := g.sprites[ingrKey]
					op := &ebiten.DrawImageOptions{}
					op.GeoM.Translate(float64(ingrX), y)
					if ok {
						g.staticImage.DrawImage(ingrImg, op)
					} else {
						// Draw placeholder: gray circle with ?
						ph := ebiten.NewImage(100, 100)
						// Fill with gray
						for py := 0; py < 100; py++ {
							for px := 0; px < 100; px++ {
								dx := px - 50
								dy := py - 50
								if dx*dx+dy*dy <= 50*50 {
									ph.Set(px, py, color.RGBA{180, 180, 180, 255})
								}
							}
						}
						g.staticImage.DrawImage(ph, op)
						// Draw ? in the center
						qW := text.BoundString(fontFaceLarge, "?").Dx()
						qH := text.BoundString(fontFaceLarge, "?").Dy()
						qx := ingrX + 50 - qW/2
						qy := int(y) + 50 + qH/2
						drawTextWithOutline(g.staticImage, "?", fontFaceLarge, qx, qy, color.Black, color.White, 2)
					}
					label := ingrKey
					// Split label into lines of max 12 chars, word-aware
					labelLines := splitLinesWordWrap(label, 12)
					labelY := int(y) + 110
					for i, line := range labelLines {
						labelWidth := text.BoundString(fontFaceSmall, line).Dx()
						labelX := ingrX + 50 - labelWidth/2
						drawTextWithOutline(g.staticImage, line, fontFaceSmall, labelX, labelY+i*18, color.White, color.Black, 2)
					}
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

func splitLinesWordWrap(s string, n int) []string {
	var lines []string
	words := []rune(s)
	start := 0
	for start < len(words) {
		end := start
		lineLen := 0
		lastSpace := -1
		for end < len(words) && lineLen < n {
			if words[end] == ' ' {
				lastSpace = end
			}
			lineLen++
			end++
		}
		if end < len(words) && lastSpace > start {
			lines = append(lines, string(words[start:lastSpace]))
			start = lastSpace + 1
		} else {
			lines = append(lines, string(words[start:end]))
			start = end
		}
	}
	return lines
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
