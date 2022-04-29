pub mod columns;
pub mod format;

use columns::*;
use diesel::{mysql::MysqlConnection, prelude::*, result::Error, sql_query};

pub fn get_table_info(
    con: &MysqlConnection,
    db_name: &String,
) -> Result<Vec<columns::TableInfo>, Error> {
    let mut sql = String::from(
        "SELECT `TABLE_NAME` FROM `information_schema`.`tables` WHERE `table_schema` = '",
    );
    sql.push_str(&db_name);
    sql.push_str("';");
    let tables_info: Vec<columns::DatabaseInfo> = match sql_query(sql).load(con) {
        Ok(v) => v,
        Err(e) => {
            return Err(e);
        }
    };

    let mut table_infos: Vec<TableInfo> = Vec::new();

    for v in tables_info {
        let mut sql = String::from("SHOW COLUMNS FROM `");
        sql.push_str(&db_name);
        sql.push_str("`.`");
        sql.push_str(&v.table_name);
        sql.push_str("`;");
        let table_info_array: Vec<columns::Record> = match sql_query(sql).load(con) {
            Ok(v) => v,
            Err(e) => {
                return Err(e);
            }
        };

        let table_info = TableInfo {
            table_name: String::from(&v.table_name),
            records: table_info_array,
        };

        table_infos.push(table_info.clone());
    }

    Ok(table_infos)
}
