package service

import (
	"context"
	"errors"
	"io"

	"github.com/appcelerator/amp/api/rpc/logs"
	"github.com/appcelerator/amp/cli"
	"github.com/spf13/cobra"
	"google.golang.org/grpc/status"
)

type logsServiceOptions struct {
	meta   bool
	follow bool
}

// NewServiceLogsCommand returns a new instance of the service logs command.
func NewServiceLogsCommand(c cli.Interface) *cobra.Command {
	opts := logsServiceOptions{}
	cmd := &cobra.Command{
		Use:     "logs [OPTIONS] SERVICE",
		Short:   "Get all logs of a given service",
		PreRunE: cli.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return getLogs(c, args, opts)
		},
	}
	flags := cmd.Flags()
	flags.BoolVarP(&opts.follow, "follow", "f", false, "Follow log output")
	flags.BoolVarP(&opts.meta, "meta", "m", false, "Display entry metadata")
	return cmd
}

func getLogs(c cli.Interface, args []string, opts logsServiceOptions) error {
	request := logs.GetRequest{IncludeAmpLogs: false}
	request.Service = args[0]

	// Get logs from amplifier
	ctx := context.Background()
	conn := c.ClientConn()
	lc := logs.NewLogsClient(conn)
	r, err := lc.Get(ctx, &request)
	if err != nil {
		if s, ok := status.FromError(err); ok {
			return errors.New(s.Message())
		}
	}
	for _, entry := range r.Entries {
		displayLogEntry(c, entry, opts.meta)
	}
	if !opts.follow {
		return nil
	}

	// If follow is requested, get subsequent logs and stream it
	stream, err := lc.GetStream(ctx, &request)
	if err != nil {
		if s, ok := status.FromError(err); ok {
			return errors.New(s.Message())
		}
	}
	for {
		entry, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			if s, ok := status.FromError(err); ok {
				return errors.New(s.Message())
			}
		}
		displayLogEntry(c, entry, opts.meta)
	}
	return nil
}

func displayLogEntry(c cli.Interface, entry *logs.LogEntry, meta bool) {
	if meta {
		c.Console().Printf("%+v\n", entry)
	} else {
		c.Console().Printf("%s\n", entry.Msg)
	}
}
