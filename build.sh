package="goaround"
platforms=("darwin/amd64"
"darwin/arm64"
"dragonfly/amd64"
"freebsd/386"
"freebsd/amd64"
"freebsd/arm"
"freebsd/arm64"
"illumos/amd64"
"linux/386"
"linux/amd64"
"linux/arm"
"linux/arm64"
"netbsd/386"
"netbsd/amd64"
"netbsd/arm"
"netbsd/arm64"
"openbsd/386"
"openbsd/amd64"
"openbsd/arm"
"openbsd/arm64"
"openbsd/mips64"
"plan9/386"
"plan9/amd64"
"plan9/arm"
"solaris/amd64"
"windows/386"
"windows/amd64"
"windows/arm"
"windows/arm64")

version="0.3"
package_split=(${package//\// })
package_name=${package_split[-1]}
for platform in "${platforms[@]}"
do
    bin_name=$package_name
    platform_split=(${platform//\// })
    GOOS=${platform_split[0]}
    GOARCH=${platform_split[1]}
    if [ $GOOS = "windows" ]; then
        bin_name+='.exe'
    fi

    env GOOS=$GOOS GOARCH=$GOARCH go build -o $bin_name main.go
    if [ $? -ne 0 ]; then
        echo 'An error has occurred! Aborting the script execution...'
        exit 1
    fi
    output_name=$package_name'-v'$version'-'$GOOS'-'$GOARCH.tar.gz
    tar -czvf 'dist/'$output_name $bin_name
    echo "Package $output_name is created."
    rm $bin_name
done