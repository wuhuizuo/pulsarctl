// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package sources

import (
	"github.com/streamnative/pulsarctl/pkg/cmdutils"
	"github.com/streamnative/pulsarctl/pkg/pulsar"

	"github.com/spf13/pflag"
)

func deleteSourcesCmd(vc *cmdutils.VerbCmd) {
	desc := pulsar.LongDescription{}
	desc.CommandUsedFor = "This command is used for deleting a Pulsar IO source connector."
	desc.CommandPermission = "This command requires namespace function permissions."

	var examples []pulsar.Example

	del := pulsar.Example{
		Desc: "Delete a Pulsar IO source connector",
		Command: "pulsarctl source delete \n" +
			"\t--tenant public\n" +
			"\t--namespace default\n" +
			"\t--name <the name of Pulsar Source>",
	}
	examples = append(examples, del)
	desc.CommandExamples = examples

	var out []pulsar.Output
	successOut := pulsar.Output{
		Desc: "normal output",
		Out:  "Deleted <the name of a Pulsar Source> successfully",
	}

	nameNotExistOut := pulsar.Output{
		Desc: "source doesn't exist",
		Out:  "code: 404 reason: Source <the name of a Pulsar Source> doesn't exist",
	}

	out = append(out, successOut, nameNotExistOut)
	desc.CommandOutput = out

	vc.SetDescription(
		"delete",
		"Delete a Pulsar IO source connector",
		desc.ToString(),
		"delete",
	)

	sourceData := &pulsar.SourceData{}

	// set the run source
	vc.SetRunFunc(func() error {
		return doDeleteSource(vc, sourceData)
	})

	// register the params
	vc.FlagSetGroup.InFlagSet("SourcesConfig", func(flagSet *pflag.FlagSet) {
		flagSet.StringVar(
			&sourceData.Tenant,
			"tenant",
			"",
			"The source's tenant")

		flagSet.StringVar(
			&sourceData.Namespace,
			"namespace",
			"",
			"The source's namespace")

		flagSet.StringVar(
			&sourceData.Name,
			"name",
			"",
			"The source's name")
	})
}

func doDeleteSource(vc *cmdutils.VerbCmd, sourceData *pulsar.SourceData) error {
	err := processBaseArguments(sourceData)
	if err != nil {
		vc.Command.Help()
		return err
	}
	admin := cmdutils.NewPulsarClientWithAPIVersion(pulsar.V3)
	err = admin.Sources().DeleteSource(sourceData.Tenant, sourceData.Namespace, sourceData.Name)
	if err != nil {
		return err
	}

	vc.Command.Printf("Deleted %s successfully", sourceData.Name)
	return nil
}
