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

// servicesCmd represents the services command
var servicesCmd = &cobra.Command{
	Use:   "services",
	Short: "Lists all services from the local Consul catalog",
	Long: `Usage:
	yoc services`,
	Run: func(cmd *cobra.Command, args []string) {
		listServices()
	},
}

func init() {
	RootCmd.AddCommand(servicesCmd)
}

func listServices() error {
	config := api.DefaultConfig()
	client, err := api.NewClient(config)
	if err != nil {
		return err
	}

	agent := client.Agent()

	services, err := agent.Services()
	if err != nil {
		return err
	}

	fmt.Println("ID\t\tName\t\tAddress")
	fmt.Println("")

	for _, serviceValues := range services {
		fmt.Printf("%s\t\t%s\t\t%s:%v\n", serviceValues.ID, serviceValues.Service, serviceValues.Address, serviceValues.Port)
	}

	return nil
}
