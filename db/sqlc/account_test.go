package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/oluwaferanmiadetunji/simple_bank/util"
	"github.com/stretchr/testify/require"
)

// The function creates a random test account and performs various assertions to ensure its
// correctness.
func createRandomTestAccount(t *testing.T) Account {
	user := createRandomTestUser(t)

	arg := CreateAccountParams{
		Owner:    user.Username,
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	account, err := testQueries.CreateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}

// The function "TestCreateAccount" is used to test the creation of a random test account.
func TestCreateAccount(t *testing.T) {
	createRandomTestAccount(t)
}

// The TestGetAccount function tests the GetAccount function by creating a random test account,
// retrieving the account using its ID, and then comparing the retrieved account with the original
// account.
func TestGetAccount(t *testing.T) {
	account1 := createRandomTestAccount(t)
	account2, err := testQueries.GetAccount(context.Background(), account1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Owner, account2.Owner)
	require.Equal(t, account1.Balance, account2.Balance)
	require.Equal(t, account1.Currency, account2.Currency)
	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)
}

// The TestUpdateAccount function tests the update functionality of an account by creating a random
// test account, updating its balance, and comparing the updated account with the original account.
func TestUpdateAccount(t *testing.T) {
	account1 := createRandomTestAccount(t)

	arg := UpdateAccountParams{
		ID:      account1.ID,
		Balance: util.RandomMoney(),
	}

	account2, err := testQueries.UpdateAccount(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, account2)

	require.Equal(t, account1.ID, account2.ID)
	require.Equal(t, account1.Owner, account2.Owner)
	require.Equal(t, arg.Balance, account2.Balance)
	require.Equal(t, account1.Currency, account2.Currency)
	require.WithinDuration(t, account1.CreatedAt, account2.CreatedAt, time.Second)
}

// The TestDeleteAccount function tests the deletion of an account and verifies that the account no
// longer exists in the database.
func TestDeleteAccount(t *testing.T) {
	account1 := createRandomTestAccount(t)

	err := testQueries.DeleteAccount(context.Background(), account1.ID)
	require.NoError(t, err)

	account2, err := testQueries.GetAccount(context.Background(), account1.ID)

	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, account2)
}

// The TestListAccounts function tests the ListAccounts method by creating random test accounts,
// setting the limit and offset parameters, and asserting that the returned accounts meet the expected
// criteria.
func TestListAccounts(t *testing.T) {
	var lastAccount Account

	for i := 0; i < 10; i++ {
		lastAccount = createRandomTestAccount(t)
	}

	arg := ListAccountsParams{
		Limit:  5,
		Offset: 0,
		Owner:  lastAccount.Owner,
	}

	accounts, err := testQueries.ListAccounts(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, accounts)

	for _, account := range accounts {
		require.NotEmpty(t, account)
		require.Equal(t, lastAccount.Owner, account.Owner)
	}

}
