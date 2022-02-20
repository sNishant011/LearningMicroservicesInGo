# Requirements

- Go compiler

# Running

- go run filename.go

# What learned?

- Parsing struct type to JSON
- using json new encoder method instead of Marshal
  -- marshall traversel recursively to a string
  -- encode traverse to a stream, it uses io writer (also a bit quicker)
- MAKING RESTFUL API
  -- seperating requests
