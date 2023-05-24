run:
	go run .


docs:
	cd docs && go generate


release VERSION: build-all
	# TODO: Make sure there are no unstaged changes
	# TODO: Automatically get the notes from the changelog file
	gh release create "{{VERSION}}" --notes "See CHANGELOG file for notes" --prerelease {{cross}}/*


cross := 'cross-platform-builds'

build-all:
	rm -rf {{cross}}
	mkdir -p {{cross}}
	GOOS=windows GOARCH=amd64 go build -o {{cross}}/localstatus_windows_amd64.exe
	GOOS=linux   GOARCH=amd64 go build -o {{cross}}/localstatus_linux_amd64
	GOOS=darwin  GOARCH=amd64 go build -o {{cross}}/localstatus_mac_amd64
	GOOS=darwin  GOARCH=arm64 go build -o {{cross}}/localstatus_mac_arm64
