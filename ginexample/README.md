# Gin Example

A small server created with Gin

## Usage

```bash
# Run app
go run .

# Get all albums
curl http://localhost:8080/albums \
    --header "Content-Type: application/json" \
    --request "GET"

# Add an album
curl http://localhost:8080/albums \
    --include \
    --header "Content-Type: application/json" \
    --request "POST" \
    --data '{"id": "4","title": "The Modern Sound of Betty Carter","artist": "Betty Carter","price": 49.99}'

# Delete an album
curl http://localhost:8080/albums/1 \
    --header "Content-Type: application/json" \
    --request "DELETE"
```
