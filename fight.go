package gobots

import(
    "errors"
)

type Fight struct{

    CurrentBot *Bot
    NextBot *Bot
    MaxRounds int
    MaxRoundTurns int
    Turn int
    Round int
    RemainingTurns int
    Winner *Bot
}

func NewFight(first, second *Bot,
            max_rounds int, max_round_turns int) *Fight {

    f := &Fight{

        CurrentBot: first,
        NextBot: second,
        MaxRounds: max_rounds,
        MaxRoundTurns: max_round_turns,
        Round: 0,
        Turn: 0,
        Winner: nil,
    }

    return f
}

func (f *Fight) CurrentBotMove() *Move {

    return f.CurrentBot.CurrentMove()
}

func (f *Fight) ComputeTurn() *Turn {

    t := &Turn {
        Round: f.Round,
        Turn: f.Turn,
        CurrentBot: f.CurrentBot,
        NextBot: f.NextBot,
        CurrentMove: *f.CurrentBotMove(),
    }

    for _, p := range f.CurrentBot.PowerUps {

        p.Update(t)
    }

    for _, p := range f.NextBot.PowerUps {

        p.Update(t)
    }

    return t
}

func (f *Fight) PlayTurn(t *Turn) int {

    move := &t.CurrentMove

    last_res := 0
    for i := 0; i < move.MovesByTurn; i++ {
        last_res = f.NextBot.Hit(move)
    }
    return last_res
}

func (f *Fight) NextTurn() (*Bot, error) {

    f.Turn ++
    if f.Turn >= f.MaxRoundTurns {

        f.Round ++
        f.Turn = 0

        if f.Round >= f.MaxRounds {

            /* fight finished */
            if f.CurrentBot.Energy > f.NextBot.Energy {
                return f.CurrentBot, errors.New("current bot wins")
            } else if f.CurrentBot.Energy < f.NextBot.Energy {
                return f.NextBot, errors.New("other bot wins")
            } else {
                return nil, errors.New("tie")
            }

        } else {
            f.CurrentBot.Rest()
            f.NextBot.Rest()
        }
    }

    f.CurrentBot, f.NextBot = f.NextBot, f.CurrentBot

    return nil, nil
}

/*
 An example game loop implementation
*/
func (f *Fight) Loop() *Bot {

    var winner *Bot = nil
    var err error
    for winner == nil {

        t := f.ComputeTurn()

        if f.PlayTurn(t) <= 0 {
            return f.CurrentBot
        } else {
            winner, err = f.NextTurn()

            if err != nil {
                return winner
            }
        }
    }

    return winner
}
