package banking

import "errors"

type Account struct {
	owner   string
	balance int
}

func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// 함수명 앞의 account는 메서드의 리시버라고 한다.
// go에선 struct 에 바로 함수를 정의할수 없어서 이렇게 정의한다.
// 메서드는 리시버를 통해서 함수를 호출하여 객체지향 프래그래밍과 유사한 형태를 가진다.
// 또한 리시버의 타입이 포인터 타입이면 객체지향 프래그래밍에서 참조에 의한 호출이 된다.
func (account *Account) AddDeposit(amount int) int {
	account.balance += amount
	return account.balance
}

func (account Account) GetBalance() int {
	return account.balance
}

func (account Account) GetOwner() string {
	return account.owner
}

func (account *Account) Withdraw(amount int) error {
	if account.balance < amount {
		return errors.New("can't withdraw, you are poor")
	}

	account.balance -= amount
	return nil
}

func (account *Account) ChangeOwner(newOwner string) {
	account.owner = newOwner
}

// String 메서드는 객체를 문자열로 표현할 때 호출되는 메서드이다. 
//미리 정의된 메서드라 이 함수를 재정의하면 fmt.PrintIn 으로 출력할 때 재정의한 내용이 출력된다.
func (account *Account) String() string {
	return "this is Override Account String Method"
}