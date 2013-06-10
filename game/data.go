package main

import(
    "math/rand"
    gobots "github.com/fpischedda/gobots"
)

func RandomizeBot(armors []gobots.Armor, moves []*gobots.Move,
    powerups []gobots.PowerUp,
    name string) *gobots.Bot {

    armor := Armors[rand.Intn(len(armors)-1)]
    energy := 30 + rand.Intn(10)
    strength := 10 + rand.Intn(5)
    defense := 5 + rand.Intn(5)
    speed := 1 + rand.Intn(5)
    rest_perc := 30 + rand.Intn(30)
    powerup := powerups[rand.Intn(len(powerups))]

    bot := &gobots.Bot{
        Name: name,
        MaxEnergy: energy,
        Energy: energy,
        MountedArmor: armor,
        Strength: strength,
        MaxDefense: defense,
        Defense: defense,
        Speed: speed,
        RestPerc: rest_perc,
        Moves: moves,
        PowerUps: []gobots.PowerUp { powerup },
    }

    return bot
}

var Armors = []gobots.Armor{
    gobots.Armor{
        Name: "Base armor",
        MaxResistance: 10,
        Resistance: 10,
    },
    gobots.Armor{
        Name: "Advanced armor",
        MaxResistance: 30,
        Resistance: 30,
    },
}

var Moves = []*gobots.Move{
    &gobots.Move{
        Name: "Attack",
        Type: "Attack",
        HitDamage: 12,
        Protection: 0,
        TurnsToRecharge: 0,
        MovesByTurn: 1,
    },
    &gobots.Move{
        Name: "Attack2",
        Type: "Attack",
        HitDamage: 14,
        Protection: 0,
        TurnsToRecharge: 0,
        MovesByTurn: 2,
    },
}

var PowerUps = []gobots.PowerUp {

    gobots.PowerUp_DoubleDefense {
        PowerUpItem: gobots.PowerUpItem{
            Name: "Double Defense",
            DurationInTurns: 1,
            RechargeInTurns: 3,
        },
    },

    gobots.PowerUp_HalveAttack {
        gobots.PowerUpItem{
            Name: "Halve Attack",
            DurationInTurns: 1,
            RechargeInTurns: 3,
        },
    },
}
