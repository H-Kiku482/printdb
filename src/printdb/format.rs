use super::columns::TableInfo;
use std::fs::File;
use std::io::{Error, ErrorKind, Write};

struct List {
    field_max: u32,
    type_max: u32,
    null_max: u32,
    key_max: u32,
    default_max: u32,
    extra_max: u32,
}

impl List {
    fn new() -> List {
        List {
            field_max: 5,
            type_max: 4,
            null_max: 4,
            key_max: 3,
            default_max: 7,
            extra_max: 5,
        }
    }

    fn get_horizonal_border(&self) -> String {
        let mut s = String::new();
        s.push_str("+-");
        for _ in 0..self.field_max {
            s.push_str("-");
        }
        s.push_str("-+-");
        for _ in 0..self.type_max {
            s.push_str("-");
        }
        s.push_str("-+-");
        for _ in 0..self.null_max {
            s.push_str("-");
        }
        s.push_str("-+-");
        for _ in 0..self.key_max {
            s.push_str("-");
        }
        s.push_str("-+-");
        for _ in 0..self.default_max {
            s.push_str("-");
        }
        s.push_str("-+-");
        for _ in 0..self.extra_max {
            s.push_str("-");
        }
        s.push_str("-+\n");
        s
    }

    fn get_header(&self) -> String {
        let mut s = String::new();
        s.push_str(&self.get_horizonal_border());
        s.push_str("| Field");
        for _ in 0..(self.field_max - 5) {
            s.push(' ');
        }
        s.push_str(" | Type");
        for _ in 0..(self.type_max - 4) {
            s.push(' ');
        }
        s.push_str(" | Null");
        for _ in 0..(self.null_max - 4) {
            s.push(' ');
        }
        s.push_str(" | Key");
        for _ in 0..(self.key_max - 3) {
            s.push(' ');
        }
        s.push_str(" | Default");
        for _ in 0..(self.default_max - 7) {
            s.push(' ');
        }
        s.push_str(" | Extra");
        for _ in 0..(self.extra_max - 5) {
            s.push(' ');
        }
        s.push_str(" |\n");
        s.push_str(&self.get_horizonal_border());
        s
    }
}

pub fn get_std_out(table_infos: &Vec<TableInfo>) -> String {
    let mut list = String::new();
    for table_info in table_infos {
        list.push_str(&get_list_string(&table_info));
        list.push('\n');
    }
    list.pop();
    list
}

pub fn write_as_text(
    table_infos: &Vec<TableInfo>,
    file_path: &String,
    overwrite_flg: bool,
) -> Result<(), Error> {
    let mut file = match File::options()
        .read(true)
        .write(true)
        .append(!overwrite_flg)
        .truncate(overwrite_flg)
        .open(&file_path)
    {
        Ok(f) => f,
        Err(ref e) if e.kind() == ErrorKind::NotFound => match File::create(&file_path) {
            Ok(created_f) => created_f,
            Err(e) => {
                return Err(e);
            }
        },
        Err(e) => {
            return Err(e);
        }
    };

    let s = get_std_out(table_infos);
    match file.write_all(s.as_bytes()) {
        Err(e) => {
            return Err(e);
        }
        _ => {}
    };

    match file.flush() {
        Ok(_) => {
            return Ok(());
        }
        Err(e) => {
            return Err(e);
        }
    };
}

