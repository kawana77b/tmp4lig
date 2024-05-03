# tmp4lig

![demo](/docs//images/demo.gif)

Template Output Tool for license and gitignore.

## Description

The tool allows licenses to be retrieved from the binary's embedded template,
and `.gitignore` files to retrieve data from [gitignore.io](https://www.toptal.com/developers/gitignore) and output to the current work directory.

Some licenses have a username and current year injected into the template that can be retrieved with `git config --get user.name`.

Note that this tool is intended to be used locally and does not print out to standard output.

## Install

Binaries are available on the [Release page](https://github.com/kawana77b/tmp4lig/releases).

or Go can be run, the `install` command is available.

```bash
go install github.com/kawana77b/tmp4lig@latest
```

## Usage

```bash
# outputs a .gitignore or LICENSE file to the current current directory.
tmp4lig

# outputs a .gitignore file to the current current directory.
tmp4lig gitignore

# outputs a LICENSE file to the current current directory.
tmp4lig license
```

## Available license templates

The generated files refer to the following:

| Name                                          | URL                                                            |
| --------------------------------------------- | -------------------------------------------------------------- |
| GNU AFFERO GENERAL PUBLIC LICENSE Version 3.0 | https://www.gnu.org/licenses/agpl-3.0.html.en                  |
| Apache License Version 2.0                    | https://www.apache.org/licenses/LICENSE-2.0                    |
| BSD 2-Clause License                          | https://opensource.org/license/bsd-2-clause                    |
| BSD 3-Clause License                          | https://opensource.org/license/bsd-3-clause                    |
| Boost Software License                        | https://www.boost.org/users/license.html                       |
| CC0 1.0 Universal                             | https://creativecommons.org/publicdomain/zero/1.0/legalcode.en |
| MIT                                           | https://opensource.org/license/mit                             |
| Mozilla Public License Version 2.0            | https://www.mozilla.org/en-US/MPL/2.0/                         |
| The Unlicense                                 | https://unlicense.org/                                         |

## Digression

The naming of this tool is intended to stand for 'Template Output Tool for license and gitignore'.
Although `tmp` may easily remind you of _temporary_, we name it so because it seems easier to say and more coherent than describing it as tmpl or tpl.
