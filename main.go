package main

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"

	"algorand.com/cli/tui/bubbles/test"
)

var buildConfigSimple = test.EventConfig{
	Title: "Building (PyTEAL)",
	Events: []test.Event{
		{
			Msg:   "contract1.py",
			Dur:   2 * time.Second,
			Final: test.Success,
		},
		{
			Msg:   " -> approve.teal",
			Dur:   750 * time.Millisecond,
			Final: test.Success,
		},
		{
			Msg:   " -> clear.teal",
			Start: 750 * time.Millisecond,
			Dur:   750 * time.Millisecond,
			Final: test.Success,
		},
		{
			Msg:   " -> abi.json",
			Start: 1500 * time.Millisecond,
			Dur:   500 * time.Millisecond,
			Final: test.Success,
		},
	},
}

var buildConfig = test.EventConfig{
	Title: "Building (PyTEAL)",
	Events: []test.Event{
		{
			Msg:   "application_create.py",
			Dur:   2 * time.Second,
			Final: test.Success,
		},
		{
			Msg:     "smart_sig.py",
			Visible: 2 * time.Second,
			Start:   2 * time.Second,
			Dur:     1 * time.Second,
			Final:   test.Success,
		},
		{
			Msg:     "app_call_new.py",
			Visible: 3 * time.Second,
			Start:   3 * time.Second,
			Dur:     250 * time.Millisecond,
			Final:   test.Success,
		},
		{
			Msg:     "app_call_delete.py",
			Visible: 3*time.Second + 250*time.Millisecond,
			Start:   3*time.Second + 250*time.Millisecond,
			Dur:     250 * time.Millisecond,
			Final:   test.Success,
		},
		{
			Msg:     "app_call_swap.py",
			Visible: 3*time.Second + 500*time.Millisecond,
			Start:   3*time.Second + 500*time.Millisecond,
			Dur:     250 * time.Millisecond,
			Final:   test.Success,
		},
		{
			Msg:     "app_call_inc.py",
			Visible: 3*time.Second + 750*time.Millisecond,
			Start:   3*time.Second + 750*time.Millisecond,
			Dur:     250 * time.Millisecond,
			Final:   test.Success,
		},
		{
			Msg:     "app_call_dec.py",
			Visible: 4 * time.Second,
			Start:   4 * time.Second,
			Dur:     250 * time.Millisecond,
			Final:   test.Success,
		},
		{
			Msg:     "app_call_mint.py",
			Visible: 4*time.Second + 250*time.Millisecond,
			Start:   4*time.Second + 250*time.Millisecond,
			Dur:     250 * time.Millisecond,
			Final:   test.Success,
		},
		{
			Msg:     "app_call_freeze.py",
			Visible: 4*time.Second + 500*time.Millisecond,
			Start:   4*time.Second + 500*time.Millisecond,
			Dur:     250 * time.Millisecond,
			Final:   test.Success,
		},
		{
			Msg:     "app_call_load.py",
			Visible: 4*time.Second + 750*time.Millisecond,
			Start:   4*time.Second + 750*time.Millisecond,
			Dur:     250 * time.Millisecond,
			Final:   test.Success,
		},
		{
			Msg:     "app_call_view.py",
			Visible: 5 * time.Second,
			Start:   5 * time.Second,
			Dur:     250 * time.Millisecond,
			Final:   test.Success,
		},
	},
}

var testConfigSimple = test.EventConfig{
	Title: "Running Tests",
	Events: []test.Event{
		{
			Msg:   "contract1_test1.lua",
			Dur:   2 * time.Second,
			Final: test.Fail,
		},
		{
			Msg:   "contract1_test2.lua",
			Dur:   time.Second,
			Final: test.Success,
		},
	},
}

var testConfig = test.EventConfig{
	Title: "Running Tests",
	Events: []test.Event{
		{
			Msg:   "Lorem ipsum dolor",
			Dur:   time.Second,
			Final: test.Fail,
		},
		{
			Msg:   "sit amet, consectetur adipiscing elit",
			Dur:   5 * time.Second,
			Final: test.Success,
		},
		{
			Msg:   "sed do eiusmod tempor",
			Start: time.Second,
			Dur:   time.Second,
			Final: test.Success,
		},
		{
			Msg:   "incididunt ut labore",
			Start: 2 * time.Second,
			Dur:   3 * time.Second,
			Final: test.Success,
		},
		{
			Msg:   "et dolore magna aliqua. Ut enim ad",
			Start: 5 * time.Second,
			Dur:   2 * time.Second,
			Final: test.Success,
		},
		{
			Msg:   "minim veniam, quis nostrud exercitation",
			Start: 5 * time.Second,
			Dur:   3 * time.Second,
			Final: test.Fail,
		},
		{
			Msg:   "ullamco laboris",
			Start: 7 * time.Second,
			Dur:   time.Second,
			Final: test.Success,
		},
		{
			Msg:   "nisi ut aliquip ex ea commodo",
			Start: 8 * time.Second,
			Dur:   5 * time.Second,
			Final: test.Fail,
		},
		{
			Msg:   "consequat. Duis aute irure",
			Start: 8 * time.Second,
			Dur:   time.Second,
			Final: test.Success,
		},
		{
			Msg:   "dolor in reprehenderit",
			Start: 9 * time.Second,
			Dur:   time.Second,
			Final: test.Success,
		},
	},
}

