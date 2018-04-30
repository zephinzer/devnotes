const Benchmark = require('benchmark');
const funcA = require('./funcA');
const funcB = require('./funcB');

const suite = new Benchmark.Suite();

suite
  .add('funcA', () => {
    funcA({name: 'hey'});
  })
  .add('funcB', () => {
    funcB({name: 'hey'});
  })
  .on('complete', function() {
    console.log(`Fastest: `, this.filter('fastest').map('name'));
  })
.run({async: true});
