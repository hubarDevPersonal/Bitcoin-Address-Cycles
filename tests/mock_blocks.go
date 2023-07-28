package tests

import "bitcoin-address-cycles/rpc"

var (
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
						Index:   0,
						Address: "addrB",
					},
					{
						TxHash:  "hash2",
						Index:   0,
						Address: "addrC",
					},
					{
						TxHash:  "hash3",
						Index:   0,
						Address: "addrB",
					},
					{
						TxHash:  "hash4",
						Index:   0,
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
						TxHash:  "hash0",
						Index:   0,
						Address: "addrA",
					},
					{
						TxHash:  "hash1",
						Index:   0,
						Address: "addrB",
					},
					{
						TxHash:  "hash2",
						Index:   0,
						Address: "addrC",
					},
				},
				Outputs: []*rpc.Vertex{
					{
						TxHash:  "hash5",
						Index:   0,
						Address: "addrA",
					},
					{
						TxHash:  "hash6",
						Index:   0,
						Address: "addrA",
					},
				},
			},
			{
				Inputs: []*rpc.Vertex{
					{
						TxHash:  "hash3",
						Index:   0,
						Address: "addrB",
					},
					{
						TxHash:  "hash4",
						Index:   0,
						Address: "addrA",
					},
				},
				Outputs: []*rpc.Vertex{
					{
						TxHash:  "hash7",
						Index:   0,
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
						TxHash:  "hash5",
						Index:   0,
						Address: "addrA",
					},
					{
						TxHash:  "hash6",
						Index:   0,
						Address: "addrA",
					},
					{
						TxHash:  "hash7",
						Index:   0,
						Address: "addrA",
					},
				},
				Outputs: []*rpc.Vertex{
					{
						TxHash:  "hash8",
						Index:   0,
						Address: "addrA",
					},
					{
						TxHash:  "hash9",
						Index:   0,
						Address: "addrB",
					},
				},
			},
		},
	}
)
