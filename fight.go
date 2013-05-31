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

func (f *Fight) PlayTurn() int {

    move := f.CurrentBotMove()

    return f.NextBot.Hit(move)
}

func (f *Fight) NextTurn() (*Bot, error) {

    f.Turn ++
    if f.Turn >= f.MaxRoundTurns {

        f.Round ++

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

    tmp := f.CurrentBot
    f.CurrentBot = f.NextBot
    f.NextBot = tmp

    return nil, nil
}

func (f *Fight) Loop() *Bot {

    var winner *Bot = nil
    var err error
    for winner == nil {

        if f.PlayTurn() <= 0 {
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
