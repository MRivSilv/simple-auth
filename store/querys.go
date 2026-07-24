package store

// REGISTER
var registerUser string = "INSERT INTO users (email, username, password_hash) VALUES($1, $2, $3) RETURNING user_id"
var registerApp string = "INSERT INTO apps (app_name, description) VALUES($1, $2) returning app_id"
var registerUserInApp string = "INSERT INTO user_apps (user_id, app_id) VALUES($1, $2)"

// GETTERS
var getUser string = "SELECT user_id, email, username, password_hash FROM users WHERE user_id = $1"
var getUserByEmail = "SELECT user_id, email, username, password_hash FROM users WHERE email = $1"

// DELETE
var deleteUser string = "DELETE from users where user_id = $1"
var deleteApp string = "DELETE from apps where app_id = $1"
var deleteUserFromApp string = "DELETE from user_apps where user_id = $1 and app_id = $2"

// CHECKERS
var checkEmail string = "SELECT EXISTS(SELECT 1 FROM users WHERE email = $1)"
var checkAppName string = "SELECT EXISTS(SELECT 1 FROM apps WHERE app_name = $1)"
var checkAppID string = "SELECT EXISTS(SELECT 1 FROM apps WHERE app_id = $1)"
var checkUserID string = "SELECT EXISTS(SELECT 1 FROM users WHERE user_id = $1)"
