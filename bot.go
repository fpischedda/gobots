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

type Bot struct{

    Name string
    MaxEnergy int
    Energy int
    MountedArmor Armor
    Strength int
    Defense int
    Speed int
    RestPerc int
    Moves []*Move
}

func (b *Bot) Rest() {

    b.Energy = min(b.MaxEnergy, b.Energy + b.MaxEnergy*b.RestPerc/100)
    b.MountedArmor.Repair(50)
}

func (b *Bot) Hit(move *Move) int {

    damage := b.Defense - b.MountedArmor.Hit(move)

    if damage < 0 {
        b.Energy += damage
    }

    return b.Energy
}

func (b *Bot) CurrentMove() *Move {

    return b.Moves[0]
}

type Armor struct {

    Name string
    Resistance int
    Damage int
}

func (a *Armor) Repair(repair_perc int) {

    a.Damage = max(0, a.Damage - a.Damage*repair_perc / 100)
}

func (a *Armor) Status() int {

    return (a.Resistance - a.Damage) * 100 / a.Resistance
}

func (a *Armor) Hit(move *Move) int {

    a.Damage += move.HitDamage

    real_damage := 0
    if a.Damage > a.Resistance {
        real_damage = a.Damage - a.Resistance
        a.Damage = a.Resistance
    }

    return real_damage
}

func (a *Armor) DefensePoints() int {

    return a.Resistance - a.Damage
}

type Move struct{

    Name string
    Type string
    HitDamage int
    Protection int
    TurnsToRecharge int //turns required to recharge this move
    MovesByTurn int //how many time this move can be used in a turn
}
