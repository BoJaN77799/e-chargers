import { VehicleDTO } from "src/modules/user/model/vehicleDTO"

export interface Report {
    date_from: number,
    date_to: number,
    chargers: Map<number, ReportItem>,
}

export interface ReportItem {
    charger: ChargerReportDTO,
    money_earned: number,
    used_energy: number,
}

export interface ChargerReportDTO {
    name: string,
    capacity: number,
    rating: number,
    price_per_hour: number,
    charging_speed: number,
}

export interface UserReportDTO {
    username: string
    email: string
    firstname: string
    lastname: string
    user_role: string
    strikes: number
    banned: boolean
    banned_at: number
    banned_until: number
}

export interface UserProfileDTO {
    username: string
    email: string
    firstname: string
    lastname: string
    strikes: number
    vechicles: VehicleDTO[]
}