#!/usr/bin/env bash
#This expects file names to be with/without underscore only and refreshes mocks which were earlier generated.
MOCK_FILES=$(find . -name "*mock.go")
for PATH_MOCK_FILE in $MOCK_FILES; do
  PACKAGE_NAME="$(echo $PATH_MOCK_FILE | cut -d '/' -f2)"
  MOCK_FILE_NAME="$(echo $PATH_MOCK_FILE | rev | cut -d '/' -f1 | rev)"
  FILE_NAME="$(echo $MOCK_FILE_NAME | sed 's/_mock//g')"
  echo "refreshing mockgen --source=$FILE_NAME --destination=$PATH_MOCK_FILE --package=$PACKAGE_NAME"
  ~/go/bin/mockgen --source=$FILE_NAME --destination=$PATH_MOCK_FILE --package=$PACKAGE_NAME
done
