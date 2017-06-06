// Copyright © 2017 Richard Price <richardprice@gmail.com>
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"

	"github.com/hashicorp/consul/api"
	"github.com/spf13/cobra"
)

// nodesCmd represents the nodes command
var nodesCmd = &cobra.Command{
	Use:   "nodes",
	Short: "Lists all nodes in the Consul cluster",
	Long: `Usage:
	yoc nodes`,
	Run: func(cmd *cobra.Command, args []string) {
		listNodes()
	},
}

func init() {
	RootCmd.AddCommand(nodesCmd)
}

func listNodes() error {
	config := api.DefaultConfig()
	client, err := api.NewClient(config)
	if err != nil {
		return err
	}

	catalog := client.Catalog()

	nodes, _, err := catalog.Nodes(&api.QueryOptions{})
	if err != nil {
		panic(err)
	}

	fmt.Println("Name\t\t\t\tDatacenter\t\tAddress\n")

	for _, node := range nodes {
		fmt.Printf("%s\t\t\t%s\t\t\t%s\n", node.Node, node.Datacenter, node.Address)
	}

	return nil
}
