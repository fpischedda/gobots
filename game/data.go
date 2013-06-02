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
    defense := 10 + rand.Intn(10)
    speed := 1 + rand.Intn(5)
    rest_perc := 30 + rand.Intn(30)
    return &gobots.Bot{
        Name: name,
        MaxEnergy: energy,
        Energy: energy,
        MountedArmor: armor,
        Strength: strength,
        Defense: defense,
        Speed: speed,
        RestPerc: rest_perc,
        Moves: moves,
    }
}

var Armors = []gobots.Armor{
    gobots.Armor{
        Name: "Base armor",
        Resistance: 10,
        Damage: 0,
    },
    gobots.Armor{
        Name: "Advanced armor",
        Resistance: 100,
        Damage: 0,
    },
}

var Moves = []*gobots.Move{
    &gobots.Move{
        Name: "Attack",
        Type: "Attack",
        HitDamage: 40,
        Protection: 0,
        TurnsToRecharge: 0,
        MovesByTurn: 1,
    },
    &gobots.Move{
        Name: "Attack2",
        Type: "Attack",
        HitDamage: 20,
        Protection: 0,
        TurnsToRecharge: 0,
        MovesByTurn: 2,
    },
}
