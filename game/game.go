package main

import (
    "fmt"
    gobots "github.com/fpischedda/gobots"
)

func main() {

    var chronicle = make(chan GameChronicle)
    bot1 := RandomizeBot(Armors, Moves, "ciccio bot")
    bot2 := RandomizeBot(Armors, Moves, "pinottobot")
    f := gobots.NewFight(bot1, bot2, 2, 10)
    go game_loop(f, chronicle)

    for {
        c := <-chronicle
        c.Print()
        if c.MatchStatus != "running" {
            break;
        }
    }

    close(chronicle)
    panic("show me the leaks")
}

func game_loop(f *gobots.Fight, chronicle chan GameChronicle) {

    var err error
    for {

        damage := f.NextBot.Energy
        status := f.PlayTurn()
        damage = damage - f.NextBot.Energy

        if status <= 0 {
            chronicle <- NewChronicle(f, "current bot wins", damage)
            break;
        } else {
            _, err = f.NextTurn()

            if err != nil {
                chronicle <- NewChronicle(f, err.Error(), damage)
                break;
            }

            chronicle <- NewChronicle(f, "running", damage)
        }
    }
}

type GameChronicle struct {

    Action string
    ActionDamage int
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

func NewChronicle(f *gobots.Fight, status string, damage int) GameChronicle {

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
        ActionDamage: damage,
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

    fmt.Println("Action ", c.Action, "Damage", c.ActionDamage)

    fmt.Println("Other Bot ", c.NextBotInfo.Name,
        "Energy ", c.NextBotInfo.Energy,
        "Armor ", c.NextBotInfo.ArmorStatus)

    fmt.Println("Match status: ", c.MatchStatus)
}
