# Mongo

mongo is a wrapper around the <a href="https://github.com/mongodb/mongo-go-driver">MongoDB driver</a>, designed to streamline the process of connecting to a MongoDB database.

## Installation

```zsh
go get github.com/gosuit/mongo
```

## Features

- Easy configuration for MongoDB connection.
- Simple interface for obtaining a MongoDB database client.

## Usage

```golang
import (
    "context"
    "log"

    "github.com/gosuit/mongo"
)

func main() {
    cfg := &mongo.Config{
        URI:      "mongodb://localhost:27017",
        Database: "yourdatabasename",
    }

    ctx := context.Background()

    db, err := mongo.New(ctx, cfg)
    if err != nil {
        log.Fatalf("Failed to connect to MongoDB: %v", err)
    }

    // Use the db variable to interact with your MongoDB database.
}
```

## Contributing

Contributions are welcome! Please feel free to submit a pull request or open an issue for any enhancements or bug fixes.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.