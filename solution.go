package bitcoin_address_cycles

import (
	"bitcoin-address-cycles/rpc"
	"context"
	"fmt"
)

type Block struct {
	Transactions []*rpc.Transaction
}

type Graph struct {
	Blocks []*Block
}

type SolutionServer struct {
	rpc.UnimplementedSolutionServer
	BlockchainDAG *Graph
}

func NewSolutionServer() *SolutionServer {
	return &SolutionServer{
		BlockchainDAG: &Graph{},
	}
}

// MineBlock (transactions) -> -1
// transactions - transactions to batch into block and append to the graph.
func (s *SolutionServer) MineBlock(ctx context.Context, request *rpc.MineBlockRequest) (*rpc.MineBlockResponse, error) {
	if len(request.Txs) == 0 {
		return &rpc.MineBlockResponse{}, fmt.Errorf("empty transactions")
	}
	// Process each transaction
	//add request.Tsx to the graph
	s.BlockchainDAG.Blocks = append(s.BlockchainDAG.Blocks, &Block{
		Transactions: request.Txs,
	})

	// Successfully added the transactions to the DAG
	return &rpc.MineBlockResponse{}, nil
}

// CountCycles counts the number of address cycles in the graph.
// maxCycleLength - maximum length of the cycle to count.
// FromBlock - block to start the search from (inclusive).
// ToBlock - block to end the search at (inclusive).
// Address cycle is a path in the graph from one vertex to the other with the same address.
func (s *SolutionServer) CountCycles(ctx context.Context, request *rpc.CountCyclesRequest) (*rpc.CountCyclesResponse, error) {
	visited := make(map[string]bool)
	var numberOfCycles int64 = 0

	// Iterate through all blocks in the specified range
	for i := request.FromBlock; i <= request.ToBlock; i++ {
		block := s.BlockchainDAG.Blocks[i]
		if block != nil {
			// Iterate through all transactions in the block
			for _, transaction := range block.Transactions {
				// Iterate through all vertices in the transaction's outputs
				for _, output := range transaction.Outputs {
					// Perform DFS to find cycles starting from the current output vertex
					numberOfCycles += s.findCycleDFS(output, output, request.MaxCycleLength, visited, 1)
				}
			}
		}
	}

	return &rpc.CountCyclesResponse{NCycles: numberOfCycles}, nil
}

// findCycleDFS is a helper function that performs Depth-First Search (DFS)
// to find cycles starting from the given vertex.
func (s *SolutionServer) findCycleDFS(startVertex *rpc.Vertex, currentVertex *rpc.Vertex, maxCycleLength int64, visited map[string]bool, currentCycleLength int64) int64 {
	// Check if the current vertex has already been visited
	vertexKey := fmt.Sprintf("%s-%d", currentVertex.TxHash, currentVertex.Index)
	if visited[vertexKey] {
		return 0
	}

	// Mark the current vertex as visited
	visited[vertexKey] = true

	// Initialize the count of cycles found from this vertex
	cycleCount := int64(0)

	// Search for transactions that have the currentVertex as an input
	for _, block := range s.BlockchainDAG.Blocks {
		for _, transaction := range block.Transactions {
			// Check if the transaction has the currentVertex as an input
			for _, input := range transaction.Inputs {
				if input.TxHash == currentVertex.TxHash {
					// The transaction has the currentVertex as an input.
					// Now, search for vertices in the transaction's outputs with the same address.
					for _, output := range transaction.Outputs {
						if output.Address == startVertex.Address {
							// Found a cycle! Increment the cycle count
							fmt.Println(output.Address, startVertex.Address, output.TxHash, startVertex.TxHash)
							cycleCount++
						}

						// Continue DFS to explore the adjacent vertex
						if currentCycleLength+1 <= maxCycleLength {
							cycleCount += s.findCycleDFS(startVertex, output, maxCycleLength, visited, currentCycleLength+1)
						}
					}
				}
			}
		}
	}

	// Mark the current vertex as unvisited before returning from the function
	visited[vertexKey] = false

	return cycleCount
}
