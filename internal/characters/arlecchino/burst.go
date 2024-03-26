package arlecchino

import (
	"github.com/genshinsim/gcsim/internal/frames"
	"github.com/genshinsim/gcsim/pkg/core/action"
	"github.com/genshinsim/gcsim/pkg/core/attacks"
	"github.com/genshinsim/gcsim/pkg/core/attributes"
	"github.com/genshinsim/gcsim/pkg/core/combat"
	"github.com/genshinsim/gcsim/pkg/core/player"
)

const burstHitmarks = 114

var (
	burstFrames []int
)

func init() {
	burstFrames = frames.InitAbilSlice(118)
}

func (c *char) Burst(p map[string]int) (action.Info, error) {
	ai := combat.AttackInfo{
		ActorIndex: c.Index,
		Abil:       "Balemoon Rising",
		AttackTag:  attacks.AttackTagElementalBurst,
		ICDTag:     attacks.ICDTagNone,
		ICDGroup:   attacks.ICDGroupDefault,
		StrikeType: attacks.StrikeTypeDefault,
		Element:    attributes.Pyro,
		Durability: 25,
		Mult:       burst[c.TalentLvlSkill()],
	}
	skillArea := combat.NewCircleHitOnTarget(c.Core.Combat.Player(), nil, 10)
	c.Core.QueueAttack(ai, skillArea, burstHitmarks, burstHitmarks)
	c.QueueCharTask(c.nourishingCinders, burstHitmarks+1)
	// add cooldown to sim
	c.SetCDWithDelay(action.ActionBurst, 15*60, 0)
	// use up energy
	c.ConsumeEnergy(8)

	return action.Info{
		Frames:          frames.NewAbilFunc(burstFrames),
		AnimationLength: burstFrames[action.InvalidAction],
		CanQueueAfter:   burstFrames[action.ActionSwap], // earliest cancel
		State:           action.BurstState,
	}, nil
}

func (c *char) nourishingCinders() {
	amt := c.CurrentHPDebt() + 3.0*c.getTotalAtk()
	c.Core.Player.Heal(player.HealInfo{
		Caller:  c.Index,
		Target:  c.Index,
		Message: "Nourishing Cinders",
		Src:     amt,
		Bonus:   c.Stat(attributes.Heal) + healMod, // cancel out the negative heal bonus we applied to her
	})
	c.ResetActionCooldown(action.ActionSkill)
}
