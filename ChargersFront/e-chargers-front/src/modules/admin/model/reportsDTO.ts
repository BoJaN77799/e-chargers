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