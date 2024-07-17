#!/bin/env bash

platforms=(
	"windows/amd64"
	"windows/386"
	"windows/arm"
	"linux/amd64"
	"linux/arm64"
	"darwin/amd64"
)

for platform in "${platforms[@]}"; do
	platform_split=(${platform//\// })
	GOOS=${platform_split[0]}
	GOARCH=${platform_split[1]}
	
	output_name="porssisahkon-hinnan-ilmoittaja-$GOOS-$GOARCH"

	if [ $GOOS = "windows" ]; then
		output_name+='.exe'
	fi

	env GOOS=$GOOS GOARCH=$GOARCH go build -o dist/$output_name

	if [ $? -ne 0 ]; then
		echo "An error has occurred! Aborting the script execution..."
		exit 1
	else
		echo "Build succesful for $GOOS/$GOARCH"
	fi
done

