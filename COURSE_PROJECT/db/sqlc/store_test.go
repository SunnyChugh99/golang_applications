package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTransferTx(t *testing.T){
	store := NewStore(testDB)


	account1 := createRandomAccount(t)
	
	account2 := createRandomAccount(t)
	fmt.Println(">> before:", account1.Balance, account2.Balance)

	n := 5 
	amount := int64(10)

	errs := make(chan error)

	results := make(chan TransferTxResult)
	for i:=0; i<n; i++{
		go func() {
		//func (store *Store) TransferTx(ctx context.Context, args TransferTxParams) (TransferTxResult,error){
			result, err :=	store.TransferTx(context.Background(), TransferTxParams{
				FromAccountId: account1.ID,
				ToAccountId: account2.ID,
				Amount: amount,
			})

			errs <- err

			results <- result

		}()
	}

	existed := make(map[int]bool)

	for i:=0; i<n; i++{
		err := <-errs
		require.NoError(t, err)
		result := <-results

		require.NotEmpty(t, result)

		//check transfer

		transfer := result.Transfer
		require.NotEmpty(t, transfer)
		require.Equal(t, amount, transfer.Amount)
		require.Equal(t, transfer.FromAccountID, account1.ID)

		require.Equal(t, transfer.ToAccountID, account2.ID)
		require.NotZero(t, transfer.ID)
		_, err = store.GetTransfer(context.Background(), transfer.ID)
		require.NoError(t, err)


		fromEntry := result.FromEntry
		require.NotEmpty(t, fromEntry)
		require.Equal(t, fromEntry.AccountID, account1.ID)
		require.Equal(t, fromEntry.Amount, -amount)
		require.NotZero(t, fromEntry.ID)


		toEntry := result.ToEntry
		require.NotEmpty(t, toEntry)
		require.Equal(t, toEntry.AccountID, account2.ID)
		require.Equal(t, toEntry.Amount, amount)
		require.NotZero(t, toEntry.ID)


		fromAccount := result.FromAccount
		require.NotEmpty(t, fromAccount)
		require.Equal(t, fromAccount.ID, account1.ID)

		toAccount := result.ToAccount
		require.NotEmpty(t, toAccount)
		require.Equal(t, toAccount.ID, account2.ID)

		fmt.Println(">> tx:", fromAccount.Balance, toAccount.Balance)

		diff1 := account1.Balance - fromAccount.Balance
		diff2 := toAccount.Balance - account2.Balance

		require.Equal(t, diff1, diff2)
		require.True(t, diff1 > 0)
		require.True(t, diff1 % amount == 0)
		k := int(diff1 / amount)
		require.True(t, k >= 1 && k<= n)
		
		require.NotContains(t, existed, k)
		existed[k] = true

	}


	updatedAccount1, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err)

	updatedAccount2, err := testQueries.GetAccount(context.Background(), account2.ID)
	require.NoError(t, err)

	fmt.Println(">> after:", updatedAccount1.Balance, updatedAccount2.Balance)
	require.Equal(t, account1.Balance - int64(n)*(amount), updatedAccount1.Balance)
	require.Equal(t, account2.Balance + int64(n)*(amount), updatedAccount2.Balance)

}



func TestTransferTxDeadlock(t *testing.T){
	store := NewStore(testDB)


	account1 := createRandomAccount(t)
	
	account2 := createRandomAccount(t)
	fmt.Println(">> before:", account1.Balance, account2.Balance)

	n := 10 
	amount := int64(10)

	errs := make(chan error)

	for i:=0; i<n; i++{
		fromAccountID := account1.ID
		toAccountID := account2.ID
		if i%2 == 1{
			fromAccountID = account2.ID
			toAccountID = account1.ID	
		}

		go func() {
			_, err :=	store.TransferTx(context.Background(), TransferTxParams{
				FromAccountId: fromAccountID,
				ToAccountId: toAccountID,
				Amount: amount,
			})

			errs <- err
		}()
	}


	for i:=0; i<n; i++{
		err := <-errs
		require.NoError(t, err)
	}


	updatedAccount1, err := testQueries.GetAccount(context.Background(), account1.ID)
	require.NoError(t, err)

	updatedAccount2, err := testQueries.GetAccount(context.Background(), account2.ID)
	require.NoError(t, err)

	fmt.Println(">> after:", updatedAccount1.Balance, updatedAccount2.Balance)
	require.Equal(t, account1.Balance, updatedAccount1.Balance)
	require.Equal(t, account2.Balance, updatedAccount2.Balance)

}