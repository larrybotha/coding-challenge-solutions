function getFloor(xs) {
  floor = 0

  for(i = 1; i <= length(xs); i++) {
    xs[i] == "(" ? floor++ : floor--
  }

  return floor
}

BEGIN {
  FS = ""
}

{
  # $0 is not available in BEGIN
  # variables don't need to be initialised
  # if delimiter is not provided, split uses FS
  split($0, xs)

  floor = getFloor(xs);
  # awk is 1-indexed
}

END { printf "Floor: %s\n", floor }
