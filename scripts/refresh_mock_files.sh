#!/usr/bin/env zsh
#This expects file names to be with/without underscore only
MOCK_FILES=$(find . -name "*mock.go")
echo $MOCK_FILES

for PATH in $MOCK_FILES; do
  PACKAGE_NAME="$(echo $PATH | cut -d '/' -f2)"
  MOCK_FILE_NAME="$(echo $PATH | rev | cut -d '/' -f1 | rev)"
  FILE_NAME="$(echo $MOCK_FILE_NAME | sed 's/_mock//g')"
  echo "Mock generation : " $PACKAGE_NAME $MOCK_FILE_NAME $FILE_NAME
  mockgen --source=$FILE_NAME --destination=$PATH --package=$PACKAGE_NAME
  #$FILE_NAME
done