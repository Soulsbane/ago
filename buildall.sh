#!/usr/bin/env bash

platforms=("windows/amd64" "linux/amd64")

for platform in "${platforms[@]}"; do
	platform_split=(${platform//\// })
	export GOOS=${platform_split[0]}
	export GOARCH=${platform_split[1]}

	if [ $GOOS = "windows" ]; then
		output_name+='.exe'
	fi

	GOOS=$GOOS GOARCH=$GOARCH go build -C ./cmd/ago/

	if [ $? -ne 0 ]; then
		echo 'An error has occurred! Aborting the script execution...'
		exit 1
	fi
done
