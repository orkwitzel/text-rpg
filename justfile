default: build

bin := "bin"
binary := bin / "rpg"

build:
    mkdir -p {{bin}}
    go build -o {{binary}} .

run: build
    {{binary}}

vet:
    go vet ./...

test:
    go test ./...

check: vet build

clean:
    rm -rf {{bin}}