pub fn write_as_markdown(
    database_name: &String,
    table_infos: &Vec<TableInfo>,
    file_path: &String,
    overwrite_flg: bool,
) -> Result<(), Error> {
    let mut file = match File::options()
        .read(true)
        .write(true)
        .append(!overwrite_flg)
        .truncate(overwrite_flg)
        .open(&file_path)
    {
        Ok(f) => f,
        Err(ref e) if e.kind() == ErrorKind::NotFound => match File::create(&file_path) {
            Ok(created_f) => created_f,
            Err(e) => {
                return Err(e);
            }
        },
        Err(e) => {
            return Err(e);
        }
    };

    let mut s = String::from("# ");
    s.push_str(&database_name);
    s.push_str("\n\n");

    for table_info in table_infos {
        s.push_str("## ");
        s.push_str(&table_info.table_name);
        s.push_str("\n\n");
        s.push_str("| Field | Type | Null | Key | Default | Extra |\n| :-- | :-- | :-- | :-- | :-- | :-- |\n");
        for r in &table_info.records {
            s.push_str("| ");
            s.push_str(match &r.field {
                Some(v) => &v,
                None => "",
            });
            s.push_str(" | ");
            s.push_str(match &r.data_type {
                Some(v) => &v,
                None => "",
            });
            s.push_str(" | ");
            s.push_str(match &r.null {
                Some(v) => &v,
                None => "",
            });
            s.push_str(" | ");
            s.push_str(match &r.key {
                Some(v) => &v,
                None => "",
            });
            s.push_str(" | ");
            s.push_str(match &r.default {
                Some(v) => &v,
                None => "NULL",
            });
            s.push_str(" | ");
            s.push_str(match &r.extra {
                Some(v) => &v,
                None => "",
            });
            s.push_str(" |\n");
        }
        s.push('\n');
    }

    s.pop();
    match file.write_all(s.as_bytes()) {
        Err(e) => {
            return Err(e);
        }
        _ => {}
    };

    match file.flush() {
        Ok(_) => {
            return Ok(());
        }
        Err(e) => {
            return Err(e);
        }
    };
}

fn get_list_string(table_info: &TableInfo) -> String {
    let mut table_list = List::new();
    let mut s = String::from(&table_info.table_name);
    s.push('\n');

    // counting string length
    for record in &table_info.records {
        let c = match &record.field {
            Some(v) => v.len() as u32,
            None => 0,
        };
        if table_list.field_max < c {
            table_list.field_max = c;
        }

        let c = match &record.data_type {
            Some(v) => v.len() as u32,
            None => 0,
        };
        if table_list.type_max < c {
            table_list.type_max = c;
        }

        let c = match &record.null {
            Some(v) => v.len() as u32,
            None => 0,
        };
        if table_list.null_max < c {
            table_list.null_max = c;
        }

        let c = match &record.key {
            Some(v) => v.len() as u32,
            None => 0,
        };
        if table_list.key_max < c {
            table_list.key_max = c;
        }

        let c = match &record.default {
            Some(v) => v.len() as u32,
            None => 0,
        };
        if table_list.default_max < c {
            table_list.default_max = c;
        }

        let c = match &record.extra {
            Some(v) => v.len() as u32,
            None => 0,
        };
        if table_list.extra_max < c {
            table_list.extra_max = c;
        }
    }

    s.push_str(&table_list.get_header());

    for record in &table_info.records {
        s.push_str("| ");
        let v = match &record.field {
            Some(v) => v,
            None => "",
        };
        s.push_str(v);
        for _ in 0..(table_list.field_max - v.len() as u32) {
            s.push(' ');
        }
        s.push_str(" | ");

        let v = match &record.data_type {
            Some(v) => v,
            None => "",
        };
        s.push_str(v);
        for _ in 0..(table_list.type_max - v.len() as u32) {
            s.push(' ');
        }
        s.push_str(" | ");

        let v = match &record.null {
            Some(v) => v,
            None => "",
        };
        s.push_str(v);
        for _ in 0..(table_list.null_max - v.len() as u32) {
            s.push(' ');
        }
        s.push_str(" | ");

        let v = match &record.key {
            Some(v) => v,
            None => "",
        };
        s.push_str(v);
        for _ in 0..(table_list.key_max - v.len() as u32) {
            s.push(' ');
        }
        s.push_str(" | ");

        let v = match &record.default {
            Some(v) => v,
            None => "NULL",
        };
        s.push_str(v);
        for _ in 0..(table_list.default_max - v.len() as u32) {
            s.push(' ');
        }
        s.push_str(" | ");

        let v = match &record.extra {
            Some(v) => v,
            None => "",
        };
        s.push_str(v);
        for _ in 0..(table_list.extra_max - v.len() as u32) {
            s.push(' ');
        }
        s.push_str(" |\n");
    }
    s.push_str(&table_list.get_horizonal_border());

    s
}
