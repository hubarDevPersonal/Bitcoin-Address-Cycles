package bitcoin_address_cycles

import (
	"bitcoin-address-cycles/rpc"
	"context"
	"fmt"
	"log"
	"os"
)

// Block is a collection of transactions.
type Block struct {
	Transactions []*rpc.Transaction
}

// Graph is a collection of blocks.
type Graph struct {
	Blocks []*Block
}

// SolutionServer is the server that implements the SolutionServer interface.
type SolutionServer struct {
	rpc.UnimplementedSolutionServer
	l             *log.Logger
	BlockchainDAG *Graph
}

// NewSolutionServer creates a new instance of the SolutionServer.
func NewSolutionServer() *SolutionServer {
	return &SolutionServer{
		BlockchainDAG: &Graph{},
		l:             log.New(os.Stdout, "bitcoin-address-cycles: ", log.LstdFlags),
	}
}

// MineBlock (transactions) -> -1
// transactions - transactions to batch into block and append to the graph.
func (s *SolutionServer) MineBlock(ctx context.Context, request *rpc.MineBlockRequest) (*rpc.MineBlockResponse, error) {
	if len(request.Txs) == 0 {
		s.l.Println("Error: empty transactions")
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
	var numberOfCycles int64 = 0
	//Check if the block numbers are valid
	if request.FromBlock < 0 || request.ToBlock < 0 {
		s.l.Println("Error: invalid block number")
		return &rpc.CountCyclesResponse{NCycles: numberOfCycles}, fmt.Errorf("invalid block number")
	}

	//Check if the block range is valid
	if request.FromBlock >= request.ToBlock {
		s.l.Println("Error: invalid block range")
		return &rpc.CountCyclesResponse{NCycles: numberOfCycles}, fmt.Errorf("invalid block range")
	}

	//Check if the max cycle length is valid
	if request.MaxCycleLength <= 0 {
		s.l.Println("Error: invalid max cycle length")
		return &rpc.CountCyclesResponse{NCycles: numberOfCycles}, fmt.Errorf("invalid max cycle length")
	}

	// Initialize a map to keep track of visited vertices
	visited := make(map[int64]bool)

	// Iterate through all blocks in the specified range
	for i := request.FromBlock; i < request.ToBlock; i++ {
		block := s.BlockchainDAG.Blocks[i]
		if block != nil {
			// Iterate through all transactions in the block
			for _, transaction := range block.Transactions {
				// Iterate through all vertices in the transaction's outputs
				for _, output := range transaction.Outputs {
					// Perform DFS to find cycles starting from the current output vertex
					numberOfCycles += s.findCycleDFS(output, output, request.MaxCycleLength, 1, request.ToBlock, request.FromBlock, visited)
				}
			}
		}
	}

	return &rpc.CountCyclesResponse{NCycles: numberOfCycles}, nil
}

// findCycleDFS is a helper function that performs Depth-First Search (DFS)
// to find cycles starting from the given vertex.
func (s *SolutionServer) findCycleDFS(startVertex *rpc.Vertex, currentVertex *rpc.Vertex, maxCycleLength, currentCycleLength, toBlock, fromBlock int64, visited map[int64]bool) int64 {
	// Check if the current vertex has already been visited
	if visited[currentVertex.Index] {
		return 0
	}

	// Mark the current vertex as visited
	visited[currentVertex.Index] = true

	// Initialize the count of cycles found from this vertex
	cycleCount := int64(0)

	// Search for transactions that have the currentVertex as an input
	for i := fromBlock; i <= toBlock; i++ {
		for _, transaction := range s.BlockchainDAG.Blocks[i].Transactions {
			// Check if the transaction has the currentVertex as an input
			for _, input := range transaction.Inputs {
				if input.Index == currentVertex.Index {
					// The transaction has the currentVertex as an input.
					// Search for vertices in the transaction's outputs with the same address.
					for _, output := range transaction.Outputs {
						if output.Address == startVertex.Address {
							// Found a cycle! Increment the cycle count
							s.l.Println("Found a cycle", output.Address, startVertex.Address, output.Index, startVertex.Index)
							cycleCount++
						}

						// Continue DFS to explore the adjacent vertex
						if currentCycleLength+1 <= maxCycleLength {
							cycleCount += s.findCycleDFS(startVertex, output, maxCycleLength, currentCycleLength+1, toBlock, fromBlock, visited)
						}
					}
				}
			}
		}
	}

	// Mark the current vertex as unvisited before returning from the function
	visited[currentVertex.Index] = false

	return cycleCount
}
