run:
	go run .


test:
	go test          # none here currently
	go test ./app    # none here currently
	go test ./checks


docs:
	# Note: The generated docs should be committed.
	cd docs && go generate


release VERSION: docs build-all _no-unstaged-changes
	# TODO: Automatically get the notes from the changelog file
	gh release create "{{VERSION}}" --draft --notes "See CHANGELOG file for notes" --prerelease {{cross}}/*


cross := 'cross-platform-builds'

build-all:
	rm -rf {{cross}}
	mkdir -p {{cross}}
	GOOS=windows GOARCH=amd64 go build -o {{cross}}/localstatus_windows_amd64.exe
	GOOS=linux   GOARCH=amd64 go build -o {{cross}}/localstatus_linux_amd64
	GOOS=darwin  GOARCH=amd64 go build -o {{cross}}/localstatus_mac_amd64
	GOOS=darwin  GOARCH=arm64 go build -o {{cross}}/localstatus_mac_arm64


_no-unstaged-changes:
	@echo Checking for unstaged changes...
	# https://stackoverflow.com/a/3879077/1076188
	git update-index --refresh
	git diff-index --quiet HEAD --


# Justfile settings
set ignore-comments := true
