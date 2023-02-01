package block

import (
	"github.com/df-mc/dragonfly/server/block/cube"
	"github.com/df-mc/dragonfly/server/internal/nbtconv"
	"github.com/df-mc/dragonfly/server/world"
)

// Moving ...
type Moving struct {
	empty
	transparent

	// Moving represents the block that is moving.
	Moving world.Block
	// Extra represents an extra block that is moving with the main block.
	Extra world.Block
	// Piston is the position of the piston that is moving the block.
	Piston cube.Pos
}

// EncodeBlock ...
func (Moving) EncodeBlock() (string, map[string]any) {
	return "minecraft:moving_block", nil
}

// PistonImmovable ...
func (Moving) PistonImmovable() bool {
	return true
}

// EncodeNBT ...
func (b Moving) EncodeNBT() map[string]any {
	return map[string]any{
		"id":               "Moving",
		"movingBlock":      nbtconv.WriteBlock(b.Moving),
		"movingBlockExtra": nbtconv.WriteBlock(b.Extra),
		"pistonPosX":       int32(b.Piston.X()),
		"pistonPosY":       int32(b.Piston.Y()),
		"pistonPosZ":       int32(b.Piston.Z()),
	}
}

// DecodeNBT ...
func (b Moving) DecodeNBT(m map[string]any) any {
	b.Moving = nbtconv.Block(m, "movingBlock")
	b.Extra = nbtconv.Block(m, "movingBlockExtra")
	b.Piston = cube.Pos{
		int(nbtconv.Int32(m, "pistonPosX")),
		int(nbtconv.Int32(m, "pistonPosY")),
		int(nbtconv.Int32(m, "pistonPosZ")),
	}
	return b
}
