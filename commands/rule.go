package commands

import (
	"errors"
	"fmt"

	"github.ibm.com/Bluemix/whisk-cli/client"

	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/cobra"
)

// ruleCmd represents the rule command
var ruleCmd = &cobra.Command{
	Use:   "rule",
	Short: "work with rules",
	Long:  `[ TODO :: add longer description here ]`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: Work your own magic here
		fmt.Println("rule called")
	},
}

var ruleEnableCmd = &cobra.Command{
	Use:   "enable <name string>",
	Short: "enable rule",
	Long:  `[ TODO :: add longer description here ]`,
	Run: func(cmd *cobra.Command, args []string) {

		var err error
		if len(args) != 1 {
			err = errors.New("Invalid argument")
			fmt.Println(err)
			return
		}

		ruleName := args[0]

		_, _, err = whisk.Rules.SetState(ruleName, "enable")
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("ok: enabled rule ", ruleName)

	},
}
var ruleDisableCmd = &cobra.Command{
	Use:   "disable <name string>",
	Short: "disable rule",
	Long:  `[ TODO :: add longer description here ]`,
	Run: func(cmd *cobra.Command, args []string) {

		var err error
		if len(args) != 1 {
			err = errors.New("Invalid argument")
			fmt.Println(err)
			return
		}

		ruleName := args[0]

		_, _, err = whisk.Rules.SetState(ruleName, "disable")
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("ok: disabled rule ", ruleName)

	},
}

var ruleStatusCmd = &cobra.Command{
	Use:   "status <name string>",
	Short: "get rule status",
	Long:  `[ TODO :: add longer description here ]`,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO: how is this different than "rule get" ??
		fmt.Println("rule status called")
	},
}

var ruleCreateCmd = &cobra.Command{
	Use:   "create <name string> <trigger string> <action string>",
	Short: "create new rule",
	Long:  `[ TODO :: add longer description here ]`,
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		if len(args) != 3 {
			err = errors.New("Invalid argument list")
			fmt.Println(err)
			return
		}

		ruleName := args[0]
		triggerName := args[1]
		actionName := args[2]

		rule := &client.Rule{
			Name:    ruleName,
			Trigger: triggerName,
			Action:  actionName,
			Publish: flags.shared,
		}

		rule, _, err = whisk.Rules.Insert(rule, false)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("ok: created rule ", ruleName)
		spew.Dump(rule)
	},
}

var ruleUpdateCmd = &cobra.Command{
	Use:   "update <name string> <trigger string> <action string>",
	Short: "update existing rule",
	Long:  `[ TODO :: add longer description here ]`,
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		if len(args) != 3 {
			err = errors.New("Invalid argument list")
			fmt.Println(err)
			return
		}

		ruleName := args[0]
		triggerName := args[1]
		actionName := args[2]

		rule := &client.Rule{
			Name:    ruleName,
			Trigger: triggerName,
			Action:  actionName,
			Publish: flags.shared,
		}

		rule, _, err = whisk.Rules.Insert(rule, true)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("ok: updated rule ", ruleName)
		spew.Dump(rule)
	},
}

var ruleGetCmd = &cobra.Command{
	Use:   "get <name string>",
	Short: "get rule",
	Long:  `[ TODO :: add longer description here ]`,
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		if len(args) != 1 {
			err = errors.New("Invalid argument")
			fmt.Println(err)
			return
		}

		ruleName := args[0]

		rule, _, err := whisk.Rules.Fetch(ruleName)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("ok: got rule ", ruleName)
		spew.Dump(rule)
	},
}

var ruleDeleteCmd = &cobra.Command{
	Use:   "delete <name string>",
	Short: "delete rule",
	Long:  `[ TODO :: add longer description here ]`,
	Run: func(cmd *cobra.Command, args []string) {
		var err error
		if len(args) != 1 {
			err = errors.New("Invalid argument")
			fmt.Println(err)
			return
		}

		ruleName := args[0]

		_, err = whisk.Rules.Delete(ruleName)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("ok: deleted rule ", ruleName)
	},
}

var ruleListCmd = &cobra.Command{
	Use:   "list",
	Short: "list all rules",
	Long:  `[ TODO :: add longer description here ]`,
	Run: func(cmd *cobra.Command, args []string) {

		ruleListOptions := &client.RuleListOptions{
			Skip:  flags.skip,
			Limit: flags.limit,
		}

		rules, _, err := whisk.Rules.List(ruleListOptions)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("rules")
		spew.Dump(rules)
	},
}

func init() {

	ruleCreateCmd.Flags().BoolVar(&flags.shared, "shared", false, "shared action (default: private)")
	ruleCreateCmd.Flags().BoolVar(&flags.auto, "auto", false, "autmatically enable rule after creating it")

	ruleUpdateCmd.Flags().BoolVar(&flags.shared, "shared", false, "shared action (default: private)")

	ruleDeleteCmd.Flags().BoolVar(&flags.auto, "auto", false, "autmatically disable rule before deleting it")

	ruleListCmd.Flags().IntVarP(&flags.skip, "skip", "s", 0, "skip this many entities from the head of the collection")
	ruleListCmd.Flags().IntVarP(&flags.limit, "limit", "l", 30, "only return this many entities from the collection")

	ruleCmd.AddCommand(
		ruleEnableCmd,
		ruleDisableCmd,
		ruleStatusCmd,
		ruleCreateCmd,
		ruleUpdateCmd,
		ruleGetCmd,
		ruleDeleteCmd,
		ruleListCmd,
	)

}
