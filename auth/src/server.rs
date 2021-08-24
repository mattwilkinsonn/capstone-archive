use tonic::{transport::Server, Request, Response, Status};

use auth_proto::auth_server::{Auth, AuthServer};

use auth_proto::{LoginReply, LoginRequest, Role, User};

pub mod auth_proto {
    tonic::include_proto!("auth"); // The string specified here must match the proto package name
}

#[derive(Debug, Default)]
pub struct Authenticator {}

#[tonic::async_trait]
impl Auth for Authenticator {
    async fn login(
        &self,
        _request: Request<LoginRequest>,
    ) -> Result<tonic::Response<LoginReply>, Status> {
        let user = User {
            id: "1".to_string(),
            created_at: "10/21/2021".to_string(),
            email: "mattwilki17@gmail.com".to_string(),
            role: Role::Admin.into(),
            updated_at: "10/21/2021".to_string(),
            username: "zireael".to_string(),
            deleted_at: None,
        };

        Ok(Response::new(LoginReply { user: Some(user) }))
    }
}

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    let addr = "127.0.0.1:50051".parse()?;
    let authenticator = Authenticator::default();

    Server::builder()
        .add_service(AuthServer::new(authenticator))
        .serve(addr)
        .await?;

    Ok(())
}
