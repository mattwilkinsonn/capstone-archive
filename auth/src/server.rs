use tonic::{transport::Server, Request, Response, Status};

use auth_proto::auth_server::{Auth, AuthServer};

use auth_proto::{LoginReply, LoginRequest, LogoutReply, Role, User};

pub mod auth_proto {
    tonic::include_proto!("auth"); // The string specified here must match the proto package name
}

#[derive(Debug, Default)]
pub struct Authenticator {}

#[tonic::async_trait]
impl Auth for Authenticator {
    async fn login(&self, _request: Request<LoginRequest>) -> Result<Response<LoginReply>, Status> {
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

    async fn logout(&self, _request: Request<()>) -> Result<Response<LogoutReply>, Status> {
        let logout = LogoutReply { logout: true };

        Ok(Response::new(logout))
    }
}

#[tokio::main]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    let addr = "localhost:50051".parse()?;

    let authenticator = Authenticator::default();

    Server::builder()
        .add_service(AuthServer::new(authenticator))
        .serve(addr)
        .await?;

    Ok(())
}
