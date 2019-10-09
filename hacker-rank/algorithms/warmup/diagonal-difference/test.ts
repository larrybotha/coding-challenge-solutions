import * as assert from 'assert';
import fc from 'fast-check';

import {diagonalDifference} from './index';

const arbSquare = fc.integer(0, 4).map(x => x ** 2);
const arbMatrix = arbSquare.chain(x => {
  return fc.array(fc.integer(-9, 9), x, x);
});

const printMatrix = xs => {
  const n = Math.sqrt(xs.length);
  const rows = xs.reduce((acc, x, i) => {
    const index = acc.length === 0 ? 0 : acc.lengrh - 1;
    let row = acc[index] && acc[index].length < n ? acc[index] : [];
    row = row.concat(x);
    acc[index] = row;

    return acc;
  }, []);

  rows.map(row => console.log(row.join(' | ')));
};

fc.assert(
  fc.property(arbMatrix, xs => {
    printMatrix(xs);

    console.log(diagonalDifference(xs));
    console.log('==============================\n');

    return true;
  })
);
