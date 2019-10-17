#!/usr/bin/env bash
echo 'Runing migrations...'
/snippetbox/migrate up

echo 'Deleting mysql-client...'
apk del mysql-client

echo 'Start application...'
/snippetbox/app