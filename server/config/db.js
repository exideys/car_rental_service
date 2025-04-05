const mysql = require('mysql2');

const connection = mysql.createConnection({
  host: '127.0.0.1',
  user: 'root',
  database: 'car_rent',
  password: '',
  port : 3306
});

connection.connect(err => {
  if (err) {
    return console.error('Erorr connection:', err);
  }
  console.log('Successfully connected to Mysql');
});