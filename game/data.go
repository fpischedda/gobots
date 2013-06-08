package main

import(
    "math/rand"
    gobots "github.com/fpischedda/gobots"
)

func RandomizeBot(armors []gobots.Armor, moves []*gobots.Move,
    name string) *gobots.Bot {

    armor := Armors[rand.Intn(len(Armors)-1)]
    energy := 30 + rand.Intn(10)
    strength := 10 + rand.Intn(5)
    defense := 5 + rand.Intn(5)
    speed := 1 + rand.Intn(5)
    rest_perc := 30 + rand.Intn(30)
    return &gobots.Bot{
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
    }
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
