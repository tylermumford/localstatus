run:
	go run .

docs:
	cd docs && go generate

release VERSION:
	# TODO: Make sure there are no unstaged changes
	# TODO: Build and upload some binaries
	# TODO: Automatically get the notes from the changelog file
	gh release create "{{VERSION}}" --notes "See CHANGELOG file for notes" --prerelease
