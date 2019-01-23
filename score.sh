#!/bin/bash

count=0
t=0
total=$(wc -l data.txt | cut -d " " -f1)
while read line
 do
  t=$((t+1))
  echo "$t from $total"
  wrong=$( echo $line | cut -d ">" -f1 )
  right=$( echo $line | cut -d ">" -f2 )
  goresult=$(go run goSpellcheck.go $wrong | sed "s/'/\'/g")
  suggestions=(${goresult})
  for ((i=0; i<${#suggestions[@]}; i++))
    do
    if [[ ${suggestions[$i]} == $right ]]
      then
      count=$((count+1))
      break
 	  fi
 	done
 done < data.txt

echo "got $count from $total"
