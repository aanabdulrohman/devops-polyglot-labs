const express = require('express');
const app = express();
const port = process.env.PORT || 3000;

app.get('/todos', (req, res) => {
  res.json([ { id: 1, task: 'Setup Docker Swarm Security', completed: false } ]);
});

app.listen(port, () => console.log(`Server is running on port ${port}`));