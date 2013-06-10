package gobots

type PowerUp interface {

    Update(t *Turn)
    Activate()

    on_activate(t *Turn)
    on_run(t *Turn)
    on_clean(t *Turn)
}

type PowerUpItem struct {

    Name string
    status string
    DurationInTurns int
    turns_left int
    turns_to_recharge int
    RechargeInTurns int
}

func (p PowerUpItem) Activate() {

    p.status = "activated"
    p.turns_left = p.DurationInTurns
}

func (p PowerUpItem) on_activate(t *Turn) {
}

func (p PowerUpItem) on_run(t *Turn) {
}

func (p PowerUpItem) on_clean(t *Turn) {
}

/*
 A powerp can be in one of these states:
 - available
 - activated
 - running
 - recharging
*/
func (p PowerUpItem) Update(t *Turn) {

    switch p.status {

    case "activated":
        p.on_activate(t)
        p.status = "running"

    case "running":
        p.on_run(t)

        p.turns_left -= 1

        if p.turns_left <= 0 {

            p.status = "recharging"
            p.turns_to_recharge = p.RechargeInTurns
            p.on_clean(t)
        }
        break;

    case "recharging":
        p.turns_to_recharge -= 1

        if p.turns_to_recharge <= 0 {

            p.status = "available"
        }
        break;
    }
}

type PowerUp_DoubleDefense struct {
    PowerUpItem
    original int
    applyed_to_bot *Bot
}

func (p PowerUp_DoubleDefense) on_activate(t *Turn) {

    b := t.CurrentBot
    p.original = b.Defense
    p.applyed_to_bot = b
    b.Defense *= 2
}

func (p PowerUp_DoubleDefense) on_clean(t *Turn) {

    p.applyed_to_bot.Defense -= p.original
}

type PowerUp_HalveAttack struct {

    PowerUpItem
}

func (p PowerUp_HalveAttack) on_run(t *Turn) {

    t.CurrentMove.HitDamage /= 2
}
