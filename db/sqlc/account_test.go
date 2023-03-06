package db

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/igua95/simplebank/util"
	"github.com/stretchr/testify/require"
)

func createRandomAccount(t *testing.T) Account {
	user := createRandomUser(t)
	arg := CreateAccountParams{
		Owner:    user.Username,
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}

	account, error := testQueries.CreateAccount(context.Background(), arg)
	require.NoError(t, error)
	require.NotEmpty(t, account)

	require.Equal(t, arg.Owner, account.Owner)
	require.Equal(t, arg.Balance, account.Balance)
	require.Equal(t, arg.Currency, account.Currency)

	require.NotZero(t, account.ID)
	require.NotZero(t, account.CreatedAt)

	return account
}
func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestGetAccount(t *testing.T) {
	account := createRandomAccount(t)
	accountFromDb, error := testQueries.GetAccount(context.Background(), account.ID)

	require.NoError(t, error)
	require.NotEmpty(t, accountFromDb)

	require.Equal(t, account.Balance, accountFromDb.Balance)
	require.Equal(t, account.Currency, accountFromDb.Currency)
	require.Equal(t, account.Owner, accountFromDb.Owner)
	require.WithinDuration(t, account.CreatedAt, accountFromDb.CreatedAt, time.Second)

}

func TestUpdateAccount(t *testing.T) {
	account := createRandomAccount(t)

	arg := UpdateAccountParams{
		ID:      account.ID,
		Balance: util.RandomMoney(),
	}

	accountFromDb, error := testQueries.UpdateAccount(context.Background(), arg)

	require.NoError(t, error)
	require.NotEmpty(t, accountFromDb)

	require.Equal(t, account.Currency, accountFromDb.Currency)
	require.Equal(t, account.Owner, accountFromDb.Owner)
	require.Equal(t, arg.Balance, accountFromDb.Balance)
	require.WithinDuration(t, account.CreatedAt, accountFromDb.CreatedAt, time.Second)

}

func TestDeleteAccount(t *testing.T) {
	account := createRandomAccount(t)
	error := testQueries.DeleteAccount(context.Background(), account.ID)

	require.NoError(t, error)

	accountFromDb, error := testQueries.GetAccount(context.Background(), account.ID)
	require.Error(t, error)
	require.EqualError(t, error, sql.ErrNoRows.Error())
	require.Empty(t, accountFromDb)
}

func TestListAccount(t *testing.T) {
	for i := 0; i < 10; i++ {
		createRandomAccount(t)
	}

	accounts, error := testQueries.ListAccounts(context.Background(), ListAccountsParams{Limit: 5, Offset: 5})

	require.NoError(t, error)
	require.Len(t, accounts, 5)

	for _, account := range accounts {
		require.NotEmpty(t, account)
	}
}
