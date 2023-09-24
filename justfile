# https://just.systems/man/en/chapter_1.html

[private]
@default:
    echo "View file 'justfile' to see internals of any recipe."
    just --list --unsorted

# Runs golangci-lint.
@lint:
    golangci-lint run

# Runs coverage adn delivers results to web browser.
@cover:
    go test -coverprofile=coverage.out ./... && go tool cover -html=coverage.out