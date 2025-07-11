import { createServer } from 'node:http';

const server = createServer((req, res) => {
    res.writeHead(200, { 'content-type': 'text/plain' });
    res.end('hello world!\n');
})

server.listen(3000, '127.0.0.1', () => {
    console.log('listening on 127.0.0.1:3000');
});