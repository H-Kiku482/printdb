use crate::config::Config;
pub use diesel::{mysql::MysqlConnection, result::ConnectionResult, Connection};

pub fn get_connection(c: &Config) -> ConnectionResult<MysqlConnection> {
    let mut database_url = String::from("mysql://");
    database_url.push_str(&c.user);
    if !c.plane_password.is_empty() {
        database_url.push_str(":");
        database_url.push_str(&c.plane_password);
    }
    database_url.push_str("@");
    database_url.push_str(&c.host);
    database_url.push_str("/");
    database_url.push_str(&c.database);
    MysqlConnection::establish(&database_url)
}
