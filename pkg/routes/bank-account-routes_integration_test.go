package routes

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/MonduCareers/-Nwuguru-Sunday-Coding-Challenge/pkg/models"
	"github.com/MonduCareers/-Nwuguru-Sunday-Coding-Challenge/pkg/utils"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	utils.TruncateDBTables()
	code := m.Run()
	utils.TruncateDBTables()
	os.Exit(code)
}

func TestHomeRoute(t *testing.T) {
	url := "http://localhost:8080/"

	t.Run("Can load homepage", func(t *testing.T) {
		resp, err := http.Get(url)
		require.NoError(t,  err)
		require.Equal(t, resp.StatusCode, http.StatusOK)
		defer resp.Body.Close()
    	data, _ := ioutil.ReadAll(resp.Body)
		require.Equal(t, string(data), "Welcome Home!")
	})
}

func TestBankAcountRoute(t *testing.T) {
	utils.TruncateDBTables()
	url := "http://localhost:8080/bank-account"
	
	t.Run("Can create bank account", func(t *testing.T) {
		payload := `
			{
				"name": "test",
				"email": "test@example.com",
				"username": "guru"
			}`
		resp, err := http.Post(url, "application/json", strings.NewReader(payload))
		NewBankAccount := &models.BankAccount{}
		utils.ParseContent(resp, NewBankAccount)
		require.NoError(t,  err)
		require.Equal(t, resp.StatusCode, http.StatusOK)
		
		require.Equal(t, NewBankAccount.Username, "guru")
	})

	t.Run("Cannot create bank account", func(t *testing.T) {
		NewBankAccount := models.BankAccount{Name: "test", Username: "guru", Email: "test@example.com"}
		accountDetails, _ := NewBankAccount.CreateBankAccount()
		var Info struct{Message string `json:"message"`}

		t.Run("when username is duplicate", func(t *testing.T) {
			payload := fmt.Sprintf(`
				{
					"name": "test",
					"email": "test@example.com",
					"username": "%s"
				}`, accountDetails.Username)
			resp, _ := http.Post(url, "application/json", strings.NewReader(payload))
			utils.ParseContent(resp, &Info)
			require.Equal(t, resp.StatusCode, http.StatusBadRequest)
			require.Equal(t, Info.Message, "username already exists")
		})

		t.Run("when email format is incorrect", func(t *testing.T) {
			payload := `
				{
					"name": "test",
					"email": "testexample.com",
					"username": "tester"
				}`
			resp, _ := http.Post(url, "application/json", strings.NewReader(payload))
			utils.ParseContent(resp, &Info)
			require.Equal(t, resp.StatusCode, http.StatusBadRequest)
			require.Equal(t, Info.Message, "invalid email address")
		})

		t.Run("when missing required field", func(t *testing.T) {
			payload := `
				{
					"email": "test@example.com",
					"username": "tester"
				}`
			resp, _ := http.Post(url, "application/json", strings.NewReader(payload))
			utils.ParseContent(resp, &Info)
			require.Equal(t, resp.StatusCode, http.StatusBadRequest)
			require.Equal(t, Info.Message, "name can't be blank")
		})
	})
}

func TestDepositRoute(t *testing.T) {
	utils.TruncateDBTables()
	url := "http://localhost:8080/bank-account/guru/deposit"
	NewBankAccount := models.BankAccount{Name: "test", Username: "guru", Email: "test@example.com"}
	NewBankAccount.CreateBankAccount()

	t.Run("Can deposit", func(t *testing.T) {
		payload := `
			{
				"amount": 10,
				"note": "test"
			}`
		resp, err := http.Post(url, "application/json", strings.NewReader(payload))
		NewBankAccountTransaction := &models.BankAccountTransaction{}
		utils.ParseContent(resp, NewBankAccountTransaction)
		require.NoError(t,  err)
		require.Equal(t, resp.StatusCode, http.StatusOK)
		
		require.Equal(t, NewBankAccountTransaction.Amount, int32(10))
	})

	t.Run("Cannot Deposit", func(t *testing.T) {
		var Info struct{Message string `json:"message"`}

		t.Run("when amount is zero", func(t *testing.T) {
			payload := `
			{
				"amount": 0,
				"note": "test"
			}`
			resp, _ := http.Post(url, "application/json", strings.NewReader(payload))
			utils.ParseContent(resp, &Info)
			require.Equal(t, resp.StatusCode, http.StatusBadRequest)
			require.Equal(t, Info.Message, "amount must be greater than zero")
		})

		t.Run("when amount is less than zero", func(t *testing.T) {
			payload := `
			{
				"amount": -10,
				"note": "test"
			}`
			resp, _ := http.Post(url, "application/json", strings.NewReader(payload))
			utils.ParseContent(resp, &Info)
			require.Equal(t, resp.StatusCode, http.StatusBadRequest)
			require.Equal(t, Info.Message, "amount must be greater than zero")
		})
	})
}

