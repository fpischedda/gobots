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

    var chronicle = make(chan GameChronicle)
    f := gobots.NewFight(bots[0], bots[1], 1, 10)
    go game_loop(f, chronicle)

    for {
        c := <-chronicle
        c.Print()
        if c.MatchStatus != "running" {
            break;
        }
    }

    panic("show me the leaks")
}

func game_loop(f *gobots.Fight, chronicle chan GameChronicle) {

    var err error
    for {

        status := f.PlayTurn()

        if status <= 0 {
            chronicle <- NewChronicle(f, "current bot wins")
            break;
        } else {
            _, err = f.NextTurn()

            if err != nil {
                chronicle <- NewChronicle(f, err.Error())
                break;
            }

            chronicle <- NewChronicle(f, "running")
        }
    }
}

type GameChronicle struct {

    Action string
    MatchStatus string
    CurrentBotInfo BotInfo
    NextBotInfo BotInfo
    Turn int
    Round int
}

type BotInfo struct {

    Name string
    Energy int
    ArmorStatus int
}

func NewChronicle(f *gobots.Fight, status string) GameChronicle {

    action := f.CurrentBotMove()

    c_bot := BotInfo {
        Name: f.CurrentBot.Name,
        Energy: f.CurrentBot.Energy,
        ArmorStatus: f.CurrentBot.MountedArmor.Status(),
    }

    n_bot := BotInfo {
        Name: f.NextBot.Name,
        Energy: f.NextBot.Energy,
        ArmorStatus: f.NextBot.MountedArmor.Status(),
    }

    return GameChronicle{
        Action: action.Name,
        MatchStatus: status,
        CurrentBotInfo: c_bot,
        NextBotInfo: n_bot,
        Turn: f.Turn,
        Round: f.Round,
    }
}

func (c *GameChronicle) Print() {

    fmt.Println("Round: ", c.Round, "Turn", c.Turn)
    fmt.Println("Active Bot ", c.CurrentBotInfo.Name,
        "Energy ", c.CurrentBotInfo.Energy,
        "Armor ", c.CurrentBotInfo.ArmorStatus)

    fmt.Println("Action ", c.Action)

    fmt.Println("Other Bot ", c.NextBotInfo.Name,
        "Energy ", c.NextBotInfo.Energy,
        "Armor ", c.NextBotInfo.ArmorStatus)

    fmt.Println("Match status: ", c.MatchStatus)
}
