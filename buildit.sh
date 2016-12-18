go_back() {
 cd - > /dev/null
}

build_protoc() {
 echo "Building proto files..."
 export PATH=$PATH:$SCRIPTDIR/bin
 cd $SCRIPTDIR/bin
 protoc -I$SCRIPTDIR/src/server/rpcdef/ --go_out=plugins=grpc:$SCRIPTDIR/src/server/rpcdef/ $SCRIPTDIR/src/server/rpcdef/serverrpc.proto
 go_back
}


build_server() {
 echo "Building server..."
 echo "Cleaning up old server..."
 rm -f $OUTDIR/server
 cd src/server/server
 go build -gcflags "-N -l" -o $OUTDIR/server
 go_back
}

# Build everything. At a later date we could give command line options to
# build only specific things
main() {
 build_protoc
 build_server
}

# Get the parent directory of this script and store all the binaries in the
# bin directory
SCRIPTDIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
OUTDIR=$SCRIPTDIR/bin/

cd $SCRIPTDIR
export GOPATH=$SCRIPTDIR
main