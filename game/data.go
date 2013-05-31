package main

import(
    gobots "github.com/fpischedda/gobots"
)

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
        HitDamage: 5,
        Protection: 0,
        TurnsToRecharge: 0,
        MovesByTurn: 1,
    },
    &gobots.Move{
        Name: "Attack2",
        Type: "Attack",
        HitDamage: 2,
        Protection: 0,
        TurnsToRecharge: 0,
        MovesByTurn: 2,
    },
}

var Bots = []*gobots.Bot{
    &gobots.Bot{
        Name: "bot1",
        Energy: 50,
        MountedArmor: Armors[0],
        Strength: 10,
        Defense: 2,
        Speed: 5,
        RestEnergy: 1,
        Moves: Moves,
    },
    &gobots.Bot{
        Name: "bot2",
        Energy: 30,
        MountedArmor: Armors[1],
        Strength: 20,
        Defense: 5,
        Speed: 1,
        RestEnergy: 3,
        Moves: Moves,
    },
}
