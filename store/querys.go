package store

var registerUser string = "INSERT INTO users (email, username, password_hash, app_ids) VALUES($1, $2, $3, $4) RETURNING user_id"
var getUser string = "SELECT user_id, email, username, password_hash, app_ids FROM users WHERE user_id = $1"
var getUserByEmail = "SELECT user_id, email, username, password_hash, app_ids FROM users WHERE email = $1"
var deleteUser string = "DELETE from users where user_id = $1"
