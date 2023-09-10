### Tasks
- What skills/knowledge that you had learned earlier in the semester were useful in this project
- What additional skills/knowledge you had to learn to complete your tasks
- The biggest challenge you faced during mini-project.

### Pre learned

I worked in the company for several years before, developed some commercial software, and used languages such as Python and Java, so when I was learning GO language to develop this project, the process and ideas of developing a software in the past can be applied to the development of this mini project.

The specific technologies involved, such as widgets creation, layout, and style adjustment under the `Fyne` framework (a GUI framework of GO language), are similar to those under Python's `tkinter` and `PyQt5`. Although the API will be slightly different, the overall UI layout, widget creation, and object-oriented development call ideas are the same.

For example, in `Fyne` framework, following codes will create a `Label` widget, and `Move` method will set it to a position of coordinate at `(200, 500)`

```
inputBookIDLabel := widget.NewLabel("input book ID")
inputBookIDLabel.Move(fyne.NewPos(200, 500))
```  

Also, `NewEntry` method will create a Text input widget as following: 

```
myEntry := widget.NewEntry()
myEntry.Move(fyne.NewPos(350, 500))
```  

### Additional skills

One of most popular GUI frameworks of GO language is Fyne. I studied basic use for this framework in order to finish this project well. 

Fyne is an open source cross-platform GUI (Graphical User Interface) toolkit and app development framework written in the Go programming language. It allows developers to easily create native applications that can run on desktop (Windows, macOS, Linux), web, and mobile platforms with a consistent look and feel.

One of the main benefits of Fyne is that it compiles natively to each platform it targets, which means that Fyne applications are fast and have a small memory footprint.

To create an entry and a layout in Fyne, you can follow these steps:

Import the Fyne package:
```
import "fyne.io/fyne/v2"
```

Create a new Fyne application:

```
app := fyne.NewApp()
```

Create a new window for the application:

```
win := app.NewWindow("My App")
```

Create a new entry widget:

```
entry := widget.NewEntry()
```

```
layout := container.NewVBox(entry)
```

```
win.SetContent(layout)
```

```
win.ShowAndRun()
```

This code will create a new Fyne application with a window titled "My App", an entry widget, and a layout that contains the entry widget. The layout is then set as the content of the window, and the window is shown.


### biggest challenge

The biggest challenge is Fyne's layout.

When I add different controls to the interface, the different controls appear in the position I want, and it is cumbersome to adjust.

When the position of appearance is finally adjusted, the size of the control will unexpectedly become very small due to different layout methods.

For example, we use a table control in our micro-project. When using some different layout methods, there are originally 10 rows of data in the table, but only 1 row is displayed, and the subsequent rows can only be seen in turn by Scroll to down.

This is the biggest challenge, and I use the layout `Move` method to solve this problem.

```
checkoutBtn.Resize(fyne.NewSize(200, 40))
checkoutBtn. Move(fyne. NewPos(600, 500))
```

Overall, this mini project is very useful and thanks so much.
