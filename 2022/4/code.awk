BEGIN {
	FS=",|-"
}
/^/ {
	lo1 = $1
	hi1 = $2
	lo2 = $3
	hi2 = $4
	if ((lo1 <= lo2 && hi1 >= hi2) || (lo2 <= lo1 && hi2 >= hi1)) {overlaps++}
	if ((hi1 >= lo2 && lo1 <= hi2) || (hi2 >= lo1 && lo2 <= hi1)) {any++}
}
END {
	print any
}
