import * as assert from 'assert';
import fc from 'fast-check';

import {solveMeFirst} from './index';

const arbIntPair = fc.tuple(fc.integer(), fc.integer());

fc.assert(
  fc.property(arbIntPair, ([a, b]) => {
    const expected = a + b;
    const actual = solveMeFirst(a, b);

    assert.strictEqual(expected, actual, 'equal');

    console.log('pass\n', JSON.stringify({expected, actual}, null, 2), '\n');
  })
);
