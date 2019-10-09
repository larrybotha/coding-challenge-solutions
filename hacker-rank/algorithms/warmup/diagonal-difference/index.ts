function sum(a, b) {
  return a + b;
}

function diagonalDifference(xs) {
  console.log(xs);
  const n = Math.sqrt(xs.length);
  const rows = xs.reduce((acc, x, i) => {
    const row = acc[i % n] || [];
    row[i % n] = x;
    acc[i % n] = row;

    return acc;
  }, []);
  const primaryDiag = rows.map((row, i) => row[i]).reduce(sum, 0);
  const secondaryDiag = rows
    .map((row, i) => {
      const v = row.slice().reverse()[i];
      console.log(row[i]);

      return v;
    })
    .reduce(sum, 0);
  console.log(primaryDiag, secondaryDiag);

  return Math.abs(primaryDiag - secondaryDiag);
}

export {diagonalDifference};
