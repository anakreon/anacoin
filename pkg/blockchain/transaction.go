package blockchain

type Transaction struct {
	In  int
	Out int
}

func (transaction Transaction) IsValid() bool {
	return true
}
