#!/bin/bash

declare -A results

for ((i = 0; i < 10; i++)); {
	while read -r line; do
		[[ $line != "Benchmark"* ]] && continue
	
		name=$(cut -d' ' -f1  <<< "$line")
		ns=$(awk '{print $3}' <<< "$line")
	
		results["$name"]+="$ns"$'\t'
	done < <(go test -bench=. -benchtime=1000000x)
}

for i in "${!results[@]}"; {
	echo -e "$i\t${results[$i]}"
}
