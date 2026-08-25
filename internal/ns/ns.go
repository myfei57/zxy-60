package ns

type Namespace struct {
	terminal string
}

func New(terminal string) *Namespace {
	return &Namespace{terminal: terminal}
}

func (n *Namespace) Terminal() string {
	return n.terminal
}

func (n *Namespace) BagID(barcode string) string {
	return n.terminal + "-" + barcode
}
