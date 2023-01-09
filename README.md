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

```
VERSION=$(cat ./version.txt)
echo "Current version: $VERSION"

# Build the image, passing in the version number as a build argument
PACKAGE_NAME=my-app:$version

docker build -t package_name .

docker push $package_name
```
