#[derive(Deserialize, Debug)]
pub struct Environment {
    pub database_url: String,
}

pub fn build_config() -> Environment {
    dotenv::dotenv().expect("Failed to read .env file");

    match envy::from_env::<Environment>() {
        Ok(cfg) => cfg,
        Err(e) => panic!(e),
    }
}
