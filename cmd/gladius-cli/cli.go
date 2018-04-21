package main

import (
	"github.com/gladiusio/gladius-node/internal/cli"
)

func main() {
	commands.Execute()
}

// func theWorks() {
//
// 	myNode.Data.Name = "celo-test-3"
// 	myNode.Data.Email = "celo@gladius.io"
// 	myNode.Data.IPAddress = "1.1.1.1"
// 	myNode.Data.Status = "active"
//
// 	fmt.Println()
// 	// always post settings
// 	node.SetSettings("ropsten", &localSettings)
// 	node.PostSettings(&localSettings)
//
// 	// create a node contract and wait for tx to finish
// 	tx, _ := node.CreateNode()
// 	node.WaitForTx(tx)
// 	fmt.Println("Node created!")
//
// 	// get the address of new node, set its data, and wait
// 	myNode.Address = node.GetNodeAddress()
// 	tx, _ = node.SetNodeData(myNode.Address, myNode)
// 	node.WaitForTx(tx)
// 	fmt.Println("Data set!")
//
// 	// apply to pool
// 	tx, _ = node.ApplyToPool(myNode.Address, poolAddress)
// 	node.WaitForTx(tx)
// 	fmt.Println("Applied to pool!")
//
// 	// check pool application status
// 	fmt.Println(node.CheckPoolApplication(myNode.Address, poolAddress))
// }
//
// func networkDaemon() {
// 	node.StartEdgeNode()
// 	node.StatusEdgeNode()
// 	node.StopEdgeNode()
// }
