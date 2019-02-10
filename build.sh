#!/bin/bash

# all files should be here
files=( 
	main.go
	parser/*.go
	tracker/*.go
	torrent/*.go
)

# script for formatting 
for i in "${files[@]}"; do
	gofmt -w "$i"
done 

# script for running test files
go test ./... -v

#checking goreportcard-cli for issues
set -e
report=$(goreportcard-cli)
startindex=$(($(echo $report | grep -b -o Issue | cut -d: -f1)+8))
endindex=$(($(echo $report|grep -b -o gofmt | cut -d: -f1)-1))
issuecount=${report:$startindex:$endindex-$startindex}
if [ $issuecount == "0" ]; then
        echo "goreportcard-cli passed"
fi
if [ $issuecount != 0 ]; then
        echo $issuecount" issues. Run \`goreportcard-cli -v\` to check. Ignore the issues from \`vendor\` directory"
	exit 1
fi

echo "Build Successful!"
