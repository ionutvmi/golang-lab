<#
.SYNOPSIS
This script contains useful commands for this project.

.DESCRIPTION
USAGE
    .\run.ps1 <command>

COMMANDS
    start       compile and run the project
    doc         generate the project documentation in the doc.md file
    help        show this help message
    -?          show full help

.EXAMPLE
./run.ps1 start

#>
param(
    [Parameter(Position = 0)]
    [ValidateSet("start", "doc", "help")]
    [string]$Command
)

function Invoke-Help { 
    $help = Get-Help $PSCommandPath
    Write-Output $help.SYNOPSIS
    Write-Output ""
    Write-Output $help.DESCRIPTION.Text
}

function Invoke-Start {
    go build -o dist/

    & "./dist/golang-lab" variables

    & "./dist/golang-lab" functions 44 55
    & "./dist/golang-lab" functions 44 0

    & "./dist/golang-lab" methods op:count start:quick stop:the file:/tmp/file1.txt
    & "./dist/golang-lab" methods op:watch stop:fox file:/tmp/file3.txt
}

function Invoke-Doc {
    gomarkdoc --output doc.md ./...
    Write-Output "File doc.md was generated"
}

if (!$Command) {
    Invoke-Help
    exit
}

switch ($Command) {
    "start" { Invoke-Start }
    "doc" { Invoke-Doc }
    "help" { Invoke-Help }
}

# Credits
# CLI Template https://kevinareed.com/2021/04/14/creating-a-command-based-cli-in-powershell/

