#!/bin/sh

echo "installing dependencies..."
go mod tidy

echo "building monolith..."
go build -v -o ./buildartifacts/monolith
echo "done!"
