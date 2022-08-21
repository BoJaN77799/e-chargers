use std::collections::HashMap;

use serde::Deserialize;
use serde::Serialize;

#[derive(Deserialize, Serialize, Debug)]
pub struct ReservationDTO {
    pub username: String,
    pub charger_id: u32,
    pub vehicle_id: u32,
    pub date_from: u64,
    pub duration: u32,
}

#[derive(Deserialize, Serialize, Debug)]
pub struct ChargerDTO {
    pub name: String,
    pub capacity: u32,
    pub rating: f32,
    pub price_per_hour: i32,
    pub charging_speed: i32,
}

#[derive(Deserialize, Serialize, Debug)]
pub struct Report {
    pub date_from: u64,
    pub date_to: u64,
    pub chargers: HashMap<u32, ReportItem>,
}

#[derive(Deserialize, Serialize, Debug)]
pub struct ReportItem {
    pub charger: ChargerDTO,
    pub money_earned: f32,
    pub used_energy: i32,
}

#[derive(Deserialize, Serialize, Debug)]
pub struct UserReportDTO {
    pub username: String,
    pub email: String,
    pub firstname: String,
    pub lastname: String,
    pub role: String,
    pub strikes: i32,
    pub banned: bool,
    pub banned_at: u64,
    pub banned_until: u64,
}
