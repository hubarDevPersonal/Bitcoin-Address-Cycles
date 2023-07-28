package tests

import "bitcoin-address-cycles/rpc"

var (
	emptyTransactions = &rpc.MineBlockRequest{
		Txs: []*rpc.Transaction{},
	}
	transactionsFromBlock0 = &rpc.MineBlockRequest{
		Txs: []*rpc.Transaction{
			{
				Outputs: []*rpc.Vertex{
					{
						TxHash:  "hash0",
						Index:   0,
						Address: "addrA",
					},
					{
						TxHash:  "hash1",
						Index:   1,
						Address: "addrB",
					},
					{
						TxHash:  "hash2",
						Index:   2,
						Address: "addrC",
					},
					{
						TxHash:  "hash3",
						Index:   3,
						Address: "addrB",
					},
					{
						TxHash:  "hash4",
						Index:   4,
						Address: "addrA",
					},
				},
			},
		},
	}

	transactionsFromBlock1 = &rpc.MineBlockRequest{
		Txs: []*rpc.Transaction{
			{
				Inputs: []*rpc.Vertex{
					{
						TxHash:  "hash",
						Index:   0,
						Address: "addrA",
					},
					{
						TxHash:  "hash",
						Index:   1,
						Address: "addrB",
					},
					{
						TxHash:  "hash",
						Index:   2,
						Address: "addrC",
					},
				},
				Outputs: []*rpc.Vertex{
					{
						TxHash:  "hash",
						Index:   5,
						Address: "addrA",
					},
					{
						TxHash:  "hash",
						Index:   6,
						Address: "addrA",
					},
				},
			},
			{
				Inputs: []*rpc.Vertex{
					{
						TxHash:  "hash",
						Index:   3,
						Address: "addr",
					},
					{
						TxHash:  "hash",
						Index:   4,
						Address: "addrA",
					},
				},
				Outputs: []*rpc.Vertex{
					{
						TxHash:  "hash",
						Index:   7,
						Address: "addrA",
					},
				},
			},
		},
	}

	transactionsFromBlock2 = &rpc.MineBlockRequest{
		Txs: []*rpc.Transaction{
			{
				Inputs: []*rpc.Vertex{
					{
						TxHash:  "hash",
						Index:   5,
						Address: "addrA",
					},
					{
						TxHash:  "hash",
						Index:   6,
						Address: "addrA",
					},
					{
						TxHash:  "hash",
						Index:   7,
						Address: "addrA",
					},
				},
				Outputs: []*rpc.Vertex{
					{
						TxHash:  "hash",
						Index:   8,
						Address: "addrA",
					},
					{
						TxHash:  "hash",
						Index:   9,
						Address: "addrB",
					},
				},
			},
		},
	}
)
