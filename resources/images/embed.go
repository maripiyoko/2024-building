package images

import (
	_ "embed"
)

var (
	//go:embed Archer-Walk.png
	ArcherWalk_png []byte

	//go:embed Archer-Idle.png
	ArcherIdle_png []byte

	//go:embed Background.png
	Background_png []byte

	//go:embed forest_items.png
	ForestItems_png []byte
)
