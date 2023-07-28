package tests

import (
	"bitcoin-address-cycles"
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"bitcoin-address-cycles/rpc"
)

func TestCountCycles(t *testing.T) {
	ctx := context.Background()
	testServer := bitcoin_address_cycles.NewSolutionServer()
	_, _ = testServer.MineBlock(ctx, transactionsFromBlock0)
	_, _ = testServer.MineBlock(ctx, transactionsFromBlock1)
	_, _ = testServer.MineBlock(ctx, transactionsFromBlock2)

	// Test data
	testCases := []struct {
		fromBlock, toBlock, maxCycleLength int64
		expectedCycles                     int64
	}{
		//{0, 0, 1, 0},
		//{0, 1, 1, 3},
		//{0, 1, 0, 0},
		//{1, 1, 1, 0},
		{0, 2, 2, 12},
		//{0, 2, 1, 6},
		//{1, 2, 1, 3},
	}

	for _, tc := range testCases {
		req := &rpc.CountCyclesRequest{
			FromBlock:      tc.fromBlock,
			ToBlock:        tc.toBlock,
			MaxCycleLength: tc.maxCycleLength,
		}

		resp, err := testServer.CountCycles(ctx, req) // Replace CountCycles with the actual function name
		if err != nil {
			t.Errorf("CountCycles error: %v", err)
		}

		if resp.NCycles != tc.expectedCycles {
			t.Errorf("Expected %d cycles, got %d", tc.expectedCycles, resp.NCycles)
		}
	}
}

func TestMineBlock(t *testing.T) {
	// Test case 1: Empty transactions
	t.Run("EmptyTransactions", func(t *testing.T) {
		// Initialize the graph or any data structure you are using
		// For this example, let's assume we have a variable `graph` of type `Graph`
		ctx := context.Background()
		testServer := bitcoin_address_cycles.NewSolutionServer()
		// Create an empty transactions slice

		// Call the MineBlock function
		_, _ = testServer.MineBlock(ctx, transactionsFromBlock0)
		_, _ = testServer.MineBlock(ctx, transactionsFromBlock1)
		_, _ = testServer.MineBlock(ctx, transactionsFromBlock2)

		// Perform the necessary checks based on your logic
		// For this example, we assume that MineBlock returns -1 when the transactions slice is empty
		fmt.Println(testServer.BlockchainDAG)
	})

	// Test case 2: Transactions with valid data
	t.Run("ValidTransactions", func(t *testing.T) {
		// Initialize the graph or any data structure you are using
		// For this example, let's assume we have a variable `graph` of type `Graph`
		ctx := context.Background()
		testServer := bitcoin_address_cycles.NewSolutionServer()

		// Create sample vertices
		vertex1 := &rpc.Vertex{TxHash: "hash1", Index: 0, Address: "addrA"}
		vertex2 := &rpc.Vertex{TxHash: "hash2", Index: 0, Address: "addrB"}

		// Create sample transaction with inputs and outputs
		transaction := &rpc.Transaction{
			Inputs:  []*rpc.Vertex{vertex1},
			Outputs: []*rpc.Vertex{vertex2},
		}

		// Add the transaction to the transactions slice
		transactions := []*rpc.Transaction{transaction}

		// Call the MineBlock function
		_, err := testServer.MineBlock(ctx, &rpc.MineBlockRequest{
			Txs: transactions,
		})

		assert.NoError(t, err)

		// Perform the necessary checks based on your logic
		// For this example, we assume that MineBlock returns 0 when successful
		// Modify the conditions below based on your actual implementation's behavior
		fmt.Println(testServer.BlockchainDAG.Blocks[0].Transactions[0].Inputs[0].Address)
	})

	// Add more test cases as needed to cover other scenarios and edge cases.
}
