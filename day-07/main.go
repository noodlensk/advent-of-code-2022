package main

import (
	"fmt"
	"strings"

	"github.com/noodlensk/advent-of-code-2022/util"
)

type FileSystemRecord struct {
	Name  string
	IsDir bool
	Size  int

	Items []*FileSystemRecord

	Parent *FileSystemRecord
}

func (f *FileSystemRecord) TotalSize() int {
	if !f.IsDir {
		return f.Size
	}

	totalSize := 0

	for _, i := range f.Items {
		totalSize += i.TotalSize()
	}

	return totalSize
}

func (f *FileSystemRecord) Print(level int) {
	for i := 0; i < level; i++ {
		fmt.Printf(" ")
	}

	if f.IsDir {
		fmt.Printf("- %s (dir)\n", f.Name)

		for _, i := range f.Items {
			i.Print(level + 1)
		}
	} else {
		fmt.Printf("- %s (file, size=%d)\n", f.Name, f.Size)
	}
}

type Command struct {
	Command string
	Arg     string
	Output  []string
}

func FSFromCommandLog(log []Command) *FileSystemRecord {
	fsRoot := &FileSystemRecord{Name: "/", IsDir: true}

	var currNode *FileSystemRecord

	for _, c := range log {
		switch c.Command {
		case "ls":
			for _, l := range c.Output {
				parts := strings.Split(l, " ")

				n := &FileSystemRecord{
					Name:   parts[1],
					Parent: currNode,
				}

				if parts[0] == "dir" {
					n.IsDir = true
				} else {
					n.Size = util.MustParseInt(parts[0])
				}

				currNode.Items = append(currNode.Items, n)
			}
		case "cd":
			switch c.Arg {
			case "..":
				currNode = currNode.Parent

				continue
			case "/":
				currNode = fsRoot
			default:
				for _, i := range currNode.Items {
					if i.Name == c.Arg {
						currNode = i

						break
					}
				}
			}
		}
	}

	return fsRoot
}

func Solution(lines []string) int {
	commands := parseCommandsLog(lines)
	fsRoot := FSFromCommandLog(commands)

	summ := 0

	queue := []*FileSystemRecord{fsRoot}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		for _, i := range curr.Items {
			if i.IsDir {
				queue = append(queue, i)
			}
		}

		if curr.TotalSize() < 100000 {
			summ += curr.TotalSize()
		}
	}

	return summ
}

func Solution2(lines []string) int {
	commands := parseCommandsLog(lines)
	fsRoot := FSFromCommandLog(commands)

	needToFreeSpaceSize := 30000000 - (70000000 - fsRoot.TotalSize())

	minDirToDeleteTotalSize := 30000000

	queue := []*FileSystemRecord{fsRoot}

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		for _, i := range curr.Items {
			if i.IsDir {
				queue = append(queue, i)
			}
		}

		if curr.TotalSize() > needToFreeSpaceSize && curr.TotalSize() < minDirToDeleteTotalSize {
			minDirToDeleteTotalSize = curr.TotalSize()
		}
	}

	return minDirToDeleteTotalSize
}

func main() {
	lines := util.MustReadInput()

	fmt.Println("Result: ", Solution(lines))
	fmt.Println("Result #2: ", Solution2(lines))
}

func parseCommandsLog(lines []string) []Command {
	var (
		commandsLog       []Command
		commandInProgress *Command
	)

	for _, l := range lines {
		if strings.HasPrefix(l, "$ ") { // new command
			if commandInProgress != nil {
				commandsLog = append(commandsLog, *commandInProgress)
			}

			cmdParts := strings.Split(l[2:], " ")
			commandInProgress = &Command{
				Command: cmdParts[0],
				Arg:     strings.Join(cmdParts[1:], " "),
			}
		} else {
			commandInProgress.Output = append(commandInProgress.Output, l)
		}
	}

	commandsLog = append(commandsLog, *commandInProgress)

	return commandsLog
}
