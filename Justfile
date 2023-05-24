run:
	go run .

docs:
	cd docs && go generate

release VERSION:
	gh release create "{{VERSION}}" --notes "See CHANGELOG file for notes" --prerelease
