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
    Owner *Bot
    DurationInTurns int
    TurnsLeft int
    TurnsToRecharge int
    RechargeInTurns int
}

func (p *PowerUpItem) Activate() {

    p.status = "activated"
    p.TurnsLeft = p.DurationInTurns
}

func (p *PowerUpItem) on_activate(t *Turn) {
}

func (p *PowerUpItem) on_run(t *Turn) {
}

func (p *PowerUpItem) on_clean(t *Turn) {
}

/*
 A powerp can be in one of these states:
 - available
 - activated
 - running
 - recharging
*/
func (p *PowerUpItem) Update(t *Turn) {

    switch p.status {

    case "activated":
        p.on_activate(t)
        p.status = "running"

    case "running":
        p.on_run(t)

        p.TurnsLeft -= 1

        if p.TurnsLeft <= 0 {

            p.status = "recharging"
            p.TurnsToRecharge = p.RechargeInTurns
            p.on_clean(t)
        }
        break;

    case "recharging":
        p.TurnsToRecharge -= 1

        if p.TurnsToRecharge <= 0 {

            p.status = "available"
        }
        break;
    }
}

type PowerUp_DoubleDefense struct {
    PowerUpItem
    original int
}

func (p *PowerUp_DoubleDefense) on_activate(t *Turn) {

    p.original = p.Owner.Defense
    p.Owner.Defense *= 2
}

func (p *PowerUp_DoubleDefense) on_clean(t *Turn) {

    p.Owner.Defense = p.original
}

type PowerUp_HalveAttack PowerUpItem

func (p *PowerUp_HalveAttack) on_run(t *Turn) {

    t.CurrentMove.HitDamage /= 2
}
