# Start MongoDB in Docker
run-store:
    docker run --name my-mongo-dev -d -p 27017:27017 -v ~/mongo-data-dev:/data/db mongo

# Build the new-character-repository application
build:
    go build -o new-character-repo ./cmd/server

# Run the new-character-repository application
run-local:
    ./new-character-repo -env local.env
