import assert from "assert";

class Money {
  amount: number;
  currency: string;
  divide(divider: number) {
    return new Money(this.amount / divider, this.currency);
  }
  times(multiplier: number): Money {
    return new Money(this.amount * multiplier, this.currency);
  }
  constructor(amount: number, currency: string) {
    this.amount = amount;
    this.currency = currency;
  }
}

/**
 *
 * Chapter 1: The money problem
 *
 * */

const fiveDollars = new Money(5, "USD");
const tenDollars = new Money(10, "USD");

assert.deepStrictEqual(fiveDollars.times(2), tenDollars);

/**
 *
 * Chapter 2: Multi currency money
 *
 * */

const fiveEuros = new Money(5, "EUR");
const tenEuros = new Money(10, "EUR");

assert.deepStrictEqual(fiveEuros.times(2), tenEuros);

/**
 *
 * Chapter 2: Division
 *
 * */

const originalMoney = new Money(4002, "KRW");
const actualMoneyAfterDivision = originalMoney.divide(4);
const expectedMoneyAfterDivision = new Money(1000.5, "KRW");

assert.deepStrictEqual(actualMoneyAfterDivision, expectedMoneyAfterDivision);
