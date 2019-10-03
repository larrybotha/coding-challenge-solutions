import * as assert from 'assert';
import fc from 'fast-check';

import {simpleArraySum} from './index';

fc.assert(
  fc.property(fc.array(fc.integer()), xs => {
    let expected = 0;

    for (let i = 0; i < xs.length; i++) {
      expected += xs[i];
    }

    const actual = simpleArraySum(xs);

    try {
      assert.strictEqual(expected, actual);

      console.log('pass\n', JSON.stringify({expected, actual}, null, 2), '\n');

      return true;
    } catch (err) {
      return false;
    }
  })
);
