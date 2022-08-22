use crate::types::*;

pub fn get_all_users() -> Result<Vec<UserReportDTO>, Box<dyn std::error::Error>> {
    let users: Vec<UserReportDTO> = reqwest::get("http://localhost:50001/api/users")?.json()?;

    return Ok(users);
}
