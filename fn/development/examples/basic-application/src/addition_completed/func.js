process.stdin.resume();
process.stdin.setEncoding('utf8');
process.stdin.on('data', (_data) => {
  const data = JSON.parse(_data);
  console.log(
    data.reduce((prev, curr) => (prev + curr), 0)
  );
});
