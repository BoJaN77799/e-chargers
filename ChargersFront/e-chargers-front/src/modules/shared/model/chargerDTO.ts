export interface AdressDTO {
    street: string
    city: string
    country: string
    postal_code: number
    latitude: number
    longitude: number
}


export interface ChargerDTO {
    name: string
    address: AdressDTO
    work_time: string,
    capacity: number,
    description: string,
    plugs: string[],
    price_per_hour: string,
    charging_speed: string
}