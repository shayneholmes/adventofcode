/^/ {
	# Win points
	if ($2 == "X") {points += 0}
	if ($2 == "Y") {points += 3}
	if ($2 == "Z") {points += 6}
	# Shape points


	if ($2 == "X" && $1 == "A") {points += 3}
	if ($2 == "X" && $1 == "B") {points += 1}
	if ($2 == "X" && $1 == "C") {points += 2}
	if ($2 == "Y" && $1 == "A") {points += 1}
	if ($2 == "Y" && $1 == "B") {points += 2}
	if ($2 == "Y" && $1 == "C") {points += 3}
	if ($2 == "Z" && $1 == "A") {points += 2}
	if ($2 == "Z" && $1 == "B") {points += 3}
	if ($2 == "Z" && $1 == "C") {points += 1}
}

END {
	print points
}
