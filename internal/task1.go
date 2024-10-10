package internal

type KeyExpander interface {
	expandKey(key []byte) error
}

type Encryptor interface {
	Encrypt(block []byte, roundKeys [][]byte) ([]byte, error)
}

type Decryptor interface {
	Decrypt(block []byte, roundKeys [][]byte) ([]byte, error)
}

type Cypherer interface {
	KeyExpander
	Encryptor
	Decryptor
}

const (
	ecb = iota
	cbc
	pcbc
	cfb
	ofb
	ctr
	randomDelta
)

const (
	zeros = iota
	ansiX923
	pkcs7
	iso10126
)

type Task1Cypherer struct {
	cryptoType  int
	paddingType int
	roundKeys   [][]byte
}

func NewTask1Cypherer(key []byte, cryptoType int, paddingType int, initVect []byte, options ...any) *Task1Cypherer {
	t1c := Task1Cypherer{
		cryptoType:  cryptoType,
		paddingType: paddingType,
	}

	t1c.expandKey(key)

	return &t1c
}

func (t1c *Task1Cypherer) expandKey(key []byte) ([][]byte, error) {
	// TODO
	return nil, nil
}

func (t1c *Task1Cypherer) Encrypt(block []byte, roundKeys [][]byte) ([]byte, error) {
	// TODO
	return nil, nil
}

func (t1c *Task1Cypherer) Decrypt(block []byte, roundKeys [][]byte) ([]byte, error) {
	// TODO
	return nil, nil
}
