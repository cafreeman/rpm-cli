### R Distribution Builder

`rpm-cli` is a command-line program that creates a new R distribution (e.g. R install + a set of packages) from scratch.

The CLI itself is built using golang's [`cli`](https://github.com/urfave/cli) package. While it currently only has one command (`build-installer`), the `cli` package should make it easy to extend the tool and add new commands/flags/etc.

To use this program, you'll need to compile the source code in this repo using `go build`. Once that's done, you can start the tool by running the `rpm-cli` executable that it produces.

`rpm-cli` currently has one command, `build-installer`. This command will download an R installer, use it to install R into a specified directory, and then use the newly installed R binary to install a set of R packages.

`build-installer` requires 3 pieces of information in order to run: the R version you want to install, the path to the directory where you want to install R, and the path to a .csv file containing your package list. Each of these are described further in the documentation below:

```
NAME:
   rpm-cli build-installer - Build a new R installer from scratch

USAGE:
   rpm-cli build-installer [command options] [arguments...]

OPTIONS:
   --release VERSION, -r VERSION	The VERSION of R you'd like to install. e.g. "3.3.0"
   --destination value, -d value	The destination directory for your R install
   --manifest value, -m value		The path to your package manifest (packages.csv)
```

Note: `rpm-cli` expects the manifest file to be a `.csv` file that conforms to the following structure:

```
"Package","Version","Status","Priority","Built"
"abind","1.4-3","ok",NA,"3.2.3"
"acepack","1.3-3.3","ok",NA,"3.2.3"
"alabama","2015.3-1","ok",NA,"3.2.3"
```
