package combat

import "github.com/genshinsim/gsim/pkg/def"

type eHook struct {
	f   func(s def.Sim) bool
	key string
	src int
}

//AddHook adds a hook to sim. Hook will be called based on the type of hook
func (s *Sim) AddEventHook(f func(s def.Sim) bool, key string, hook def.EventHookType) {

	a := s.eventHooks[hook]

	//check if override first
	ind := len(a)
	for i, v := range a {
		if v.key == key {
			ind = i
		}
	}
	if ind != 0 && ind != len(a) {
		s.log.Debugw("hook added", "frame", s.f, "event", def.LogHookEvent, "overwrite", true, "key", key, "type", hook)
		a[ind] = eHook{
			f:   f,
			key: key,
			src: s.f,
		}
	} else {
		a = append(a, eHook{
			f:   f,
			key: key,
			src: s.f,
		})
		s.log.Debugw("hook added", "frame", s.f, "event", def.LogHookEvent, "overwrite", true, "key", key, "type", hook)
	}
	s.eventHooks[hook] = a
}

func (s *Sim) executeEventHook(t def.EventHookType) {
	n := 0
	for i, v := range s.eventHooks[t] {
		if v.f(s) {
			s.log.Debugw("event hook ended", "frame", s.f, "event", def.LogHookEvent, "key", i, "src", v.src)

		} else {
			s.eventHooks[t][n] = v
			n++
		}
	}
	s.eventHooks[t] = s.eventHooks[t][:n]
}

type attackWillLandHook struct {
	f   func(t def.Target, ds *def.Snapshot)
	key string
	src int
}

func (s *Sim) AddOnAttackWillLand(f func(t def.Target, ds *def.Snapshot), key string) {
	//check if override first
	ind := -1
	for i, v := range s.onAttackWillLand {
		if v.key == key {
			ind = i
		}
	}
	if ind != -1 {
		s.log.Debugw("on attack landed hook added", "frame", s.f, "event", def.LogHookEvent, "overwrite", true, "key", key)
		s.onAttackWillLand[ind] = attackWillLandHook{
			f:   f,
			key: key,
			src: s.f,
		}
	} else {
		s.onAttackWillLand = append(s.onAttackWillLand, attackWillLandHook{
			f:   f,
			key: key,
			src: s.f,
		})
		s.log.Debugw("hook added", "frame", s.f, "event", def.LogHookEvent, "overwrite", true, "key", key)
	}
}
func (s *Sim) OnAttackWillLand(t def.Target, ds *def.Snapshot) {
	for _, v := range s.onAttackWillLand {
		v.f(t, ds)
	}
}

type attackLandedHook struct {
	f   func(t def.Target, ds *def.Snapshot, dmg float64, crit bool)
	key string
	src int
}

func (s *Sim) OnAttackLanded(t def.Target, ds *def.Snapshot, dmg float64, crit bool) {
	for _, v := range s.onAttackLanded {
		v.f(t, ds, dmg, crit)
	}
}

func (s *Sim) AddOnAttackLanded(f func(t def.Target, ds *def.Snapshot, dmg float64, crit bool), key string) {
	//check if override first
	ind := -1
	for i, v := range s.onAttackLanded {
		if v.key == key {
			ind = i
		}
	}
	if ind != -1 {
		s.log.Debugw("on attack landed hook added", "frame", s.f, "event", def.LogHookEvent, "overwrite", true, "key", key)
		s.onAttackLanded[ind] = attackLandedHook{
			f:   f,
			key: key,
			src: s.f,
		}
	} else {
		s.onAttackLanded = append(s.onAttackLanded, attackLandedHook{
			f:   f,
			key: key,
			src: s.f,
		})
		s.log.Debugw("hook added", "frame", s.f, "event", def.LogHookEvent, "overwrite", true, "key", key)
	}
}

type onReactionHook struct {
	f   func(t def.Target, ds *def.Snapshot)
	key string
	src int
}

