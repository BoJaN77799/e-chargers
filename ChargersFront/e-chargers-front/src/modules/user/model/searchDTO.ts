export interface SearchDTO {
    searchField: string
    workTimeFrom: number,
    workTimeTo: number,
    capacity: number,
    pricePerHourFrom: number,
    pricePerHourTo: number,
    type: string,
    chargingSpeedFrom: number,
    chargingSpeedTo: number,
}