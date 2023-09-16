const http = require('http');

const server = http.createServer((req, res) => {
  res.writeHead(200, { 'Content-Type': 'text/plain' });
  res.end('Hello, express!\n');
});

const port = process.env.PORT || 8080;
server.listen(port, () => {
  console.log(`Server running on http://localhost:${port}`);
});
