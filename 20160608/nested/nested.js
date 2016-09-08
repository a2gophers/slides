var express = require('express');

var app = express();
var level0 = express.Router();
var level1 = express.Router();


level1.get('/yolo/:bolo', (req, res) => {
    res.send(req.params.bolo);
});

level0.use('/nested', level1);

app.use('/api', level0);


app.listen("18015")

console.log('%s - starting up', new Date());
