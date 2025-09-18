// index.js
const express = require('express');
const _ = require('lodash');

const app = express();
const port = process.env.PORT || 3000;

app.get('/', (req, res) => { res.send('Hello, World!'); });

app.get('/lodash-example', (req, res) => {
  const numbers = [1, 2, 3, 4, 5];
  const doubled = _.map(numbers, n => n * 2);
  res.send(`Doubled Numbers: ${doubled}`);
});

app.listen(port, () => {
  console.log(`Server is running on http://localhost:${port}`);
});
