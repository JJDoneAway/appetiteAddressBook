#!/bin/sh

# Start the first process
/app &

# Start the second process
/usr/share/env_injector /usr/share/nginx/html/ /usr/share/.env.odj && exec nginx -g 'daemon off;' &

# Wait for any process to exit
wait -n

# Exit with status of process that exited first
exit $?
