use sqlx::postgres::PgPoolOptions;

use async_graphql::http::{playground_source, GraphQLPlaygroundConfig};
use async_graphql::{Context, EmptyMutation, EmptySubscription, Object, Schema};
use async_graphql_warp::{graphql, BadRequest, Response};
use std::convert::Infallible;
use warp::http::StatusCode;
use warp::{http::Response as HttpResponse, Filter, Rejection};

#[macro_use]
extern crate serde_derive;

#[macro_use]
extern crate sqlx;

mod config;

struct User {
    id: i32,
    name: String,
}

struct QueryRoot;

#[Object]
impl QueryRoot {
    async fn value(&self, ctx: &Context<'_>) -> i32 {
        unimplemented!()
    }

    async fn user(&self, ctx: &Context<'_>) -> User {
        query_as!(User, "SELECT * FROM users WHERE id = $1", id)
            .fetch_all(&pool)
            .await?
    }
}

type GQLSchema = Schema<QueryRoot, EmptyMutation, EmptySubscription>;

#[tokio::main]
async fn main() -> Result<(), sqlx::Error> {
    println!("Hello, world!");

    let schema = Schema::new(QueryRoot, EmptyMutation, EmptySubscription);

    let cfg = config::build_config();

    let pool = PgPoolOptions::new()
        .max_connections(5)
        .connect(&cfg.database_url)
        .await?;

    // migrate!().run(&pool).await?;

    let id = 1;

    let user = query_as!(User, "SELECT * FROM users WHERE id = $1", id)
        .fetch_all(&pool)
        .await?;

    println!("{}", user[0].name);

    let graphql = graphql(schema).and_then(
        |(schema, request): (
            Schema<QueryRoot, EmptyMutation, EmptySubscription>,
            async_graphql::Request,
        )| async move { Ok::<_, Infallible>(Response::from(schema.execute(request).await)) },
    );

    let playground = warp::path::end().and(warp::get()).map(|| {
        HttpResponse::builder()
            .header("content-type", "text/html")
            .body(playground_source(GraphQLPlaygroundConfig::new("/")))
    });

    let routes = playground.or(graphql).recover(|err: Rejection| async move {
        if let Some(BadRequest(err)) = err.find() {
            return Ok::<_, Infallible>(warp::reply::with_status(
                err.to_string(),
                StatusCode::BAD_REQUEST,
            ));
        }

        Ok(warp::reply::with_status(
            "INTERNAL_SERVER_ERROR".to_string(),
            StatusCode::INTERNAL_SERVER_ERROR,
        ))
    });

    warp::serve(routes).run(([127, 0, 0, 1], 4000)).await;

    Ok(())
}
