mod config;
mod db_connect;
mod printdb;

use config::Config;
use db_connect::*;
use printdb::{
    format::{get_std_out, write_as_markdown, write_as_text},
    get_table_info,
};

pub fn run() {
    let config = Config::new();

    let dbcon: MysqlConnection = match get_connection(&config) {
        Ok(v) => v,
        Err(e) => {
            eprintln!("failed to connect: {}", e);
            return;
        }
    };

    let table_infos = match get_table_info(&dbcon, &config.database) {
        Ok(v) => v,
        Err(e) => {
            eprintln!("failed to get table information: {}", e);
            return;
        }
    };

    if config.print_on_console {
        print!("{}", get_std_out(&table_infos));
    }

    if !config.print_text_file.is_empty() {
        match write_as_text(&table_infos, &config.print_text_file, config.file_overwrite) {
            Err(e) => {
                eprintln!("failed to write: {}", e);
            }
            _ => (),
        }
    }

    if !config.print_markdown.is_empty() {
        match write_as_markdown(
            &config.database,
            &table_infos,
            &config.print_markdown,
            config.file_overwrite,
        ) {
            Err(e) => {
                eprintln!("failed to write: {}", e);
            }
            _ => (),
        }
    }
}
