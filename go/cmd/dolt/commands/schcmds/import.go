// Copyright 2019 Liquidata, Inc.
//
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

package schcmds

import (
	"context"

	"github.com/liquidata-inc/dolt/go/cmd/dolt/cli"
	"github.com/liquidata-inc/dolt/go/cmd/dolt/commands"
	"github.com/liquidata-inc/dolt/go/libraries/doltcore/env"
	"github.com/liquidata-inc/dolt/go/libraries/utils/argparser"
)

var schImportShortDesc = "Creates a new table with an inferred schema."
var schImportLongDesc = ""
var schImportSynopsis = []string{
	"[--force] [--dry-run] [--file-type <type>] [--pks <field...>] <table> <file>",
}

func Import(ctx context.Context, commandStr string, args []string, dEnv *env.DoltEnv) int {
	ap := argparser.NewArgParser()
	ap.ArgListHelp["table"] = "Name of the table to be created."
	ap.ArgListHelp["file"] = "The file being used to infer the schema."

	help, usage := cli.HelpAndUsagePrinters(commandStr, schImportShortDesc, schImportLongDesc, schImportSynopsis, ap)
	cli.ParseArgs(ap, args, help)

	return commands.HandleVErrAndExitCode(nil, usage)
}