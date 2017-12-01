package diffiehellman

import (
	"math/big"
	"math/rand"
	"time"
)

func PrivateKey(p *big.Int) *big.Int {
	key := new(big.Int) //allocate
	seed := rand.New(rand.NewSource(time.Now().UnixNano()))

	//between 2 and P == 2+ [between 0 and p-2]
	limit := new(big.Int).Sub(p, big.NewInt(2)) //p-2
	key.Rand(seed, limit) //[between 0 and p-2]

	key.Add(key, big.NewInt(2))//[between 2 and P]

	return key
}

func PublicKey(private, p *big.Int, g int64) *big.Int{

	return new(big.Int).Exp(
		big.NewInt(g), 
		private, 
		p,
	) // A = g**private mod p	
}

func NewPair(p *big.Int, g int64) (*big.Int, *big.Int){
	
	private := PrivateKey(p)
	public := PublicKey(private, p, g)

	return private, public
}

func SecretKey(private1, public2, p *big.Int) *big.Int{
	
	return new(big.Int).Exp(
		public2,
		private1,
		p,
	)
}