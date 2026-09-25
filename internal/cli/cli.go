package cli

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/xamma/pit/internal/postgres"
)

func Run(args []string) error {
	if len(args) == 0 {
		usage()
		return fmt.Errorf("no command given")
	}

	switch args[0] {
	case "init":
		return runInit(args[1:])
	case "help", "-h", "--help":
		usage()
		return nil
	default:
		usage()
		return fmt.Errorf("unknown command %q", args[0])
	}
}

func usage() {
	fmt.Fprint(os.Stderr, `pit - postgres init tool

usage:
  pit init --sqlfile <path> [--conn <url>]   run a SQL file against the database
  pit help                                   show this message

The connection string comes from --conn, or $DATABASE_URL.
`)
}

func runInit(args []string) error {
	fs := flag.NewFlagSet("init", flag.ExitOnError)

	sqlFile := fs.String("sqlfile", "", "path to the SQL file to run (required)")
	conn := fs.String("conn", "", "postgres connection string (default $DATABASE_URL)")

	if err := fs.Parse(args); err != nil {
		return err
	}

	if *sqlFile == "" {
		return fmt.Errorf("init: --sqlfile is required")
	}

	connString := *conn
	if connString == "" {
		connString = os.Getenv("DATABASE_URL")
	}
	if connString == "" {
		return fmt.Errorf("init: no connection string: pass --conn or set DATABASE_URL")
	}

	if err := postgres.Init(context.Background(), connString, *sqlFile); err != nil {
		return fmt.Errorf("init: %w", err)
	}

	fmt.Printf("applied %s\n", *sqlFile)
	return nil
}
