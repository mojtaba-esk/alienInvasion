# Alien Invasion

Mad aliens are about to invade the earth and this code simulates the invasion.

The app will be given a map containing the names of cities in the non-existent world X. The map is in a file, with one city per line. The city name is first, followed by 1-4 directions (north, south, east, or west). Each one represents a road to another city that lies in that direction.

For example:

```
Foo north=Bar west=Baz south=Qu-ux
Bar south=Foo west=Bee
```

The city and each of the pairs are separated by a single space, and the directions are separated from their respective cities with an equals (=) sign.

## How to build and run the app

```
git clone https://github.com/mojtaba-esk/alienInvasion
cd alienInvasion
go build -o alienInvasion .
./alienInvasion
```

## Command line arguments

### `-a`: Action selector

This argument has two options:

- `start` Starts the invasion normally (this is the default value)
- `generate-world` Generates a new world

The app is capable of generating new random world with real names for cities.

### `-c`: Number of cities to generate a new world

This argument goes with the `generate-world` action where we can specify how large our world should be. The default value is `10`.

Example on how to generate a new world with `1000` cities and store it in a file:

```
./alienInvasion -a generate-world -c 1000 > myDatabase.txt
```

### `-d`: Path to the database file

```
./alienInvasion -d myDatabase.txt
```

The default path for database is `database.txt`

### `-n`: Number of aliens

We can specify the number of aliens with this argument. Example:

```
./alienInvasion -d myDatabase.txt -n 500
```

## Test

To perform a unit test, please execute the following command:

```
go test ./...
```

Verbose test:

```
go test ./... -v
```

## Source code documentation:

To see the source code documentation first run the following command:

```
godoc
```

Then head over to this URL: http://localhost:6060/pkg/alienInvasion/
