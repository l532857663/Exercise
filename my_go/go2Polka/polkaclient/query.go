package polkaclient

import "fmt"

func (n *Node) QueryTimestamp() error {
	fmt.Printf("wch-------test\n")
	err := n.Client.Call()
	return nil
}
