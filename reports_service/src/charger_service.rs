use std::collections::HashMap;

use crate::types::*;

pub fn get_all_reservations_in_period(
    date_from: u64,
    date_to: u64,
) -> Result<Vec<ReservationDTO>, Box<dyn std::error::Error>> {
    let mut url: String = "http://localhost:50003/api/reservations/".to_owned();
    let date_from_str = date_from.to_string();
    let date_to_str = date_to.to_string();

    url.push_str(&date_from_str);
    url.push_str("/");
    url.push_str(&date_to_str);

    let reservations: Vec<ReservationDTO> = reqwest::get(&url)?.json()?;

    return Ok(reservations);
}

pub fn get_charger(charger_id: u32) -> Result<ChargerDTO, Box<dyn std::error::Error>> {
    let mut url: String = "http://localhost:50002/api/chargers/report/".to_owned();
    let charger_id_str = charger_id.to_string();

    url.push_str(&charger_id_str);

    let charger: ChargerDTO = reqwest::get(&url)?.json()?;

    return Ok(charger);
}

pub fn get_all_chargers(reservations: &[ReservationDTO]) -> HashMap<u32, ChargerDTO> {
    let mut map_chargers: HashMap<u32, ChargerDTO> = HashMap::new();
    
    reservations.iter().for_each(|el| {
        if !map_chargers.contains_key(&el.charger_id) {
            let charger: ChargerDTO = get_charger(el.charger_id).ok().unwrap();
            map_chargers.insert(el.charger_id, charger);
        }
    });

    map_chargers
}

pub fn create_report_items(
    reservations: &[ReservationDTO],
    map_chargers: HashMap<u32, ChargerDTO>,
) -> HashMap<u32, ReportItem> {
    let mut map_report_items: HashMap<u32, ReportItem> = HashMap::new();

    reservations.iter().for_each(|el| {
        let temp_charger = map_chargers.get(&el.charger_id).unwrap();
        if !map_report_items.contains_key(&el.charger_id) {
            map_report_items.insert(
                el.charger_id,
                ReportItem {
                    charger: ChargerDTO {
                        name: temp_charger.name.to_owned(),
                        capacity: temp_charger.capacity,
                        rating: temp_charger.rating,
                        price_per_hour: temp_charger.price_per_hour,
                        charging_speed: temp_charger.charging_speed,
                    },
                    money_earned: (temp_charger.price_per_hour as f32) / 60.0 * el.duration as f32,
                    used_energy: temp_charger.charging_speed * el.duration as i32,
                },
            );
        } else {
            map_report_items.insert(
                el.charger_id,
                ReportItem {
                    charger: ChargerDTO {
                        name: temp_charger.name.to_owned(),
                        capacity: temp_charger.capacity,
                        rating: temp_charger.rating,
                        price_per_hour: temp_charger.price_per_hour,
                        charging_speed: temp_charger.charging_speed,
                    },
                    money_earned: map_report_items.get(&el.charger_id).unwrap().money_earned
                        + (temp_charger.price_per_hour as f32) / 60.0 * el.duration as f32,
                    used_energy: map_report_items.get(&el.charger_id).unwrap().used_energy
                        + temp_charger.charging_speed * el.duration as i32,
                },
            );
        }
    });

    map_report_items
}
