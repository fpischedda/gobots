package gobots

type Bot struct{

    Name string
    Energy int
    MountedArmor Armor
    Strength int
    Defense int
    Speed int
    RestEnergy int
    Moves []*Move
}

func (b *Bot) Rest() {

    b.Energy += b.RestEnergy
}

func (b *Bot) Hit(move *Move) {

    damage := b.Defense - b.MountedArmor.Hit(move)

    if damage < 0 {
        b.Energy += damage
    }
}

func (b *Bot) CurrentMove() *Move {

    return b.Moves[0]
}

type Armor struct {

    Name string
    Resistance int
    Damage int
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
