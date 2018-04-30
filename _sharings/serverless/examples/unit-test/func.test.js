const {expect} = require('chai');
const func = require('./func');

describe('unit-test', () => {
  it('works', () => {
    expect(func()).to.have.key('message');
    expect(func().message).to.eql('Hello World');
  });

  context('with arguments', () => {
    it('works too', () => {
      expect(func()).to.have.key('message');
      expect(func({name: 'Gandalf'}).message).to.eql('Hello Gandalf');
    });
  });
});
