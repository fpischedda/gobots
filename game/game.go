package main

import (
    "fmt"
    gobots "github.com/fpischedda/gobots"
)

func main() {

    armors := []gobots.Armor{
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

    moves := []*gobots.Move{
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

    bots := []*gobots.Bot{
        &gobots.Bot{
            Name: "bot1",
            Energy: 50,
            MountedArmor: armors[0],
            Strength: 10,
            Defense: 2,
            Speed: 5,
            RestEnergy: 1,
            Moves: moves,
        },
        &gobots.Bot{
            Name: "bot2",
            Energy: 30,
            MountedArmor: armors[1],
            Strength: 20,
            Defense: 5,
            Speed: 1,
            RestEnergy: 3,
            Moves: moves,
        },
    }

    f := gobots.NewFight(bots[0], bots[1], 1, 10)
    winner := f.Loop()
    fmt.Println("and the winner is ", winner.Name)
    panic("show me the leaks")
}
