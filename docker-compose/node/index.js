const mongoose = require("mongoose");

mongoose.connect("mongodb://mongo:27017/test", {}).then(() => console.log('Conectado!!')).catch((err => console.log(err)));