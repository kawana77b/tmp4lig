package license

import (
	"path"
)

type License struct {
	Name        string
	Description string
	FileName    string
}

func (l License) Path() string {
	return path.Join("templates", "license", l.FileName)
}

var licenses = []License{
	{
		Name:        "The Unlicense",
		Description: "A license with no conditions whatsoever which dedicates works to the public domain. Unlicensed works, modifications, and larger works may be distributed under different terms and without source code.",
		FileName:    "unlicense.tpl",
	},
	{
		Name:        "Creative Commons Zero v1.0 Universal",
		Description: "A public domain dedication used for works that are already in the public domain or that are too simple to be eligible for copyright.",
		FileName:    "cc0-1.0.tpl",
	},
	{
		Name:        "Boost Software License 1.0",
		Description: "A permissive license that is fairly short and accommodates larger projects well. It lets people do anything with your code with proper attribution and without warranty, but it prevents them from using your name to promote their products.",
		FileName:    "bsl-1.0.tpl",
	},
	{
		Name:        "MIT",
		Description: "A permissive license that is short and to the point. It lets people do anything with your code with proper attribution and without warranty.",
		FileName:    "mit.tpl",
	},
	{
		Name:        "BSD 2-Clause",
		Description: "A permissive license that comes in two variants, the New or Revised variant and the Simplified variant. It is used by many popular open source projects, e.g., it is the license used by GitHub.",
		FileName:    "bsd-2.0-clause.tpl",
	},
	{
		Name:        "BSD 3-Clause",
		Description: "A permissive license similar to the BSD 2-Clause License, but with a 3rd clause that prohibits others from using the name of the project or its contributors to promote derived products without written consent.",
		FileName:    "bsd-3.0-clause.tpl",
	},
	{
		Name:        "Apache 2.0",
		Description: "A permissive license similar to the BSD 2-Clause License, but with a 3rd clause that prohibits others from using the name of the project or its contributors to promote derived products without written consent.",
		FileName:    "apache-2.0.tpl",
	},
	{
		Name:        "Mozilla Public License 2.0",
		Description: "A permissive license that also contains clauses on how to deal with the source code of licensed files. It is often used for projects related to web technology.",
		FileName:    "mpl-2.0.tpl",
	},
	{
		Name:        "GNU LGPL v3",
		Description: "This is the GNU Lesser General Public License and LGPL is a modified version of GPL that allows software libraries to be linked with proprietary software.",
		FileName:    "lgpl-3.0.tpl",
	},
	{
		Name:        "GNU GPL v3",
		Description: "The GNU GPL is the most widely used free software license and has a strong copyleft requirement. When distributing derived works, the source code of the work must be made available under the same license.",
		FileName:    "gpl-3.0.tpl",
	},
	{
		Name:        "GNU AGPL v3",
		Description: "This is the GNU Affero General Public License and Affero General Public License is a modified version of GPL that requires source code to be made available to any network users of the AGPL-licensed work, typically a web application.",
		FileName:    "agpl-3.0.tpl",
	},
}

// returns a list of licenses
func GetLicenses() []License {
	ls := make([]License, len(licenses))
	copy(ls, licenses)
	return ls
}

// returns a list of template names
func GetLicNames() []string {
	names := make([]string, len(licenses))
	for i, t := range licenses {
		names[i] = t.Name
	}
	return names
}

// returns a info if it exists, and a boolean indicating if it was found
func GetLicInfoByName(name string) (*License, bool) {
	for _, t := range licenses {
		if t.Name == name {
			return &t, true
		}
	}

	return nil, false
}
