package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/BurntSushi/toml"
)

type ApplicationConfig struct {
	Name            string `toml:"name"`
	Version         string `toml:"version"`
	Url             string `toml:"url"`
	License         string `toml:"license"`
	Description     string `toml:"description"`
	LongDescription string `toml:"long_description"`
}

type BuildConfig struct {
	Target    string   `toml:"target"`
	Flags     string   `toml:"flags"`
	Platforms []string `toml:"platforms"`
}

type MaintainerConfig struct {
	Name  string `toml:"name"`
	Email string `toml:"email"`
}

type SimplePackagingConfig struct {
	Package       bool     `toml:"package"`
	Architectures []string `toml:"architectures"`
}

type PackagingConfig struct {
	Package       bool     `toml:"package"`
	BuildSource   bool     `toml:"build_src"`
	Architectures []string `toml:"architectures"`
}

type Config struct {
	Application ApplicationConfig     `toml:"application"`
	Build       BuildConfig           `toml:"build"`
	Maintainer  MaintainerConfig      `toml:"maintainer"`
	Deb         SimplePackagingConfig `toml:"deb"`
	RPM         PackagingConfig       `toml:"rpm"`
	Pkg         SimplePackagingConfig `toml:"pkg"`
}

func loadConfig() {
	_, err := toml.DecodeFile(configFile, &config)

	if err != nil {
		fatal(fmt.Sprintf("Failed to load make config \"%s\": %s", configFile, strings.Split(err.Error(), ":")[1][1:]))
	}

	countPackageFormats()
}

func writeDefaultConfig() {
	configText := CONFIG_DEFAULT
	if generateTarget == "all" {
		configText = CONFIG_ALL
	} else if generateTarget == "empty" {
		configText = CONFIG_EMPTY
	} else {
		generateTarget = "default"
	}

	info(time.Now(), "Writing config template "+generateTarget+" to "+configFile+".")

	file, err := os.Create(configFile)
	if err != nil {
		fatal("Failed to create config: " + err.Error())
		return
	}
	defer file.Close()

	_, err = file.WriteString(configText)
	if err != nil {
		stepError("Failed to write config: "+err.Error(), 1, packageFormatCount, 0)
		return
	}
}

const CONFIG_EMPTY = `[application]
name = ""
version = ""
description = ""
long_description = ""
url = ""
license = ""

[maintainer]
name = ""
email = ""

[build]
target = "."
flags = ""
platforms = [ ]

[deb]
package = false
architectures = [ ]

[rpm]
package = false
build_src = false
architectures = [ ]

[pkg]
package = false
architectures = [ ]
`

const CONFIG_DEFAULT = `[application]
name = "app"
version = "1.0.0"
description = "My cool application."
long_description = "My cool application."
url = "https://github.com/Username/app"
license = ""

[maintainer]
name = "Name Surname"
email = "name.surname@email.com"

[build]
target = "."
flags = "-ldflags=\"-w -s\""
platforms = [ "linux/amd64", "windows/amd64", "darwin/arm64" ]

[deb]
package = false
architectures = [ "amd64" ]

[rpm]
package = false
build_src = true
architectures = [ "amd64" ]

[pkg]
package = true
architectures = [ "amd64" ]
`

const CONFIG_ALL = `[application]
name = "app"
version = "1.0.0"
description = "My cool application."
long_description = "My cool application."
url = "https://github.com/Username/app"
license = ""

[maintainer]
name = "Name Surname"
email = "name.surname@email.com"

[build]
target = "."
flags = "-ldflags=\\"-w -s\\""
platforms = [ "linux/amd64", "linux/386", "linux/arm", "linux/arm64",
"windows/amd64", "windows/386", "windows/arm", "windows/arm64",
"darwin/amd64", "darwin/arm64" ]

[deb]
package = true
architectures = [ "amd64", "386", "arm", "arm64" ]

[rpm]
package = true
build_src = true
architectures = [ "amd64", "386", "arm", "arm64" ]

[pkg]
package = true
architectures = [ "amd64", "386", "arm", "arm64" ]\
`
