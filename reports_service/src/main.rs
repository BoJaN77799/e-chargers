#![feature(proc_macro_hygiene, decl_macro)]

extern crate rocket;

extern crate reqwest;

use reqwest::Error;

use rocket::{get, routes};
use rocket_contrib::json::Json;

use std::collections::HashMap;

pub mod charger_service;
pub mod types;
pub mod users_service;

use charger_service::*;
use types::*;
use users_service::*;

#[get("/chargers/<date_from>/<date_to>")]
fn get_reservations(date_from: u64, date_to: u64) -> Result<Json<Report>, Error> {
    let reservations: Vec<ReservationDTO> = get_all_reservations_in_period(date_from, date_to)
        .ok()
        .unwrap();

    let map_chargers: HashMap<u32, ChargerDTO> = get_all_chargers(&reservations);

    let map_report_items: HashMap<u32, ReportItem> =
        create_report_items(&reservations, map_chargers);

    let report: Report = Report {
        date_from: date_from,
        date_to: date_to,
        chargers: map_report_items,
    };

    Ok(Json(report))
}

#[get("/users")]
fn get_users() -> Result<Json<Vec<UserReportDTO>>, Error> {
    let users: Vec<UserReportDTO> = get_all_users().ok().unwrap();
    Ok(Json(users))
}

#[get("/")]
fn index() -> &'static str {
    "Hello, world."
}

fn main() {
    rocket::ignite()
        .mount("/api/reports", routes![index, get_reservations, get_users])
        .launch();
}
