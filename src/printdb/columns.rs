use diesel::{
    deserialize::{QueryableByName, Result},
    mysql::Mysql,
    row::NamedRow,
};

pub struct DatabaseInfo {
    pub table_name: String,
}

impl QueryableByName<Mysql> for DatabaseInfo {
    fn build<R: NamedRow<Mysql>>(row: &R) -> Result<Self> {
        Ok(DatabaseInfo {
            table_name: row.get("TABLE_NAME")?,
        })
    }
}

#[derive(Clone, Debug)]
pub struct TableInfo {
    pub table_name: String,
    pub records: Vec<Record>,
}

#[derive(Clone, Debug)]
pub struct Record {
    pub field: Option<String>,
    pub data_type: Option<String>,
    pub null: Option<String>,
    pub key: Option<String>,
    pub default: Option<String>,
    pub extra: Option<String>,
}

impl QueryableByName<Mysql> for Record {
    fn build<R: NamedRow<Mysql>>(row: &R) -> Result<Self> {
        Ok(Record {
            field: row.get("Field")?,
            data_type: row.get("Type")?,
            null: row.get("Null")?,
            key: row.get("Key")?,
            default: row.get("Default")?,
            extra: row.get("Extra")?,
        })
    }
}
