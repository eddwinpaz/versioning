### Versioning

### How to use it?

Simple just install go and  execute

```
./versioning minor
./versioning patch
./versioning major
```

### Important
- Remember to place the version.txt at the level of the compiled versioning code.

### Want to collaborate?
- Make a pull request


### Increment version using Docker
(remember you can only use major, minor, patch)
run: 
```
sh version.sh minor
```

version.sh
```
#!/usr/bin/env bash

./versioning $1

VERSION=$(cat ./version.txt)
echo "🆚 New version: $VERSION"

# Build the image, passing in the version number as a build argument
PACKAGE_NAME=my-app:$version

echo "🏗️  Building $PACKAGE_NAME"
docker build -t package_name .

echo "📍 Pushing $PACKAGE_NAME" 
docker push $package_name
```
