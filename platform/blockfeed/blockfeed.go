package blockfeed

type Getter interface {
	GetAll() []Block
}

type Adder interface {
	Add(block Block)
}

type Block struct {
	Height  string `json:"number"`
	Minedby string `json:"minedby"`
}

type DB struct {
	Blocks []Block
}

func New() *DB {
	return &DB{
		Blocks: []Block{},
	}
}

func (r *DB) Add(block Block) {
	r.Blocks = append(r.Blocks, block)
}

func (r *DB) GetAll() []Block {
	return r.Blocks
}
