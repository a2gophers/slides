var express = require('express');

var app = express();

app.get('/yolo/:bolo', (req, res) => {
    res.send(req.params.bolo);
});

app.listen("18015")

console.log('%s - starting up', new Date());
