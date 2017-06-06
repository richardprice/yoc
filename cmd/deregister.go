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

// deregisterCmd represents the deregister command
var deregisterCmd = &cobra.Command{
	Use:   "deregister",
	Short: "Deregister a service from the local Consul catalog",
	Long: `Usage:
	yoc deregister myserviceid`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("Service ID is a required argument")
			return
		}

		if len(args) > 1 {
			fmt.Println("yoc deregister only takes one argument")
			return
		}

		deregisterService(args[0])
	},
}

func init() {
	RootCmd.AddCommand(deregisterCmd)
}

func deregisterService(id string) error {
	config := api.DefaultConfig()
	client, err := api.NewClient(config)
	if err != nil {
		return err
	}

	agent := client.Agent()

	if err := agent.ServiceDeregister(id); err != nil {
		return err
	}

	return nil
}
