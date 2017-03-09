package main

import (
	"fmt"
	"io"
	"log"
	"strconv"

	"github.com/appcelerator/amp/api/rpc/logs"
	"github.com/appcelerator/amp/cmd/amp/cli"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var logsCmd = &cobra.Command{
	Use:     "logs",
	Short:   "Fetch log entries matching provided criteria",
	Example: "amp logs -n 150 \namp logs -m",
	RunE: func(cmd *cobra.Command, args []string) error {
		err := AMP.Connect()
		if err != nil {
			return err
		}
		return Logs(AMP, cmd, args)
	},
}

func init() {
	// TODO logsCmd.Flags().String("timestamp", "", "filter by the given timestamp")
	logsCmd.Flags().String("container", "", "Filter by the given container")
	logsCmd.Flags().BoolP("follow", "f", false, "Follow log output")
	logsCmd.Flags().BoolP("infra", "i", false, "Include infrastructure logs")
	logsCmd.Flags().String("message", "", "Filter the message content by the given pattern")
	logsCmd.Flags().BoolP("meta", "m", false, "Display entry metadata")
	logsCmd.Flags().String("node", "", "Filter by the given node")
	logsCmd.Flags().StringP("number", "n", "1000", "Number of results")
	logsCmd.Flags().String("stack", "", "Filter by the given stack")

	RootCmd.AddCommand(logsCmd)
}

// Logs fetches the logs
func Logs(amp *cli.AMP, cmd *cobra.Command, args []string) (err error) {
	ctx, er := amp.GetAuthorizedContext()
	if er != nil {
		manager.fatalf(grpc.ErrorDesc(er))
		return
	}
	if amp.Verbose() {
		fmt.Println("Log flags:")
		fmt.Printf("container: %v\n", cmd.Flag("container").Value)
		fmt.Printf("follow: %v\n", cmd.Flag("follow").Value)
		fmt.Printf("infra: %v\n", cmd.Flag("infra").Value)
		fmt.Printf("message: %v\n", cmd.Flag("message").Value)
		fmt.Printf("meta: %v\n", cmd.Flag("meta").Value)
		fmt.Printf("node: %v\n", cmd.Flag("node").Value)
		fmt.Printf("number: %v\n", cmd.Flag("number").Value)
		fmt.Printf("stack: %v\n", cmd.Flag("stack").Value)
	}

	request := logs.GetRequest{}
	if len(args) > 0 {
		request.Service = args[0]
	}
	request.Container = cmd.Flag("container").Value.String()
	request.Node = cmd.Flag("node").Value.String()
	request.Message = cmd.Flag("message").Value.String()
	request.Stack = cmd.Flag("stack").Value.String()
	if request.Size, err = strconv.ParseInt(cmd.Flag("number").Value.String(), 10, 64); err != nil {
		log.Fatalf("Unable to convert number parameter: %v\n", cmd.Flag("n").Value.String())
	}
	meta := false
	if meta, err = strconv.ParseBool(cmd.Flag("meta").Value.String()); err != nil {
		log.Fatalf("Unable to convert meta parameter: %v\n", cmd.Flag("meta").Value.String())
	}
	follow := false
	if follow, err = strconv.ParseBool(cmd.Flag("follow").Value.String()); err != nil {
		log.Fatalf("Unable to convert follow parameter: %v\n", cmd.Flag("f").Value.String())
	}
	if request.Infra, err = strconv.ParseBool(cmd.Flag("infra").Value.String()); err != nil {
		log.Fatalf("Unable to convert infra parameter: %v\n", cmd.Flag("f").Value.String())
	}

	// Get logs from amplifier
	c := logs.NewLogsClient(amp.Conn)
	r, er := c.Get(ctx, &request)
	if er != nil {
		manager.fatalf(grpc.ErrorDesc(er))
		return
	}
	for _, entry := range r.Entries {
		displayLogEntry(entry, meta)
	}
	if !follow {
		return nil
	}

	// If follow is requested, get subsequent logs and stream it
	stream, er := c.GetStream(ctx, &request)
	if er != nil {
		manager.fatalf(grpc.ErrorDesc(er))
		return
	}
	for {
		entry, er := stream.Recv()
		if er == io.EOF {
			break
		}
		if er != nil {
			manager.fatalf(grpc.ErrorDesc(er))
			return
		}
		displayLogEntry(entry, meta)
	}
	return nil
}

func displayLogEntry(entry *logs.LogEntry, meta bool) {
	if meta {
		fmt.Printf("%+v\n", entry)
	} else {
		fmt.Printf("%s\n", entry.Message)
	}
}
