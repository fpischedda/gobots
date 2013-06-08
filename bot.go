package gobots

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func gain(src, max_val, perc int) int {

    return min(max_val, src+max_val*perc/100)
}

type Bot struct{

    Name string
    MaxEnergy int
    Energy int
    MountedArmor Armor
    Strength int
    MaxDefense int
    Defense int
    Speed int
    RestPerc int
    Moves []*Move
}

func (b *Bot) Rest() {

    b.Energy = gain(b.Energy, b.MaxEnergy, b.RestPerc)
    b.Defense = gain(b.Defense, b.MaxDefense, b.RestPerc)
    b.MountedArmor.Repair(50)
}

/*
    every hit the bot defense decrease by half the damage
    that passes through the armor
*/
func (b *Bot) Hit(move *Move) int {

    damage := b.Defense - b.MountedArmor.Hit(move)

    if damage < 0 {
        b.Defense = max(0, b.Defense + damage/2)
        b.Energy += damage
    }

    return b.Energy
}

func (b *Bot) CurrentMove() *Move {

    return b.Moves[0]
}

type Armor struct {

    Name string
    MaxResistance int
    Resistance int
}

func (a *Armor) Repair(repair_perc int) {

    a.Resistance = gain(a.Resistance, a.MaxResistance, repair_perc)
}

func (a *Armor) Status() int {

    return a.Resistance * 100 / a.MaxResistance
}

func (a *Armor) Hit(move *Move) int {

    a.Resistance -= move.HitDamage

    damage := 0
    if a.Resistance < 0 {
        damage = -a.Resistance
        a.Resistance = 0
    }

    return damage
}

func (a *Armor) DefensePoints() int {

    return a.Resistance
}

type Move struct{

    Name string
    Type string
    HitDamage int
    Protection int
    TurnsToRecharge int //turns required to recharge this move
    MovesByTurn int //how many time this move can be used in a turn
}
