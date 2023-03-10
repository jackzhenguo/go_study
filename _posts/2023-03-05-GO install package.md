
Today at this blog let's begin to implement a simple calculator with GUI interface in GO language.

### install third-party package

First, I define the project name and create a folder named `calculator_gui`
```
mkdir calculator_gui
```

Then, run command `cd calculator_gui` to change directory to this folder:
```
cd calculator_gui
```

Finally, it is the most important command. In GO language, we must run command following to initilize a GO project.

```
go mod init calculator_gui
```

This will automatically create a `go.mod` file that lists the module and its dependencies.

To install a third-party package in Go, we can use the `go get` command followed by the package import path. 

The `go get` command downloads the package and its dependencies, compiles them, and installs them in your Go workspace.

Here's an example of installing the `fyne.io/fyne/v2` package, which is used at calculator GUI design.

```
go get fyne.io/fyne/v2
```

Then GO will download the package.

### import third-party package

Import the package in your Go code. We should first create a GO file, such as calculator_gui.go file.

Then write this line of code to this file:

```
import (
    "fyne.io/fyne/v2"
)
```

After adding the package as a dependency and importing it in our Go code, we can use its functionality the package provides us.

### Challenges

Because the calculator GUI design needs to use fyne.io/fyne/v2, I encountered some challenges when downloading third-party packages for the first time in GO language. Initially, I tried to use "go get fyne.io/fyne/v2" directly, but found that it was still unavailable after downloading. 

After careful research, I found that I had to create a separate project file first, and then use "go mod init project_name" command. This command will automatically create a default "go.mod" file. 

Only when this file exists can I install the package successfully using "go get" command.