func (s *Sim) AddOnReaction(f func(t def.Target, ds *def.Snapshot), key string) {
	//check if override first
	ind := -1
	for i, v := range s.onReaction {
		if v.key == key {
			ind = i
		}
	}
	if ind != -1 {
		s.log.Debugw("on attack landed hook added", "frame", s.f, "event", def.LogHookEvent, "overwrite", true, "key", key)
		s.onReaction[ind] = onReactionHook{
			f:   f,
			key: key,
			src: s.f,
		}
	} else {
		s.onReaction = append(s.onReaction, onReactionHook{
			f:   f,
			key: key,
			src: s.f,
		})
		s.log.Debugw("hook added", "frame", s.f, "event", def.LogHookEvent, "overwrite", true, "key", key)
	}
}

func (s *Sim) OnReaction(t def.Target, ds *def.Snapshot) {
	if ds.ReactionType != def.NoReaction {
		s.stats.ReactionsTriggered[ds.ReactionType]++
	}
	for _, v := range s.onReaction {
		v.f(t, ds)
	}
}

type onReactionDamageHook struct {
	f   func(t def.Target, ds *def.Snapshot)
	key string
	src int
}

func (s *Sim) AddOnAmpReaction(f func(t def.Target, ds *def.Snapshot), key string) {
	//check if override first
	ind := -1
	for i, v := range s.onAmpReaction {
		if v.key == key {
			ind = i
		}
	}
	if ind != -1 {
		s.log.Debugw("on attack landed hook added", "frame", s.f, "event", def.LogHookEvent, "overwrite", true, "key", key)
		s.onAmpReaction[ind] = onReactionDamageHook{
			f:   f,
			key: key,
			src: s.f,
		}
	} else {
		s.onAmpReaction = append(s.onAmpReaction, onReactionDamageHook{
			f:   f,
			key: key,
			src: s.f,
		})
		s.log.Debugw("hook added", "frame", s.f, "event", def.LogHookEvent, "overwrite", true, "key", key)
	}
}

func (s *Sim) OnAmpReaction(t def.Target, ds *def.Snapshot) {
	for _, v := range s.onAmpReaction {
		v.f(t, ds)
	}
}

func (s *Sim) AddOnTransReaction(f func(t def.Target, ds *def.Snapshot), key string) {
	//check if override first
	ind := -1
	for i, v := range s.onTransReaction {
		if v.key == key {
			ind = i
		}
	}
	if ind != -1 {
		s.log.Debugw("on attack landed hook added", "frame", s.f, "event", def.LogHookEvent, "overwrite", true, "key", key)
		s.onTransReaction[ind] = onReactionDamageHook{
			f:   f,
			key: key,
			src: s.f,
		}
	} else {
		s.onTransReaction = append(s.onTransReaction, onReactionDamageHook{
			f:   f,
			key: key,
			src: s.f,
		})
		s.log.Debugw("hook added", "frame", s.f, "event", def.LogHookEvent, "overwrite", true, "key", key)
	}
}

func (s *Sim) OnTransReaction(t def.Target, ds *def.Snapshot) {
	for _, v := range s.onTransReaction {
		v.f(t, ds)
	}
}

func (s *Sim) AddInitHook(f func()) {
	s.initHooks = append(s.initHooks, f)
}

type defeatHook struct {
	f   func(t def.Target)
	key string
	src int
}

func (s *Sim) OnTargetDefeated(t def.Target) {
	for _, v := range s.onTargetDefeated {
		v.f(t)
	}
}
func (s *Sim) AddOnTargetDefeated(f func(t def.Target), key string) {
	//check if override first
	ind := -1
	for i, v := range s.onTargetDefeated {
		if v.key == key {
			ind = i
		}
	}
	if ind != -1 {
		s.log.Debugw("on attack landed hook added", "frame", s.f, "event", def.LogHookEvent, "overwrite", true, "key", key)
		s.onTargetDefeated[ind] = defeatHook{
			f:   f,
			key: key,
			src: s.f,
		}
	} else {
		s.onTargetDefeated = append(s.onTargetDefeated, defeatHook{
			f:   f,
			key: key,
			src: s.f,
		})
		s.log.Debugw("hook added", "frame", s.f, "event", def.LogHookEvent, "overwrite", true, "key", key)
	}
}
