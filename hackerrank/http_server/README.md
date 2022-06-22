# System init
1. Change input stream reader to a local file
2. export OUTPUT_PATH="tests/results.txt"

## Specs
### POST
post json lake info to the db.
{
    "type": "post",
    "payload": {"id": "id00", "name": "Name of the lake", "area": 452001}
}

### DELETE
delete the lake by ID, return 404 if not present.
{
    "type": "delete",
    "payload": "id00"
}

### getHandler
take a Lake from the DB by ID and returns it to the caller.
If it is not found it returns "404 Not Found"
{
    "type": "delete",
    "payload": "id00"
}

## Possible improvements
- Concurrent client requests
- Protect DB behind a Mutex and some handling functions
- Separate client and server modules