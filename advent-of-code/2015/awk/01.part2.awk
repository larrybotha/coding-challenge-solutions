function getPos(xs) {
  pos = 1
  floor = 0

  # awk is 1-indexed
  for(i = 1; i <= length(xs); i++) {
    xs[i] == "(" ? floor++ : floor--

    if (floor == -1) {
      pos = i
      break
    }
  }

  return pos
}

BEGIN {
  FS = ""
}

{
  # $0 is not available in BEGIN
  # variables don't need to be initialised
  # if delimiter is not provided, split uses FS
  split($0, xs)

  pos = getPos(xs)
}

END {
  printf "Character position: %s\n", pos
}
