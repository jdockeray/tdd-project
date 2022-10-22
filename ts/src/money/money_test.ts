import assert from "assert";

class Dollar {
  times(multiplier: number): Dollar {
    return new Dollar(this.amount * multiplier);
  }
  amount: number;
  constructor(amount: number) {
    this.amount = amount;
  }
}

const fiver = new Dollar(5);
const tenner = fiver.times(2);

assert.strictEqual(tenner.amount, 10);
