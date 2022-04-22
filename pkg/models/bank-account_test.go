package models

import (
	"os"
	"testing"

	"github.com/MonduCareers/-Nwuguru-Sunday-Coding-Challenge/pkg/utils"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	utils.TruncateDBTables()
	code := m.Run()
	utils.TruncateDBTables()
	os.Exit(code)
}

func TestCreateBankAccount(t *testing.T) {
	t.Run("cannot create bank account without required fields", func(t *testing.T) {
		NewBankAccount := BankAccount{}
		_, err := NewBankAccount.CreateBankAccount()

		require.EqualError(t, err, "name can't be blank; email can't be blank; username can't be blank")
	})

	t.Run("can create bank account", func(t *testing.T) {
		NewBankAccount := BankAccount{Name: "test", Username: "guru", Email: "test@example.com"}
		accountDetails, err := NewBankAccount.CreateBankAccount()
		require.NoError(t, err)
		require.NotEqual(t, 0, accountDetails.ID)
	})

	t.Run("cannot create bank account with incorrect email format", func(t *testing.T) {
		NewBankAccount := BankAccount{Name: "test2", Username: "guru2", Email: "testexample.com"}
		_, err := NewBankAccount.CreateBankAccount()
		require.EqualError(t, err, "invalid email address")
	})

	t.Run("cannot create bank account with duplicate username", func(t *testing.T) {
		NewBankAccount := BankAccount{Name: "test", Username: "guru", Email: "testexample.com"}
		_, err := NewBankAccount.CreateBankAccount()
		require.EqualError(t, err, "username already exists")
	})

}

func TestCreateBankAccountTransaction(t *testing.T) {
	NewBankAccount := BankAccount{Name: "test", Username: "guru", Email: "test@example.com"}
	accountDetails, _ := NewBankAccount.CreateBankAccount()

	t.Run("cannot create transaction without bank account", func(t *testing.T) {
		NewBankAccountTransaction := BankAccountTransaction{}
		_, err := NewBankAccountTransaction.CreateBankAccountTransaction(utils.TransactionTypeCredit, "fakeuser")
		require.EqualError(t, err, "invalid username")
	})

	t.Run("cannot create transaction without valid amount", func(t *testing.T) {
		t.Run("when amount is zero", func(t *testing.T) {
			NewBankAccountTransaction := BankAccountTransaction{Amount: 0}
			_, err := NewBankAccountTransaction.CreateBankAccountTransaction(utils.TransactionTypeCredit, accountDetails.Username)

			require.EqualError(t, err, "amount must be greater than zero")
		})
		t.Run("when amount is negative", func(t *testing.T) {
			NewBankAccountTransaction := BankAccountTransaction{Amount: -10}
			_, err := NewBankAccountTransaction.CreateBankAccountTransaction(utils.TransactionTypeCredit, accountDetails.Username)

			require.EqualError(t, err, "amount must be greater than zero")
		})
    })

	t.Run("cannot create debit transaction with amount greater than account balance ", func(t *testing.T) {
		NewBankAccountTransaction := BankAccountTransaction{Amount: 10}
		_, err := NewBankAccountTransaction.CreateBankAccountTransaction(utils.TransactionTypeDebit, accountDetails.Username)

		require.EqualError(t, err, "balance insufficient")
	})

	t.Run("cannot create credit transaction ", func(t *testing.T) {
		t.Run("with only amount", func(t *testing.T) {
			NewBankAccountTransaction := BankAccountTransaction{Amount: 10}
			_, err := NewBankAccountTransaction.CreateBankAccountTransaction(utils.TransactionTypeCredit, accountDetails.Username)

			require.NoError(t, err)
		})
		t.Run("with amount and note", func(t *testing.T) {
			NewBankAccountTransaction := BankAccountTransaction{Amount: 10, Note: "test"}
			_, err := NewBankAccountTransaction.CreateBankAccountTransaction(utils.TransactionTypeCredit, accountDetails.Username)

			require.NoError(t, err)
		})
	})

	t.Run("cannot create debit transaction ", func(t *testing.T) {
		t.Run("with only amount", func(t *testing.T) {
			NewBankAccountTransaction := BankAccountTransaction{Amount: 10}
			d, err := NewBankAccountTransaction.CreateBankAccountTransaction(utils.TransactionTypeDebit, accountDetails.Username)

			require.NoError(t, err)
			require.Equal(t, d.Amount, int32(-10))
		})
		t.Run("with amount and note", func(t *testing.T) {
			NewBankAccountTransaction := BankAccountTransaction{Amount: 10, Note: "test"}
			d, err := NewBankAccountTransaction.CreateBankAccountTransaction(utils.TransactionTypeDebit, accountDetails.Username)

			require.NoError(t, err)
			require.Equal(t, d.Amount, int32(-10))
		})
	})

}

func TestGetBankAccountBalance(t *testing.T) {
	NewBankAccount := BankAccount{Name: "test", Username: "guru", Email: "test@example.com"}
	accountDetails, _ := NewBankAccount.CreateBankAccount()

	t.Run("zero account balance if no account", func(t *testing.T) {
		balance := GetBankAccountBalance("fakeuser")
		require.Equal(t, balance, int32(0))
	})

	t.Run("zero account balance if no transactions", func(t *testing.T) {
		balance := GetBankAccountBalance(accountDetails.Username)
		require.Equal(t, balance, int32(0))
	})

	t.Run("account balance is calculated correctly", func(t *testing.T) {
		NewBankAccountTransaction1 := BankAccountTransaction{Amount: 10}
		NewBankAccountTransaction2 := BankAccountTransaction{Amount: 10}
		NewBankAccountTransaction3 := BankAccountTransaction{Amount: 10}
		NewBankAccountTransaction4 := BankAccountTransaction{Amount: 10}
		NewBankAccountTransaction5 := BankAccountTransaction{Amount: 10}
		NewBankAccountTransaction6 := BankAccountTransaction{Amount: 10}
		NewBankAccountTransaction1.CreateBankAccountTransaction(utils.TransactionTypeCredit, accountDetails.Username)
		NewBankAccountTransaction2.CreateBankAccountTransaction(utils.TransactionTypeCredit, accountDetails.Username)
		NewBankAccountTransaction3.CreateBankAccountTransaction(utils.TransactionTypeCredit, accountDetails.Username)
		NewBankAccountTransaction4.CreateBankAccountTransaction(utils.TransactionTypeCredit, accountDetails.Username)
		NewBankAccountTransaction5.CreateBankAccountTransaction(utils.TransactionTypeDebit, accountDetails.Username)
		NewBankAccountTransaction6.CreateBankAccountTransaction(utils.TransactionTypeDebit, accountDetails.Username)
		balance := GetBankAccountBalance(accountDetails.Username)
		require.Equal(t, balance, int32(20))
	})
}