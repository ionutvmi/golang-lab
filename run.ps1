<#
.SYNOPSIS
This script contains useful commands for this project.

.DESCRIPTION
USAGE
    .\run.ps1 <command>

COMMANDS
    start       compile and run the project
    doc         generate the project documentation in the doc.md file
    help, -?    show this help message
#>
param(
    [Parameter(Position = 0)]
    [ValidateSet("start", "doc", "help")]
    [string]$Command
)

function Invoke-Help { 
    Get-Help $PSCommandPath 
}

function Invoke-Start {
    go run main.go
}

function Invoke-Doc {
    gomarkdoc --output doc.md ./...
    Write-Host "File doc.md was generated"
}

if (!$Command) {
    Invoke-Help
    exit
}

switch ($Command) {
    "start" { Invoke-Start }
    "doc" { Invoke-Doc }
    "help" { Command-Help }
}