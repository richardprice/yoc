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
	"os"
	"strconv"

	"github.com/hashicorp/consul/api"
	"github.com/spf13/cobra"
)

var id string
var name string
var address string
var port string
var tags []string

// registerCmd represents the register command
var registerCmd = &cobra.Command{
	Use:   "register",
	Short: "Register a service with the local Consul catalog",
	Long: `Usage:
	yoc register --id myserviceid --name myservice --address 10.0.0.1 --port 8080`,
	Run: func(cmd *cobra.Command, args []string) {
		if id == "" {
			fmt.Println("Service ID is a required parameter, specify using --id or -i")
			return
		}
		if id == "" {
			fmt.Println("Service Name is a required parameter, specify using --name or -n")
			return
		}
		if address == "" {
			fmt.Println("Address is a required parameter, specify using --address or -a")
			return
		}
		if port == "" {
			fmt.Println("Port is a required parameter, specify using --port or -p")
			return
		}

		i, err := strconv.Atoi(port)
		if err != nil {
			fmt.Println("Port needs to be a valid unix port number, e.g. 8080")
			os.Exit(2)
		}

		err = registerService(id, name, address, i, tags)
		if err != nil {
			fmt.Println(err.Error())
		}
	},
}

func init() {
	RootCmd.AddCommand(registerCmd)

	registerCmd.Flags().StringVarP(&id, "id", "i", "", "Service ID")
	registerCmd.Flags().StringVarP(&name, "name", "n", "", "Service Name")
	registerCmd.Flags().StringVarP(&address, "address", "a", "", "Address")
	registerCmd.Flags().StringVarP(&port, "port", "p", "", "Port")
	registerCmd.Flags().StringArrayVarP(&tags, "tag", "t", tags, "Tag")
}

func registerService(id string, name string, address string, port int, tags []string) error {
	config := api.DefaultConfig()
	client, err := api.NewClient(config)
	if err != nil {
		return err
	}

	agent := client.Agent()

	service := &api.AgentServiceRegistration{
		ID:      id,
		Name:    name,
		Address: address,
		Tags:    tags,
		Port:    port,
	}

	err = agent.ServiceRegister(service)
	if err != nil {
		return err
	}

	return nil
}
