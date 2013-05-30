package gobots

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

func (f *Fight) NextTurn() *Bot {

    f.Turn ++
    if f.Turn >= f.MaxRoundTurns {

        f.Round ++

        /* fight finished */
        if f.Round >= f.MaxRounds {

            if f.CurrentBot.Energy >= f.NextBot.Energy {
                return f.CurrentBot
            } else {
                return f.NextBot
            }
        } else {
            f.CurrentBot.Rest()
            f.NextBot.Rest()
        }
    }

    tmp := f.CurrentBot
    f.CurrentBot = f.NextBot
    f.NextBot = tmp

    return nil
}

func (f *Fight) Loop() *Bot {

    var winner *Bot = nil
    for winner == nil {

        move := f.CurrentBotMove()

        f.NextBot.Hit(move)
        winner = f.NextTurn()
    }

    return winner
}
