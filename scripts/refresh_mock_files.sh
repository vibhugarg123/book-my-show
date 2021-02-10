#!/usr/bin/env zsh
#This expects file names to be with/without underscore only and refreshes mocks which were earlier generated.
MOCK_FILES=$(find . -name "*mock.go")
echo "Mocks to be refreshed for "$MOCK_FILES

for PATH in $MOCK_FILES; do
  PACKAGE_NAME="$(echo $PATH | cut -d '/' -f2)"
  MOCK_FILE_NAME="$(echo $PATH | rev | cut -d '/' -f1 | rev)"
  FILE_NAME="$(echo $MOCK_FILE_NAME | sed 's/_mock//g')"
  echo "Mock generation : " " for package name: " $PACKAGE_NAME " mock file name: " $MOCK_FILE_NAME " mock generated for : " $FILE_NAME
  echo "Executing...mockgen --source=$FILE_NAME --destination=$PATH --package=$PACKAGE_NAME"
  mockgen --source=$FILE_NAME --destination=$PATH --package=$PACKAGE_NAME
done
