package cli

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/user"
	"github.com/klauern/wlsrest/wls"
)

func WlsRestCmd(cmd *cobra.Command, args []string) {
	fmt.Printf("Args passed to WlsRestCmd: %v\n", args)
}

func Servers(cmd *cobra.Command, args []string) {
	fmt.Printf("Args passed to Servers: %v\n", args)
}

func Clusters(cmd *cobra.Command, args []string) {
	fmt.Printf("Args passed to Clusters: %v\n", args)
}

func DataSources(cmd *cobra.Command, args []string) {
	fmt.Printf("Args passed to DataSources: %v\n", args)
}

func Applications(cmd *cobra.Command, args []string) {
	fmt.Printf("Args passed to Applications: %v\n", args)
}

// Configure default credentials to use when making REST queries to an AdminServer
func Configure(cmd *cobra.Command, args []string) {
	cfg, err := findConfiguration()
	if err != nil {
		fmt.Printf("Error found: %v", err)
	}
	fmt.Printf("Current Working Directory: %v", cfg.ServerUrl)
}

// Finds a configuration setting for your login.  Looks for the following configuration file, processed in the following
// order:
//   - command-line flags --username, --password and --server <host:port>
//   - WLSREST_CONFIG (environment variable)
//   - wlsrest.cfg (in the current directory)
//   - .wlsrest.cfg (in the $HOME directory)
//
// This is borrowed lovingly from Ansible's similar setup for it's configuration (http://docs.ansible.com/ansible/intro_configuration.html)
func findConfiguration() (*wls.Environment, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return nil, err
	}
	fmt.Println(cwd)
	u, _ := user.Current()
	fmt.Printf("Home Directory %v", u.HomeDir)
	return &wls.Environment{"", "", ""}, nil
}