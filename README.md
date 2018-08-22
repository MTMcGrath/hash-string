# go-hash-string
Simple web server that takes json POST with a key and value. it return the sha256 digest of the payload.

Example usage:

    curl -X POST --data '{"Name":"gouki"}' http://localhost:3000/messages
