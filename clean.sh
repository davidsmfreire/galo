for d in */ ; do
    docker container rm -f "galo_${d%/}" 2>/dev/null
    docker image rm -f "galo_${d%/}" 2>/dev/null
done
