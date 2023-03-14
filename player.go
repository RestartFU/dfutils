package dfutil

import (
	"github.com/df-mc/dragonfly/server/entity/physics"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/go-gl/mathgl/mgl64"
)

func PlayersInRadius(p *player.Player, radius float64) (players []*player.Player) {
	p1 := mgl64.Vec3{-radius, -radius, -radius}.Add(p.Position())
	p2 := mgl64.Vec3{radius, radius, radius}.Add(p.Position())

	physic := physics.NewAABB(p1, p2)
	p.World().EntitiesWithin(physic, func(e world.Entity) bool {
		pl, ok := e.(*player.Player)
		if ok {
			players = append(players, pl)
		}
		return false
	})
	return
}
