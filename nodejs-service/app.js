const express = require('express');
const bodyParser = require('body-parser');
const app = express();
const PORT = 3000; // You can choose any available port

app.use(bodyParser.json());

app.post('/initiate-project', (req, res) => {
  const requestBody = req.body;
  res.json(requestBody);
});

app.listen(PORT, () => {
  console.log(`Server is running on http://localhost:${PORT}`);
});
