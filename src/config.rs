use clap::{Arg, Command};
use rpassword::prompt_password;

const APP_NAME: &str = "printdb";
const VERSION: &str = "2.0.1";
const AUTHOR: &str = "H-Kiku482 <h.kikuchi482@gmail.com>";
const ABOUT: &str = "Print your MySQL database";

pub struct Config {
    pub host: String,
    pub port: String,
    pub user: String,
    pub password: bool,
    pub print_text_file: String,
    pub print_markdown: String,
    pub file_overwrite: bool,
    pub print_on_console: bool,
    pub database: String,
    pub plane_password: String,
}

impl Config {
    pub fn new() -> Config {
        let app = Command::new(APP_NAME)
            .version(VERSION)
            .author(AUTHOR)
            .about(ABOUT)
            .arg(
                Arg::new("host")
                    .help("Connect to host.")
                    .short('h')
                    .long("host")
                    .default_value("127.0.0.1")
                    .takes_value(true),
            )
            .arg(
                Arg::new("port")
                    .help("Port number to connect to server.")
                    .short('P')
                    .long("port")
                    .default_value("3306")
                    .takes_value(true),
            )
            .arg(
                Arg::new("user")
                    .help("User for login.")
                    .short('u')
                    .long("user")
                    .default_value("root")
                    .takes_value(true),
            )
            .arg(
                Arg::new("password")
                    .help("Password to use connecting to server.")
                    .short('p')
                    .long("password"),
            )
            .arg(
                Arg::new("text file path")
                    .help("Output to a given file as plain text.")
                    .short('t')
                    .long("text")
                    .default_value("")
                    .takes_value(true),
            )
            .arg(
                Arg::new("markdown file path")
                    .help("Output to a given file as markdown file format.")
                    .short('m')
                    .long("markdown")
                    .default_value("")
                    .takes_value(true),
            )
            .arg(
                Arg::new("overwrite")
                    .help("Overwriting text and markdown files")
                    .short('o')
                    .long("overwrite"),
            )
            .arg(Arg::new("cli").help("Print on CLI").short('c').long("cli"))
            .arg(
                Arg::new("database")
                    .help("Target database name to printing")
                    .required(true),
            );

        let parsed_args = app.get_matches();

        let host = match parsed_args.value_of("host") {
            Some(v) => v,
            None => panic!(),
        };

        let port = match parsed_args.value_of("port") {
            Some(v) => v,
            None => panic!(),
        };

        let user = match parsed_args.value_of("user") {
            Some(v) => v,
            None => panic!(),
        };

        let password = parsed_args.is_present("password");

        let text = match parsed_args.value_of("text file path") {
            Some(v) => v,
            None => panic!(),
        };

        let markdown = match parsed_args.value_of("markdown file path") {
            Some(v) => v,
            None => panic!(),
        };

        let overwrite = parsed_args.is_present("overwrite");

        let mut cli = parsed_args.is_present("cli");

        let database = match parsed_args.value_of("database") {
            Some(v) => v,
            None => panic!(),
        };

        if text == "" && markdown == "" {
            cli = true;
        }

        Config {
            host: String::from(host),
            port: String::from(port),
            user: String::from(user),
            password: password,
            print_text_file: String::from(text),
            print_markdown: String::from(markdown),
            file_overwrite: overwrite,
            print_on_console: cli,
            database: String::from(database),
            plane_password: if password {
                prompt_password("Enter password: ").unwrap()
            } else {
                String::new()
            },
        }
    }
}
