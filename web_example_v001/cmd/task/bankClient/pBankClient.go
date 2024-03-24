package bankClient

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type ClientInfo struct {
	ID        int
	FirstName string
	LastName  string
}

type Transfer struct {
	Amount   int
	Time     time.Time
	Success  bool
	Messages string
	Type     int
}

type BankClientImpl struct {
	AccountBalance  int
	ClientInfo      ClientInfo
	TransferHistory []Transfer
	mu              sync.Mutex
}

var bankClientCurr *BankClientImpl

type BankClient interface {
	Deposit(amount int)
	Withdrawal(amount int) error
	Balance() int
	GetClientInfo() ClientInfo
	GetTransferHistory() []Transfer
	AddBankClientData()
	CopyBankClient() BankClientImpl
	ClearAccountBalance()
}

func (bc *BankClientImpl) Deposit(amount int) {
	bc.mu.Lock()
	defer bc.mu.Unlock()
	bc.AccountBalance += amount
	bc.TransferHistory = append(bc.TransferHistory, Transfer{
		Amount:   amount,
		Time:     time.Now(),
		Success:  true,
		Messages: "Deposit successful",
		Type:     1,
	})
}

func (bc *BankClientImpl) Withdrawal(amount int) error {
	bc.mu.Lock()
	defer bc.mu.Unlock()
	if bc.AccountBalance < amount {
		bc.TransferHistory = append(bc.TransferHistory, Transfer{
			Amount:   amount,
			Time:     time.Now(),
			Success:  false,
			Messages: "Error withdrawal: Insufficient balance",
			Type:     3,
		})
		return fmt.Errorf("insufficient balance")
	}
	bc.AccountBalance -= amount
	bc.TransferHistory = append(bc.TransferHistory, Transfer{
		Amount:   amount,
		Time:     time.Now(),
		Success:  true,
		Messages: "Withdrawal successful",
		Type:     2,
	})
	return nil
}

func (bc *BankClientImpl) Balance() int {
	bc.mu.Lock()
	defer bc.mu.Unlock()
	return bc.AccountBalance
}

func (bc *BankClientImpl) GetClientInfo() ClientInfo {
	bc.mu.Lock()
	defer bc.mu.Unlock()
	return bc.ClientInfo
}

func (bc *BankClientImpl) GetTransferHistory() []Transfer {
	bc.mu.Lock()
	defer bc.mu.Unlock()
	return bc.TransferHistory
}

func (bc *BankClientImpl) ClearAccountBalance() {
	bc.mu.Lock()
	defer bc.mu.Unlock()
	bc.AccountBalance = 0
	bc.TransferHistory = nil
}

func (bc *BankClientImpl) AddBankClientData() {
	bc.ClearAccountBalance()

	for i := 0; i < 10; i++ {
		go func(id int) {
			timeRand := time.Duration(rand.Intn(500)+500) * time.Millisecond
			amount := rand.Intn(10) + 1
			time.Sleep(timeRand)
			bc.Deposit(amount)
		}(i)
	}

	for i := 0; i < 5; i++ {
		go func(id int) {
			timeRand := time.Duration(rand.Intn(500)+500) * time.Millisecond
			amount := rand.Intn(5) + 1
			time.Sleep(timeRand)
			err := bc.Withdrawal(amount)
			if err != nil {
				fmt.Println(err)
			}
		}(i)
	}
}

func (bc *BankClientImpl) CopyBankClient() BankClientImpl {
	bc.mu.Lock()
	defer bc.mu.Unlock()

	bankClientCopy := BankClientImpl{
		AccountBalance:  bc.AccountBalance,
		ClientInfo:      bc.ClientInfo,
		TransferHistory: make([]Transfer, len(bc.TransferHistory)),
		mu:              sync.Mutex{},
	}

	for i, transfer := range bc.TransferHistory {
		bankClientCopy.TransferHistory[len(bc.TransferHistory)-1-i] = Transfer{
			Amount:   transfer.Amount,
			Time:     transfer.Time,
			Success:  transfer.Success,
			Messages: transfer.Messages,
			Type:     transfer.Type,
		}
	}

	return bankClientCopy
}

func BankClientOperation(nameOperation string, amount int) (BankClientImpl, string) {

	if bankClientCurr == nil {
		bankClientCurr = &BankClientImpl{
			ClientInfo: ClientInfo{
				ID:        1,
				FirstName: "John",
				LastName:  "Doe",
			},
		}
	}

	switch nameOperation {
	case "balance":
		bankClientCurr.Balance()
		mess := fmt.Sprintf("Balance: %d", bankClientCurr.Balance())
		return bankClientCurr.CopyBankClient(), mess
	case "deposit":
		bankClientCurr.Deposit(amount)
		mess := fmt.Sprintf("Deposit: %d", amount)
		return bankClientCurr.CopyBankClient(), mess
	case "withdrawal":
		err := bankClientCurr.Withdrawal(amount)
		mess := fmt.Sprintf("Withdrawal: %d", amount)
		if err != nil {
			mess = fmt.Sprintf("Withdrawal: %d, Error: %v", amount, err)
		}
		return bankClientCurr.CopyBankClient(), mess
	case "addRandomData":
		bankClientCurr.AddBankClientData()
		mess := "addRandomData"
		time.Sleep(2 * time.Second)
		return bankClientCurr.CopyBankClient(), mess
	default:
		return bankClientCurr.CopyBankClient(), "Unsupported command. You can use commands: balance, deposit, withdrawal, addRandomData"
	}

}
