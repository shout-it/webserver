#!/bin/sh
CURRENT_PROJECT_ROOT=`git rev-parse --show-toplevel`
HOOKS=('pre-commit')
for i in "${HOOKS[@]}"
do
  cp $CURRENT_PROJECT_ROOT/scripts/$i $CURRENT_PROJECT_ROOT/.git/hooks/ &> /dev/null
  chmod +x $CURRENT_PROJECT_ROOT/.git/hooks/$i
  echo "$i hook successfully setup :)"
done