func TestWithdrawRoute(t *testing.T) {
	utils.TruncateDBTables()
	url := "http://localhost:8080/bank-account/guru/withdraw"
	NewBankAccount := models.BankAccount{Name: "test", Username: "guru", Email: "test@example.com"}
	NewBankAccount.CreateBankAccount()
	NewBankAccountTransaction := models.BankAccountTransaction{Amount: 10}
	NewBankAccountTransaction.CreateBankAccountTransaction(utils.TransactionTypeCredit, "guru")

	t.Run("Can withdraw", func(t *testing.T) {
		payload := `
			{
				"amount": 10,
				"note": "test"
			}`
		resp, err := http.Post(url, "application/json", strings.NewReader(payload))
		NewBankAccountTransaction := &models.BankAccountTransaction{}
		utils.ParseContent(resp, NewBankAccountTransaction)
		require.NoError(t,  err)
		require.Equal(t, resp.StatusCode, http.StatusOK)
		
		require.Equal(t, NewBankAccountTransaction.Amount, int32(-10))
	})

	t.Run("Cannot Withdraw", func(t *testing.T) {
		var Info struct{Message string `json:"message"`}

		t.Run("when amount is zero", func(t *testing.T) {
			payload := `
			{
				"amount": 0,
				"note": "test"
			}`
			resp, _ := http.Post(url, "application/json", strings.NewReader(payload))
			utils.ParseContent(resp, &Info)
			require.Equal(t, resp.StatusCode, http.StatusBadRequest)
			require.Equal(t, Info.Message, "amount must be greater than zero")
		})

		t.Run("when amount is less than zero", func(t *testing.T) {
			payload := `
			{
				"amount": -10,
				"note": "test"
			}`
			resp, _ := http.Post(url, "application/json", strings.NewReader(payload))
			utils.ParseContent(resp, &Info)
			require.Equal(t, resp.StatusCode, http.StatusBadRequest)
			require.Equal(t, Info.Message, "amount must be greater than zero")
		})

		t.Run("when balance is insufficient", func(t *testing.T) {
			payload := `
			{
				"amount": 100,
				"note": "test"
			}`
			resp, _ := http.Post(url, "application/json", strings.NewReader(payload))
			utils.ParseContent(resp, &Info)
			require.Equal(t, resp.StatusCode, http.StatusBadRequest)
			require.Equal(t, Info.Message, "balance insufficient")
		})
	})
}

func TestAccountBalanceRoute(t *testing.T) {
	url := "http://localhost:8080/bank-account/guru/balance"
	utils.TruncateDBTables()
	NewBankAccount := models.BankAccount{Name: "test", Username: "guru", Email: "test@example.com"}
	NewBankAccount.CreateBankAccount()
	NewBankAccountTransaction := models.BankAccountTransaction{Amount: 10}
	NewBankAccountTransaction.CreateBankAccountTransaction(utils.TransactionTypeCredit, "guru")

	t.Run("Can retrieve account balance", func(t *testing.T) {
		resp, err := http.Get(url)
		var  Account struct {Balance int32 `json:"balance"`}
		utils.ParseContent(resp, &Account)
		require.NoError(t,  err)
		require.Equal(t, resp.StatusCode, http.StatusOK)
		
		require.Equal(t, Account.Balance, int32(10))
	})
}