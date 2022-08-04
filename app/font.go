// Created by EldersJavas(EldersJavas&gmail.com)

package app

import (
	"embed"
	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
	"io"
	"io/fs"
)

//go:embed res
var resourceRootFS embed.FS

var ResourceFS fs.FS

var (
	DebugFace   font.Face
	TimeFace    font.Face
	ContextFace font.Face
)

func init() {
	f, err := ResourceFS.Open("fusion-pixel.otf")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	bs, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}

	fonts, err := opentype.Parse(bs)
	if err != nil {
		panic(err)
	}
	{
		face, err := opentype.NewFace(fonts, &opentype.FaceOptions{
			Size: 12,
			DPI:  72,
		})
		if err != nil {
			panic(err)
		}
		DebugFace = face
	}
	{
		face, err := opentype.NewFace(fonts, &opentype.FaceOptions{
			Size: 36,
			DPI:  72,
		})
		if err != nil {
			panic(err)
		}
		TimeFace = face
	}
	{
		face, err := opentype.NewFace(fonts, &opentype.FaceOptions{
			Size: 24,
			DPI:  72,
		})
		if err != nil {
			panic(err)
		}
		ContextFace = face
	}
}
