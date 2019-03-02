// Require the serialport node module
var SerialPort = require('serialport');
const ReadLine = require('@serialport/parser-readline')
//var SerialPort = serialport.SerialPort;
// Open the port
var port = new SerialPort('/dev/ttyACM0', {
    baudRate: 9600
});

const parser = port.pipe(new ReadLine())
// Read the port data

parser.on('data', function(data) {
    console.log(data);
    if (data < 100) {
        updateClient(data)
    }
});


var request = require('request');
 function updateClient(postData){
            var clientServerOptions = {
                uri: 'http://10.35.64.127:8080/setValue?Id=truck2&Value='+postData,
                //body: JSON.stringify({Id: 'truck1', Value: postData}),
                method: 'GET',
                headers: {
                    'Content-Type': 'application/json'
                }
            }
            request(clientServerOptions, function (error, response) {
                //console.log(error);
                return;
            });
        }