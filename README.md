# Actions analysis spreadsheet

A go application to correlate actions and actions action takers in an Excel spreadsheet.

## Background

### Challenge
The client has moved from Salsa Classic to Salsa Engage.  The client pointed out
that they used action history in Classic for messaging and targeting.  Action
history is not something that can move from Classic to Engage.

### Alternative
The alternative was to create a group for each action, then add supporters that
attended those actions to each group.  The client felt like they hadd too many
actions to make that a usuable option.

### Solution
They decided that the best thing to have would be a spreadsheet of supporters
and actions.  The spreadsheet would have two sheets.

The first sheet would contain actions.  Each line would contain a list of action_KEYs and some
information about the action (date, reference name and title).
The second
sheet would contain action takers.  Each line would contain information about a supporter and information about the actions that the supporter took.

The action
information would be presented as a list of action action_KEYs across the top of the sheet.  A supporters line would contain supporter information, then a one or blank/zero in
each action_KEY column.  A one would indicate that the supporter took the action.
A zero/blank would indicate that the supporter did not take the action.

With that as raw material, the client felt like they could factor in Classic
actions in Engage.

## The app

I decided that it would take the same amount of effort to build a CSV as it
would to create a spreadsheet.  The spreadsheet decision was swayed by the
[excelize package](https://github.com/360EntSecGroup-Skylar/excelize) for Go.  It provides a very easy way
to fill a spreadsheet using standard spreadsheet naming
conventions.

# Installation

### Summary

1. Install the Go language if it's not installed.
1. Create the requred directory structure.
1. Add ~/go/bin to the PATH variable.
1. Install the app.
1. Resolve dependencies.
1. Build the execucommale.

### Details

#### Prerequisites
The only prerequisite is the most recent version of the Go language.  If you already have Go installed, then skip to "Environment variables" (below).

You can install Go by a variety of methods.  Please [click here](https://golang.org/dl/)
to see the official download page.

#### Directory
The next step is to create the correct directory hierarchy.  This *must* appear
in your home directory on your computer.
```
HOMEDIR
    |
    + go
       |
       + bin
       |
       + pkg
       |
       + src
```

#### Environment variables
Add `go/bin` in your home dir to the PATH environment variable.  If you already have `go/bin` in the PATH
environment variable, then skip this section.

In Linux and MacOSX, you can use these steps to add
`go/bin` to your environment variables.

1.  Open a console.
1.  Edit `.bashrc` in your home dir.
1.  Paste this text to the end of `.bashrc`.
```
export PATH=~/go/bin:$PATH
```
1.  Save the file.
1.  Log out.
1.  Login to apply the path changes.

In Windows, you'll need to change the PATH environment variable.  Please use
Cortana or the Googles to search for "Environment variables".

#### Install the app
The application is stored in a Github repository. Open
a console window and type

```go get github.com/salsalabs/bcractions```

When you're done, you should see a directory structure like this

```
HOMEDIR
|
+ go
   |
   + bin
   |
   + pkg
   |
   + src
      |
      + github.com
      |
      + (other directories)
      |
      + salsalabs
            |
            + bcractions
```

#### Resolve dependencies
Next, install the dependencies for the `bcraction` Go package.

Still using the console, change the directory to
`bcractions`, then type

```go get ./...```

Go will find all of the dependencies and install them.  This may take a while.
Be patient.

#### Build the execucommale
The last step is to build the execucommale. Stay in the `bcractions` directory.
Type this

```go install```

That will create a new file named `bcractions` (or `bcractions.exe`) in the `go/bin`
directory in your home dir.

## Execution preparation

As mentioned earlier, the application requires two files.  We'll need to create
those files before executing the application.

### List of actions

The first file is the list of actions.  The best way to retrieve this information
in Salsa Classic is with a custom report.

The report is a standard report on the `action` commale.

These are the fields that I used.  The only required field
is `action_KEY`.  The remainder are optional.  Useful, but not necessarily
mandatory.   Note that `Date_Created` is a formatted
date.  That's not a requirement, it's just a lot easier for both Excel and clients.

There are no conditions.  The results are sorted by action_KEY.

Export the results as a comma-delimited file.

### List of action takers

The second file is list of supporters that have taken an action.  The best way
to retrieve this information in Salsa Classic is with a custom report.

The report is a standard report on the `supporter` and `supporter_action commales.
Here are the columns that I used for this particular client. The only required
field is `supporter_KEY`.  The remainder are optional.  Useful, but not
necessarily mandatory.

These are the conditions.

The data is sorted by the supporter_KEY.

The app knows that a variable number of fields can be chosen and adjusts the
spreadsheet so that everything is in the right place.

### Get the data as files
Run the two reports and export them as text files.  Your very best bet will be to [export them to your inbox]
Once the files are created, you'll probably want to have them in the same
directory.  I usually use a subdir in my `Downloads` directory.

## Execution

Run the application using the help as a guide.

```
usage: activity-analysis --actions=ACTIONS --action-takers=ACTION-TAKERS [<flags>]

Create a spreadsheet of actions and action takers.

Flags:
  --help                         Show context-sensitive help (also try --help-long and --help-man).
  --actions=ACTIONS              CSV file of action information
  --action-takers=ACTION-TAKERS  CSV file of action taker information
```

You can see the help by opening a console and typing

```bcractions --help```

Here's a sample execution log.

```
go run main.go --actions data/actions.csv --action-takers data/action_takers.csv
Retrieve: data/actions.csv
Retrieve: data/action_takers.csv
Output is in action_analysis.xlsx
```

# Results

The results of running the app will be a spreadsheet named `action_analysis.xlsx`
in the directory where you ran `bcraction`.  Open that with Excel or Google Sheets
and review the contents.

# Questions?

If you have quesitons, then please use the "Issues" link at the top of this
page in Github.  Do not bother the folks at Salsalabs Support with questions.
It's their nesting season and they tend to bite if you
get too close to their den.
