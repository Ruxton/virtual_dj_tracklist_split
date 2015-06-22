package main;

import (
  "github.com/mattn/go-colorable"
  flag "launchpad.net/gnuflag"
  "os"
  "bufio"
  "github.com/ruxton/mixcloud/term"
  "fmt"
  "regexp"
  "strings"
)

var VERSION string
var MINVERSION string

var STD_OUT = bufio.NewWriter(colorable.NewColorableStdout())
var STD_ERR = bufio.NewWriter(colorable.NewColorableStderr())
var STD_IN = bufio.NewReader(os.Stdin)

var VDJ_TRACKLIST_HEADER_MATCH = `^VirtualDJ History - ((20\d{2})\/(\d{2})\/(\d{2}))$`
var VDJ_TRACKLIST_SPLIT_MATCH = "------------------------------"

var FILE_OUTPUT_PRE = "vdj"

var aboutFlag = flag.Bool("about", false, "About the application")
var onlyLastFlag = flag.Bool("last", false, "Only output the most recent tracklist")
var outputPreFlag = flag.String("filepre", "", "A pre-text for the output filename, defaults to 'vdj' eg. vdj-20150102.txt")
var outputFolderFlag = flag.String("outputfolder", "", "The folder to output individual files to, defaults to current directory")

func main() {
  flag.Parse(true)

  showWelcomeMessage()
  if *aboutFlag == true {
    showAboutMessage()
    os.Exit(0)
  }
  if *outputPreFlag != "" {
    FILE_OUTPUT_PRE = *outputPreFlag
  }

  args := os.Args[1:]
  tracklistFile := args[len(args)-1:][0]

  tracklists,readError := readLines(tracklistFile)
  if(readError != nil) {
    OutputError(fmt.Sprintf("Unable to read file %s",tracklistFile))
    os.Exit(2)
  }

  tracklistCollection,tracklistDateKeys := parseTracklists(tracklists)


  if *onlyLastFlag == true {
    OutputMessage(term.Green+"Outputting last tracklist"+term.Reset+"\n")
    outputLastTracklist(tracklistCollection,tracklistDateKeys)
  } else {
    outputAllTracklists(tracklistCollection,tracklistDateKeys)
  }

}

func outputAllTracklists(tracklistCollection map[string][]string, keys []string) {
  for _,date := range keys {
    outputTracklist(date,tracklistCollection[date])
  }
}

func outputLastTracklist(tracklistCollection map[string][]string, keys []string) {
  last_date := keys[len(keys)-1:][0]
  lines := tracklistCollection[last_date]
  outputTracklist(last_date,lines)
}

func outputTracklist(date string, lines []string) {
  OutputMessage(fmt.Sprintf("Outputting %s\n",date))
  filename := fmt.Sprintf("./%s-%s.txt",FILE_OUTPUT_PRE,date)
  err := writeLines(lines, filename)
  if(err != nil) {
    OutputError(fmt.Sprintf("Error: %v",err))
  }
}

func parseTracklists(tracklists []string) (map[string][]string,[]string) {
  current_date := ""
  matcher := regexp.MustCompile(VDJ_TRACKLIST_HEADER_MATCH)
  lists := map[string][]string{}
  key_order := []string{}

  iterator := 0

  for _,line := range tracklists {
    location := matcher.FindStringSubmatchIndex(line)
    if(location != nil) {
      new_date := strings.Replace(line[location[2]:location[3]],"/","",-1)
      if(new_date == current_date) {
        iterator++
        current_date = fmt.Sprintf("%s-%d",new_date,iterator)
      } else {
        iterator = 0
        current_date = new_date
      }
      key_order =  append(key_order, current_date)
    } else {
      if(current_date != "") {
        if (line != VDJ_TRACKLIST_SPLIT_MATCH && line != "") {
          lists[current_date] = append(lists[current_date], line)
        }
      }
    }
  }

  return lists,key_order
}

func showWelcomeMessage() {
	OutputMessage(term.Green + "Virtual DJ Tracklist Split v" + VERSION + term.Reset + "\n\n")
}

func showAboutMessage() {
	OutputMessage(fmt.Sprintf("Build Number: %s\n", MINVERSION))
	OutputMessage("Created by: Greg Tangey (http://ignite.digitalignition.net/)\n")
	OutputMessage("Website: http://www.rhythmandpoetry.net/\n")
  OutputMessage("\nParses an splits Virtual DJ's tracklist.txt into individual files")
}

func OutputError(message string) {
	STD_ERR.WriteString(term.Bold + term.Red + message + term.Reset + "\n")
	STD_ERR.Flush()
}

func OutputMessage(message string) {
	STD_OUT.WriteString(message)
	STD_OUT.Flush()
}

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
  file, err := os.Open(path)
  if err != nil {
    return nil, err
  }
  defer file.Close()

  var lines []string
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    lines = append(lines, scanner.Text())
  }
  return lines, scanner.Err()
}

// writeLines writes the lines to the given file.
func writeLines(lines []string, path string) error {
  file, err := os.Create(path)
  if err != nil {
    return err
  }
  defer file.Close()

  w := bufio.NewWriter(file)
  for _, line := range lines {
    fmt.Fprintln(w, line)
  }
  return w.Flush()
}
