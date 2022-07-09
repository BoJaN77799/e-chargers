export interface AdressDTO {
    street: string
    city: string
    country: string
    postal_code: number
    latitude: number
    longitude: number
}

export interface PlugDTO {
    price_per_hour: string
    type: string
    charging_speed: string
}

export interface ChargerDTO {
    name: string
    address: AdressDTO
    work_time: string,
    capacity: number,
    description: string,
    plugs: PlugDTO[]
}