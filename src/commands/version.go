// Copyright 2015 The Hugo Authors. All rights reserved.
// Copyright 2020 The Gotham Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package commands

import (
	"errors"

	"github.com/strawberryssg/strawberry-v0/common/hugo"
	"github.com/spf13/cobra"

	jww "github.com/spf13/jwalterweatherman"
)

var _ cmder = (*versionCmd)(nil)

type versionCmd struct {
	*baseCmd
}

var vType string

func newVersionCmd() *versionCmd {

	vc := &versionCmd{
		newBaseCmd(&cobra.Command{
			Use:   "version",
			Short: "Print the version number of Gotham",
			Long: `This will print the Gotham and Hugo version numbers. There
are flags available to print just the Gotham version for scripting.`,
			RunE: func(cmd *cobra.Command, args []string) error {

				var theType hugo.VersionType

				if vType == "regular" {
					theType = hugo.VersionRegular
				} else if vType == "short" {
					theType = hugo.VersionShort
				} else if vType == "detailed" {
					theType = hugo.VersionDetailed
				} else {
					return errors.New("Invalid value for --type.")
				}

				jww.FEEDBACK.Println(hugo.PrintGothamVersion(theType))
				return nil
			},
		}),
	}

	vc.cmd.Flags().StringVarP(&vType, "type", "", "regular", "level of information to display: short, regular, or detailed")

	return vc
}