var deployConfig = test.EventConfig{
	Title: "Deploy (PyTEAL)",
	Events: []test.Event{
		{
			Msg:   "Constructing transaction group.",
			Dur:   time.Second,
			Final: test.Success,
		},
		{
			Msg:   "Submitting transaction group to ledger for signing (Press OK).",
			Start: time.Second,
			Dur:   5 * time.Second,
			Final: test.Success,
		},
		{
			Msg:   "Ledger signing complete.",
			Start: 6 * time.Second,
			Dur:   0,
			Final: test.Success,
		},
		{
			Msg:   "Connecting to network.",
			Start: 6 * time.Second,
			Dur:   6 * time.Second,
			Final: test.Success,
		},
		{
			Msg:   "Submitting signed transaction group.",
			Start: 12 * time.Second,
			Dur:   1 * time.Second,
			Final: test.Success,
		},
		{
			Msg:   "Waiting for confirmation.",
			Start: 13 * time.Second,
			Dur:   3 * time.Second,
			Final: test.Success,
		},
		{
			Msg:     "Application created with ID: 85919201",
			Visible: 16 * time.Second,
			Start:   16 * time.Second,
			Dur:     0 * time.Second,
			Final:   test.Success,
		},
	},
}

// directory layout
/*
project1
	purr.yml
	src
		contract1
			contract1.py
			approval.teal
			clear.teal
			abi.json
		contract2
			contract2.py
			approval.teal
			clear.teal
			abi.json
	test
		test1.lua
		test2.lua
*/

// why would you copy when you can call?
// ->

// config file
/*
name: MyAlgoApp

test_dir: test
contract_dir: source

contracts:
  - name: contract1
    type: PyTEAL
    main: contract1.py
    global-bytes: 4
    global-ints: 2
    local-bytes: 2
    local-ints: 0
    # not set because they match "{name}.*(approve|clear)"
    #approve: contract1_approve.teal
    #clear: contract1_clear.teal
  - name: contract2
    type: PyTEAL
    main: contract2.py
    approve: contract_2_approve.teal
    clear: contract_2_clear.teal
    global-bytes: 1
    global-ints: 1
    local-bytes: 0
    local-ints: 0

package:
  - name: testnet
    creator: <addr>
  - name: mainnet

*/

func main() {
	rootCmd := &cobra.Command{
		Use:   `purr`,
		Short: `purr is a build automation tool for Algorand smart contract development.`,
	}

	var projectName string
	initCmd := &cobra.Command{
		Use:   `init`,
		Short: `Initialize a smart contract project template in a new directory.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Printf("Creating project '%s'...\n", projectName)
			time.Sleep(100 * time.Millisecond)
			fmt.Printf("%s/purr.yml\n", projectName)
			time.Sleep(100 * time.Millisecond)
			fmt.Printf("%s/src/contract1/contract1.py\n", projectName)
			time.Sleep(100 * time.Millisecond)
			fmt.Printf("%s/src/contract1/approval.teal\n", projectName)
			time.Sleep(100 * time.Millisecond)
			fmt.Printf("%s/src/contract1/clear.teal\n", projectName)
			time.Sleep(100 * time.Millisecond)
			fmt.Printf("%s/test/contract1_test1.lua\n", projectName)
			time.Sleep(100 * time.Millisecond)
			fmt.Printf("%s/test/contract1_test2.lua\n", projectName)
		},
	}
	initCmd.Flags().StringVarP(&projectName, "project-name", "n", "", "Name of the new project.")
	//projectName := initCmd.Flags().StringP("project-name", "n", "", "Name of the new project.")
	initCmd.Flags().StringP("environment", "e", "", "Development environment for development [PyTEAL, Reach, AlgoML]")

	buildCmd := &cobra.Command{
		Use: `build`,
		//Short: `Compile high level language to TEAL files and ABI JSON spec.`,
		Long: `Compile high level language to TEAL files and ABI JSON spec.`,
		Run: func(cmd *cobra.Command, args []string) {
			//build.BuildCmd()
			test.Test(buildConfigSimple)
		},
	}

	testCmd := &cobra.Command{
		Use:   `test`,
		Short: `Run automated tests.`,
		Run: func(cmd *cobra.Command, args []string) {
			test.Test(testConfigSimple)
		},
	}
	testCmd.Flags().StringP("filter", "f", "", "An expression to select which tests to run.")
	testCmd.Flags().BoolP("debug", "d", false, "Run tests in debug mode. This produces a run trace for detailed debugging.")
	testCmd.Flags().BoolP("verbose", "v", false, "Run in verbose mode with additional progress output.")

	var network string
	packageCmd := &cobra.Command{
		Use:   `package`,
		Short: `Prepare an unsigned transaction group ready for final signing and deployment.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Preparing transaction group...")
			time.Sleep(100 * time.Millisecond)
			fmt.Printf("created: app_transactions_%s.txnrn", network)

		},
	}
	packageCmd.Flags().StringVarP(&network, "network", "n", "", "Network where the contract will be deployed.")

	/*
		deployCmd := &cobra.Command{
			Use:   `deploy`,
			Short: `Securely deploy the package.`,
			Run: func(cmd *cobra.Command, args []string) {
				//test.Test(deployConfig)
			},
		}
	*/

	rootCmd.CompletionOptions.HiddenDefaultCmd = true
	rootCmd.AddCommand(initCmd)
	rootCmd.AddCommand(buildCmd)
	rootCmd.AddCommand(testCmd)
	rootCmd.AddCommand(packageCmd)
	//rootCmd.AddCommand(deployCmd)

	rootCmd.Execute()
}
